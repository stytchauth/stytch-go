package b2b

// !!!
// WARNING: This file is autogenerated
// Only modify code within MANUAL() sections
// or your changes may be overwritten later!
// !!!

import (
	"context"
	"encoding/json"

	"github.com/stytchauth/stytch-go/v15/stytch"
	"github.com/stytchauth/stytch-go/v15/stytch/b2b/passwords/discovery"
	"github.com/stytchauth/stytch-go/v15/stytch/stytcherror"
)

type PasswordsDiscoveryClient struct {
	C     stytch.Client
	Email *PasswordsDiscoveryEmailClient
}

func NewPasswordsDiscoveryClient(c stytch.Client) *PasswordsDiscoveryClient {
	return &PasswordsDiscoveryClient{
		C: c,

		Email: NewPasswordsDiscoveryEmailClient(c),
	}
}

// Authenticate an email/password combination in the discovery flow. This authenticate flow is only valid
// for cross-org passwords use cases, and is not tied to a specific organization.
//
// If you have breach detection during authentication enabled in your
// [password strength policy](https://stytch.com/docs/b2b/guides/passwords/strength-policies) and the
// member's credentials have appeared in the HaveIBeenPwned dataset, this endpoint will return a
// `member_reset_password` error even if the member enters a correct password. We force a password reset in
// this case to ensure that the member is the legitimate owner of the email address and not a malicious
// actor abusing the compromised credentials.
//
// If successful, this endpoint will create a new intermediate session and return a list of discovered
// organizations that can be session exchanged into.
func (c *PasswordsDiscoveryClient) Authenticate(
	ctx context.Context,
	body *discovery.AuthenticateParams,
) (*discovery.AuthenticateResponse, error) {
	var jsonBody []byte
	var err error
	if body != nil {
		jsonBody, err = json.Marshal(body)
		if err != nil {
			return nil, stytcherror.NewClientLibraryError("error marshaling request body")
		}
	}

	headers := make(map[string][]string)

	var retVal discovery.AuthenticateResponse
	err = c.C.NewRequest(
		ctx,
		"POST",
		"/v1/b2b/passwords/discovery/authenticate",
		nil,
		jsonBody,
		&retVal,
		headers,
	)
	return &retVal, err
}
