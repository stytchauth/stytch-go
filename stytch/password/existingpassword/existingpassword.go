package existingpassword

import (
	"encoding/json"
	"github.com/stytchauth/stytch-go/v5/stytch"
	"github.com/stytchauth/stytch-go/v5/stytch/stytcherror"
)

type Client struct {
	C *stytch.Client
}

func (c *Client) Reset(
	body *stytch.PasswordExistingPasswordResetParams,
) (*stytch.PasswordExistingPasswordResetResponse, error) {
	path := "/passwords/existing_password/reset"

	var jsonBody []byte
	var err error
	if body != nil {
		jsonBody, err = json.Marshal(body)
		if err != nil {
			return nil, stytcherror.NewClientLibraryError(
				"Oops, something seems to have gone wrong " +
					"marshalling the /passwords/existing_password/reset request body")
		}
	}

	var retVal stytch.PasswordExistingPasswordResetResponse
	err = c.C.NewRequest("POST", path, nil, jsonBody, &retVal)
	return &retVal, err
}
