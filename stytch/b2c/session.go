package b2c

import (
	"strings"


	"github.com/stytchauth/stytch-go/v8/stytch/shared"

	"github.com/golang-jwt/jwt/v4"
)

type SessionsGetParams struct {
	UserID string `json:"user_id"`
}

type SessionsGetResponse struct {
	RequestID string    `json:"request_id,omitempty"`
	Sessions  []Session `json:"sessions,omitempty"`
}

type SessionsGetJWKSParams struct {
	ProjectID string `json:"project_id"`
}

type SessionsGetJWKSResponse struct {
	RequestID string       `json:"request_id,omitempty"`
	Keys      []shared.Key `json:"keys,omitempty"`
}

type SessionsAuthenticateParams struct {
	SessionToken           string                 `json:"session_token,omitempty"`
	SessionDurationMinutes int32                  `json:"session_duration_minutes,omitempty"`
	SessionJWT             string                 `json:"session_jwt,omitempty"`
	SessionCustomClaims    map[string]interface{} `json:"session_custom_claims,omitempty"`
}

type SessionsAuthenticateResponse struct {
	RequestID    string  `json:"request_id,omitempty"`
	Session      Session `json:"session,omitempty"`
	SessionToken string  `json:"session_token,omitempty"`
	SessionJWT   string  `json:"session_jwt,omitempty"`
	User         User    `json:"user,omitempty"`
}

type SessionsRevokeParams struct {
	SessionID    string `json:"session_id,omitempty"`
	SessionToken string `json:"session_token,omitempty"`
	SessionJWT   string `json:"session_jwt,omitempty"`
}

type SessionsRevokeResponse struct {
	RequestID string `json:"request_id,omitempty"`
}

type SessionClaim struct {
	ID                    string                 `json:"id"`
	StartedAt             string                 `json:"started_at"`
	LastAccessedAt        string                 `json:"last_accessed_at"`
	ExpiresAt             string                 `json:"expires_at"`
	Attributes            Attributes             `json:"attributes"`
	AuthenticationFactors []AuthenticationFactor `json:"authentication_factors"`
}

type Claims struct {
	StytchSession SessionClaim `json:"https://stytch.com/session"`
	jwt.RegisteredClaims
}

// Validation options in GoJWT are currently unexported. Once they're exported, we
// can define this as a Valid() function, see
// https://github.com/golang-jwt/jwt/blob/1096e506e671d6d6fe134cc997bbd475937392c8/validator_option.go#L9-L11 //nolint:lll
func (c Claims) IsValid(projectID string) error {
	vErr := new(jwt.ValidationError)
	if !c.verifyIssuer(projectID) {
		vErr.Inner = jwt.ErrTokenInvalidIssuer
		vErr.Errors |= jwt.ValidationErrorIssuer
	}

	if !c.verifyAudience(projectID) {
		vErr.Inner = jwt.ErrTokenInvalidAudience
		vErr.Errors |= jwt.ValidationErrorAudience
	}

	if vErr.Errors == 0 {
		return nil
	}
	return vErr
}

func (c *Claims) verifyIssuer(cmp string) bool {
	issuerSplit := strings.Split(c.RegisteredClaims.Issuer, "/")
	return len(issuerSplit) == 2 && issuerSplit[1] == cmp
}

func (c *Claims) verifyAudience(cmp string) bool {
	return len(c.RegisteredClaims.Audience) == 1 && c.RegisteredClaims.Audience[0] == cmp
}
