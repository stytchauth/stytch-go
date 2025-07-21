package idp

import (
	"github.com/golang-jwt/jwt/v5"
	"time"

	"github.com/stytchauth/stytch-go/v16/stytch/b2b/sessions"
)

type OrganizationClaim struct {
	OrganizationID string `json:"organization_id"`
	Slug           string `json:"slug"`
}

type IntrospectTokenResponse struct {
	Active       bool              `json:"active"`
	TokenType    string            `json:"token_type"`
	Issuer       string            `json:"iss"`
	Subject      string            `json:"sub"`
	Audience     []string          `json:"aud"`
	Scope        string            `json:"scope"`
	ClientID     string            `json:"client_id"`
	Expiry       *jwt.NumericDate  `json:"exp"`
	IssuedAt     *jwt.NumericDate  `json:"iat"`
	Organization OrganizationClaim `json:"https://stytch.com/organization"`
	CustomClaims map[string]any
}

type IntrospectTokenClaims struct {
	ClientID     string            `json:"client_id"`
	Scope        string            `json:"scope"`
	Organization OrganizationClaim `json:"https://stytch.com/organization"`
	JTI          string            `json:"jti"`
	jwt.RegisteredClaims
}

type IntrospectTokenLocalParams struct {
	Token              string
	MaxTokenAge        time.Duration
	AuthorizationCheck *sessions.AuthorizationCheck
}

type IntrospectTokenNetworkParams struct {
	Token              string
	ClientID           string
	ClientSecret       *string
	TokenTypeHint      *string
	AuthorizationCheck *sessions.AuthorizationCheck
}
