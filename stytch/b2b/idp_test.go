package b2b_test

import (
	"context"
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
	"github.com/stytchauth/stytch-go/v16/stytch"
	"github.com/stytchauth/stytch-go/v16/stytch/b2b"
	"github.com/stytchauth/stytch-go/v16/stytch/b2b/idp"
	"github.com/stytchauth/stytch-go/v16/stytch/config"
	"github.com/stytchauth/stytch-go/v16/stytch/stytcherror"
)

func TestIDP_IntrospectTokenNetwork(t *testing.T) {
	client := &stytch.DefaultClient{
		Config: &config.Config{
			ProjectID: "some-project-id-0000-000-000-0000",
		},
		HTTPClient: http.DefaultClient,
	}
	expectedToken := "expectedToken"
	expectedClientID := "mock_client_id"
	expectedClientSecret := "mock_client_id"

	t.Run("sends token request", func(t *testing.T) {
		svr := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			assert.Equal(t, "/v1/public/some-project-id-0000-000-000-0000/oauth2/introspect", r.URL.Path)
			assert.Equal(t, "application/x-www-form-urlencoded", r.Header.Get("content-type"))
			bytes, err := io.ReadAll(r.Body)
			require.NoError(t, err)

			params, err := url.ParseQuery(string(bytes))
			require.NoError(t, err)

			assert.Equal(t, expectedClientID, params.Get("client_id"))
			assert.Equal(t, expectedClientSecret, params.Get("client_secret"))
			assert.Equal(t, expectedToken, params.Get("token"))

			w.WriteHeader(http.StatusOK)
			_, err = w.Write([]byte(`{
				"active": true,
				"aud": ["project-test-0000-000-000-0000"],
				"client_id": "connected-app-test-0000-000-000-0000",
				"exp": 1738848103,
				"iat": 1738844503,
				"iss": "https://upn.customers.stytch.dev",
				"scope": "openid email profile",
				"sub": "member-test-0000-000-000-0000",
				"token_type": "access_token",
				"https://stytch.com/organization": {
					"organization_id": "org-test-0000-000-000-0000",
					"slug": "some-slug"
				}
			}`))
			assert.NoError(t, err)
		}))

		client.Config.BaseURI = config.BaseURI(svr.URL)

		res, err := b2b.NewIDPClient(client, nil, nil).
			IntrospectTokenNetwork(context.Background(), &idp.IntrospectTokenNetworkParams{
				Token:        expectedToken,
				ClientID:     expectedClientID,
				ClientSecret: &expectedClientSecret,
			})
		require.NoError(t, err)

		expected := &idp.IntrospectTokenResponse{
			Active:       true,
			TokenType:    "access_token",
			Issuer:       "https://upn.customers.stytch.dev",
			Subject:      "member-test-0000-000-000-0000",
			Audience:     []string{"project-test-0000-000-000-0000"},
			Scope:        "openid email profile",
			CustomClaims: nil,
			ClientID:     "connected-app-test-0000-000-000-0000",
			Organization: idp.OrganizationClaim{
				OrganizationID: "org-test-0000-000-000-0000",
				Slug:           "some-slug",
			},
			Expiry:   res.Expiry,
			IssuedAt: res.IssuedAt,
		}
		assert.Equal(t, expected, res)
	})

	t.Run("handles inactive token response", func(t *testing.T) {
		svr := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.WriteHeader(http.StatusOK)
			_, err := w.Write([]byte(`{
				"active": false
			}`))
			assert.NoError(t, err)
		}))

		client.Config.BaseURI = config.BaseURI(svr.URL)

		_, err := b2b.NewIDPClient(client, nil, nil).
			IntrospectTokenNetwork(context.Background(), &idp.IntrospectTokenNetworkParams{
				Token:        expectedToken,
				ClientID:     expectedClientID,
				ClientSecret: &expectedClientSecret,
			})
		assert.EqualError(t, err, stytcherror.NewInvalidOAuth2TokenError().Error())
	})

	t.Run("handles error response", func(t *testing.T) {
		svr := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.WriteHeader(http.StatusBadRequest)
			_, err := w.Write([]byte(`{ "error":"invalid_client" }`))
			assert.NoError(t, err)
		}))

		client.Config.BaseURI = config.BaseURI(svr.URL)

		res, err := b2b.NewIDPClient(client, nil, nil).
			IntrospectTokenNetwork(context.Background(), &idp.IntrospectTokenNetworkParams{
				Token:        expectedToken,
				ClientID:     expectedClientID,
				ClientSecret: &expectedClientSecret,
			})
		assert.Nil(t, res)
		var stytchErr stytcherror.OAuth2Error
		assert.ErrorAs(t, err, &stytchErr)
	})
}

func TestIDP_IntrospectTokenLocal(t *testing.T) {
	client := &stytch.DefaultClient{
		Config: &config.Config{
			ProjectID: "project-test-0000-000-000-0000",
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
	idpClient := b2b.NewIDPClient(client, jwks, policyCache)

	t.Run("valid JWT", func(t *testing.T) {
		token := signJWT(t, keyID, key, jwt.MapClaims{
			"iss":       "https://upn.customers.stytch.dev",
			"sub":       "member-test-0000-000-000-0000",
			"aud":       []string{"project-test-0000-000-000-0000"},
			"iat":       time.Now().Add(-time.Minute).Unix(),
			"exp":       time.Now().Add(time.Minute).Unix(),
			"client_id": "connected-app-test-0000-000-000-0000",
			"scope":     "openid email profile",
			"https://stytch.com/organization": map[string]string{
				"organization_id": "org-test-0000-000-000-0000",
				"slug":            "some-slug",
			},
		})

		ctx := context.Background()
		res, err := idpClient.IntrospectTokenLocal(ctx, &idp.IntrospectTokenLocalParams{
			Token: token,
		})
		require.NoError(t, err)

		expected := &idp.IntrospectTokenResponse{
			Active:       true,
			TokenType:    "access_token",
			Issuer:       "https://upn.customers.stytch.dev",
			Subject:      "member-test-0000-000-000-0000",
			Audience:     []string{"project-test-0000-000-000-0000"},
			Scope:        "openid email profile",
			CustomClaims: map[string]any{},
			ClientID:     "connected-app-test-0000-000-000-0000",
			Organization: idp.OrganizationClaim{
				OrganizationID: "org-test-0000-000-000-0000",
				Slug:           "some-slug",
			},
			Expiry:   res.Expiry,
			IssuedAt: res.IssuedAt,
		}
		assert.Equal(t, expected, res)
	})
}
