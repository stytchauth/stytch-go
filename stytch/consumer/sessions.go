package consumer

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
	"github.com/stytchauth/stytch-go/v12/stytch"
	"github.com/stytchauth/stytch-go/v12/stytch/consumer/sessions"
	"github.com/stytchauth/stytch-go/v12/stytch/shared"
	"github.com/stytchauth/stytch-go/v12/stytch/stytcherror"
)

type SessionsClient struct {
	C    stytch.Client
	JWKS *keyfunc.JWKS
}

func NewSessionsClient(c stytch.Client, jwks *keyfunc.JWKS) *SessionsClient {
	return &SessionsClient{
		C:    c,
		JWKS: jwks,
	}
}

// Get: List all active Sessions for a given `user_id`. All timestamps are formatted according to the RFC
// 3339 standard and are expressed in UTC, e.g. `2021-12-29T12:33:09Z`.
func (c *SessionsClient) Get(
	ctx context.Context,
	body *sessions.GetParams,
) (*sessions.GetResponse, error) {
	queryParams := make(map[string]string)
	if body != nil {
		queryParams["user_id"] = body.UserID
	}

	headers := make(map[string][]string)

	var retVal sessions.GetResponse
	err := c.C.NewRequest(
		ctx,
		"GET",
		"/v1/sessions",
		queryParams,
		nil,
		&retVal,
		headers,
	)
	return &retVal, err
}

// Authenticate a session token and retrieve associated session data. If `session_duration_minutes` is
// included, update the lifetime of the session to be that many minutes from now. All timestamps are
// formatted according to the RFC 3339 standard and are expressed in UTC, e.g. `2021-12-29T12:33:09Z`. This
// endpoint requires exactly one `session_jwt` or `session_token` as part of the request. If both are
// included you will receive a `too_many_session_arguments` error.
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
		"POST",
		"/v1/sessions/authenticate",
		nil,
		jsonBody,
		&retVal,
		headers,
	)
	return &retVal, err
}

// AuthenticateWithClaims fills in the claims pointer with custom claims from the response.
//
// The value for claims must be one of these types:
//   - A pointer to a map (*map[string]any), which will be overwritten with the custom claims.
//   - A pointer to a struct (*T), which will be populated using its "json" struct tags.
//
// See ExampleClient_AuthenticateWithClaims_map, ExampleClient_AuthenticateWithClaims_struct for
// examples.
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
		"POST",
		"/v1/sessions/authenticate",
		nil,
		jsonBody,
		headers,
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
		*m = retVal.Session.CustomClaims
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

	err = decoder.Decode(retVal.Session.CustomClaims)
	if err != nil {
		return nil, err
	}

	return &retVal, err
}

// Revoke a Session, immediately invalidating all of its session tokens. You can revoke a session in three
// ways: using its ID, or using one of its session tokens, or one of its JWTs. This endpoint requires
// exactly one of those to be included in the request. It will return an error if multiple are present.
func (c *SessionsClient) Revoke(
	ctx context.Context,
	body *sessions.RevokeParams,
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

	var retVal sessions.RevokeResponse
	err = c.C.NewRequest(
		ctx,
		"POST",
		"/v1/sessions/revoke",
		nil,
		jsonBody,
		&retVal,
		headers,
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
// If you're using one of our [backend SDKs](https://stytch.com/docs/sdks), the JWKS roll will be handled
// for you.
//
// If you're using your own JWT validation library, many have built-in support for JWKS rotation, and
// you'll just need to supply this API endpoint. If not, your application should decide which JWKS to use
// for validation by inspecting the `kid` value.
func (c *SessionsClient) GetJWKS(
	ctx context.Context,
	body *sessions.GetJWKSParams,
) (*sessions.GetJWKSResponse, error) {
	headers := make(map[string][]string)

	var retVal sessions.GetJWKSResponse
	err := c.C.NewRequest(
		ctx,
		"GET",
		fmt.Sprintf("/v1/sessions/jwks/%s", body.ProjectID),
		nil,
		nil,
		&retVal,
		headers,
	)
	return &retVal, err
}

// MANUAL(AuthenticateJWT)(SERVICE_METHOD)
// ADDIMPORT: "encoding/json"
// ADDIMPORT: "time"
// ADDIMPORT: "github.com/golang-jwt/jwt/v5"
// ADDIMPORT: "github.com/MicahParks/keyfunc/v2"
// ADDIMPORT: "github.com/stytchauth/stytch-go/v12/stytch/stytcherror"

func (c *SessionsClient) AuthenticateJWT(
	ctx context.Context,
	maxTokenAge time.Duration,
	body *sessions.AuthenticateParams,
) (*sessions.AuthenticateResponse, error) {
	if body.SessionJWT == "" || maxTokenAge == time.Duration(0) {
		return c.Authenticate(ctx, body)
	}

	session, err := c.AuthenticateJWTLocal(body.SessionJWT, maxTokenAge)
	if err != nil {
		// JWT couldn't be verified locally. Check with the Stytch API.
		return c.Authenticate(ctx, body)
	}

	return &sessions.AuthenticateResponse{
		Session: *session,
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
	// TODO(v13): Remove support for populating a pre-existing map this way.

	var resp *sessions.AuthenticateResponse

	// Some special cases can force remote authentication. Otherwise, prefer local validation.
	if body.SessionJWT == "" || maxTokenAge == time.Duration(0) {
		var err error
		resp, err = c.AuthenticateWithClaims(ctx, body, nil)
		if err != nil {
			return nil, err
		}
	} else if session, err := c.AuthenticateJWTLocal(body.SessionJWT, maxTokenAge); err == nil {
		resp = &sessions.AuthenticateResponse{
			Session: *session,
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
		for key, val := range resp.Session.CustomClaims {
			claims[key] = val
		}
	}

	return resp, nil
}

func (c *SessionsClient) AuthenticateJWTLocal(
	token string,
	maxTokenAge time.Duration,
) (*sessions.Session, error) {
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

	return marshalJWTIntoSession(staticClaims, customClaims)
}

func marshalJWTIntoSession(claims sessions.Claims, customClaims map[string]any) (*sessions.Session, error) {
	// For JWTs that include it, prefer the inner expires_at claim.
	expiresAt := claims.StytchSession.ExpiresAt
	if expiresAt == "" {
		expiresAt = claims.RegisteredClaims.ExpiresAt.Time.Format(time.RFC3339)
	}

	started, err := time.Parse(time.RFC3339, claims.StytchSession.StartedAt)
	if err != nil {
		return nil, err
	}
	started = started.UTC()

	accessed, err := time.Parse(time.RFC3339, claims.StytchSession.LastAccessedAt)
	if err != nil {
		return nil, err
	}
	accessed = accessed.UTC()

	expires, err := time.Parse(time.RFC3339, expiresAt)
	if err != nil {
		return nil, err
	}
	expires = expires.UTC()

	return &sessions.Session{
		SessionID:             claims.StytchSession.ID,
		UserID:                claims.RegisteredClaims.Subject,
		StartedAt:             &started,
		LastAccessedAt:        &accessed,
		ExpiresAt:             &expires,
		Attributes:            &claims.StytchSession.Attributes,
		AuthenticationFactors: claims.StytchSession.AuthenticationFactors,
		CustomClaims:          customClaims,
	}, nil
}

// ENDMANUAL(AuthenticateJWT)
