package consumer

// !!!
// WARNING: This file is autogenerated
// Only modify code within MANUAL() sections
// or your changes may be overwritten later!
// !!!

import (
	"context"
	"encoding/json"

	"github.com/stytchauth/stytch-go/v14/stytch"
	"github.com/stytchauth/stytch-go/v14/stytch/consumer/passwords/session"
	"github.com/stytchauth/stytch-go/v14/stytch/stytcherror"
)

type PasswordsSessionsClient struct {
	C stytch.Client
}

func NewPasswordsSessionsClient(c stytch.Client) *PasswordsSessionsClient {
	return &PasswordsSessionsClient{
		C: c,
	}
}

// Reset the user’s password using their existing session. The endpoint will error if the session does not
// have a password, email magic link, or email OTP authentication factor that has been issued within the
// last 5 minutes. This endpoint requires either a `session_jwt` or `session_token` be included in the
// request.
//
// Note that a successful password reset via an existing session will revoke all active sessions for the
// `user_id`, except for the one used during the reset flow.
func (c *PasswordsSessionsClient) Reset(
	ctx context.Context,
	body *session.ResetParams,
) (*session.ResetResponse, error) {
	var jsonBody []byte
	var err error
	if body != nil {
		jsonBody, err = json.Marshal(body)
		if err != nil {
			return nil, stytcherror.NewClientLibraryError("error marshaling request body")
		}
	}

	headers := make(map[string][]string)

	var retVal session.ResetResponse
	err = c.C.NewRequest(
		ctx,
		"POST",
		"/v1/passwords/session/reset",
		nil,
		jsonBody,
		&retVal,
		headers,
	)
	return &retVal, err
}
