package magiclink

import (
	"context"
	"encoding/json"

	"github.com/stytchauth/stytch-go/v7/stytch/b2b"
	"github.com/stytchauth/stytch-go/v7/stytch/b2b/magiclink/email"

	"github.com/stytchauth/stytch-go/v7/stytch"
	"github.com/stytchauth/stytch-go/v7/stytch/stytcherror"
)

type Client struct {
	C     *stytch.Client
	Email *email.Client
}

func (c *Client) Authenticate(
	ctx context.Context,
	body *b2b.MagicLinksAuthenticateParams,
) (*b2b.MagicLinksAuthenticateResponse, error) {
	path := "/b2b/magic_links/authenticate"

	var jsonBody []byte
	var err error
	if body != nil {
		jsonBody, err = json.Marshal(body)
		if err != nil {
			return nil, stytcherror.NewClientLibraryError(
				"Oops, something seems to have gone wrong " +
					"marshalling the /b2b/magic_links/authenticate request body")
		}
	}

	var retVal b2b.MagicLinksAuthenticateResponse
	err = c.C.NewRequest(ctx, "POST", path, nil, jsonBody, &retVal)
	return &retVal, err
}
