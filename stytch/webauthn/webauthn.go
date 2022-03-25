package webauthn

import (
	"encoding/json"

	"github.com/stytchauth/stytch-go/v4/stytch"
	"github.com/stytchauth/stytch-go/v4/stytch/stytcherror"
)

type Client struct {
	C *stytch.Client
}

func (c *Client) RegisterStart(body *stytch.WebAuthnRegisterStartParams,
) (*stytch.WebAuthnRegisterStartResponse, error) {
	path := "/webauthn/register/start"

	var jsonBody []byte
	var err error
	if body != nil {
		jsonBody, err = json.Marshal(body)
		if err != nil {
			return nil, stytcherror.NewClientLibraryError("Oops, something seems to have gone wrong " +
				"marshalling the register/start request body")
		}
	}

	var retVal stytch.WebAuthnRegisterStartResponse
	err = c.C.NewRequest("POST", path, nil, jsonBody, &retVal)
	return &retVal, err
}

func (c *Client) Register(body *stytch.WebAuthnRegisterParams,
) (*stytch.WebAuthnRegisterResponse, error) {
	path := "/webauthn/register"

	var jsonBody []byte
	var err error
	if body != nil {
		jsonBody, err = json.Marshal(body)
		if err != nil {
			return nil, stytcherror.NewClientLibraryError("Oops, something seems to have gone wrong " +
				"marshalling the register request body")
		}
	}

	var retVal stytch.WebAuthnRegisterResponse
	err = c.C.NewRequest("POST", path, nil, jsonBody, &retVal)
	return &retVal, err
}

func (c *Client) AuthenticateStart(body *stytch.WebAuthnAuthenticateStartParams,
) (*stytch.WebAuthnAuthenticateStartResponse, error) {
	path := "/webauthn/authenticate/start"

	var jsonBody []byte
	var err error
	if body != nil {
		jsonBody, err = json.Marshal(body)
		if err != nil {
			return nil, stytcherror.NewClientLibraryError("Oops, something seems to have gone wrong " +
				"marshalling the authenticate/start request body")
		}
	}

	var retVal stytch.WebAuthnAuthenticateStartResponse
	err = c.C.NewRequest("POST", path, nil, jsonBody, &retVal)
	return &retVal, err
}

func (c *Client) Authenticate(body *stytch.WebAuthnAuthenticateParams,
) (*stytch.WebAuthnAuthenticateResponse, error) {
	path := "/webauthn/authenticate"

	var jsonBody []byte
	var err error
	if body != nil {
		jsonBody, err = json.Marshal(body)
		if err != nil {
			return nil, stytcherror.NewClientLibraryError("Oops, something seems to have gone wrong " +
				"marshalling the authenticate request body")
		}
	}

	var retVal stytch.WebAuthnAuthenticateResponse
	err = c.C.NewRequest("POST", path, nil, jsonBody, &retVal)
	return &retVal, err
}
