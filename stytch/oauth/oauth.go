package oauth

import (
	"encoding/json"
	"fmt"

	"github.com/stytchauth/stytch-go/v5/stytch"
	"github.com/stytchauth/stytch-go/v5/stytch/magiclink/email"
	"github.com/stytchauth/stytch-go/v5/stytch/stytcherror"
)

type Client struct {
	C     *stytch.Client
	Email *email.Client
}

func (c *Client) Authenticate(
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
	err = c.C.NewRequest("POST", path, nil, jsonBody, &retVal)
	return &retVal, err
}

// AuthenticateWithClaims fills in the claims pointer with custom claims from the response.
// Pass in a map with the types of values you're expecting so that this function can marshal
// the claims from the response. See session_test.go for an example
func (c *Client) AuthenticateWithClaims(
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

	b, err := c.C.RawRequest("POST", path, nil, jsonBody)
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
