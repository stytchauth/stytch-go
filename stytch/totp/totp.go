package totp

import (
	"encoding/json"

	"github.com/stytchauth/stytch-go/v3/stytch"
	"github.com/stytchauth/stytch-go/v3/stytch/stytcherror"
)

type Client struct {
	C *stytch.Client
}

func (c *Client) Create(body *stytch.TOTPsCreateParams) (
	*stytch.TOTPsCreateResponse, error) {
	path := "/totps"

	var jsonBody []byte
	var err error
	if body != nil {
		jsonBody, err = json.Marshal(body)
		if err != nil {
			return nil, stytcherror.NewClientLibraryError(
				"Oops, something seems to have gone wrong " +
					"marshalling the /totps request body")
		}
	}

	var retVal stytch.TOTPsCreateResponse
	err = c.C.NewRequest("POST", path, nil, jsonBody, &retVal)
	return &retVal, err
}

func (c *Client) Authenticate(body *stytch.TOTPsAuthenticateParams) (
	*stytch.TOTPsAuthenticateResponse, error) {
	path := "/totps/authenticate"

	var jsonBody []byte
	var err error
	if body != nil {
		jsonBody, err = json.Marshal(body)
		if err != nil {
			return nil, stytcherror.NewClientLibraryError(
				"Oops, something seems to have gone wrong " +
					"marshalling the /totps/authenticate request body")
		}
	}

	var retVal stytch.TOTPsAuthenticateResponse
	err = c.C.NewRequest("POST", path, nil, jsonBody, &retVal)
	return &retVal, err
}

func (c *Client) RecoveryCodes(body *stytch.TOTPsRecoveryCodesParams) (
	*stytch.TOTPsRecoveryCodesResponse, error) {
	path := "/totps/recovery_codes"

	var jsonBody []byte
	var err error
	if body != nil {
		jsonBody, err = json.Marshal(body)
		if err != nil {
			return nil, stytcherror.NewClientLibraryError(
				"Oops, something seems to have gone wrong " +
					"marshalling the /totps/recovery_codes request body")
		}
	}

	var retVal stytch.TOTPsRecoveryCodesResponse
	err = c.C.NewRequest("POST", path, nil, jsonBody, &retVal)
	return &retVal, err
}

func (c *Client) Recover(body *stytch.TOTPsRecoverParams) (
	*stytch.TOTPsRecoverResponse, error) {
	path := "/totps/recover"

	var jsonBody []byte
	var err error
	if body != nil {
		jsonBody, err = json.Marshal(body)
		if err != nil {
			return nil, stytcherror.NewClientLibraryError(
				"Oops, something seems to have gone wrong " +
					"marshalling the /totps/recover request body")
		}
	}

	var retVal stytch.TOTPsRecoverResponse
	err = c.C.NewRequest("POST", path, nil, jsonBody, &retVal)
	return &retVal, err
}
