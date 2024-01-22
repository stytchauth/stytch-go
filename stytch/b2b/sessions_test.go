package b2b_test

import (
	"context"
	"crypto/rand"
	"crypto/rsa"
	"fmt"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"time"

	"github.com/stytchauth/stytch-go/v12/stytch/b2b"
	"github.com/stytchauth/stytch-go/v12/stytch/b2b/b2bstytchapi"
	"github.com/stytchauth/stytch-go/v12/stytch/b2b/sessions"
	consumersessions "github.com/stytchauth/stytch-go/v12/stytch/consumer/sessions"

	"github.com/MicahParks/keyfunc/v2"
	"github.com/golang-jwt/jwt/v5"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/stytchauth/stytch-go/v12/stytch"
	"github.com/stytchauth/stytch-go/v12/stytch/config"
)

func TestAuthenticateJWTLocal(t *testing.T) {
	client := &stytch.DefaultClient{
		Config: &config.Config{
			Env:       config.EnvTest,
			BaseURI:   "https://example.test/v1/b2b/",
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
		keyID: keyfunc.NewGivenRSA(&key.PublicKey, keyfunc.GivenKeyOptions{Algorithm: "RS256"}),
	})

	policyCache := b2b.NewPolicyCache(b2b.NewRBACClient(client))
	sessionClient := b2b.NewSessionsClient(client, jwks, policyCache)

	t.Run("expired JWT", func(t *testing.T) {
		iat := time.Now().UTC().Add(-time.Hour).Truncate(time.Second)
		exp := iat.Add(time.Minute)

		claims := sandboxClaims(t, iat, exp)
		token := signJWT(t, keyID, key, claims)

		ctx := context.Background()
		s, err := sessionClient.AuthenticateJWTLocal(ctx, token, 10*time.Minute, nil)
		assert.ErrorIs(t, err, jwt.ErrTokenExpired)
		assert.Nil(t, s)
	})

	t.Run("stale JWT", func(t *testing.T) {
		iat := time.Now().UTC().Add(-3 * time.Minute).Truncate(time.Second)
		exp := iat.Add(time.Hour)

		claims := sandboxClaims(t, iat, exp)
		token := signJWT(t, keyID, key, claims)

		ctx := context.Background()
		s, err := sessionClient.AuthenticateJWTLocal(ctx, token, 1*time.Minute, nil)
		assert.ErrorIs(t, err, sessions.ErrJWTTooOld)
		assert.Nil(t, s)
	})

	t.Run("incorrect audience", func(t *testing.T) {
		iat := time.Now().UTC().Truncate(time.Second)
		exp := iat.Add(time.Hour)

		claims := sandboxClaims(t, iat, exp)
		claims.Audience = jwt.ClaimStrings{"not this project"}

		token := signJWT(t, keyID, key, claims)

		ctx := context.Background()
		s, err := sessionClient.AuthenticateJWTLocal(ctx, token, 1*time.Minute, nil)
		assert.ErrorIs(t, err, jwt.ErrTokenInvalidAudience)
		assert.Nil(t, s)
	})

	t.Run("incorrect issuer", func(t *testing.T) {
		iat := time.Now().UTC().Truncate(time.Second)
		exp := iat.Add(time.Hour)

		claims := sandboxClaims(t, iat, exp)
		claims.Issuer = "not this project"

		token := signJWT(t, keyID, key, claims)

		ctx := context.Background()
		s, err := sessionClient.AuthenticateJWTLocal(ctx, token, 1*time.Minute, nil)
		assert.ErrorIs(t, err, jwt.ErrTokenInvalidIssuer)
		assert.Nil(t, s)
	})

	t.Run("valid JWT", func(t *testing.T) {
		iat := time.Now().UTC().Truncate(time.Second)
		exp := iat.Add(time.Hour)

		claims := sandboxClaims(t, iat, exp)
		token := signJWT(t, keyID, key, claims)

		ctx := context.Background()
		session, err := sessionClient.AuthenticateJWTLocal(ctx, token, 3*time.Minute, nil)
		require.NoError(t, err)

		expected := &sessions.MemberSession{
			MemberSessionID: "session-live-e26a0ccb-0dc0-4edb-a4bb-e70210f43555",
			MemberID:        "member-live-fde03dd1-fff7-4b3c-9b31-ead3fbc224de",
			StartedAt:       &iat,
			LastAccessedAt:  &iat,
			ExpiresAt:       &exp,
			AuthenticationFactors: []consumersessions.AuthenticationFactor{
				{
					Type:                "magic_link",
					DeliveryMethod:      "email",
					LastAuthenticatedAt: &iat,
					EmailFactor: &consumersessions.EmailFactor{
						EmailAddress: "sandbox@stytch.com",
						EmailID:      "email-live-cca9d7d0-11b6-4167-9385-d7e0c9a77418",
					},
				},
			},
		}
		assert.Equal(t, expected, session)
	})
}

func TestAuthenticateWithClaims(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Handle the async JWKS fetch.
		if strings.HasPrefix(r.URL.Path, "/v1/b2b/sessions/jwks/") {
			_, _ = w.Write([]byte(`{"keys": []}`))
			return
		}

		// This is the test request
		if r.URL.Path == "/v1/b2b/sessions/authenticate" {
			// There are  many other fields in this response, but these are the only ones we need
			// for this test.
			_, _ = w.Write([]byte(`{
			  "member_session": {
			    "expires_at": "2022-06-29T19:53:48Z",
			    "last_accessed_at": "2022-06-29T17:54:13Z",
			    "member_session_id": "session-test-aaaaaaaa-aaaa-aaaa-aaaa-aaaaaaaaaaaa",
			    "started_at": "2022-06-29T17:53:48Z",
			    "member_id": "user-test-00000000-0000-0000-0000-000000000000",

			    "custom_claims": {
			      "https://my-app.example.net/custom-claim": {
			        "number": 1,
			        "array": [1, "foo", null],
			        "nested": {
			          "data": "here"
			        }
			      }
			    }
			  },
				"member": {},
				"organization": {}
			}`))
			return
		}

		http.Error(w, "Bad Request", http.StatusBadRequest)
	}))

	client, err := b2bstytchapi.NewClient(
		"project-test-00000000-0000-0000-0000-000000000000",
		"secret-test-11111111-1111-1111-1111-111111111111",
		b2bstytchapi.WithBaseURI(srv.URL),
	)
	require.NoError(t, err)

	req := &sessions.AuthenticateParams{
		SessionToken: "fake session token",
	}

	t.Run("marshaling claims into a map", func(t *testing.T) {
		var claims map[string]any
		_, err := client.Sessions.AuthenticateWithClaims(context.Background(), req, &claims)
		require.NoError(t, err)

		type object = map[string]any
		expected := object{
			"https://my-app.example.net/custom-claim": object{
				// Remember that numbers without specified types unmarshal as float64.
				"number": float64(1),
				"array":  []interface{}{float64(1), "foo", nil},
				"nested": object{
					"data": "here",
				},
			},
		}
		assert.Equal(t, expected, claims)
	})
	t.Run("marshaling claims into a struct", func(t *testing.T) {
		type MyAppClaims struct {
			Number int
			Array  []interface{}
			Nested struct {
				Data string
			}
		}

		type Claims struct {
			MyApp MyAppClaims `json:"https://my-app.example.net/custom-claim"`
		}

		{
			var claims Claims
			_, err = client.Sessions.AuthenticateWithClaims(context.Background(), req, &claims)
			require.NoError(t, err)
			expected := Claims{
				MyApp: MyAppClaims{
					Number: 1,
					// Remember that numbers without specified types unmarshal as float64.
					Array:  []interface{}{float64(1), "foo", nil},
					Nested: struct{ Data string }{Data: "here"},
				},
			}
			assert.Equal(t, expected, claims)
		}
	})
}

func ExampleSessionsClient_AuthenticateWithClaims_map() {
	// If we know that our claims will follow this exact map structure, we can marshal the
	// custom claims from the response into it
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Handle the async JWKS fetch.
		if strings.HasPrefix(r.URL.Path, "/v1/b2b/sessions/jwks/") {
			_, _ = w.Write([]byte(`{"keys": []}`))
			return
		}

		// This is the test request
		if r.URL.Path == "/v1/b2b/sessions/authenticate" {
			// There are  many other fields in this response, but these are the only ones we need
			// for this test.
			_, _ = w.Write([]byte(`{
			  "member_session": {
			    "expires_at": "2022-06-29T19:53:48Z",
			    "last_accessed_at": "2022-06-29T17:54:13Z",
			    "member_session_id": "session-test-aaaaaaaa-aaaa-aaaa-aaaa-aaaaaaaaaaaa",
			    "started_at": "2022-06-29T17:53:48Z",
			    "member_id": "user-test-00000000-0000-0000-0000-000000000000",

			    "custom_claims": {
			      "https://my-app.example.net/custom-claim": {
			        "claim1": 1,
			        "claim2": 2,
			        "claim3": 3
			      }
			    }
			  },
				"member": {},
				"organization": {}
			}`))
			return
		}

		http.Error(w, "Bad Request", http.StatusBadRequest)
	}))

	client, _ := b2bstytchapi.NewClient(
		"project-test-00000000-0000-0000-0000-000000000000",
		"secret-test-11111111-1111-1111-1111-111111111111",
		b2bstytchapi.WithBaseURI(srv.URL),
	)

	// Expecting a map where all the values are maps from strings to integers
	var mapClaims map[string]map[string]int32
	_, _ = client.Sessions.AuthenticateWithClaims(
		context.Background(),
		&sessions.AuthenticateParams{
			SessionToken: "fake session token",
		},
		&mapClaims,
	)

	fmt.Println(mapClaims)
	// Output: map[https://my-app.example.net/custom-claim:map[claim1:1 claim2:2 claim3:3]]
}

func ExampleSessionsClient_AuthenticateWithClaims_struct() {
	// When we define a struct that follows the shape of our claims, we can marshal the
	// custom claims from the response into it
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Handle the async JWKS fetch.
		if strings.HasPrefix(r.URL.Path, "/v1/b2b/sessions/jwks/") {
			_, _ = w.Write([]byte(`{"keys": []}`))
			return
		}

		// This is the test request
		if r.URL.Path == "/v1/b2b/sessions/authenticate" {
			// There are  many other fields in this response, but these are the only ones we need
			// for this test.
			_, _ = w.Write([]byte(`{
			  "member_session": {
			    "expires_at": "2022-06-29T19:53:48Z",
			    "last_accessed_at": "2022-06-29T17:54:13Z",
			    "member_session_id": "session-test-aaaaaaaa-aaaa-aaaa-aaaa-aaaaaaaaaaaa",
			    "started_at": "2022-06-29T17:53:48Z",
			    "member_id": "user-test-00000000-0000-0000-0000-000000000000",

			    "custom_claims": {
			      "https://my-app.example.net/custom-claim": {
			        "number": 1,
			        "array": [1, "foo", null],
			        "nested": {
			          "data": "here"
			        }
			      }
			    }
			  },
				"member": {},
				"organization": {}
			}`))
			return
		}

		http.Error(w, "Bad Request", http.StatusBadRequest)
	}))

	client, _ := b2bstytchapi.NewClient(
		"project-test-00000000-0000-0000-0000-000000000000",
		"secret-test-11111111-1111-1111-1111-111111111111",
		b2bstytchapi.WithBaseURI(srv.URL),
	)

	// Expecting claims to follow this exact data structure
	type MyAppClaims struct {
		Number int
		Array  []interface{}
		Nested struct {
			Data string
		}
	}
	type StructClaims struct {
		MyApp MyAppClaims `json:"https://my-app.example.net/custom-claim"`
	}

	var structClaims StructClaims
	_, _ = client.Sessions.AuthenticateWithClaims(
		context.Background(),
		&sessions.AuthenticateParams{
			SessionToken: "fake session token",
		},
		&structClaims,
	)

	fmt.Println(structClaims)
	// Output: {{1 [1 foo <nil>] {here}}}
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

func sandboxClaims(t *testing.T, iat, exp time.Time) sessions.Claims {
	return sessions.Claims{
		Session: consumersessions.SessionClaim{
			ID:             "session-live-e26a0ccb-0dc0-4edb-a4bb-e70210f43555",
			StartedAt:      iat.Format(time.RFC3339),
			LastAccessedAt: iat.Format(time.RFC3339),
			ExpiresAt:      exp.Format(time.RFC3339),
			AuthenticationFactors: []consumersessions.AuthenticationFactor{
				{
					Type:                "magic_link",
					DeliveryMethod:      "email",
					LastAuthenticatedAt: &iat,
					EmailFactor: &consumersessions.EmailFactor{
						EmailAddress: "sandbox@stytch.com",
						EmailID:      "email-live-cca9d7d0-11b6-4167-9385-d7e0c9a77418",
					},
				},
			},
		},
		RegisteredClaims: jwt.RegisteredClaims{
			Issuer:    "stytch.com/project-test-00000000-0000-0000-0000-000000000000",
			Audience:  []string{"project-test-00000000-0000-0000-0000-000000000000"},
			Subject:   "member-live-fde03dd1-fff7-4b3c-9b31-ead3fbc224de",
			IssuedAt:  jwt.NewNumericDate(iat),
			NotBefore: jwt.NewNumericDate(iat),
			ExpiresAt: jwt.NewNumericDate(iat.Add(5 * time.Minute)),
		},
	}
}
