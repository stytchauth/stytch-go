package stytchapi

import (
	"github.com/stytchauth/stytch-go/stytch"
	"github.com/stytchauth/stytch-go/stytch/config"
	"github.com/stytchauth/stytch-go/stytch/magiclink"
	"github.com/stytchauth/stytch-go/stytch/magiclink/email"
	"github.com/stytchauth/stytch-go/stytch/otp"
	"github.com/stytchauth/stytch-go/stytch/otp/sms"
	"github.com/stytchauth/stytch-go/stytch/user"
)

type API struct {
	MagicLinks *magiclink.Client
	OTPs       *otp.Client
	Users      *user.Client
}

func NewAPIClient(env config.Env, projectID string, secret string) *API {
	a := &API{}
	client := stytch.New(env, projectID, secret)

	a.MagicLinks = &magiclink.Client{C: client, Email: &email.Client{C: client}}
	a.OTPs = &otp.Client{C: client, SMS: &sms.Client{C: client}}
	a.Users = &user.Client{C: client}
	return a
}
