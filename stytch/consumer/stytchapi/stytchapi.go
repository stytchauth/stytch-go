package stytchapi

// !!!
// WARNING: This file is autogenerated
// Only modify code within MANUAL() sections
// or your changes may be overwritten later!
// !!!

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/MicahParks/keyfunc/v2"
	"github.com/stytchauth/stytch-go/v10/stytch"
	"github.com/stytchauth/stytch-go/v10/stytch/config"
	"github.com/stytchauth/stytch-go/v10/stytch/consumer"
)

type Logger interface {
	Printf(format string, v ...any)
}

type API struct {
	projectID string
	secret    string
	baseURI   config.BaseURI

	client                stytch.Client
	initializationContext context.Context
	logger                Logger

	CryptoWallets *consumer.CryptoWalletsClient
	MagicLinks    *consumer.MagicLinksClient
	OAuth         *consumer.OAuthClient
	OTPs          *consumer.OTPsClient
	Passwords     *consumer.PasswordsClient
	Sessions      *consumer.SessionsClient
	TOTPs         *consumer.TOTPsClient
	Users         *consumer.UsersClient
	WebAuthn      *consumer.WebAuthnClient
}

type Option func(*API)

func WithLogger(logger Logger) Option {
	return func(api *API) { api.logger = logger }
}

// WithClient overrides the stytch.Client used by the API client. This is useful for completely mocking out requests by
// using something like GoMock against the stytch.Client interface to customize the responses you receive from API
// methods.
//
// NOTE: You should not use this in conjunction with WithHTTPClient or WithBaseURI since the latter two assume usage of
// the default stytch.DefaultClient and this method completely overrides it to use anything conforming to the interface.
func WithClient(client stytch.Client) Option {
	return func(api *API) { api.client = client }
}

// WithHTTPClient overrides the HTTP client used by the API client. The default value is &http.Client{}.
//
// NOTE: You should not use this in conjunction with the WithClient option since WithClient completely overrides the
// stytch.Client with one that may not be a stytch.DefaultClient.
func WithHTTPClient(client *http.Client) Option {
	return func(api *API) {
		if defaultClient, ok := api.client.(*stytch.DefaultClient); ok {
			defaultClient.HTTPClient = client
		}
	}
}

// WithBaseURI overrides the client base URI determined by the environment.
//
// The value derived from stytch.EnvLive or stytch.EnvTest is already correct for production use
// in the Live or Test environment, respectively. This is implemented to make it easier to use
// this client to access internal development versions of the API.
//
// NOTE: You should not use this in conjunction with the WithClient option since WithClient completely overrides the
// stytch.Client with one that may not be a stytch.DefaultClient.
func WithBaseURI(uri string) Option {
	return func(api *API) {
		api.baseURI = config.BaseURI(uri)
		if defaultClient, ok := api.client.(*stytch.DefaultClient); ok {
			defaultClient.Config.BaseURI = config.BaseURI(uri)
		}
	}
}

// WithInitializationContext overrides the context used during initialization.
//
// The context argument is used only during client setup and can be used to cancel client
// creation. After the client is created and returned, canceling the context has no effect.
// It is preferred to use this function over the less flexible NewAPIClientWithContext function,
// which will be deprecated in a future MAJOR release.
func WithInitializationContext(ctx context.Context) Option {
	return func(api *API) { api.initializationContext = ctx }
}

// NewClient returns a Stytch API client that uses the provided credentials.
//
// It detects the environment from the given projectID. You are still free to pass WithBaseURI as an option if you wish
// to override this behavior, but the intention is to provide a simpler interface for creating a client since it's
// extremely rare that developers would want to use something other than the detected environment.
func NewClient(projectID string, secret string, opts ...Option) (*API, error) {
	defaultClient := stytch.New(projectID, secret)
	a := &API{
		projectID: projectID,
		secret:    secret,
		baseURI:   defaultClient.Config.BaseURI,

		client:                defaultClient,
		initializationContext: context.Background(),
	}
	for _, o := range opts {
		o(a)
	}

	a.CryptoWallets = consumer.NewCryptoWalletsClient(a.client)
	a.MagicLinks = consumer.NewMagicLinksClient(a.client)
	a.OAuth = consumer.NewOAuthClient(a.client)
	a.OTPs = consumer.NewOTPsClient(a.client)
	a.Passwords = consumer.NewPasswordsClient(a.client)
	a.Sessions = consumer.NewSessionsClient(a.client)
	a.TOTPs = consumer.NewTOTPsClient(a.client)
	a.Users = consumer.NewUsersClient(a.client)
	a.WebAuthn = consumer.NewWebAuthnClient(a.client)
	// Set up JWKS for local session authentication
	httpClient := defaultClient.HTTPClient
	if realClient, ok := a.client.(*stytch.DefaultClient); ok {
		httpClient = realClient.HTTPClient
	}
	jwks, err := a.instantiateJWKSClient(httpClient)
	if err != nil {
		return nil, fmt.Errorf("fetch JWKS from URL: %w", err)
	}
	a.Sessions.JWKS = jwks

	return a, nil
}

func (a *API) instantiateJWKSClient(httpClient *http.Client) (*keyfunc.JWKS, error) {
	// The context given in the keyfunc Options applies throughout the lifetime of the JWKS
	// fetcher. The context we were given here is _only_ for init, so we arrange to cancel the
	// JWKS context manually if we couldn't start in time.
	jwksCtx, jwksCancel := context.WithCancel(a.initializationContext)

	jwkOptions := keyfunc.Options{
		Client: httpClient,

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

	jwksURL := fmt.Sprintf("%s/v1/sessions/jwks/%s", a.baseURI, a.projectID)

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
	case <-a.initializationContext.Done():
		// Couldn't start in time, clean up the JWKS.
		jwksCancel()
		return nil, a.initializationContext.Err()
	case res := <-res:
		// JWKS setup finished first, _do not_ cancel its context. Let it continue fetching in the
		// background.
		_ = jwksCancel // lostcancel

		return res.jwks, res.err
	}
}
