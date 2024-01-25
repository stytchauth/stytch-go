package shared

import "strings"

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
