package magiclink

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/stytchauth/stytch-go/v9/stytch/b2c"
	"github.com/stytchauth/stytch-go/v9/stytch/b2c/magiclink/email"

	"github.com/stytchauth/stytch-go/v9/stytch"
	"github.com/stytchauth/stytch-go/v9/stytch/stytcherror"
)

type Client struct {
	C     *stytch.Client
	Email *email.Client
}

func (c *Client) Create(
	ctx context.Context,
	body *b2c.MagicLinksCreateParams,
) (*b2c.MagicLinksCreateResponse, error) {
	path := "/magic_links"

	var jsonBody []byte
	var err error
	if body != nil {
		jsonBody, err = json.Marshal(body)
		if err != nil {
			return nil, stytcherror.NewClientLibraryError(
				"Oops, something seems to have gone wrong " +
					"marshalling the /magic_links request body")
		}
	}

	var retVal b2c.MagicLinksCreateResponse
	err = c.C.NewRequest(ctx, "POST", path, nil, jsonBody, &retVal)
	return &retVal, err
}

func (c *Client) Authenticate(
	ctx context.Context,
	body *b2c.MagicLinksAuthenticateParams,
) (*b2c.MagicLinksAuthenticateResponse, error) {
	path := "/magic_links/authenticate"

	var jsonBody []byte
	var err error
	if body != nil {
		jsonBody, err = json.Marshal(body)
		if err != nil {
			return nil, stytcherror.NewClientLibraryError(
				"Oops, something seems to have gone wrong " +
					"marshalling the /magic_links/authenticate request body")
		}
	}

	var retVal b2c.MagicLinksAuthenticateResponse
	err = c.C.NewRequest(ctx, "POST", path, nil, jsonBody, &retVal)
	return &retVal, err
}

// AuthenticateWithClaims fills in the claims pointer with custom claims from the response.
// Pass in a map with the types of values you're expecting so that this function can marshal
// the claims from the response. See ExampleClient_AuthenticateWithClaims_map,
// ExampleClient_AuthenticateWithClaims_struct for examples
func (c *Client) AuthenticateWithClaims(
	ctx context.Context,
	body *b2c.MagicLinksAuthenticateParams,
	claims interface{},
) (*b2c.MagicLinksAuthenticateResponse, error) {
	path := "/magic_links/authenticate"

	var jsonBody []byte
	var err error
	if body != nil {
		jsonBody, err = json.Marshal(body)
		if err != nil {
			return nil, stytcherror.NewClientLibraryError(
				"Oops, something seems to have gone wrong " +
					"marshalling the /magic_links/authenticate request body")
		}
	}

	b, err := c.C.RawRequest(ctx, "POST", path, nil, jsonBody)
	if err != nil {
		return nil, err
	}

	// First extract the Stytch data.
	var retVal b2c.MagicLinksAuthenticateResponse
	if err := json.Unmarshal(b, &retVal); err != nil {
		return nil, fmt.Errorf("unmarshal MagicLinksAuthenticateResponse: %w", err)
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
