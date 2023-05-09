package session

import (
	"context"
	"encoding/json"

	"github.com/stytchauth/stytch-go/v9/stytch/b2b"

	"github.com/stytchauth/stytch-go/v9/stytch"
	"github.com/stytchauth/stytch-go/v9/stytch/stytcherror"
)

type Client struct {
	C *stytch.Client
}

func (c *Client) Reset(
	ctx context.Context,
	body *b2b.PasswordSessionResetParams,
) (*b2b.PasswordSessionResetResponse, error) {
	path := "/b2b/passwords/session/reset"

	var jsonBody []byte
	var err error
	if body != nil {
		jsonBody, err = json.Marshal(body)
		if err != nil {
			return nil, stytcherror.NewClientLibraryError(
				"Oops, something seems to have gone wrong " +
					"marshalling the /b2b/passwords/session/reset request body")
		}
	}

	var retVal b2b.PasswordSessionResetResponse
	err = c.C.NewRequest(ctx, "POST", path, nil, jsonBody, &retVal)
	return &retVal, err
}
