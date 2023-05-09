package stytchapi_test

import (
	"context"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/stytchauth/stytch-go/v8/stytch/b2c"
	"github.com/stytchauth/stytch-go/v8/stytch/b2c/stytchapi"

	"github.com/stretchr/testify/assert"
	"github.com/stytchauth/stytch-go/v8/stytch"
	"github.com/stytchauth/stytch-go/v8/stytch/config"
)

func TestNewClient(t *testing.T) {
	t.Run("internal development override", func(t *testing.T) {
		srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			// Handle the JWKS fetch that happens during setup.
			if strings.HasPrefix(r.URL.Path, "/sessions/jwks/") {
				_, _ = w.Write([]byte(`{"keys": []}`))
				return
			}

			// This is the test request.
			if r.URL.Path == "/magic_links/authenticate" {
				_, _ = w.Write([]byte(`{}`))
				return
			}

			http.Error(w, "Bad Request", http.StatusBadRequest)
		}))

		client, err := stytchapi.NewAPIClient(
			config.Env("anything"),
			"project-test-00000000-0000-0000-0000-000000000000",
			"secret-test-11111111-1111-1111-1111-111111111111",
			stytchapi.WithBaseURI(srv.URL),
		)
		assert.NoError(t, err)

		_, err = client.MagicLinks.Authenticate(
			context.Background(),
			&b2c.MagicLinksAuthenticateParams{
				Token: "fake-token",
			},
		)
		assert.NoError(t, err)
	})
}
