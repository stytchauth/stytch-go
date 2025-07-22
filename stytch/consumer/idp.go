package consumer

import (
	"github.com/MicahParks/keyfunc/v2"
	"github.com/stytchauth/stytch-go/v16/stytch"
)

type IDPClient struct {
	C           stytch.Client
	JWKS        *keyfunc.JWKS
	PolicyCache any
}

func NewIDPClient(c stytch.Client, jwks *keyfunc.JWKS, policyCache any) *IDPClient {
	return &IDPClient{
		C:           c,
		JWKS:        jwks,
		PolicyCache: policyCache,
	}
}
