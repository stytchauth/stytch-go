package shared

import (
	"fmt"
	"github.com/golang-jwt/jwt/v5"
	"strings"
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

func ValidateJWTToken(token string, staticClaims jwt.Claims, keyFunc jwt.Keyfunc, audience string, issuer string, fallbackIssuer string) error {
	// If we fail the first parse, we try again with the fallback issuer.
	// We need to do this because when a customer is using a custom base URI (usually a CNAME pointing at the Stytch API),
	// instead of returning the usual `stytch.com/<project_id>` issuer, we return the URL they used to
	// make the request. This is in line with the OIDC spec. Ideally we would've always returned the baseURI issuer, but
	// that would've broken existing customers who are dependent on the `stytch.com/<project_id>`.
	_, err := jwt.ParseWithClaims(token, staticClaims, keyFunc, jwt.WithAudience(audience), jwt.WithIssuer(issuer))
	if err != nil {
		_, err := jwt.ParseWithClaims(token, staticClaims, keyFunc, jwt.WithAudience(audience), jwt.WithIssuer(fallbackIssuer))
		if err != nil {
			return fmt.Errorf("failed to parse JWT: %w", err)
		}
	}
	return nil
}
