package stytchapi

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/stytchauth/stytch-go/v8/stytch/b2c/cryptowallet"
	"github.com/stytchauth/stytch-go/v8/stytch/b2c/magiclink"
	mle "github.com/stytchauth/stytch-go/v8/stytch/b2c/magiclink/email"
	"github.com/stytchauth/stytch-go/v8/stytch/b2c/oauth"
	"github.com/stytchauth/stytch-go/v8/stytch/b2c/otp"
	otpe "github.com/stytchauth/stytch-go/v8/stytch/b2c/otp/email"
	"github.com/stytchauth/stytch-go/v8/stytch/b2c/otp/sms"
	"github.com/stytchauth/stytch-go/v8/stytch/b2c/otp/whatsapp"
	"github.com/stytchauth/stytch-go/v8/stytch/b2c/password"
	pe "github.com/stytchauth/stytch-go/v8/stytch/b2c/password/email"
	pep "github.com/stytchauth/stytch-go/v8/stytch/b2c/password/existingpassword"
	ps "github.com/stytchauth/stytch-go/v8/stytch/b2c/password/session"
	"github.com/stytchauth/stytch-go/v8/stytch/b2c/session"
	"github.com/stytchauth/stytch-go/v8/stytch/b2c/totp"
	"github.com/stytchauth/stytch-go/v8/stytch/b2c/user"
	"github.com/stytchauth/stytch-go/v8/stytch/b2c/webauthn"

	"github.com/MicahParks/keyfunc/v2"
	"github.com/stytchauth/stytch-go/v8/stytch"
	"github.com/stytchauth/stytch-go/v8/stytch/config"
)

type Logger interface {
	Printf(format string, v ...interface{})
}

type API struct {
	client        *stytch.Client
	logger        Logger
	CryptoWallets *cryptowallet.Client
	MagicLinks    *magiclink.Client
	OAuth         *oauth.Client
	OTPs          *otp.Client
	Passwords     *password.Client
	Sessions      *session.Client
	TOTPs         *totp.Client
	Users         *user.Client
	WebAuthn      *webauthn.Client
}

type Option func(*API)

func WithLogger(logger Logger) Option {
	return func(api *API) { api.logger = logger }
}

// WithHTTPClient overrides the HTTP client used by the API client. The default value is
// &http.Client{}.
func WithHTTPClient(client *http.Client) Option {
	return func(api *API) { api.client.HTTPClient = client }
}

// WithBaseURI overrides the client base URI determined by the environment.
//
// The value derived from stytch.EnvLive or stytch.EnvTest is already correct for production use
// in the Live or Test environment, respectively. This is implemented to make it easier to use
// this client to access internal development versions of the API.
func WithBaseURI(uri string) Option {
	return func(api *API) { api.client.Config.BaseURI = config.BaseURI(uri) }
}

// NewAPIClient returns a Stytch API client that uses the provided credentials.
func NewAPIClient(env config.Env, projectID string, secret string, opts ...Option) (*API, error) {
	return NewAPIClientWithContext(context.Background(), env, projectID, secret, opts...)
}

// NewAPIClientWithContext is like NewAPIClient but with extra control over the initialization
// context.
//
// The context argument is used only during client setup and can be used to cancel client
// creation. After the client is created and returned by this function, canceling the context has
// no effect.
func NewAPIClientWithContext(ctx context.Context, env config.Env, projectID string, secret string, opts ...Option) (*API, error) {
	a := &API{
		client: stytch.New(env, projectID, secret),
	}
	for _, o := range opts {
		o(a)
	}

	a.CryptoWallets = &cryptowallet.Client{C: a.client}
	a.MagicLinks = &magiclink.Client{C: a.client, Email: &mle.Client{C: a.client}}
	a.OTPs = &otp.Client{
		C:        a.client,
		Email:    &otpe.Client{C: a.client},
		SMS:      &sms.Client{C: a.client},
		WhatsApp: &whatsapp.Client{C: a.client},
	}
	a.OAuth = &oauth.Client{C: a.client}
	a.Passwords = &password.Client{
		C:                a.client,
		Email:            &pe.Client{C: a.client},
		ExistingPassword: &pep.Client{C: a.client},
		Session:          &ps.Client{C: a.client},
	}
	a.TOTPs = &totp.Client{C: a.client}
	a.Users = &user.Client{C: a.client}
	a.WebAuthn = &webauthn.Client{C: a.client}
	jwks, err := a.instantiateJWKSClient(ctx, a.client)
	if err != nil {
		return nil, fmt.Errorf("fetch JWKS from URL: %w", err)
	}
	a.Sessions = &session.Client{C: a.client, JWKS: jwks}
	return a, nil
}

func (a *API) instantiateJWKSClient(ctx context.Context, client *stytch.Client) (*keyfunc.JWKS, error) {
	// The context given in the keyfunc Options applies throughout the lifetime of the JWKS
	// fetcher. The context we were given here is _only_ for init, so we arrange to cancel the
	// JWKS context manually if we couldn't start in time.
	jwksCtx, jwksCancel := context.WithCancel(ctx)

	jwkOptions := keyfunc.Options{
		Client: client.HTTPClient,

		// This is the context for ongoing background JWKS fetches. If the keyfunc starts in time,
		// it should run until API.Close is called.
		Ctx: jwksCtx,

		RefreshErrorHandler: func(err error) {
			if a.logger != nil {
				a.logger.Printf("There was an error with the jwt.Keyfunc\nError: %s", err.Error())
			}
		},
		RefreshInterval:   time.Hour,
		RefreshRateLimit:  5 * time.Minute,
		RefreshTimeout:    10 * time.Second,
		RefreshUnknownKID: true,
	}

	baseURI := client.Config.GetBaseURI()
	projectID := client.Config.BasicAuthProjectID()
	jwksURL := fmt.Sprintf("%s/sessions/jwks/%s", baseURI, projectID)

	type Res struct {
		jwks *keyfunc.JWKS
		err  error
	}
	res := make(chan Res)
	go func() {
		jwks, err := keyfunc.Get(jwksURL, jwkOptions)
		res <- Res{jwks, err}
	}()

	select {
	case <-ctx.Done():
		// Couldn't start in time, clean up the JWKS.
		jwksCancel()
		return nil, ctx.Err()
	case res := <-res:
		// JWKS setup finished first, _do not_ cancel its context. Let it continue fetching in the
		// background.
		_ = jwksCancel // lostcancel

		return res.jwks, res.err
	}
}
