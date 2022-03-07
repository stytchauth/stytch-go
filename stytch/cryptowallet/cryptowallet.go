package cryptowallet

import (
	"encoding/json"

	"github.com/stytchauth/stytch-go/v3/stytch"
	"github.com/stytchauth/stytch-go/v3/stytch/stytcherror"
)

type Client struct {
	C *stytch.Client
}

func (c *Client) AuthenticateStart(body *stytch.CryptoWalletAuthenticateStartParams,
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
	err = c.C.NewRequest("POST", path, nil, jsonBody, &retVal)
	return &retVal, err
}

func (c *Client) Authenticate(body *stytch.CryptoWalletAuthenticateParams,
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
	err = c.C.NewRequest("POST", path, nil, jsonBody, &retVal)
	return &retVal, err
}
