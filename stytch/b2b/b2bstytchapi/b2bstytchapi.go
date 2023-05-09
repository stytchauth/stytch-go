package b2bstytchapi

import (
	"github.com/stytchauth/stytch-go/v9/stytch"
	"github.com/stytchauth/stytch-go/v9/stytch/b2b/discovery"
	"github.com/stytchauth/stytch-go/v9/stytch/b2b/magiclink"
	mle "github.com/stytchauth/stytch-go/v9/stytch/b2b/magiclink/email"
	"github.com/stytchauth/stytch-go/v9/stytch/b2b/organization"
	"github.com/stytchauth/stytch-go/v9/stytch/b2b/password"
	pe "github.com/stytchauth/stytch-go/v9/stytch/b2b/password/email"
	pep "github.com/stytchauth/stytch-go/v9/stytch/b2b/password/existingpassword"
	ps "github.com/stytchauth/stytch-go/v9/stytch/b2b/password/session"
	"github.com/stytchauth/stytch-go/v9/stytch/b2b/session"
	"github.com/stytchauth/stytch-go/v9/stytch/b2b/sso"
	"github.com/stytchauth/stytch-go/v9/stytch/b2b/sso/oidc"
	"github.com/stytchauth/stytch-go/v9/stytch/b2b/sso/saml"
	"github.com/stytchauth/stytch-go/v9/stytch/config"
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
