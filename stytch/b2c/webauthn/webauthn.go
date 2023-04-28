package webauthn

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/stytchauth/stytch-go/v8/stytch/b2c"

	"github.com/stytchauth/stytch-go/v8/stytch"
	"github.com/stytchauth/stytch-go/v8/stytch/stytcherror"
)

type Client struct {
	C *stytch.Client
}

func (c *Client) RegisterStart(
	ctx context.Context,
	body *b2c.WebAuthnRegisterStartParams,
) (*b2c.WebAuthnRegisterStartResponse, error) {
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

	var retVal b2c.WebAuthnRegisterStartResponse
	err = c.C.NewRequest(ctx, "POST", path, nil, jsonBody, &retVal)
	return &retVal, err
}

func (c *Client) Register(
	ctx context.Context,
	body *b2c.WebAuthnRegisterParams,
) (*b2c.WebAuthnRegisterResponse, error) {
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

	var retVal b2c.WebAuthnRegisterResponse
	err = c.C.NewRequest(ctx, "POST", path, nil, jsonBody, &retVal)
	return &retVal, err
}

func (c *Client) AuthenticateStart(
	ctx context.Context,
	body *b2c.WebAuthnAuthenticateStartParams,
) (*b2c.WebAuthnAuthenticateStartResponse, error) {
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

	var retVal b2c.WebAuthnAuthenticateStartResponse
	err = c.C.NewRequest(ctx, "POST", path, nil, jsonBody, &retVal)
	return &retVal, err
}

func (c *Client) Authenticate(
	ctx context.Context,
	body *b2c.WebAuthnAuthenticateParams,
) (*b2c.WebAuthnAuthenticateResponse, error) {
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

	var retVal b2c.WebAuthnAuthenticateResponse
	err = c.C.NewRequest(ctx, "POST", path, nil, jsonBody, &retVal)
	return &retVal, err
}

// AuthenticateWithClaims fills in the claims pointer with custom claims from the response.
// Pass in a map with the types of values you're expecting so that this function can marshal
// the claims from the response. See ExampleClient_AuthenticateWithClaims_map,
// ExampleClient_AuthenticateWithClaims_struct for examples
func (c *Client) AuthenticateWithClaims(
	ctx context.Context,
	body *b2c.WebAuthnAuthenticateParams,
	claims interface{},
) (*b2c.WebAuthnAuthenticateResponse, error) {
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

	b, err := c.C.RawRequest(ctx, "POST", path, nil, jsonBody)
	if err != nil {
		return nil, err
	}

	// First extract the Stytch data.
	var retVal b2c.WebAuthnAuthenticateResponse
	if err := json.Unmarshal(b, &retVal); err != nil {
		return nil, fmt.Errorf("unmarshal WebAuthnAuthenticateResponse: %w", err)
	}

	// Then extract the custom claims. Build a claims wrapper using the caller's `claims` value so
	// the unmarshal fills it.
	wrapper := b2c.SessionWrapper{
		Session: b2c.ClaimsWrapper{
			Claims: claims,
		},
	}
	if err := json.Unmarshal(b, &wrapper); err != nil {
		return nil, fmt.Errorf("unmarshal custom claims: %w", err)
	}
	retVal.Session.CustomClaims = wrapper.Session.Claims
	return &retVal, err
}
