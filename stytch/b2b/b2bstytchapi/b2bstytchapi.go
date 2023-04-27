package b2bstytchapi

import (
	"github.com/stytchauth/stytch-go/v7/stytch"
	"github.com/stytchauth/stytch-go/v7/stytch/b2b/discovery"
	"github.com/stytchauth/stytch-go/v7/stytch/b2b/magiclink"
	mle "github.com/stytchauth/stytch-go/v7/stytch/b2b/magiclink/email"
	"github.com/stytchauth/stytch-go/v7/stytch/b2b/organization"
	"github.com/stytchauth/stytch-go/v7/stytch/b2b/session"
	"github.com/stytchauth/stytch-go/v7/stytch/config"
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
	Discovery    *discovery.Client
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
	a.Discovery = &discovery.Client{C: a.client}
	return a, nil
}
