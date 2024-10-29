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
	"github.com/stytchauth/stytch-go/v15/stytch/b2b/passwords/email"
	"github.com/stytchauth/stytch-go/v15/stytch/stytcherror"
)

type PasswordsEmailClient struct {
	C stytch.Client
}

func NewPasswordsEmailClient(c stytch.Client) *PasswordsEmailClient {
	return &PasswordsEmailClient{
		C: c,
	}
}

// ResetStart: Initiates a password reset for the email address provided. This will trigger an email to be
// sent to the address, containing a magic link that will allow them to set a new password and authenticate.
//
// This endpoint adapts to your Project's password strength configuration.
// If you're using [zxcvbn](https://stytch.com/docs/guides/passwords/strength-policy), the default, your
// passwords are considered valid
// if the strength score is >= 3. If you're using
// [LUDS](https://stytch.com/docs/guides/passwords/strength-policy), your passwords are
// considered valid if they meet the requirements that you've set with Stytch.
// You may update your password strength configuration in the
// [stytch dashboard](https://stytch.com/dashboard/password-strength-config).
func (c *PasswordsEmailClient) ResetStart(
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
		"POST",
		"/v1/b2b/passwords/email/reset/start",
		nil,
		jsonBody,
		&retVal,
		headers,
	)
	return &retVal, err
}

// Reset the's password and authenticate them. This endpoint checks that the password reset token is valid,
// hasn’t expired, or already been used.
//
// The provided password needs to meet our password strength requirements, which can be checked in advance
// with the password strength endpoint. If the token and password are accepted, the password is securely
// stored for future authentication and the user is authenticated.
//
// If the Member is required to complete MFA to log in to the Organization, the returned value of
// `member_authenticated` will be `false`, and an `intermediate_session_token` will be returned.
// The `intermediate_session_token` can be passed into the
// [OTP SMS Authenticate endpoint](https://stytch.com/docs/b2b/api/authenticate-otp-sms) to complete the
// MFA step and acquire a full member session.
// The `session_duration_minutes` and `session_custom_claims` parameters will be ignored.
//
// If a valid `session_token` or `session_jwt` is passed in, the Member will not be required to complete an
// MFA step.
//
// Note that a successful password reset by email will revoke all active sessions for the `member_id`.
func (c *PasswordsEmailClient) Reset(
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
		"POST",
		"/v1/b2b/passwords/email/reset",
		nil,
		jsonBody,
		&retVal,
		headers,
	)
	return &retVal, err
}

// RequireReset: Require a password be reset by the associated email address. This endpoint is only
// functional for cross-org password use cases.
func (c *PasswordsEmailClient) RequireReset(
	ctx context.Context,
	body *email.RequireResetParams,
	methodOptions ...*email.RequireResetRequestOptions,
) (*email.RequireResetResponse, error) {
	var jsonBody []byte
	var err error
	if body != nil {
		jsonBody, err = json.Marshal(body)
		if err != nil {
			return nil, stytcherror.NewClientLibraryError("error marshaling request body")
		}
	}

	headers := make(map[string][]string)
	for _, methodOption := range methodOptions {
		headers = methodOption.AddHeaders(headers)
	}

	var retVal email.RequireResetResponse
	err = c.C.NewRequest(
		ctx,
		"POST",
		"/v1/b2b/passwords/email/require_reset",
		nil,
		jsonBody,
		&retVal,
		headers,
	)
	return &retVal, err
}
