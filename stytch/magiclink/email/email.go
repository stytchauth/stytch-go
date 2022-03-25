package email

import (
	"encoding/json"

	"github.com/stytchauth/stytch-go/v4/stytch"
	"github.com/stytchauth/stytch-go/v4/stytch/stytcherror"
)

type Client struct {
	C *stytch.Client
}

func (c *Client) Send(
	body *stytch.MagicLinksEmailSendParams) (*stytch.MagicLinksEmailSendResponse, error) {
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

	var retVal stytch.MagicLinksEmailSendResponse
	err = c.C.NewRequest("POST", path, nil, jsonBody, &retVal)
	return &retVal, err
}

func (c *Client) LoginOrCreate(body *stytch.MagicLinksEmailLoginOrCreateParams) (
	*stytch.MagicLinksEmailLoginOrCreateResponse, error) {
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

	var retVal stytch.MagicLinksEmailLoginOrCreateResponse
	err = c.C.NewRequest("POST", path, nil, jsonBody, &retVal)
	return &retVal, err
}

func (c *Client) Invite(
	body *stytch.MagicLinksEmailInviteParams) (*stytch.MagicLinksEmailInviteResponse, error) {
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

	var retVal stytch.MagicLinksEmailInviteResponse
	err = c.C.NewRequest("POST", path, nil, jsonBody, &retVal)
	return &retVal, err
}

func (c *Client) RevokeInvite(body *stytch.MagicLinksEmailRevokeInviteParams) (
	*stytch.MagicLinksEmailRevokeInviteResponse, error) {
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

	var retVal stytch.MagicLinksEmailRevokeInviteResponse
	err = c.C.NewRequest("POST", path, nil, jsonBody, &retVal)
	return &retVal, err
}
