package consumer_test

import (
	"context"
	"github.com/stytchauth/stytch-go/v9/stytch/stytcherror"
	"io"
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"
	"time"

	"github.com/MicahParks/keyfunc/v2"
	"github.com/golang-jwt/jwt/v5"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/stytchauth/stytch-go/v9/stytch"
	"github.com/stytchauth/stytch-go/v9/stytch/config"
	"github.com/stytchauth/stytch-go/v9/stytch/consumer"
	"github.com/stytchauth/stytch-go/v9/stytch/consumer/m2m"
)

func TestM2MClient_Token(t *testing.T) {
	client := &stytch.DefaultClient{
		Config: &config.Config{
			ProjectID: "some-project-id-0000-000-000-0000",
		},
		// In these tests, the keyset has already been downloaded, so no other network requests
		// should be made.
		HTTPClient: http.DefaultClient,
	}
	expectedClientID := "mock_client_id"
	expectedClientSecret := "mock_client_id"
	scopes := []string{"read:users", "write:users"}
	expectedScope := "read:users write:users"

	t.Run("sends token request", func(t *testing.T) {
		svr := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			assert.Equal(t, "/v1/public/some-project-id-0000-000-000-0000/oauth2/token", r.URL.Path)
			assert.Equal(t, "application/x-www-form-urlencoded", r.Header.Get("content-type"))
			bytes, err := io.ReadAll(r.Body)
			require.NoError(t, err)

			params, err := url.ParseQuery(string(bytes))
			require.NoError(t, err)

			assert.Equal(t, expectedClientID, params.Get("client_id"))
			assert.Equal(t, expectedClientSecret, params.Get("client_secret"))
			assert.Equal(t, "client_credentials", params.Get("grant_type"))
			assert.Equal(t, expectedScope, params.Get("scope"))

			w.WriteHeader(http.StatusOK)
			_, err = w.Write([]byte(`{ "access_token": "mock-token", "expires_in": 3600, "token_type": "bearer" }`))
			assert.NoError(t, err)
		}))

		client.Config.BaseURI = config.BaseURI(svr.URL)

		res, err := consumer.NewM2MClient(client).Token(context.Background(), &m2m.TokenParams{
			ClientID:     expectedClientID,
			ClientSecret: expectedClientSecret,
			Scopes:       scopes,
		})
		require.NoError(t, err)

		expected := &m2m.TokenResponse{
			AccessToken: "mock-token",
			ExpiresIn:   3600,
			TokenType:   "bearer",
		}
		assert.Equal(t, expected, res)
	})

	t.Run("handles error response", func(t *testing.T) {
		svr := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.WriteHeader(http.StatusBadRequest)
			_, err := w.Write([]byte(`{ "error":"invalid_client" }`))
			assert.NoError(t, err)
		}))

		client.Config.BaseURI = config.BaseURI(svr.URL)

		res, err := consumer.NewM2MClient(client).Token(context.Background(), &m2m.TokenParams{
			ClientID:     expectedClientID,
			ClientSecret: expectedClientSecret,
			Scopes:       scopes,
		})
		assert.Nil(t, res)
		var stytchErr stytcherror.OAuth2Error
		assert.ErrorAs(t, err, &stytchErr)
	})
}

func TestM2MClient_AuthenticateM2MToken(t *testing.T) {
	client := &stytch.DefaultClient{
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
		keyID: keyfunc.NewGivenRSA(&key.PublicKey, keyfunc.GivenKeyOptions{Algorithm: "RS256"}),
	})

	m2mClient := consumer.NewM2MClient(client)
	m2mClient.JWKS = jwks

	t.Run("expired JWT", func(t *testing.T) {
		iat := time.Now().UTC().Add(-time.Hour).Truncate(time.Second)
		exp := iat.Add(time.Minute)

		claims := sandboxM2MClaims(t, iat, exp, "read:users")
		token := signJWT(t, keyID, key, claims)

		s, err := m2mClient.AuthenticateM2MToken(token)
		assert.ErrorIs(t, err, jwt.ErrTokenExpired)
		assert.Nil(t, s)
	})

	t.Run("stale JWT", func(t *testing.T) {
		iat := time.Now().UTC().Add(-3 * time.Minute).Truncate(time.Second)
		exp := iat.Add(time.Hour)

		claims := sandboxM2MClaims(t, iat, exp, "read:users")
		token := signJWT(t, keyID, key, claims)

		s, err := m2mClient.AuthenticateM2MToken(token, m2m.WithMaxTokenAge(time.Minute))
		assert.ErrorIs(t, err, m2m.ErrJWTTooOld)
		assert.Nil(t, s)
	})

	t.Run("incorrect audience", func(t *testing.T) {
		iat := time.Now().UTC().Truncate(time.Second)
		exp := iat.Add(time.Hour)

		claims := sandboxM2MClaims(t, iat, exp, "read:users")
		claims.MapClaims["aud"] = "not this project"

		token := signJWT(t, keyID, key, claims)

		s, err := m2mClient.AuthenticateM2MToken(token)
		assert.ErrorIs(t, err, jwt.ErrTokenInvalidAudience)
		assert.Nil(t, s)
	})

	t.Run("incorrect issuer", func(t *testing.T) {
		iat := time.Now().UTC().Truncate(time.Second)
		exp := iat.Add(time.Hour)

		claims := sandboxM2MClaims(t, iat, exp, "read:users")
		claims.MapClaims["iss"] = "not this project"

		token := signJWT(t, keyID, key, claims)

		s, err := m2mClient.AuthenticateM2MToken(token)
		assert.ErrorIs(t, err, jwt.ErrTokenInvalidIssuer)
		assert.Nil(t, s)
	})

	t.Run("missing scopes", func(t *testing.T) {
		iat := time.Now().UTC().Truncate(time.Second)
		exp := iat.Add(time.Hour)

		claims := sandboxM2MClaims(t, iat, exp, "read:users")

		token := signJWT(t, keyID, key, claims)

		s, err := m2mClient.AuthenticateM2MToken(token, m2m.WithRequiredScopes("write:users"))
		assert.ErrorIs(t, err, m2m.ErrMissingScope)
		assert.Nil(t, s)
	})

	t.Run("valid JWT", func(t *testing.T) {
		iat := time.Now().UTC().Truncate(time.Second)
		exp := iat.Add(time.Hour)

		claims := sandboxM2MClaims(t, iat, exp, "read:users read:books write:penguins")

		token := signJWT(t, keyID, key, claims)

		tok, err := m2mClient.AuthenticateM2MToken(token, m2m.WithRequiredScopes("write:penguins", "read:books"))
		require.NoError(t, err)

		expected := &m2m.AuthenticateM2MTokenResponse{
			Scopes:   []string{"read:users", "read:books", "write:penguins"},
			ClientID: "m2m-client-live-63532f0c-b600-425b-a5b5-3a42ead94a8e",
			CustomClaims: map[string]any{
				"custom key": "custom value",
			},
		}

		assert.Equal(t, expected, tok)
	})
}

func sandboxM2MClaims(t *testing.T, iat, exp time.Time, scope string) m2m.Claims {
	return m2m.Claims{
		Scope: scope,
		MapClaims: jwt.MapClaims{
			"iss":        "stytch.com/project-test-00000000-0000-0000-0000-000000000000",
			"aud":        []string{"project-test-00000000-0000-0000-0000-000000000000"},
			"sub":        "m2m-client-live-63532f0c-b600-425b-a5b5-3a42ead94a8e",
			"iat":        jwt.NewNumericDate(iat),
			"nbf":        jwt.NewNumericDate(iat),
			"exp":        jwt.NewNumericDate(iat.Add(5 * time.Minute)),
			"custom key": "custom value",
		},
	}
}
