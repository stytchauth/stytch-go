package b2c

import (
	"errors"
	"strings"

	"github.com/stytchauth/stytch-go/v8/stytch/shared"

	"github.com/golang-jwt/jwt/v5"
)

type Key = shared.Key

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
	RequestID string `json:"request_id,omitempty"`
	Keys      []Key  `json:"keys,omitempty"`
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

// IsValid returns an error if there is an issuer or audience mismatch in the claims.
//
// Deprecated: JWT claims are validated when the token is parsed. There is no need to call this method.
func (c Claims) IsValid(projectID string) error {
	var errs []error

	if !c.verifyIssuer(projectID) {
		errs = append(errs, jwt.ErrTokenInvalidIssuer)
	}

	if !c.verifyAudience(projectID) {
		errs = append(errs, jwt.ErrTokenInvalidAudience)
	}

	if len(errs) == 0 {
		return nil
	}
	return multiError{errs}
}

type multiError struct {
	errs []error
}

func (me multiError) Error() string {
	var msgs []string
	for _, err := range me.errs {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, ", ")
}

func (me multiError) Is(target error) bool {
	for _, err := range me.errs {
		if errors.Is(err, target) {
			return true
		}
	}
	return false
}

func (c *Claims) verifyIssuer(cmp string) bool {
	issuerSplit := strings.Split(c.RegisteredClaims.Issuer, "/")
	return len(issuerSplit) == 2 && issuerSplit[1] == cmp
}

func (c *Claims) verifyAudience(cmp string) bool {
	return len(c.RegisteredClaims.Audience) == 1 && c.RegisteredClaims.Audience[0] == cmp
}
