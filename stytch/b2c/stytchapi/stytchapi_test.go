package stytchapi_test

import (
	"context"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/stytchauth/stytch-go/v8/stytch/b2c"
	"github.com/stytchauth/stytch-go/v8/stytch/b2c/stytchapi"
	"github.com/stytchauth/stytch-go/v8/stytch/stytcherror"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
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

	t.Run("custom HTTP client", func(t *testing.T) {
		// This custom HTTP client intercepts all outbound requests and responds to them with a
		// fictional Stytch-like error.
		httpClient := &http.Client{
			Transport: RoundTripperFunc(func(r *http.Request) (*http.Response, error) {
				// Handle the JWKS fetch that happens during setup.
				if r.URL.Path == "/v1/sessions/jwks/project-test-00000000-0000-0000-0000-000000000000" {
					resp := &http.Response{
						StatusCode: http.StatusOK,
						Body:       io.NopCloser(strings.NewReader(`{"keys": []}`)),
					}
					return resp, nil
				}

				// Everything else is an error.
				resp := &http.Response{
					StatusCode: http.StatusTeapot,
					Body: io.NopCloser(strings.NewReader(
						`{"status_code": 418, "error_type": "teapot", "error_message": "I'm a teapot!"}`,
					)),
				}
				return resp, nil
			}),
		}

		client, err := stytchapi.NewAPIClient(
			stytch.EnvTest,
			"project-test-00000000-0000-0000-0000-000000000000",
			"secret-test-11111111-1111-1111-1111-111111111111",
			stytchapi.WithHTTPClient(httpClient),
		)
		require.NoError(t, err)

		_, err = client.MagicLinks.Authenticate(
			context.Background(),
			&b2c.MagicLinksAuthenticateParams{},
		)

		var stytchErr stytcherror.Error
		if assert.ErrorAs(t, err, &stytchErr) {
			assert.Equal(t, stytcherror.Type("teapot"), stytchErr.ErrorType)
			assert.Equal(t, stytcherror.Message("I'm a teapot!"), stytchErr.ErrorMessage)
		}
	})

	t.Run("JWKS URL", func(t *testing.T) {
		httpClient := &http.Client{
			Transport: RoundTripperFunc(func(r *http.Request) (*http.Response, error) {
				// Make sure the path got built correctly.
				if assert.Equal(t, "/v1/sessions/jwks/project-test-00000000-0000-0000-0000-000000000000", r.URL.Path) {
					resp := &http.Response{
						StatusCode: http.StatusOK,
						Body:       io.NopCloser(strings.NewReader(`{"keys": []}`)),
					}
					return resp, nil
				}

				// And then fail it anyway so this doesn't hit the real internet.
				resp := &http.Response{
					StatusCode: http.StatusTeapot,
					Body: io.NopCloser(strings.NewReader(
						`{"status_code": 418, "error_type": "teapot", "error_message": "I'm a teapot!"}`,
					)),
				}
				return resp, nil
			}),
		}

		_, err := stytchapi.NewAPIClient(
			stytch.EnvTest,
			"project-test-00000000-0000-0000-0000-000000000000",
			"secret-test-11111111-1111-1111-1111-111111111111",
			stytchapi.WithHTTPClient(httpClient),
		)
		require.NoError(t, err)
	})
}

type RoundTripperFunc func(*http.Request) (*http.Response, error)

func (r RoundTripperFunc) RoundTrip(req *http.Request) (*http.Response, error) {
	return r(req)
}
