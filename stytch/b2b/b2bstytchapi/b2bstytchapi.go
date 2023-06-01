package b2bstytchapi

import (
	"strings"

	"github.com/stytchauth/stytch-go/v8/stytch"
	"github.com/stytchauth/stytch-go/v8/stytch/b2b/discovery"
	"github.com/stytchauth/stytch-go/v8/stytch/b2b/magiclink"
	mle "github.com/stytchauth/stytch-go/v8/stytch/b2b/magiclink/email"
	"github.com/stytchauth/stytch-go/v8/stytch/b2b/organization"
	"github.com/stytchauth/stytch-go/v8/stytch/b2b/password"
	pe "github.com/stytchauth/stytch-go/v8/stytch/b2b/password/email"
	pep "github.com/stytchauth/stytch-go/v8/stytch/b2b/password/existingpassword"
	ps "github.com/stytchauth/stytch-go/v8/stytch/b2b/password/session"
	"github.com/stytchauth/stytch-go/v8/stytch/b2b/session"
	"github.com/stytchauth/stytch-go/v8/stytch/b2b/sso"
	"github.com/stytchauth/stytch-go/v8/stytch/b2b/sso/oidc"
	"github.com/stytchauth/stytch-go/v8/stytch/b2b/sso/saml"
	"github.com/stytchauth/stytch-go/v8/stytch/config"
)

type Logger interface {
	Printf(format string, v ...interface{})
}

type API struct {
	client       *stytch.Client
	logger       Logger
	MagicLinks   *magiclink.Client
	Organization *organization.Client
	Session      *session.Client
	Passwords    *password.Client
	Discovery    *discovery.Client
	SSO          *sso.Client
}

type Option func(*API)

func WithLogger(logger Logger) Option {
	return func(api *API) { api.logger = logger }
}

// WithBaseURI overrides the client base URI determined by the environment.
//
// The value derived from stytch.EnvLive or stytch.EnvTest is already correct for production use
// in the Live or Test environment, respectively. This is implemented to make it easier to use
// this client to access internal development versions of the API.
func WithBaseURI(uri string) Option {
	return func(api *API) { api.client.Config.BaseURI = config.BaseURI(uri) }
}

// NewClient returns a Stytch API client that uses the provided credentials.
//
// It detects the environment from the given projectID. You are still free to pass WithBaseURI as an option if you wish
// to override this behavior, but the intention is to provide a simpler interface for creating a client since it's
// extremely rare that developers would want to use something other than the detected environment.
func NewClient(projectID string, secret string, opts ...Option) (*API, error) {
	var detectedEnv config.Env
	if strings.HasPrefix(projectID, "project-live-") {
		detectedEnv = config.EnvLive
	} else {
		detectedEnv = config.EnvTest
	}

	return NewAPIClient(detectedEnv, projectID, secret, opts...)
}

// NewAPIClient returns a Stytch API client that uses the provided credentials.
//
// Deprecated: This method requires explicitly supplying a config.Env instead of detecting it automatically, which can
// lead to bugs when the env does not match what's expected from the given projectID. Use NewClient instead and supply a
// WithBaseURI if you need to explicitly override the client's BaseURI (typically only done for internal development).
func NewAPIClient(env config.Env, projectID string, secret string, opts ...Option) (*API, error) {
	a := &API{
		client: stytch.New(env, projectID, secret),
	}
	for _, o := range opts {
		o(a)
	}

	a.MagicLinks = &magiclink.Client{C: a.client, Email: &mle.Client{C: a.client}}
	a.Organization = &organization.Client{C: a.client}
	a.Session = &session.Client{C: a.client}
	a.Passwords = &password.Client{
		C:                a.client,
		Email:            &pe.Client{C: a.client},
		Session:          &ps.Client{C: a.client},
		ExistingPassword: &pep.Client{C: a.client},
	}
	a.Discovery = &discovery.Client{C: a.client}
	a.SSO = &sso.Client{C: a.client, OIDC: &oidc.Client{C: a.client}, SAML: &saml.Client{C: a.client}}
	return a, nil
}
