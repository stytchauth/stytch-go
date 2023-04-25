package session

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"time"

	"github.com/stytchauth/stytch-go/v7/stytch/b2c"

	"github.com/MicahParks/keyfunc"
	"github.com/golang-jwt/jwt/v4"
	"github.com/stytchauth/stytch-go/v7/stytch"
	"github.com/stytchauth/stytch-go/v7/stytch/stytcherror"
)

var ErrJWTTooOld = errors.New("JWT too old")

type Client struct {
	C    *stytch.Client
	JWKS *keyfunc.JWKS
}

func (c *Client) Get(
	ctx context.Context,
	body *b2c.SessionsGetParams,
) (*b2c.SessionsGetResponse, error) {
	queryParams := make(map[string]string)
	if body != nil {
		queryParams["user_id"] = body.UserID
	}

	var retVal b2c.SessionsGetResponse
	err := c.C.NewRequest(ctx, "GET", "/sessions", queryParams, nil, &retVal)
	return &retVal, err
}

func (c *Client) GetJWKS(
	ctx context.Context, body *b2c.SessionsGetJWKSParams,
) (*b2c.SessionsGetJWKSResponse, error) {
	path := "/sessions/jwks/" + body.ProjectID

	var retVal b2c.SessionsGetJWKSResponse
	err := c.C.NewRequest(ctx, "GET", path, nil, nil, &retVal)
	return &retVal, err
}

func (c *Client) AuthenticateJWT(
	ctx context.Context,
	maxTokenAge time.Duration,
	body *b2c.SessionsAuthenticateParams,
) (*b2c.SessionsAuthenticateResponse, error) {
	if body.SessionJWT == "" || maxTokenAge == time.Duration(0) {
		return c.Authenticate(ctx, body)
	}

	session, err := c.AuthenticateJWTLocal(body.SessionJWT, maxTokenAge)
	if err != nil {
		// JWT couldn't be verified locally. Check with the Stytch API.
		return c.Authenticate(ctx, body)
	}

	return &b2c.SessionsAuthenticateResponse{
		Session: *session,
	}, nil
}

func (c *Client) AuthenticateJWTWithClaims(
	ctx context.Context,
	maxTokenAge time.Duration,
	body *b2c.SessionsAuthenticateParams,
	claims interface{},
) (*b2c.SessionsAuthenticateResponse, error) {
	if body.SessionJWT == "" || maxTokenAge == time.Duration(0) {
		return c.AuthenticateWithClaims(ctx, body, claims)
	}

	session, err := c.AuthenticateJWTLocal(body.SessionJWT, maxTokenAge)
	if err != nil {
		// JWT couldn't be verified locally. Check with the Stytch API.
		return c.Authenticate(ctx, body)
	}

	return &b2c.SessionsAuthenticateResponse{
		Session: *session,
	}, nil
}

func (c *Client) AuthenticateJWTLocal(
	token string,
	maxTokenAge time.Duration,
) (*b2c.Session, error) {
	var claims b2c.Claims
	_, err := jwt.ParseWithClaims(token, &claims, c.JWKS.Keyfunc)
	if err != nil {
		return nil, fmt.Errorf("failed to parse JWT: %w", err)
	}

	if err := claims.IsValid(c.C.Config.BasicAuthProjectID()); err != nil {
		// If JWT is invalid, return error
		return nil, fmt.Errorf("authenticate JWT: %w", err)
	}
	if claims.RegisteredClaims.IssuedAt.Add(maxTokenAge).Before(time.Now()) {
		// The JWT is valid, but older than the tolerable maximum age.
		return nil, ErrJWTTooOld
	}

	session := marshalJWTIntoSession(claims)
	return &session, nil
}

func marshalJWTIntoSession(claims b2c.Claims) b2c.Session {
	var authFactorPtrs []*b2c.AuthenticationFactor
	for _, factor := range claims.StytchSession.AuthenticationFactors {
		factor := factor
		authFactorPtrs = append(authFactorPtrs, &factor)
	}

	// For JWTs that include it, prefer the inner expires_at claim.
	expiresAt := claims.StytchSession.ExpiresAt
	if expiresAt == "" {
		expiresAt = claims.RegisteredClaims.ExpiresAt.Time.Format(time.RFC3339)
	}

	return b2c.Session{
		SessionID:             claims.StytchSession.ID,
		UserID:                claims.RegisteredClaims.Subject,
		StartedAt:             claims.StytchSession.StartedAt,
		LastAccessedAt:        claims.StytchSession.LastAccessedAt,
		ExpiresAt:             expiresAt,
		Attributes:            claims.StytchSession.Attributes,
		AuthenticationFactors: authFactorPtrs,
	}
}

func (c *Client) Authenticate(
	ctx context.Context,
	body *b2c.SessionsAuthenticateParams,
) (*b2c.SessionsAuthenticateResponse, error) {
	path := "/sessions/authenticate"

	var jsonBody []byte
	var err error
	if body != nil {
		jsonBody, err = json.Marshal(body)
		if err != nil {
			return nil, stytcherror.NewClientLibraryError("Oops, something seems to have gone wrong " +
				"marshalling the /sessions/authenticate request body")
		}
	}

	var retVal b2c.SessionsAuthenticateResponse
	err = c.C.NewRequest(ctx, "POST", path, nil, jsonBody, &retVal)
	return &retVal, err
}

// AuthenticateWithClaims fills in the claims pointer with custom claims from the response.
// Pass in a map with the types of values you're expecting so that this function can marshal
// the claims from the response. See ExampleClient_AuthenticateWithClaims_map,
// ExampleClient_AuthenticateWithClaims_struct for examples
func (c *Client) AuthenticateWithClaims(
	ctx context.Context,
	body *b2c.SessionsAuthenticateParams,
	claims interface{},
) (*b2c.SessionsAuthenticateResponse, error) {
	path := "/sessions/authenticate"

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
	var retVal b2c.SessionsAuthenticateResponse
	if err := json.Unmarshal(b, &retVal); err != nil {
		return nil, fmt.Errorf("unmarshal SessionsAuthenticateResponse: %w", err)
	}

	// Then extract the custom claims. Build a claims wrapper using the caller's `claims` value so
	// the unmarshal fills it.
	wrapper := b2c.SessionWrapper{
		Session: b2c.ClaimsWrapper{
			Claims: claims,
		},
	}
	if err := json.Unmarshal(b, &wrapper); err != nil {
		return nil, fmt.Errorf("unmarshal custom claims: %w", err)
	}
	retVal.Session.CustomClaims = wrapper.Session.Claims
	return &retVal, err
}

func (c *Client) Revoke(
	ctx context.Context,
	body *b2c.SessionsRevokeParams,
) (*b2c.SessionsRevokeResponse, error) {
	path := "/sessions/revoke"

	var jsonBody []byte
	var err error
	if body != nil {
		jsonBody, err = json.Marshal(body)
		if err != nil {
			return nil, stytcherror.NewClientLibraryError("Oops, something seems to have gone wrong " +
				"marshalling the /sessions/revoke request body")
		}
	}

	var retVal b2c.SessionsRevokeResponse
	err = c.C.NewRequest(ctx, "POST", path, nil, jsonBody, &retVal)
	return &retVal, err
}
