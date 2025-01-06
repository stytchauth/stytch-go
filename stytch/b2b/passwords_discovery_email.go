package b2b

// !!!
// WARNING: This file is autogenerated
// Only modify code within MANUAL() sections
// or your changes may be overwritten later!
// !!!

import (
	"context"
	"encoding/json"

	"github.com/stytchauth/stytch-go/v16/stytch"
	"github.com/stytchauth/stytch-go/v16/stytch/b2b/passwords/discovery/email"
	"github.com/stytchauth/stytch-go/v16/stytch/stytcherror"
)

type PasswordsDiscoveryEmailClient struct {
	C stytch.Client
}

func NewPasswordsDiscoveryEmailClient(c stytch.Client) *PasswordsDiscoveryEmailClient {
	return &PasswordsDiscoveryEmailClient{
		C: c,
	}
}

// ResetStart: Initiates a password reset for the email address provided, when cross-org passwords are
// enabled. This will trigger an email to be sent to the address, containing a magic link that will allow
// them to set a new password and authenticate.
//
// This endpoint adapts to your Project's password strength configuration.
// If you're using [zxcvbn](https://stytch.com/docs/guides/passwords/strength-policy), the default, your
// passwords are considered valid
// if the strength score is >= 3. If you're using
// [LUDS](https://stytch.com/docs/guides/passwords/strength-policy), your passwords are
// considered valid if they meet the requirements that you've set with Stytch.
// You may update your password strength configuration in the
// [stytch dashboard](https://stytch.com/dashboard/password-strength-config).
func (c *PasswordsDiscoveryEmailClient) ResetStart(
	ctx context.Context,
	body *email.ResetStartParams,
) (*email.ResetStartResponse, error) {
	var jsonBody []byte
	var err error
	if body != nil {
		jsonBody, err = json.Marshal(body)
		if err != nil {
			return nil, stytcherror.NewClientLibraryError("error marshaling request body")
		}
	}

	headers := make(map[string][]string)

	var retVal email.ResetStartResponse
	err = c.C.NewRequest(
		ctx,
		stytch.RequestParams{
			Method:      "POST",
			Path:        "/v1/b2b/passwords/discovery/email/reset/start",
			QueryParams: nil,
			Body:        jsonBody,
			V:           &retVal,
			Headers:     headers,
		},
	)
	return &retVal, err
}

// Reset the password associated with an email and start an intermediate session. This endpoint checks that
// the password reset token is valid, hasn’t expired, or already been used.
//
// The provided password needs to meet the project's password strength requirements, which can be checked
// in advance with the password strength endpoint. If the token and password are accepted, the password is
// securely stored for future authentication and the user is authenticated.
//
// Resetting a password will start an intermediate session and return a list of discovered organizations
// the session can be exchanged into.
func (c *PasswordsDiscoveryEmailClient) Reset(
	ctx context.Context,
	body *email.ResetParams,
) (*email.ResetResponse, error) {
	var jsonBody []byte
	var err error
	if body != nil {
		jsonBody, err = json.Marshal(body)
		if err != nil {
			return nil, stytcherror.NewClientLibraryError("error marshaling request body")
		}
	}

	headers := make(map[string][]string)

	var retVal email.ResetResponse
	err = c.C.NewRequest(
		ctx,
		stytch.RequestParams{
			Method:      "POST",
			Path:        "/v1/b2b/passwords/discovery/email/reset",
			QueryParams: nil,
			Body:        jsonBody,
			V:           &retVal,
			Headers:     headers,
		},
	)
	return &retVal, err
}
