package shared_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stytchauth/stytch-go/v17/stytch/shared"
)

func TestReservedClaim(t *testing.T) {
	testcases := map[string]bool{
		// Standard claims
		"iss": true,
		"aud": true,
		"sub": true,
		"iat": true,
		"nbf": true,
		"exp": true,

		// Stytch claims
		"https://stytch.com/session":                true,
		"https://stytch.com/organization":           true,
		"https://stytch.com/some/hypothetical/path": true,

		// Custom claims
		"any":                         false,
		"https://example.com/session": false,
	}

	for key, expected := range testcases {
		t.Run(key, func(t *testing.T) {
			assert.Equal(t, expected, shared.ReservedClaim(key))
		})
	}
}
