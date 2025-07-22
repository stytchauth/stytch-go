package idp

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type IntrospectTokenResponse struct {
	Active       bool             `json:"active"`
	TokenType    string           `json:"token_type"`
	Issuer       string           `json:"iss"`
	Subject      string           `json:"sub"`
	Audience     []string         `json:"aud"`
	Scope        string           `json:"scope"`
	ClientID     string           `json:"client_id"`
	Expiry       *jwt.NumericDate `json:"exp"`
	IssuedAt     *jwt.NumericDate `json:"iat"`
	CustomClaims map[string]any
}

type IntrospectTokenClaims struct {
	ClientID string `json:"client_id"`
	Scope    string `json:"scope"`
	JTI      string `json:"jti"`
	jwt.RegisteredClaims
}

type IntrospectTokenLocalParams struct {
	Token       string
	MaxTokenAge time.Duration
	// AuthorizationCheck *sessions.AuthorizationCheck
}

type IntrospectTokenNetworkParams struct {
	Token         string
	ClientID      string
	ClientSecret  *string
	TokenTypeHint *string
	// AuthorizationCheck *sessions.AuthorizationCheck
}
