package b2b

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/stytchauth/stytch-go/v16/stytch/config"
	"io"
	"net/http"
	"net/url"
	"strings"
	"time"

	"github.com/MicahParks/keyfunc/v2"
	"github.com/golang-jwt/jwt/v5"
	"github.com/stytchauth/stytch-go/v16/stytch"
	"github.com/stytchauth/stytch-go/v16/stytch/b2b/idp"
	"github.com/stytchauth/stytch-go/v16/stytch/b2b/rbac"
	"github.com/stytchauth/stytch-go/v16/stytch/b2b/sessions"
	"github.com/stytchauth/stytch-go/v16/stytch/shared"
	"github.com/stytchauth/stytch-go/v16/stytch/stytcherror"
)

type IDPClient struct {
	C           stytch.Client
	JWKS        *keyfunc.JWKS
	PolicyCache *PolicyCache
}

func NewIDPClient(c stytch.Client, jwks *keyfunc.JWKS, policyCache *PolicyCache) *IDPClient {
	return &IDPClient{
		C:           c,
		JWKS:        jwks,
		PolicyCache: policyCache,
	}
}

func (c *IDPClient) IntrospectTokenNetwork(
	ctx context.Context,
	body *idp.IntrospectTokenNetworkParams,
) (*idp.IntrospectTokenResponse, error) {
	cfg := c.C.GetConfig()
	client := c.C.GetHTTPClient()
	path := string(cfg.BaseURI) + "/v1/public/" + cfg.ProjectID + "/oauth2/introspect"

	data := url.Values{}
	data.Add("token", body.Token)
	data.Add("client_id", body.ClientID)
	if body.ClientSecret != nil {
		data.Add("client_secret", *body.ClientSecret)
	}
	if body.TokenTypeHint != nil {
		data.Add("token_type_hint", *body.TokenTypeHint)
	}

	req, err := http.NewRequestWithContext(ctx, "POST", path, strings.NewReader(data.Encode()))
	if err != nil {
		return nil, fmt.Errorf("error creating http request: %w", err)
	}

	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Add("User-Agent", "Stytch Go v"+config.APIVersion)

	res, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("error sending http request: %w", err)
	}
	defer func() {
		res.Body.Close()
	}()

	bytes, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, fmt.Errorf("error reading http request: %w", err)
	}

	if res.StatusCode != 200 {
		// Attempt to unmarshal into Stytch error format
		var stytchErr stytcherror.OAuth2Error
		if err = json.Unmarshal(bytes, &stytchErr); err != nil {
			return nil, fmt.Errorf("error decoding http request: %w", err)
		}
		stytchErr.StatusCode = res.StatusCode
		return nil, stytchErr
	}

	var tokenRes idp.IntrospectTokenResponse
	if err = json.Unmarshal(bytes, &tokenRes); err != nil {
		return nil, fmt.Errorf("error decoding http request: %w", err)
	}
	if !tokenRes.Active {
		return nil, stytcherror.NewInvalidOAuth2TokenError()
	}
	if body.AuthorizationCheck != nil {
		policy, err := c.PolicyCache.Get(ctx)
		if err != nil {
			return nil, fmt.Errorf("failed to get cached policy: %w", err)
		}

		tokenScopes := strings.Split(strings.TrimSpace(tokenRes.Scope), " ")

		err = shared.PerformScopeAuthorizationCheck(policy, tokenScopes, tokenRes.Organization.OrganizationID, body.AuthorizationCheck)
		if err != nil {
			return nil, err
		}
	}

	return &tokenRes, nil
}

func (c *IDPClient) IntrospectTokenLocal(
	ctx context.Context,
	req *idp.IntrospectTokenLocalParams,
) (*idp.IntrospectTokenResponse, error) {
	if c.JWKS == nil {
		return nil, stytcherror.ErrJWKSNotInitialized
	}

	// It's difficult to extract all sets of claims (standard/registered, Stytch, custom) all at
	// once. So we parse the token twice.
	//
	// The first parse is for validating and extracting the statically-known claims. It will fail
	// if the token is invalid for any reason.
	//
	// The second parse is for extracting the custom claims.
	var staticClaims idp.IntrospectTokenClaims
	err := shared.ValidateJWTToken(shared.ValidateJWTTokenParams{
		Token:          req.Token,
		StaticClaims:   &staticClaims,
		KeyFunc:        c.JWKS.Keyfunc,
		Audience:       c.C.GetConfig().ProjectID,
		Issuer:         fmt.Sprintf("stytch.com/%s", c.C.GetConfig().ProjectID),
		FallbackIssuer: string(c.C.GetConfig().BaseURI),
	})
	if err != nil {
		return nil, err
	}

	if req.MaxTokenAge != 0 {
		iat, err := staticClaims.GetIssuedAt()
		if err != nil {
			return nil, err
		}
		if iat.Add(req.MaxTokenAge).Before(time.Now()) {
			// The JWT is valid, but older than the tolerable maximum age.
			return nil, sessions.ErrJWTTooOld
		}
	}

	// The token has already been verified at this point, so its claims and signature are all
	// valid. This call to ParseUnverified is _only_ for extracting the remaining custom claims.
	//
	// Using ParseWithClaims again would cause this to re-validate the token's timestamps and
	// signature, which fail if it was very close to its expiration on the previous parse.
	var customClaims jwt.MapClaims
	_, _, err = jwt.NewParser().ParseUnverified(req.Token, &customClaims)
	if err != nil {
		return nil, fmt.Errorf("failed to parse access token: %w", err)
	}

	// Remove all the reserved claims that are already present in staticClaims.
	for key := range customClaims {
		if shared.ReservedClaim(key) || key == "jti" || key == "scope" || key == "client_id" {
			delete(customClaims, key)
		}
	}

	var policy *rbac.Policy
	if req.AuthorizationCheck != nil {
		policy, err = c.PolicyCache.Get(ctx)
		if err != nil {
			return nil, fmt.Errorf("failed to get cached policy: %w", err)
		}

		tokenScopes := strings.Split(strings.TrimSpace(staticClaims.Scope), " ")

		err = shared.PerformScopeAuthorizationCheck(policy, tokenScopes, staticClaims.Organization.OrganizationID, req.AuthorizationCheck)
		if err != nil {
			return nil, err
		}
	}

	return marshalJWTIntoResponse(staticClaims, customClaims)
}

func marshalJWTIntoResponse(staticClaims idp.IntrospectTokenClaims, customClaims jwt.MapClaims) (*idp.IntrospectTokenResponse, error) {
	return &idp.IntrospectTokenResponse{
		Active:       true,
		TokenType:    "access_token",
		Issuer:       staticClaims.Issuer,
		Subject:      staticClaims.Subject,
		Audience:     staticClaims.Audience,
		Scope:        staticClaims.Scope,
		ClientID:     staticClaims.ClientID,
		Expiry:       staticClaims.ExpiresAt,
		IssuedAt:     staticClaims.IssuedAt,
		Organization: staticClaims.Organization,
		CustomClaims: customClaims,
	}, nil
}
