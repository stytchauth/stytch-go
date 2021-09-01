package stytchapi

import (
	"github.com/stytchauth/stytch-go/v3/stytch"
	"github.com/stytchauth/stytch-go/v3/stytch/config"
	"github.com/stytchauth/stytch-go/v3/stytch/magiclink"
	mle "github.com/stytchauth/stytch-go/v3/stytch/magiclink/email"
	"github.com/stytchauth/stytch-go/v3/stytch/otp"
	otpe "github.com/stytchauth/stytch-go/v3/stytch/otp/email"
	"github.com/stytchauth/stytch-go/v3/stytch/otp/sms"
	"github.com/stytchauth/stytch-go/v3/stytch/otp/whatsapp"
	"github.com/stytchauth/stytch-go/v3/stytch/session"
	"github.com/stytchauth/stytch-go/v3/stytch/user"
)

type API struct {
	MagicLinks *magiclink.Client
	OTPs       *otp.Client
	Sessions   *session.Client
	Users      *user.Client
}

func NewAPIClient(env config.Env, projectID string, secret string) *API {
	a := &API{}
	client := stytch.New(env, projectID, secret)

	a.MagicLinks = &magiclink.Client{C: client, Email: &mle.Client{C: client}}
	a.OTPs = &otp.Client{
		C:        client,
		Email:    &otpe.Client{C: client},
		SMS:      &sms.Client{C: client},
		WhatsApp: &whatsapp.Client{C: client},
	}
	a.Sessions = &session.Client{C: client}
	a.Users = &user.Client{C: client}
	return a
}
