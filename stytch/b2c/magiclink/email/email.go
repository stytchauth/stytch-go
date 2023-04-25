package email

import (
	"context"
	"encoding/json"
	"github.com/stytchauth/stytch-go/v7/stytch/b2c"

	"github.com/stytchauth/stytch-go/v7/stytch"
	"github.com/stytchauth/stytch-go/v7/stytch/stytcherror"
)

type Client struct {
	C *stytch.Client
}

func (c *Client) Send(
	ctx context.Context,
	body *b2c.MagicLinksEmailSendParams,
) (*b2c.MagicLinksEmailSendResponse, error) {
	path := "/magic_links/email/send"

	var jsonBody []byte
	var err error
	if body != nil {
		jsonBody, err = json.Marshal(body)
		if err != nil {
			return nil, stytcherror.NewClientLibraryError("Oops, something seems to have gone wrong " +
				"marshalling the /magic_links/email/send request body")
		}
	}

	var retVal b2c.MagicLinksEmailSendResponse
	err = c.C.NewRequest(ctx, "POST", path, nil, jsonBody, &retVal)
	return &retVal, err
}

func (c *Client) LoginOrCreate(
	ctx context.Context,
	body *b2c.MagicLinksEmailLoginOrCreateParams,
) (
	*b2c.MagicLinksEmailLoginOrCreateResponse, error,
) {
	path := "/magic_links/email/login_or_create"

	var jsonBody []byte
	var err error
	if body != nil {
		jsonBody, err = json.Marshal(body)
		if err != nil {
			return nil, stytcherror.NewClientLibraryError("Oops, something seems to have gone wrong " +
				"marshalling the /magic_links/email/login_or_create request body")
		}
	}

	var retVal b2c.MagicLinksEmailLoginOrCreateResponse
	err = c.C.NewRequest(ctx, "POST", path, nil, jsonBody, &retVal)
	return &retVal, err
}

func (c *Client) Invite(
	ctx context.Context,
	body *b2c.MagicLinksEmailInviteParams,
) (*b2c.MagicLinksEmailInviteResponse, error) {
	path := "/magic_links/email/invite"

	var jsonBody []byte
	var err error
	if body != nil {
		jsonBody, err = json.Marshal(body)
		if err != nil {
			return nil, stytcherror.NewClientLibraryError("Oops, something seems to have gone wrong " +
				"marshalling the /magic_links/email/invite request body")
		}
	}

	var retVal b2c.MagicLinksEmailInviteResponse
	err = c.C.NewRequest(ctx, "POST", path, nil, jsonBody, &retVal)
	return &retVal, err
}

func (c *Client) RevokeInvite(
	ctx context.Context,
	body *b2c.MagicLinksEmailRevokeInviteParams,
) (
	*b2c.MagicLinksEmailRevokeInviteResponse, error,
) {
	path := "/magic_links/email/revoke_invite"

	var jsonBody []byte
	var err error
	if body != nil {
		jsonBody, err = json.Marshal(body)
		if err != nil {
			return nil, stytcherror.NewClientLibraryError("Oops, something seems to have gone wrong " +
				"marshalling the /magic_links/email/revoke_invite request body")
		}
	}

	var retVal b2c.MagicLinksEmailRevokeInviteResponse
	err = c.C.NewRequest(ctx, "POST", path, nil, jsonBody, &retVal)
	return &retVal, err
}
