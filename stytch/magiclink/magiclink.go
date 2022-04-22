package magiclink

import (
	"encoding/json"

	"github.com/stytchauth/stytch-go/v5/stytch"
	"github.com/stytchauth/stytch-go/v5/stytch/magiclink/email"
	"github.com/stytchauth/stytch-go/v5/stytch/stytcherror"
)

type Client struct {
	C     *stytch.Client
	Email *email.Client
}

func (c *Client) Create(
	body *stytch.MagicLinksCreateParams,
) (*stytch.MagicLinksCreateResponse, error) {
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

	var retVal stytch.MagicLinksCreateResponse
	err = c.C.NewRequest("POST", path, nil, jsonBody, &retVal)
	return &retVal, err
}

func (c *Client) Authenticate(
	body *stytch.MagicLinksAuthenticateParams,
) (*stytch.MagicLinksAuthenticateResponse, error) {
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

	var retVal stytch.MagicLinksAuthenticateResponse
	err = c.C.NewRequest("POST", path, nil, jsonBody, &retVal)
	return &retVal, err
}
