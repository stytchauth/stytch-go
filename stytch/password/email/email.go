package email

import (
	"encoding/json"
	"github.com/stytchauth/stytch-go/v5/stytch"
	"github.com/stytchauth/stytch-go/v5/stytch/stytcherror"
)

type Client struct {
	C *stytch.Client
}

func (c *Client) ResetStart(
	body *stytch.PasswordEmailResetStartParams,
) (*stytch.PasswordEmailResetStartResponse, error) {
	path := "/passwords/email/reset/start"

	var jsonBody []byte
	var err error
	if body != nil {
		jsonBody, err = json.Marshal(body)
		if err != nil {
			return nil, stytcherror.NewClientLibraryError(
				"Oops, something seems to have gone wrong " +
					"marshalling the /passwords/email/reset/start request body")
		}
	}

	var retVal stytch.PasswordEmailResetStartResponse
	err = c.C.NewRequest("POST", path, nil, jsonBody, &retVal)
	return &retVal, err
}

func (c *Client) Reset(
	body *stytch.PasswordEmailResetParams,
) (*stytch.PasswordEmailResetResponse, error) {
	path := "/passwords/email/reset"

	var jsonBody []byte
	var err error
	if body != nil {
		jsonBody, err = json.Marshal(body)
		if err != nil {
			return nil, stytcherror.NewClientLibraryError(
				"Oops, something seems to have gone wrong " +
					"marshalling the /passwords/email/reset request body")
		}
	}

	var retVal stytch.PasswordEmailResetResponse
	err = c.C.NewRequest("POST", path, nil, jsonBody, &retVal)
	return &retVal, err
}
