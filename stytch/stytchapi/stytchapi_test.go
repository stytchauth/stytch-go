package stytchapi_test

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stytchauth/stytch-go/v5/stytch"
	"github.com/stytchauth/stytch-go/v5/stytch/config"
	"github.com/stytchauth/stytch-go/v5/stytch/stytchapi"
)

func TestNewClient(t *testing.T) {
	t.Run("live environment", func(t *testing.T) {
		_, err := stytchapi.NewAPIClient(
			stytch.EnvLive,
			"project-live-00000000-0000-0000-0000-000000000000",
			"secret-live-11111111-1111-1111-1111-111111111111",
		)
		assert.NoError(t, err)
	})

	t.Run("test environment", func(t *testing.T) {
		_, err := stytchapi.NewAPIClient(
			stytch.EnvTest,
			"project-test-00000000-0000-0000-0000-000000000000",
			"secret-test-11111111-1111-1111-1111-111111111111",
		)
		assert.NoError(t, err)
	})

	t.Run("internal development override", func(t *testing.T) {
		srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			// Handle the async JWKS fetch
			if strings.HasPrefix(r.URL.Path, "/sessions/jwks/") {
				w.Write([]byte(`{"keys": []}`))
				return
			}

			// This is the test request
			if r.URL.Path == "/magic_links/authenticate" {
				w.Write([]byte(`{}`))
				return
			}

			http.Error(w, "Bad Request", http.StatusBadRequest)
			return
		}))

		client, err := stytchapi.NewAPIClient(
			config.Env("anything"),
			"project-test-00000000-0000-0000-0000-000000000000",
			"secret-test-11111111-1111-1111-1111-111111111111",
			stytchapi.WithBaseURI(srv.URL),
		)
		assert.NoError(t, err)

		_, err = client.MagicLinks.Authenticate(&stytch.MagicLinksAuthenticateParams{
			Token: "fake-token",
		})
		assert.NoError(t, err)
	})
}
