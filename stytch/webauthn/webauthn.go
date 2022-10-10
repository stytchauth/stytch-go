package webauthn

import (
	"encoding/json"
	"fmt"

	"github.com/stytchauth/stytch-go/v6/stytch"
	"github.com/stytchauth/stytch-go/v6/stytch/stytcherror"
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

// AuthenticateWithClaims fills in the claims pointer with custom claims from the response.
// Pass in a map with the types of values you're expecting so that this function can marshal
// the claims from the response. See ExampleClient_AuthenticateWithClaims_map,
// ExampleClient_AuthenticateWithClaims_struct for examples
func (c *Client) AuthenticateWithClaims(
	body *stytch.WebAuthnAuthenticateParams,
	claims interface{},
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

	b, err := c.C.RawRequest("POST", path, nil, jsonBody)
	if err != nil {
		return nil, err
	}

	// First extract the Stytch data.
	var retVal stytch.WebAuthnAuthenticateResponse
	if err := json.Unmarshal(b, &retVal); err != nil {
		return nil, fmt.Errorf("unmarshal WebAuthnAuthenticateResponse: %w", err)
	}

	// Then extract the custom claims. Build a claims wrapper using the caller's `claims` value so
	// the unmarshal fills it.
	wrapper := stytch.SessionWrapper{
		Session: stytch.ClaimsWrapper{
			Claims: claims,
		},
	}
	if err := json.Unmarshal(b, &wrapper); err != nil {
		return nil, fmt.Errorf("unmarshal custom claims: %w", err)
	}
	retVal.Session.CustomClaims = wrapper.Session.Claims
	return &retVal, err
}
