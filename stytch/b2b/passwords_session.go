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
	"github.com/stytchauth/stytch-go/v16/stytch/b2b/passwords/session"
	"github.com/stytchauth/stytch-go/v16/stytch/stytcherror"
)

type PasswordsSessionsClient struct {
	C stytch.Client
}

func NewPasswordsSessionsClient(c stytch.Client) *PasswordsSessionsClient {
	return &PasswordsSessionsClient{
		C: c,
	}
}

// Reset the's password using their existing session. The endpoint will error if the session does not
// contain an authentication factor that has been issued within the last 5 minutes. Either `session_token`
// or `session_jwt` should be provided.
//
// Note that a successful password reset via an existing session will revoke all active sessions for the
// `member_id`, except for the one used during the reset flow.
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
		stytch.RequestParams{
			Method:      "POST",
			Path:        "/v1/b2b/passwords/session/reset",
			QueryParams: nil,
			Body:        jsonBody,
			V:           &retVal,
			Headers:     headers,
		},
	)
	return &retVal, err
}
