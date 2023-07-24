package b2b

// !!!
// WARNING: This file is autogenerated
// Only modify code within MANUAL() sections
// or your changes may be overwritten later!
// !!!

import (
	"github.com/stytchauth/stytch-go/v9/stytch"
)

type OTPsClient struct {
	C   stytch.Client
	Sms *OTPsSmsClient
}

func NewOTPsClient(c stytch.Client) *OTPsClient {
	return &OTPsClient{
		C:   c,
		Sms: NewOTPsSmsClient(c),
	}
}
