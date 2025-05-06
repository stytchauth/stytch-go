package idp

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/stytchauth/stytch-go/v16/stytch/b2b/sessions"
)

type OrganizationClaim struct {
	OrganizationID string `json:"organization_id"`
	Slug           string `json:"slug"`
}

type IntrospectTokenClaims struct {
	Subject      string            `json:"sub"`
	Scope        string            `json:"scope"`
	CustomClaims map[string]any    `json:"custom_claims"`
	ExpiresAt    int32             `json:"exp"`
	IssuedAt     int32             `json:"iat"`
	NotBefore    int32             `json:"nbf"`
	TokenType    string            `json:"token_type"`
	Organization OrganizationClaim `json:"https://stytch.com/organization"`
	jwt.RegisteredClaims
}

type IntrospectTokenLocalParams struct {
	Token              string
	MaxTokenAge        time.Duration
	AuthorizationCheck *sessions.AuthorizationCheck
}
