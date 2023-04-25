package stytchapi

import (
	"context"
	"fmt"
	"time"

	"github.com/stytchauth/stytch-go/v7/stytch/b2c/cryptowallet"
	"github.com/stytchauth/stytch-go/v7/stytch/b2c/magiclink"
	mle "github.com/stytchauth/stytch-go/v7/stytch/b2c/magiclink/email"
	"github.com/stytchauth/stytch-go/v7/stytch/b2c/oauth"
	"github.com/stytchauth/stytch-go/v7/stytch/b2c/otp"
	otpe "github.com/stytchauth/stytch-go/v7/stytch/b2c/otp/email"
	"github.com/stytchauth/stytch-go/v7/stytch/b2c/otp/sms"
	"github.com/stytchauth/stytch-go/v7/stytch/b2c/otp/whatsapp"
	"github.com/stytchauth/stytch-go/v7/stytch/b2c/password"
	pe "github.com/stytchauth/stytch-go/v7/stytch/b2c/password/email"
	pep "github.com/stytchauth/stytch-go/v7/stytch/b2c/password/existingpassword"
	ps "github.com/stytchauth/stytch-go/v7/stytch/b2c/password/session"
	"github.com/stytchauth/stytch-go/v7/stytch/b2c/session"
	"github.com/stytchauth/stytch-go/v7/stytch/b2c/totp"
	"github.com/stytchauth/stytch-go/v7/stytch/b2c/user"
	"github.com/stytchauth/stytch-go/v7/stytch/b2c/webauthn"

	"github.com/MicahParks/keyfunc"
	"github.com/stytchauth/stytch-go/v7/stytch"
	"github.com/stytchauth/stytch-go/v7/stytch/config"
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
	jwks, err := a.instantiateJWKSClient(a.client)
	if err != nil {
		return nil, fmt.Errorf("fetch JWKS from URL: %w", err)
	}
	a.Sessions = &session.Client{C: a.client, JWKS: jwks}
	return a, nil
}

func (a *API) instantiateJWKSClient(client *stytch.Client) (*keyfunc.JWKS, error) {
	jwkOptions := keyfunc.Options{
		Client: client.HTTPClient,
		Ctx:    context.Background(),
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
	jwkURL := string(client.Config.GetBaseURI()) +
		fmt.Sprintf("/sessions/jwks/%s", client.Config.BasicAuthProjectID())
	return keyfunc.Get(jwkURL, jwkOptions)
}
