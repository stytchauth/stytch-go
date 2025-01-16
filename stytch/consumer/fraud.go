package consumer

// !!!
// WARNING: This file is autogenerated
// Only modify code within MANUAL() sections
// or your changes may be overwritten later!
// !!!

import (
	"github.com/stytchauth/stytch-go/v16/stytch"
)

type FraudClient struct {
	C           stytch.Client
	Fingerprint *FraudFingerprintClient
	Rules       *FraudRulesClient
}

func NewFraudClient(c stytch.Client) *FraudClient {
	return &FraudClient{
		C: c,

		Fingerprint: NewFraudFingerprintClient(c),
		Rules:       NewFraudRulesClient(c),
	}
}