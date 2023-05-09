package email

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/stytchauth/stytch-go/v9/stytch/b2c"

	"github.com/stytchauth/stytch-go/v9/stytch"
	"github.com/stytchauth/stytch-go/v9/stytch/stytcherror"
)

type Client struct {
	C *stytch.Client
}

func (c *Client) ResetStart(
	ctx context.Context,
	body *b2c.PasswordEmailResetStartParams,
) (*b2c.PasswordEmailResetStartResponse, error) {
	path := "/passwords/email/reset/start"

	var jsonBody []byte
	var err error
	if body != nil {
		jsonBody, err = json.Marshal(body)
		if err != nil {
			return nil, stytcherror.NewClientLibraryError(
				"Oops, something seems to have gone wrong " +
					"marshalling the /passwords/email/reset/start request body")
		}
	}

	var retVal b2c.PasswordEmailResetStartResponse
	err = c.C.NewRequest(ctx, "POST", path, nil, jsonBody, &retVal)
	return &retVal, err
}

func (c *Client) Reset(
	ctx context.Context,
	body *b2c.PasswordEmailResetParams,
) (*b2c.PasswordEmailResetResponse, error) {
	path := "/passwords/email/reset"

	var jsonBody []byte
	var err error
	if body != nil {
		jsonBody, err = json.Marshal(body)
		if err != nil {
			return nil, stytcherror.NewClientLibraryError(
				"Oops, something seems to have gone wrong " +
					"marshalling the /passwords/email/reset request body")
		}
	}

	var retVal b2c.PasswordEmailResetResponse
	err = c.C.NewRequest(ctx, "POST", path, nil, jsonBody, &retVal)
	return &retVal, err
}

// ResetWithClaims fills in the claims pointer with custom claims from the response.
// Pass in a map with the types of values you're expecting so that this function can marshal
// the claims from the response. See ExampleClient_AuthenticateWithClaims_map,
// ExampleClient_AuthenticateWithClaims_struct for examples
func (c *Client) ResetWithClaims(
	ctx context.Context,
	body *b2c.PasswordEmailResetParams,
	claims interface{},
) (*b2c.PasswordEmailResetResponse, error) {
	path := "/passwords/email/reset"

	var jsonBody []byte
	var err error
	if body != nil {
		jsonBody, err = json.Marshal(body)
		if err != nil {
			return nil, stytcherror.NewClientLibraryError(
				"Oops, something seems to have gone wrong " +
					"marshalling the /passwords/email/reset request body")
		}
	}

	b, err := c.C.RawRequest(ctx, "POST", path, nil, jsonBody)
	if err != nil {
		return nil, err
	}

	// First extract the Stytch data.
	var retVal b2c.PasswordEmailResetResponse
	if err := json.Unmarshal(b, &retVal); err != nil {
		return nil, fmt.Errorf("unmarshal PasswordEmailResetResponse: %w", err)
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
