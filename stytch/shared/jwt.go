package shared

import (
	"fmt"
	"strings"

	"github.com/golang-jwt/jwt/v5"
)

// ReservedClaim returns true if the key is reserved by the JWT standard or the Stytch platform.
func ReservedClaim(key string) bool {
	// Standard claims
	switch key {
	case
		"iss",
		"aud",
		"sub",
		"iat",
		"nbf",
		"exp":
		return true
	}

	// Stytch-specific claims are scoped by a URL prefix.
	return strings.HasPrefix(key, "https://stytch.com/")
}

type ValidateJWTTokenParams struct {
	Token          string
	StaticClaims   jwt.Claims
	KeyFunc        jwt.Keyfunc
	Audience       string
	Issuer         string
	FallbackIssuer string
}

func ValidateJWTToken(params ValidateJWTTokenParams) error {
	// If we fail the first parse, we try again with the fallback issuer.
	// We need to do this because when a customer is using a custom base URI (usually a CNAME pointing at the Stytch API),
	// instead of returning the usual `stytch.com/<project_id>` issuer, we return the URL they used to
	// make the request. This is in line with the OIDC spec. Ideally we would've always returned the baseURI issuer, but
	// that would've broken existing customers who are dependent on the `stytch.com/<project_id>`.
	_, err := jwt.ParseWithClaims(params.Token, params.StaticClaims, params.KeyFunc, jwt.WithAudience(params.Audience), jwt.WithIssuer(params.Issuer))
	if err != nil {
		_, err := jwt.ParseWithClaims(params.Token, params.StaticClaims, params.KeyFunc, jwt.WithAudience(params.Audience), jwt.WithIssuer(params.FallbackIssuer))
		if err != nil {
			return fmt.Errorf("failed to parse JWT: %w", err)
		}
	}
	return nil
}
