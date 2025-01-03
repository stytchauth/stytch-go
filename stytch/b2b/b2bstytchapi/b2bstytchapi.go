package b2bstytchapi

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
	"github.com/stytchauth/stytch-go/v15/stytch"
	"github.com/stytchauth/stytch-go/v15/stytch/b2b"
	"github.com/stytchauth/stytch-go/v15/stytch/config"
	"github.com/stytchauth/stytch-go/v15/stytch/consumer"
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

	shouldSkipJWKSInitialization bool

	Discovery     *b2b.DiscoveryClient
	Fraud         *consumer.FraudClient
	M2M           *consumer.M2MClient
	MagicLinks    *b2b.MagicLinksClient
	OAuth         *b2b.OAuthClient
	OTPs          *b2b.OTPsClient
	Organizations *b2b.OrganizationsClient
	Passwords     *b2b.PasswordsClient
	Project       *consumer.ProjectClient
	RBAC          *b2b.RBACClient
	RecoveryCodes *b2b.RecoveryCodesClient
	SCIM          *b2b.SCIMClient
	SSO           *b2b.SSOClient
	Sessions      *b2b.SessionsClient
	TOTPs         *b2b.TOTPsClient
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

// WithFraudBaseURI overrides the client base fraud URI. This is implemented to make it easier to use
// this client to access internal development versions of the API.
//
// NOTE: You should not use this in conjunction with the WithClient option since WithClient completely overrides the
// stytch.Client with one that may not be a stytch.DefaultClient.
func WithFraudBaseURI(uri string) Option {
	return func(api *API) {
		if defaultClient, ok := api.client.(*stytch.DefaultClient); ok {
			defaultClient.Config.FraudBaseURI = config.BaseURI(uri)
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

// WithSkipJWKSInitialization skips the initialization of the JWKS client. This can be useful for testing purposes.
// Please note that if you utilize this option, any API method that makes use of the JWKS client will raise a
// stytcherror.JWKSNotInitialized error. If you need to call such a method, you should use WithClient or WithHTTPClient
// with a client that is capable of ininitalizing the JWKS keyfunc.
func WithSkipJWKSInitialization() Option {
	return func(api *API) { api.shouldSkipJWKSInitialization = true }
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

		shouldSkipJWKSInitialization: false,
	}
	for _, o := range opts {
		o(a)
	}

	policyCache := b2b.NewPolicyCache(b2b.NewRBACClient(a.client))

	// Set up JWKS for local session authentication
	jwks, err := a.instantiateJWKSClient(a.client.GetHTTPClient())
	if err != nil {
		return nil, fmt.Errorf("fetch JWKS from URL: %w", err)
	}

	a.Discovery = b2b.NewDiscoveryClient(a.client)
	a.Fraud = consumer.NewFraudClient(a.client)
	a.M2M = consumer.NewM2MClient(a.client, jwks)
	a.MagicLinks = b2b.NewMagicLinksClient(a.client)
	a.OAuth = b2b.NewOAuthClient(a.client)
	a.OTPs = b2b.NewOTPsClient(a.client)
	a.Organizations = b2b.NewOrganizationsClient(a.client)
	a.Passwords = b2b.NewPasswordsClient(a.client)
	a.Project = consumer.NewProjectClient(a.client)
	a.RBAC = b2b.NewRBACClient(a.client)
	a.RecoveryCodes = b2b.NewRecoveryCodesClient(a.client)
	a.SCIM = b2b.NewSCIMClient(a.client)
	a.SSO = b2b.NewSSOClient(a.client)
	a.Sessions = b2b.NewSessionsClient(a.client, jwks, policyCache)
	a.TOTPs = b2b.NewTOTPsClient(a.client)

	return a, nil
}

func (a *API) instantiateJWKSClient(httpClient *http.Client) (*keyfunc.JWKS, error) {
	if a.shouldSkipJWKSInitialization {
		if a.logger != nil {
			a.logger.Printf("Skipping JWKS initialization")
		}
		return nil, nil
	}
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

	jwksURL := fmt.Sprintf("%s/v1/b2b/sessions/jwks/%s", a.baseURI, a.projectID)

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
