package existingpassword

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

func (c *Client) Reset(
	ctx context.Context,
	body *b2b.PasswordExistingPasswordResetParams,
) (*b2b.PasswordExistingPasswordResetResponse, error) {
	path := "/b2b/passwords/existing_password/reset"

	var jsonBody []byte
	var err error
	if body != nil {
		jsonBody, err = json.Marshal(body)
		if err != nil {
			return nil, stytcherror.NewClientLibraryError(
				"Oops, something seems to have gone wrong " +
					"marshalling the /b2b/passwords/existing_password/reset request body")
		}
	}

	var retVal b2b.PasswordExistingPasswordResetResponse
	err = c.C.NewRequest(ctx, "POST", path, nil, jsonBody, &retVal)
	return &retVal, err
}
