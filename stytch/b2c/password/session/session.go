package session

import (
	"context"
	"encoding/json"
	"github.com/stytchauth/stytch-go/v7/stytch/b2c"

	"github.com/stytchauth/stytch-go/v7/stytch"
	"github.com/stytchauth/stytch-go/v7/stytch/stytcherror"
)

type Client struct {
	C *stytch.Client
}

func (c *Client) Reset(
	ctx context.Context,
	body *b2c.PasswordSessionResetParams,
) (*b2c.PasswordSessionResetResponse, error) {
	path := "/passwords/session/reset"

	var jsonBody []byte
	var err error
	if body != nil {
		jsonBody, err = json.Marshal(body)
		if err != nil {
			return nil, stytcherror.NewClientLibraryError(
				"Oops, something seems to have gone wrong " +
					"marshalling the /passwords/session/reset request body")
		}
	}

	var retVal b2c.PasswordSessionResetResponse
	err = c.C.NewRequest(ctx, "POST", path, nil, jsonBody, &retVal)
	return &retVal, err
}
