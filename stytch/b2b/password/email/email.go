package email

import (
	"context"
	"encoding/json"

	"github.com/stytchauth/stytch-go/v8/stytch/b2b"

	"github.com/stytchauth/stytch-go/v8/stytch"
	"github.com/stytchauth/stytch-go/v8/stytch/stytcherror"
)

type Client struct {
	C *stytch.Client
}

func (c *Client) ResetStart(
	ctx context.Context,
	body *b2b.PasswordEmailResetStartParams,
) (*b2b.PasswordEmailResetStartResponse, error) {
	path := "/b2b/passwords/email/reset/start"

	var jsonBody []byte
	var err error
	if body != nil {
		jsonBody, err = json.Marshal(body)
		if err != nil {
			return nil, stytcherror.NewClientLibraryError(
				"Oops, something seems to have gone wrong " +
					"marshalling the /b2b/passwords/email/reset/start request body")
		}
	}

	var retVal b2b.PasswordEmailResetStartResponse
	err = c.C.NewRequest(ctx, "POST", path, nil, jsonBody, &retVal)
	return &retVal, err
}

func (c *Client) Reset(
	ctx context.Context,
	body *b2b.PasswordEmailResetParams,
) (*b2b.PasswordEmailResetResponse, error) {
	path := "/b2b/passwords/email/reset"

	var jsonBody []byte
	var err error
	if body != nil {
		jsonBody, err = json.Marshal(body)
		if err != nil {
			return nil, stytcherror.NewClientLibraryError(
				"Oops, something seems to have gone wrong " +
					"marshalling the /b2b/passwords/email/reset request body")
		}
	}

	var retVal b2b.PasswordEmailResetResponse
	err = c.C.NewRequest(ctx, "POST", path, nil, jsonBody, &retVal)
	return &retVal, err
}
