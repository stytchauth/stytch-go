package b2b

// !!!
// WARNING: This file is autogenerated
// Only modify code within MANUAL() sections
// or your changes may be overwritten later!
// !!!

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	"github.com/MicahParks/keyfunc/v2"
	"github.com/golang-jwt/jwt/v5"
	"github.com/mitchellh/mapstructure"
	"github.com/stytchauth/stytch-go/v17/stytch"
	"github.com/stytchauth/stytch-go/v17/stytch/b2b/rbac"
	"github.com/stytchauth/stytch-go/v17/stytch/b2b/sessions"
	"github.com/stytchauth/stytch-go/v17/stytch/shared"
	"github.com/stytchauth/stytch-go/v17/stytch/stytcherror"
)

type SessionsClient struct {
	C           stytch.Client
	JWKS        *keyfunc.JWKS
	PolicyCache *PolicyCache
}

func NewSessionsClient(c stytch.Client, jwks *keyfunc.JWKS, policyCache *PolicyCache) *SessionsClient {
	return &SessionsClient{
		C:           c,
		JWKS:        jwks,
		PolicyCache: policyCache,
	}
}

// Get: Retrieves all active Sessions for a Member.
func (c *SessionsClient) Get(
	ctx context.Context,
	body *sessions.GetParams,
) (*sessions.GetResponse, error) {
	queryParams := make(map[string]string)
	if body != nil {
		queryParams["organization_id"] = body.OrganizationID
		queryParams["member_id"] = body.MemberID
	}

	headers := make(map[string][]string)

	var retVal sessions.GetResponse
	err := c.C.NewRequest(
		ctx,
		stytch.RequestParams{
			Method:      "GET",
			Path:        "/v1/b2b/sessions",
			QueryParams: queryParams,
			Body:        nil,
			V:           &retVal,
			Headers:     headers,
		},
	)
	return &retVal, err
}

// Authenticate: Authenticates a Session and updates its lifetime by the specified
// `session_duration_minutes`. If the `session_duration_minutes` is not specified, a Session will not be
// extended. This endpoint requires either a `session_jwt` or `session_token` be included in the request.
// It will return an error if both are present.
//
// You may provide a JWT that needs to be refreshed and is expired according to its `exp` claim. A new JWT
// will be returned if both the signature and the underlying Session are still valid. See our
// [How to use Stytch Session JWTs](https://stytch.com/docs/b2b/guides/sessions/resources/using-jwts) guide
// for more information.
//
// If an `authorization_check` object is passed in, this method will also check if the Member is authorized
// to perform the given action on the given Resource in the specified. A is authorized if their Member
// Session contains a Role, assigned
// [explicitly or implicitly](https://stytch.com/docs/b2b/guides/rbac/role-assignment), with adequate
// permissions.
// In addition, the `organization_id` passed in the authorization check must match the Member's
// Organization.
//
// If the Member is not authorized to perform the specified action on the specified Resource, or if the
// `organization_id` does not match the Member's Organization, a 403 error will be thrown.
// Otherwise, the response will contain a list of Roles that satisfied the authorization check.
func (c *SessionsClient) Authenticate(
	ctx context.Context,
	body *sessions.AuthenticateParams,
) (*sessions.AuthenticateResponse, error) {
	var jsonBody []byte
	var err error
	if body != nil {
		jsonBody, err = json.Marshal(body)
		if err != nil {
			return nil, stytcherror.NewClientLibraryError("error marshaling request body")
		}
	}

	headers := make(map[string][]string)

	var retVal sessions.AuthenticateResponse
	err = c.C.NewRequest(
		ctx,
		stytch.RequestParams{
			Method:      "POST",
			Path:        "/v1/b2b/sessions/authenticate",
			QueryParams: nil,
			Body:        jsonBody,
			V:           &retVal,
			Headers:     headers,
		},
	)
	return &retVal, err
}

// AuthenticateWithClaims fills in the claims pointer with custom claims from the response.
// Pass in a map with the types of values you're expecting so that this function can marshal
// the claims from the response. See ExampleClient_AuthenticateWithClaims_map,
// ExampleClient_AuthenticateWithClaims_struct for examples
func (c *SessionsClient) AuthenticateWithClaims(
	ctx context.Context,
	body *sessions.AuthenticateParams,
	claims any,
) (*sessions.AuthenticateResponse, error) {
	var jsonBody []byte
	var err error
	if body != nil {
		jsonBody, err = json.Marshal(body)
		if err != nil {
			return nil, stytcherror.NewClientLibraryError("error marshaling request body")
		}
	}

	headers := make(map[string][]string)

	b, err := c.C.RawRequest(
		ctx,
		stytch.RequestParams{
			Method:      "POST",
			Path:        "/v1/b2b/sessions/authenticate",
			QueryParams: nil,
			Body:        jsonBody,
			Headers:     headers,
		},
	)
	if err != nil {
		return nil, err
	}

	// First extract the Stytch data.
	var retVal sessions.AuthenticateResponse
	if err := json.Unmarshal(b, &retVal); err != nil {
		return nil, fmt.Errorf("unmarshal sessions.AuthenticateResponse: %w", err)
	}

	if claims == nil {
		return &retVal, nil
	}

	if m, ok := claims.(*map[string]any); ok {
		*m = retVal.MemberSession.CustomClaims
		return &retVal, nil
	}

	// This is where we need to convert claims into a claimsMap
	decoder, err := mapstructure.NewDecoder(&mapstructure.DecoderConfig{
		Result:  &claims,
		TagName: "json",
	})
	if err != nil {
		return nil, err
	}

	err = decoder.Decode(retVal.MemberSession.CustomClaims)
	if err != nil {
		return nil, err
	}

	return &retVal, err
}

// Revoke a Session and immediately invalidate all its tokens. To revoke a specific Session, pass either
// the `member_session_id`, `session_token`, or `session_jwt`. To revoke all Sessions for a Member, pass
// the `member_id`.
func (c *SessionsClient) Revoke(
	ctx context.Context,
	body *sessions.RevokeParams,
	methodOptions ...*sessions.RevokeRequestOptions,
) (*sessions.RevokeResponse, error) {
	var jsonBody []byte
	var err error
	if body != nil {
		jsonBody, err = json.Marshal(body)
		if err != nil {
			return nil, stytcherror.NewClientLibraryError("error marshaling request body")
		}
	}

	headers := make(map[string][]string)
	for _, methodOption := range methodOptions {
		headers = methodOption.AddHeaders(headers)
	}

	var retVal sessions.RevokeResponse
	err = c.C.NewRequest(
		ctx,
		stytch.RequestParams{
			Method:      "POST",
			Path:        "/v1/b2b/sessions/revoke",
			QueryParams: nil,
			Body:        jsonBody,
			V:           &retVal,
			Headers:     headers,
		},
	)
	return &retVal, err
}

// Exchange: Use this endpoint to exchange a's existing session for another session in a different. This
// can be used to accept an invite, but not to create a new member via domain matching.
//
// To create a new member via domain matching, use the
// [Exchange Intermediate Session](https://stytch.com/docs/b2b/api/exchange-intermediate-session) flow
// instead.
//
// Only Email Magic Link, OAuth, and SMS OTP factors can be transferred between sessions. Other
// authentication factors, such as password factors, will not be transferred to the new session.
// Any OAuth Tokens owned by the Member will not be transferred to the new Organization.
// SMS OTP factors can be used to fulfill MFA requirements for the target Organization if both the original
// and target Member have the same phone number and the phone number is verified for both Members.
// HubSpot and Slack OAuth registrations will not be transferred between sessions. Instead, you will
// receive a corresponding factor with type `"oauth_exchange_slack"` or `"oauth_exchange_hubspot"`
//
// If the Member is required to complete MFA to log in to the Organization, the returned value of
// `member_authenticated` will be `false`, and an `intermediate_session_token` will be returned.
// The `intermediate_session_token` can be passed into the
// [OTP SMS Authenticate endpoint](https://stytch.com/docs/b2b/api/authenticate-otp-sms) to complete the
// MFA step and acquire a full member session.
// The `intermediate_session_token` can also be used with the
// [Exchange Intermediate Session endpoint](https://stytch.com/docs/b2b/api/exchange-intermediate-session)
// or the
// [Create Organization via Discovery endpoint](https://stytch.com/docs/b2b/api/create-organization-via-discovery) to join a different Organization or create a new one.
// The `session_duration_minutes` and `session_custom_claims` parameters will be ignored.
func (c *SessionsClient) Exchange(
	ctx context.Context,
	body *sessions.ExchangeParams,
) (*sessions.ExchangeResponse, error) {
	var jsonBody []byte
	var err error
	if body != nil {
		jsonBody, err = json.Marshal(body)
		if err != nil {
			return nil, stytcherror.NewClientLibraryError("error marshaling request body")
		}
	}

	headers := make(map[string][]string)

	var retVal sessions.ExchangeResponse
	err = c.C.NewRequest(
		ctx,
		stytch.RequestParams{
			Method:      "POST",
			Path:        "/v1/b2b/sessions/exchange",
			QueryParams: nil,
			Body:        jsonBody,
			V:           &retVal,
			Headers:     headers,
		},
	)
	return &retVal, err
}

// Migrate a session from an external OIDC compliant endpoint. Stytch will call the external UserInfo
// endpoint defined in your Stytch Project settings in the [Dashboard](https://stytch.com/docs/dashboard),
// and then perform a lookup using the `session_token`. If the response contains a valid email address,
// Stytch will attempt to match that email address with an existing in your and create a Stytch Session.
// You will need to create the member before using this endpoint.
func (c *SessionsClient) Migrate(
	ctx context.Context,
	body *sessions.MigrateParams,
) (*sessions.MigrateResponse, error) {
	var jsonBody []byte
	var err error
	if body != nil {
		jsonBody, err = json.Marshal(body)
		if err != nil {
			return nil, stytcherror.NewClientLibraryError("error marshaling request body")
		}
	}

	headers := make(map[string][]string)

	var retVal sessions.MigrateResponse
	err = c.C.NewRequest(
		ctx,
		stytch.RequestParams{
			Method:      "POST",
			Path:        "/v1/b2b/sessions/migrate",
			QueryParams: nil,
			Body:        jsonBody,
			V:           &retVal,
			Headers:     headers,
		},
	)
	return &retVal, err
}

// GetJWKS: Get the JSON Web Key Set (JWKS) for a project.
//
// JWKS are rotated every ~6 months. Upon rotation, new JWTs will be signed using the new key set, and both
// key sets will be returned by this endpoint for a period of 1 month.
//
// JWTs have a set lifetime of 5 minutes, so there will be a 5 minute period where some JWTs will be signed
// by the old JWKS, and some JWTs will be signed by the new JWKS. The correct JWKS to use for validation is
// determined by matching the `kid` value of the JWT and JWKS.
//
// If you're using one of our [backend SDKs](https://stytch.com/docs/b2b/sdks), the JWKS roll will be
// handled for you.
//
// If you're using your own JWT validation library, many have built-in support for JWKS rotation, and
// you'll just need to supply this API endpoint. If not, your application should decide which JWKS to use
// for validation by inspecting the `kid` value.
//
// See our
// [How to use Stytch Session JWTs](https://stytch.com/docs/b2b/guides/sessions/resources/using-jwts) guide
// for more information.
func (c *SessionsClient) GetJWKS(
	ctx context.Context,
	body *sessions.GetJWKSParams,
) (*sessions.GetJWKSResponse, error) {
	headers := make(map[string][]string)

	var retVal sessions.GetJWKSResponse
	err := c.C.NewRequest(
		ctx,
		stytch.RequestParams{
			Method:      "GET",
			Path:        fmt.Sprintf("/v1/b2b/sessions/jwks/%s", body.ProjectID),
			QueryParams: nil,
			Body:        nil,
			V:           &retVal,
			Headers:     headers,
		},
	)
	return &retVal, err
}

// MANUAL(AuthenticateJWT)(SERVICE_METHOD)
// ADDIMPORT: "encoding/json"
// ADDIMPORT: "time"
// ADDIMPORT: "github.com/golang-jwt/jwt/v5"
// ADDIMPORT: "github.com/MicahParks/keyfunc/v2"
// ADDIMPORT: "github.com/stytchauth/stytch-go/v17/stytch/stytcherror"

func (c *SessionsClient) AuthenticateJWT(
	ctx context.Context,
	params *sessions.AuthenticateJWTParams,
) (*sessions.AuthenticateResponse, error) {
	if params.Body.SessionJWT == "" || params.MaxTokenAge == time.Duration(0) {
		return c.Authenticate(ctx, params.Body)
	}

	session, err := c.AuthenticateJWTLocal(ctx, params.Body.SessionJWT, params.MaxTokenAge, params.Body.AuthorizationCheck)
	if err != nil {
		// JWT couldn't be verified locally. Check with the Stytch API.
		return c.Authenticate(ctx, params.Body)
	}

	return &sessions.AuthenticateResponse{
		MemberSession: *session,
	}, nil
}

func (c *SessionsClient) AuthenticateJWTWithClaims(
	ctx context.Context,
	maxTokenAge time.Duration,
	body *sessions.AuthenticateParams,
	claims map[string]any,
) (*sessions.AuthenticateResponse, error) {
	// This method has a different signature than AuthenticateWithClaims, which we can't change in
	// this version of the library. For backward compatibility, populate the claims map by
	// mutating it instead of replacing it like the non-JWT version does.
	//
	// TODO(v12.x): Change claims to `any`, also allow pointer-to-map and pointer-to-struct.
	// TODO(v14): Remove support for populating a pre-existing map this way.

	var resp *sessions.AuthenticateResponse

	// Some special cases can force remote authentication. Otherwise, prefer local validation.
	if body.SessionJWT == "" || maxTokenAge == time.Duration(0) {
		var err error
		resp, err = c.AuthenticateWithClaims(ctx, body, claims)
		if err != nil {
			return nil, err
		}
	} else if session, err := c.AuthenticateJWTLocal(ctx, body.SessionJWT, maxTokenAge, body.AuthorizationCheck); err == nil {
		resp = &sessions.AuthenticateResponse{
			MemberSession: *session,
		}
	} else {
		// JWT couldn't be verified locally. Check with the Stytch API.
		resp, err = c.Authenticate(ctx, body)
		if err != nil {
			return nil, err
		}
	}

	// Populate claims if possible.
	if claims != nil {
		for key, val := range resp.MemberSession.CustomClaims {
			claims[key] = val
		}
	}

	return resp, nil
}

// ADDIMPORT: "github.com/stytchauth/stytch-go/v17/stytch/shared"
func (c *SessionsClient) AuthenticateJWTLocal(
	ctx context.Context,
	token string,
	maxTokenAge time.Duration,
	authorizationCheck *sessions.AuthorizationCheck,
) (*sessions.MemberSession, error) {
	if c.JWKS == nil {
		return nil, stytcherror.ErrJWKSNotInitialized
	}

	aud := c.C.GetConfig().ProjectID
	iss := fmt.Sprintf("stytch.com/%s", c.C.GetConfig().ProjectID)

	// It's difficult to extract all sets of claims (standard/registered, Stytch, custom) all at
	// once. So we parse the token twice.
	//
	// The first parse is for validating and extracting the statically-known claims. It will fail
	// if the token is invalid for any reason.
	//
	// The second parse is for extracting the custom claims.
	var staticClaims sessions.Claims
	_, err := jwt.ParseWithClaims(token, &staticClaims, c.JWKS.Keyfunc, jwt.WithAudience(aud), jwt.WithIssuer(iss))
	if err != nil {
		return nil, fmt.Errorf("failed to parse JWT: %w", err)
	}

	if staticClaims.RegisteredClaims.IssuedAt.Add(maxTokenAge).Before(time.Now()) {
		// The JWT is valid, but older than the tolerable maximum age.
		return nil, sessions.ErrJWTTooOld
	}

	// The token has already been verified at this point, so its claims and signature are all
	// valid. This call to ParseUnverified is _only_ for extracting the remaining custom claims.
	//
	// Using ParseWithClaims again would cause this to re-validate the token's timestamps and
	// signature, which fail if it was very close to its expiration on the previous parse.
	var customClaims jwt.MapClaims
	_, _, err = jwt.NewParser().ParseUnverified(token, &customClaims)
	if err != nil {
		return nil, fmt.Errorf("failed to parse JWT: %w", err)
	}

	// Remove all the reserved claims that are already present in staticClaims.
	for key := range customClaims {
		if shared.ReservedClaim(key) {
			delete(customClaims, key)
		}
	}

	memberSession, err := marshalJWTIntoSession(staticClaims, customClaims)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal JWT into session: %w", err)
	}

	var policy *rbac.Policy
	if authorizationCheck != nil {
		policy, err = c.PolicyCache.Get(ctx)
		if err != nil {
			return nil, fmt.Errorf("failed to get cached policy: %w", err)
		}

		err = shared.PerformAuthorizationCheck(policy, staticClaims.Session.Roles, memberSession.OrganizationID, authorizationCheck)
		if err != nil {
			return nil, err
		}
	}

	return memberSession, nil
}

func marshalJWTIntoSession(claims sessions.Claims, customClaims map[string]any) (*sessions.MemberSession, error) {
	// For JWTs that include it, prefer the inner expires_at claim.
	expiresAt := claims.Session.ExpiresAt
	if expiresAt == "" {
		expiresAt = claims.RegisteredClaims.ExpiresAt.Time.Format(time.RFC3339)
	}

	started, err := time.Parse(time.RFC3339, claims.Session.StartedAt)
	if err != nil {
		return nil, err
	}
	started = started.UTC()

	accessed, err := time.Parse(time.RFC3339, claims.Session.LastAccessedAt)
	if err != nil {
		return nil, err
	}
	accessed = accessed.UTC()

	expires, err := time.Parse(time.RFC3339, expiresAt)
	if err != nil {
		return nil, err
	}
	expires = expires.UTC()

	return &sessions.MemberSession{
		MemberSessionID:       claims.Session.ID,
		MemberID:              claims.RegisteredClaims.Subject,
		StartedAt:             &started,
		LastAccessedAt:        &accessed,
		ExpiresAt:             &expires,
		AuthenticationFactors: claims.Session.AuthenticationFactors,
		OrganizationID:        claims.Organization.ID,
		CustomClaims:          customClaims,
	}, nil
}

// ENDMANUAL(AuthenticateJWT)
