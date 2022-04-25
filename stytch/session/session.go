package session

import (
	"encoding/json"
	"errors"
	"fmt"
	"time"

	"github.com/MicahParks/keyfunc"
	"github.com/golang-jwt/jwt/v4"
	"github.com/stytchauth/stytch-go/v5/stytch"
	"github.com/stytchauth/stytch-go/v5/stytch/stytcherror"
)

var ErrJWTTooOld = errors.New("JWT too old")

type Client struct {
	C    *stytch.Client
	JWKS *keyfunc.JWKS
}

func (c *Client) Get(
	body *stytch.SessionsGetParams,
) (*stytch.SessionsGetResponse, error) {
	queryParams := make(map[string]string)
	if body != nil {
		queryParams["user_id"] = body.UserID
	}

	var retVal stytch.SessionsGetResponse
	err := c.C.NewRequest("GET", "/sessions", queryParams, nil, &retVal)
	return &retVal, err
}

func (c *Client) GetJWKS(
	body *stytch.SessionsGetJWKSParams,
) (*stytch.SessionsGetJWKSResponse, error) {
	path := "/sessions/jwks/" + body.ProjectID

	var retVal stytch.SessionsGetJWKSResponse
	err := c.C.NewRequest("GET", path, nil, nil, &retVal)
	return &retVal, err
}

func (c *Client) AuthenticateJWT(
	maxTokenAge time.Duration,
	body *stytch.SessionsAuthenticateParams,
) (*stytch.SessionsAuthenticateResponse, error) {
	if body.SessionJWT == "" || maxTokenAge == time.Duration(0) {
		return c.Authenticate(body)
	}

	session, err := c.AuthenticateJWTLocal(body.SessionJWT, maxTokenAge)
	if errors.Is(err, ErrJWTTooOld) {
		// If JWT is valid and the token is more than maxTokenAge old,
		// check with the API to make sure that the session hasn't been revoked
		return c.Authenticate(body)
	}
	if err != nil {
		return nil, err
	}

	return &stytch.SessionsAuthenticateResponse{
		Session: *session,
	}, nil
}

func (c *Client) AuthenticateJWTLocal(
	token string,
	maxTokenAge time.Duration,
) (*stytch.Session, error) {
	var claims stytch.Claims
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

func marshalJWTIntoSession(claims stytch.Claims) stytch.Session {
	var authFactorPtrs []*stytch.AuthenticationFactor
	for _, factor := range claims.StytchSession.AuthenticationFactors {
		factor := factor
		authFactorPtrs = append(authFactorPtrs, &factor)
	}
	return stytch.Session{
		SessionID:             claims.RegisteredClaims.ID,
		UserID:                claims.RegisteredClaims.Subject,
		StartedAt:             claims.StytchSession.StartedAt,
		LastAccessedAt:        claims.StytchSession.LastAccessedAt,
		ExpiresAt:             claims.RegisteredClaims.ExpiresAt.String(),
		Attributes:            claims.StytchSession.Attributes,
		AuthenticationFactors: authFactorPtrs,
	}
}

func (c *Client) Authenticate(
	body *stytch.SessionsAuthenticateParams,
) (*stytch.SessionsAuthenticateResponse, error) {
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

	var retVal stytch.SessionsAuthenticateResponse
	err = c.C.NewRequest("POST", path, nil, jsonBody, &retVal)
	return &retVal, err
}

func (c *Client) Revoke(
	body *stytch.SessionsRevokeParams,
) (*stytch.SessionsRevokeResponse, error) {
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

	var retVal stytch.SessionsRevokeResponse
	err = c.C.NewRequest("POST", path, nil, jsonBody, &retVal)
	return &retVal, err
}
