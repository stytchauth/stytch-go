package otp

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/stytchauth/stytch-go/v7/stytch/b2c"
	"github.com/stytchauth/stytch-go/v7/stytch/b2c/otp/email"
	"github.com/stytchauth/stytch-go/v7/stytch/b2c/otp/sms"
	"github.com/stytchauth/stytch-go/v7/stytch/b2c/otp/whatsapp"

	"github.com/stytchauth/stytch-go/v7/stytch"
	"github.com/stytchauth/stytch-go/v7/stytch/stytcherror"
)

type Client struct {
	C        *stytch.Client
	SMS      *sms.Client
	WhatsApp *whatsapp.Client
	Email    *email.Client
}

func (c *Client) Authenticate(
	ctx context.Context,
	body *b2c.OTPsAuthenticateParams,
) (*b2c.OTPsAuthenticateResponse, error) {
	path := "/otps/authenticate"

	var jsonBody []byte
	var err error
	if body != nil {
		jsonBody, err = json.Marshal(body)
		if err != nil {
			return nil, stytcherror.NewClientLibraryError("Oops, something seems to have gone wrong " +
				"marshalling the /otps/authenticate request body")
		}
	}

	var retVal b2c.OTPsAuthenticateResponse
	err = c.C.NewRequest(ctx, "POST", path, nil, jsonBody, &retVal)
	return &retVal, err
}

// AuthenticateWithClaims fills in the claims pointer with custom claims from the response.
// Pass in a map with the types of values you're expecting so that this function can marshal
// the claims from the response. See ExampleClient_AuthenticateWithClaims_map,
// ExampleClient_AuthenticateWithClaims_struct for examples
func (c *Client) AuthenticateWithClaims(
	ctx context.Context,
	body *b2c.OTPsAuthenticateParams,
	claims interface{},
) (*b2c.OTPsAuthenticateResponse, error) {
	path := "/otps/authenticate"

	var jsonBody []byte
	var err error
	if body != nil {
		jsonBody, err = json.Marshal(body)
		if err != nil {
			return nil, stytcherror.NewClientLibraryError("Oops, something seems to have gone wrong " +
				"marshalling the /otps/authenticate request body")
		}
	}

	b, err := c.C.RawRequest(ctx, "POST", path, nil, jsonBody)
	if err != nil {
		return nil, err
	}

	// First extract the Stytch data.
	var retVal b2c.OTPsAuthenticateResponse
	if err := json.Unmarshal(b, &retVal); err != nil {
		return nil, fmt.Errorf("unmarshal OTPsAuthenticateResponse: %w", err)
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
