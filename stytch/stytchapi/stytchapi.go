package stytchapi

import (
	"context"
	"fmt"
	"time"

	"github.com/MicahParks/keyfunc"
	"github.com/stytchauth/stytch-go/v5/stytch"
	"github.com/stytchauth/stytch-go/v5/stytch/config"
	"github.com/stytchauth/stytch-go/v5/stytch/cryptowallet"
	"github.com/stytchauth/stytch-go/v5/stytch/magiclink"
	mle "github.com/stytchauth/stytch-go/v5/stytch/magiclink/email"
	"github.com/stytchauth/stytch-go/v5/stytch/oauth"
	"github.com/stytchauth/stytch-go/v5/stytch/otp"
	otpe "github.com/stytchauth/stytch-go/v5/stytch/otp/email"
	"github.com/stytchauth/stytch-go/v5/stytch/otp/sms"
	"github.com/stytchauth/stytch-go/v5/stytch/otp/whatsapp"
	"github.com/stytchauth/stytch-go/v5/stytch/session"
	"github.com/stytchauth/stytch-go/v5/stytch/totp"
	"github.com/stytchauth/stytch-go/v5/stytch/user"
	"github.com/stytchauth/stytch-go/v5/stytch/webauthn"
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
	Sessions      *session.Client
	TOTPs         *totp.Client
	Users         *user.Client
	WebAuthn      *webauthn.Client
}

type Option func(*API)

func WithLogger(logger Logger) Option {
	return func(api *API) { api.logger = logger }
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
