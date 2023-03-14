package oauth

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/stytchauth/stytch-go/v6/stytch"
	"github.com/stytchauth/stytch-go/v6/stytch/magiclink/email"
	"github.com/stytchauth/stytch-go/v6/stytch/stytcherror"
)

type Client struct {
	C     *stytch.Client
	Email *email.Client
}

func (c *Client) Authenticate(
	ctx context.Context,
	body *stytch.OAuthAuthenticateParams,
) (*stytch.OAuthAuthenticateResponse, error) {
	path := "/oauth/authenticate"

	var jsonBody []byte
	var err error
	if body != nil {
		jsonBody, err = json.Marshal(body)
		if err != nil {
			return nil, stytcherror.NewClientLibraryError(
				"Oops, something seems to have gone wrong " +
					"marshalling the /oauth/authenticate request body")
		}
	}

	var retVal stytch.OAuthAuthenticateResponse
	err = c.C.NewRequest(ctx, "POST", path, nil, jsonBody, &retVal)
	return &retVal, err
}

// AuthenticateWithClaims fills in the claims pointer with custom claims from the response.
// Pass in a map with the types of values you're expecting so that this function can marshal
// the claims from the response. See ExampleClient_AuthenticateWithClaims_map,
// ExampleClient_AuthenticateWithClaims_struct for examples
func (c *Client) AuthenticateWithClaims(
	ctx context.Context,
	body *stytch.OAuthAuthenticateParams,
	claims interface{},
) (*stytch.OAuthAuthenticateResponse, error) {
	path := "/oauth/authenticate"

	var jsonBody []byte
	var err error
	if body != nil {
		jsonBody, err = json.Marshal(body)
		if err != nil {
			return nil, stytcherror.NewClientLibraryError(
				"Oops, something seems to have gone wrong " +
					"marshalling the /oauth/authenticate request body")
		}
	}

	b, err := c.C.RawRequest(ctx, "POST", path, nil, jsonBody)
	if err != nil {
		return nil, err
	}

	// First extract the Stytch data.
	var retVal stytch.OAuthAuthenticateResponse
	if err := json.Unmarshal(b, &retVal); err != nil {
		return nil, fmt.Errorf("unmarshal OAuthAuthenticateResponse: %w", err)
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

func (c *Client) Attach(
	ctx context.Context,
	body *stytch.OAuthAttachParams,
) (*stytch.OAuthAttachResponse, error) {
	path := "/oauth/attach"

	var jsonBody []byte
	var err error
	if body != nil {
		jsonBody, err = json.Marshal(body)
		if err != nil {
			return nil, stytcherror.NewClientLibraryError(
				"Oops, something seems to have gone wrong " +
					"marshalling the /oauth/attach request body")
		}
	}

	var retVal stytch.OAuthAttachResponse
	err = c.C.NewRequest(ctx, "POST", path, nil, jsonBody, &retVal)
	return &retVal, err
}
