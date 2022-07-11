package password

import (
	"encoding/json"
	"github.com/stytchauth/stytch-go/v5/stytch"
	"github.com/stytchauth/stytch-go/v5/stytch/password/email"
	"github.com/stytchauth/stytch-go/v5/stytch/stytcherror"
)

type Client struct {
	C     *stytch.Client
	Email *email.Client
}

func (c *Client) Create(
	body *stytch.PasswordsCreateParams,
) (*stytch.PasswordsCreateResponse, error) {
	path := "/passwords"

	var jsonBody []byte
	var err error
	if body != nil {
		jsonBody, err = json.Marshal(body)
		if err != nil {
			return nil, stytcherror.NewClientLibraryError(
				"Oops, something seems to have gone wrong " +
					"marshalling the /passwords request body")
		}
	}

	var retVal stytch.PasswordsCreateResponse
	err = c.C.NewRequest("POST", path, nil, jsonBody, &retVal)
	return &retVal, err
}

func (c *Client) Authenticate(
	body *stytch.PasswordsAuthenticateParams,
) (*stytch.PasswordsAuthenticateResponse, error) {
	path := "/passwords/authenticate"

	var jsonBody []byte
	var err error
	if body != nil {
		jsonBody, err = json.Marshal(body)
		if err != nil {
			return nil, stytcherror.NewClientLibraryError(
				"Oops, something seems to have gone wrong " +
					"marshalling the /passwords/authenticate request body")
		}
	}

	var retVal stytch.PasswordsAuthenticateResponse
	err = c.C.NewRequest("POST", path, nil, jsonBody, &retVal)
	return &retVal, err
}

func (c *Client) StrengthCheck(
	body *stytch.PasswordsStrengthCheckParams,
) (*stytch.PasswordsStrengthCheckResponse, error) {
	path := "/passwords/strength_check"

	var jsonBody []byte
	var err error
	if body != nil {
		jsonBody, err = json.Marshal(body)
		if err != nil {
			return nil, stytcherror.NewClientLibraryError(
				"Oops, something seems to have gone wrong " +
					"marshalling the /passwords/strength_check request body")
		}
	}

	var retVal stytch.PasswordsStrengthCheckResponse
	err = c.C.NewRequest("POST", path, nil, jsonBody, &retVal)
	return &retVal, err
}

func (c *Client) Migreate(
	body *stytch.PasswordsMigrateParams,
) (*stytch.PasswordsMigrateResponse, error) {
	path := "/passwords/migrate"

	var jsonBody []byte
	var err error
	if body != nil {
		jsonBody, err = json.Marshal(body)
		if err != nil {
			return nil, stytcherror.NewClientLibraryError(
				"Oops, something seems to have gone wrong " +
					"marshalling the /passwords/migrate request body")
		}
	}

	var retVal stytch.PasswordsMigrateResponse
	err = c.C.NewRequest("POST", path, nil, jsonBody, &retVal)
	return &retVal, err
}
