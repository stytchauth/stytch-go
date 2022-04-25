package session_test

import (
	"crypto/rand"
	"crypto/rsa"
	"testing"
	"time"

	"github.com/MicahParks/keyfunc"
	"github.com/golang-jwt/jwt/v4"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/stytchauth/stytch-go/v5/stytch"
	"github.com/stytchauth/stytch-go/v5/stytch/config"
	"github.com/stytchauth/stytch-go/v5/stytch/session"
)

func TestAuthenticateJWTLocal(t *testing.T) {
	client := &stytch.Client{
		Config: &config.Config{
			Env:       config.EnvTest,
			BaseURI:   "https://example.test/v1/",
			ProjectID: "project-test-00000000-0000-0000-0000-000000000000",
			Secret:    "secret-test-11111111-1111-1111-1111-111111111111",
		},
		// In these tests, the keyset has already been downloaded, so no other network requests
		// should be made.
		HTTPClient: nil,
	}

	key := rsaKey(t)
	keyID := "jwk-test-22222222-2222-2222-2222-222222222222"
	jwks := keyfunc.NewGiven(map[string]keyfunc.GivenKey{
		keyID: keyfunc.NewGivenRSA(&key.PublicKey),
	})

	sessions := &session.Client{
		C:    client,
		JWKS: jwks,
	}

	t.Run("expired JWT", func(t *testing.T) {
		iat := time.Now().Add(-time.Hour).Truncate(time.Second)
		exp := iat.Add(time.Minute)

		claims := sandboxClaims(t, iat, exp)
		token := signJWT(t, keyID, key, claims)

		s, err := sessions.AuthenticateJWTLocal(token, 5*time.Minute)
		assert.ErrorIs(t, err, jwt.ErrTokenExpired)
		assert.Nil(t, s)
	})

	t.Run("stale JWT", func(t *testing.T) {
		iat := time.Now().Add(-10 * time.Minute).Truncate(time.Second)
		exp := iat.Add(time.Hour)

		claims := sandboxClaims(t, iat, exp)
		token := signJWT(t, keyID, key, claims)

		s, err := sessions.AuthenticateJWTLocal(token, 5*time.Minute)
		assert.ErrorIs(t, err, session.ErrJWTTooOld)
		assert.Nil(t, s)
	})

	t.Run("valid JWT", func(t *testing.T) {
		iat := time.Now().Truncate(time.Second)
		exp := iat.Add(time.Hour)

		claims := sandboxClaims(t, iat, exp)
		token := signJWT(t, keyID, key, claims)

		session, err := sessions.AuthenticateJWTLocal(token, 5*time.Minute)
		require.NoError(t, err)

		expected := &stytch.Session{
			SessionID:      "session-live-e26a0ccb-0dc0-4edb-a4bb-e70210f43555",
			UserID:         "user-live-fde03dd1-fff7-4b3c-9b31-ead3fbc224de",
			StartedAt:      iat.Format(time.RFC3339),
			LastAccessedAt: iat.Format(time.RFC3339),
			ExpiresAt:      exp.String(),
			Attributes: stytch.Attributes{
				IPAddress: "",
				UserAgent: "",
			},
			AuthenticationFactors: []*stytch.AuthenticationFactor{
				{
					Type:                "magic_link",
					DeliveryMethod:      "email",
					LastAuthenticatedAt: iat.Format(time.RFC3339),
					EmailFactor: stytch.EmailFactor{
						EmailAddress: "sandbox@stytch.com",
						EmailID:      "email-live-cca9d7d0-11b6-4167-9385-d7e0c9a77418",
					},
				},
			},
		}
		assert.Equal(t, expected, session)
	})
}

func rsaKey(t *testing.T) *rsa.PrivateKey {
	// This short key length is fine for test data. We won't actually use the keys for anything.
	//
	// #nosec G403
	key, err := rsa.GenerateKey(rand.Reader, 1024)
	if err != nil {
		t.Fatalf("generate test RSA key: %s", err)
	}
	return key
}

func signJWT(t *testing.T, keyID string, key *rsa.PrivateKey, claims jwt.Claims) string {
	token := jwt.NewWithClaims(jwt.SigningMethodRS256, claims)
	token.Header["kid"] = keyID

	signed, err := token.SignedString(key)
	if err != nil {
		t.Fatalf("sign JWT: %s", err)
	}
	return signed
}

func sandboxClaims(t *testing.T, iat, exp time.Time) stytch.Claims {
	return stytch.Claims{
		StytchSession: stytch.SessionClaim{
			ID:             "session-live-e26a0ccb-0dc0-4edb-a4bb-e70210f43555",
			StartedAt:      iat.Format(time.RFC3339),
			LastAccessedAt: iat.Format(time.RFC3339),
			Attributes:     stytch.Attributes{},
			AuthenticationFactors: []stytch.AuthenticationFactor{
				{
					Type:                "magic_link",
					DeliveryMethod:      "email",
					LastAuthenticatedAt: iat.Format(time.RFC3339),
					EmailFactor: stytch.EmailFactor{
						EmailAddress: "sandbox@stytch.com",
						EmailID:      "email-live-cca9d7d0-11b6-4167-9385-d7e0c9a77418",
					},
				},
			},
		},
		RegisteredClaims: jwt.RegisteredClaims{
			ID:        "session-live-e26a0ccb-0dc0-4edb-a4bb-e70210f43555",
			Issuer:    "stytch.com/project-test-00000000-0000-0000-0000-000000000000",
			Audience:  []string{"project-test-00000000-0000-0000-0000-000000000000"},
			Subject:   "user-live-fde03dd1-fff7-4b3c-9b31-ead3fbc224de",
			IssuedAt:  jwt.NewNumericDate(iat),
			NotBefore: jwt.NewNumericDate(iat),
			ExpiresAt: jwt.NewNumericDate(exp),
		},
	}
}
