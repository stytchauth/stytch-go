package b2b

import (
	"context"
	"fmt"
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

func (c *IDPClient) IntrospectTokenLocal(
	ctx context.Context,
	req *idp.IntrospectTokenLocalParams,
) (*idp.IntrospectTokenClaims, error) {
	if c.JWKS == nil {
		return nil, stytcherror.ErrJWKSNotInitialized
	}

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
		if shared.ReservedClaim(key) {
			delete(customClaims, key)
		}
	}
	staticClaims.CustomClaims = customClaims

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

	return &staticClaims, nil
}
