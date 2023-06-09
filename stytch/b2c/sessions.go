package b2c

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
	"github.com/stytchauth/stytch-go/v8/stytch"
	"github.com/stytchauth/stytch-go/v8/stytch/b2c/sessions"
	"github.com/stytchauth/stytch-go/v8/stytch/stytcherror"
)

type SessionsClient struct {
	C    *stytch.Client
	JWKS *keyfunc.JWKS
}

func NewSessionsClient(c *stytch.Client) *SessionsClient {
	return &SessionsClient{
		C: c,
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

	var retVal sessions.GetResponse
	err := c.C.NewRequest(
		ctx,
		"GET",
		"/v1/sessions",
		queryParams,
		nil,
		&retVal,
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

	var retVal sessions.AuthenticateResponse
	err = c.C.NewRequest(
		ctx,
		"POST",
		"/v1/sessions/authenticate",
		nil,
		jsonBody,
		&retVal,
	)
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

	var retVal sessions.RevokeResponse
	err = c.C.NewRequest(
		ctx,
		"POST",
		"/v1/sessions/revoke",
		nil,
		jsonBody,
		&retVal,
	)
	return &retVal, err
}

// GetJWKS: Get the JSON Web Key Set (JWKS) for a Stytch Project.
func (c *SessionsClient) GetJWKS(
	ctx context.Context,
	body *sessions.JWKSParams,
) (*sessions.JWKSResponse, error) {
	var retVal sessions.JWKSResponse
	err := c.C.NewRequest(
		ctx,
		"GET",
		fmt.Sprintf("/v1/sessions/jwks/%s", body.ProjectID),
		nil,
		nil,
		&retVal,
	)
	return &retVal, err
}

// MANUAL(AuthenticateJWT)(SERVICE_METHOD)
// ADDIMPORT: "encoding/json"
// ADDIMPORT: "time"
// ADDIMPORT: "github.com/golang-jwt/jwt/v5"
// ADDIMPORT: "github.com/MicahParks/keyfunc/v2"

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
	if body.SessionJWT == "" || maxTokenAge == time.Duration(0) {
		return c.AuthenticateWithClaims(ctx, body, claims)
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

func (c *SessionsClient) AuthenticateJWTLocal(
	token string,
	maxTokenAge time.Duration,
) (*sessions.Session, error) {
	var claims sessions.Claims

	aud := c.C.Config.BasicAuthProjectID()
	iss := fmt.Sprintf("stytch.com/%s", c.C.Config.BasicAuthProjectID())

	_, err := jwt.ParseWithClaims(token, &claims, c.JWKS.Keyfunc, jwt.WithAudience(aud), jwt.WithIssuer(iss))
	if err != nil {
		return nil, fmt.Errorf("failed to parse JWT: %w", err)
	}

	if claims.RegisteredClaims.IssuedAt.Add(maxTokenAge).Before(time.Now()) {
		// The JWT is valid, but older than the tolerable maximum age.
		return nil, sessions.ErrJWTTooOld
	}

	return marshalJWTIntoSession(claims)
}

func marshalJWTIntoSession(claims sessions.Claims) (*sessions.Session, error) {
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
		Attributes:            claims.StytchSession.Attributes,
		AuthenticationFactors: claims.StytchSession.AuthenticationFactors,
	}, nil
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
	path := "/v1/sessions/authenticate"

	var jsonBody []byte
	var err error
	if body != nil {
		jsonBody, err = json.Marshal(body)
		if err != nil {
			return nil, stytcherror.NewClientLibraryError("Oops, something seems to have gone wrong " +
				"marshalling the /sessions/authenticate request body")
		}
	}

	b, err := c.C.RawRequest(ctx, "POST", path, nil, jsonBody)
	if err != nil {
		return nil, err
	}

	// First extract the Stytch data.
	var retVal sessions.AuthenticateResponse
	if err := json.Unmarshal(b, &retVal); err != nil {
		return nil, fmt.Errorf("unmarshal SessionsAuthenticateResponse: %w", err)
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

// ENDMANUAL(AuthenticateJWT)
