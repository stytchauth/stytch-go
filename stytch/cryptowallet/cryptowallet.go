package cryptowallet

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

func (c *Client) AuthenticateStart(
	ctx context.Context,
	body *stytch.CryptoWalletAuthenticateStartParams,
) (*stytch.CryptoWalletAuthenticateStartResponse, error) {
	path := "/crypto_wallets/authenticate/start"

	var jsonBody []byte
	var err error
	if body != nil {
		jsonBody, err = json.Marshal(body)
		if err != nil {
			return nil, stytcherror.NewClientLibraryError("Oops, something seems to have gone wrong " +
				"marshalling the authenticate/start request body")
		}
	}

	var retVal stytch.CryptoWalletAuthenticateStartResponse
	err = c.C.NewRequest(ctx, "POST", path, nil, jsonBody, &retVal)
	return &retVal, err
}

func (c *Client) Authenticate(
	ctx context.Context,
	body *stytch.CryptoWalletAuthenticateParams,
) (*stytch.CryptoWalletAuthenticateResponse, error) {
	path := "/crypto_wallets/authenticate"

	var jsonBody []byte
	var err error
	if body != nil {
		jsonBody, err = json.Marshal(body)
		if err != nil {
			return nil, stytcherror.NewClientLibraryError("Oops, something seems to have gone wrong " +
				"marshalling the authenticate request body")
		}
	}

	var retVal stytch.CryptoWalletAuthenticateResponse
	err = c.C.NewRequest(ctx, "POST", path, nil, jsonBody, &retVal)
	return &retVal, err
}

// AuthenticateWithClaims fills in the claims pointer with custom claims from the response.
// Pass in a map with the types of values you're expecting so that this function can marshal
// the claims from the response. See ExampleClient_AuthenticateWithClaims_map,
// ExampleClient_AuthenticateWithClaims_struct for examples
func (c *Client) AuthenticateWithClaims(
	ctx context.Context,
	body *stytch.CryptoWalletAuthenticateParams,
	claims interface{},
) (*stytch.CryptoWalletAuthenticateResponse, error) {
	path := "/crypto_wallets/authenticate"

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
	var retVal stytch.CryptoWalletAuthenticateResponse
	if err := json.Unmarshal(b, &retVal); err != nil {
		return nil, fmt.Errorf("unmarshal CryptoWalletAuthenticateResponse: %w", err)
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
