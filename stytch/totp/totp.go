package totp

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/stytchauth/stytch-go/v6/stytch"
	"github.com/stytchauth/stytch-go/v6/stytch/stytcherror"
)

type Client struct {
	C *stytch.Client
}

func (c *Client) Create(ctx context.Context, body *stytch.TOTPsCreateParams) (
	*stytch.TOTPsCreateResponse, error,
) {
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
	err = c.C.NewRequest(ctx, "POST", path, nil, jsonBody, &retVal)
	return &retVal, err
}

func (c *Client) Authenticate(ctx context.Context, body *stytch.TOTPsAuthenticateParams) (
	*stytch.TOTPsAuthenticateResponse, error,
) {
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
	err = c.C.NewRequest(ctx, "POST", path, nil, jsonBody, &retVal)
	return &retVal, err
}

// AuthenticateWithClaims fills in the claims pointer with custom claims from the response.
// Pass in a map with the types of values you're expecting so that this function can marshal
// the claims from the response. See ExampleClient_AuthenticateWithClaims_map,
// ExampleClient_AuthenticateWithClaims_struct for examples
func (c *Client) AuthenticateWithClaims(
	ctx context.Context,
	body *stytch.TOTPsAuthenticateParams,
	claims interface{},
) (*stytch.TOTPsAuthenticateResponse, error) {
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

	b, err := c.C.RawRequest(ctx, "POST", path, nil, jsonBody)
	if err != nil {
		return nil, err
	}

	// First extract the Stytch data.
	var retVal stytch.TOTPsAuthenticateResponse
	if err := json.Unmarshal(b, &retVal); err != nil {
		return nil, fmt.Errorf("unmarshal TOTPsAuthenticateResponse: %w", err)
	}

	// Then extract the custom claims. Build a claims wrapper using the caller's `claims` value so
	// the unmarshal fills it.
	wrapper := stytch.SessionWrapper{
		Session: struct {
			Claims interface{} `json:"custom_claims"`
		}{
			Claims: claims,
		},
	}
	if err := json.Unmarshal(b, &wrapper); err != nil {
		return nil, fmt.Errorf("unmarshal custom claims: %w", err)
	}
	retVal.Session.CustomClaims = wrapper.Session.Claims
	return &retVal, err
}

func (c *Client) RecoveryCodes(ctx context.Context, body *stytch.TOTPsRecoveryCodesParams) (
	*stytch.TOTPsRecoveryCodesResponse, error,
) {
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
	err = c.C.NewRequest(ctx, "POST", path, nil, jsonBody, &retVal)
	return &retVal, err
}

func (c *Client) Recover(ctx context.Context, body *stytch.TOTPsRecoverParams) (
	*stytch.TOTPsRecoverResponse, error,
) {
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
	err = c.C.NewRequest(ctx, "POST", path, nil, jsonBody, &retVal)
	return &retVal, err
}

// RecoverWithClaims fills in the claims pointer with custom claims from the response.
// Pass in a map with the types of values you're expecting so that this function can marshal
// the claims from the response. See ExampleClient_AuthenticateWithClaims_map,
// ExampleClient_AuthenticateWithClaims_struct for examples
func (c *Client) RecoverWithClaims(
	ctx context.Context,
	body *stytch.TOTPsRecoverParams,
	claims interface{},
) (*stytch.TOTPsRecoverResponse, error) {
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

	b, err := c.C.RawRequest(ctx, "POST", path, nil, jsonBody)
	if err != nil {
		return nil, err
	}

	// First extract the Stytch data.
	var retVal stytch.TOTPsRecoverResponse
	if err := json.Unmarshal(b, &retVal); err != nil {
		return nil, fmt.Errorf("unmarshal TOTPsRecoverResponse: %w", err)
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
