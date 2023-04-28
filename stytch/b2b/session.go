package b2b

import "github.com/stytchauth/stytch-go/v8/stytch/b2c"

type SessionGetParams struct {
	OrganizationID string `json:"organization_id"`
	MemberID       string `json:"member_id"`
}

type SessionGetResponse struct {
	RequestID      string          `json:"request_id,omitempty"`
	StatusCode     int             `json:"status_code,omitempty"`
	MemberSessions []MemberSession `json:"member_sessions,omitempty"`
}

type SessionAuthenticateParams struct {
	SessionToken           string                 `json:"session_token,omitempty"`
	SessionDurationMinutes int32                  `json:"session_duration_minutes,omitempty"`
	SessionJWT             string                 `json:"session_jwt,omitempty"`
	SessionCustomClaims    map[string]interface{} `json:"session_custom_claims,omitempty"`
}

type SessionAuthenticateResponse struct {
	RequestID     string        `json:"request_id,omitempty"`
	StatusCode    int           `json:"status_code,omitempty"`
	MemberSession MemberSession `json:"member_session,omitempty"`
	SessionToken  string        `json:"session_token,omitempty"`
	SessionJWT    string        `json:"session_jwt,omitempty"`
	Member        Member        `json:"member,omitempty"`
	Organization  Organization  `json:"organization,omitempty"`
}

type SessionRevokeParams struct {
	MemberSessionID string `json:"member_session_id,omitempty"`
	SessionToken    string `json:"session_token,omitempty"`
	SessionJWT      string `json:"session_jwt,omitempty"`
	MemberID        string `json:"member_id,omitempty"`
}

type SessionRevokeResponse struct {
	RequestID  string `json:"request_id,omitempty"`
	StatusCode int    `json:"status_code,omitempty"`
}

type SessionExchangeParams struct {
	OrganizationID         string                 `json:"organization_id,omitempty"`
	SessionToken           string                 `json:"session_token,omitempty"`
	SessionJWT             string                 `json:"session_jwt,omitempty"`
	SessionDurationMinutes int32                  `json:"session_duration_minutes,omitempty"`
	SessionCustomClaims    map[string]interface{} `json:"session_custom_claims,omitempty"`
}

type SessionExchangeResponse struct {
	RequestID     string        `json:"request_id,omitempty"`
	StatusCode    int           `json:"status_code,omitempty"`
	MemberID      string        `json:"member_id,omitempty"`
	MemberSession MemberSession `json:"member_session,omitempty"`
	SessionToken  string        `json:"session_token,omitempty"`
	SessionJWT    string        `json:"session_jwt,omitempty"`
	Member        Member        `json:"member,omitempty"`
	Organization  Organization  `json:"organization,omitempty"`
}

type MemberSession struct {
	MemberSessionID       string                      `json:"member_session_id,omitempty"`
	MemberID              string                      `json:"member_id,omitempty"`
	StartedAt             string                      `json:"started_at,omitempty"`
	LastAccessedAt        string                      `json:"last_accessed_at,omitempty"`
	ExpiresAt             string                      `json:"expires_at,omitempty"`
	AuthenticationFactors []*b2c.AuthenticationFactor `json:"authentication_factors,omitempty"`
	CustomClaims          interface{}                 `json:"custom_claims,omitempty"`
	OrganizationID        string                      `json:"organization_id,omitempty"`
}
