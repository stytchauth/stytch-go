package email

import (
	"context"
	"encoding/json"
	"github.com/stytchauth/stytch-go/v7/stytch"
	"github.com/stytchauth/stytch-go/v7/stytch/b2b"
	"github.com/stytchauth/stytch-go/v7/stytch/stytcherror"
)

type Client struct {
	C *stytch.Client
}

func (c *Client) LoginOrSignup(
	ctx context.Context,
	body *b2b.MagicLinksEmailLoginOrSignupParams,
) (
	*b2b.MagicLinksEmailLoginOrSignupResponse, error,
) {
	path := "/b2b/magic_links/email/login_or_signup"

	var jsonBody []byte
	var err error
	if body != nil {
		jsonBody, err = json.Marshal(body)
		if err != nil {
			return nil, stytcherror.NewClientLibraryError("Oops, something seems to have gone wrong " +
				"marshalling the /b2b/magic_links/email/login_or_signup request body")
		}
	}

	var retVal b2b.MagicLinksEmailLoginOrSignupResponse
	err = c.C.NewRequest(ctx, "POST", path, nil, jsonBody, &retVal)
	return &retVal, err
}
