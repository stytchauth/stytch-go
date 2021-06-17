package stytch

import (
	"encoding/json"
)

func (c *Client) MagicLinksEmailSend(body *MagicLinksEmailSend) (*MagicLinksEmailSendResponse,
	error) {
	path := "/magic_links/email/send"

	var jsonBody []byte
	var err error
	if body != nil {
		jsonBody, err = json.Marshal(body)
		if err != nil {
			return nil, newInternalServerError("Oops, something seems to have gone wrong " +
				"marshalling the MagicLinksEmailSend request body")
		}
	}

	var retVal MagicLinksEmailSendResponse
	err = c.newRequest("POST", path, nil, jsonBody, &retVal)
	return &retVal, err
}

func (c *Client) MagicLinksEmailLoginOrCreate(body *MagicLinksEmailLoginOrCreate) (
	*MagicLinksEmailLoginOrCreateResponse, error) {
	path := "/magic_links/email/login_or_create"

	var jsonBody []byte
	var err error
	if body != nil {
		jsonBody, err = json.Marshal(body)
		if err != nil {
			return nil, newInternalServerError("Oops, something seems to have gone wrong " +
				"marshalling the MagicLinksEmailLoginOrCreate request body")
		}
	}

	var retVal MagicLinksEmailLoginOrCreateResponse
	err = c.newRequest("POST", path, nil, jsonBody, &retVal)
	return &retVal, err
}

func (c *Client) MagicLinksEmailInvite(body *MagicLinksEmailInvite) (
	*MagicLinksEmailInviteResponse, error) {
	path := "/magic_links/email/invite"

	var jsonBody []byte
	var err error
	if body != nil {
		jsonBody, err = json.Marshal(body)
		if err != nil {
			return nil, newInternalServerError("Oops, something seems to have gone wrong " +
				"marshalling the MagicLinksEmailInvite request body")
		}
	}

	var retVal MagicLinksEmailInviteResponse
	err = c.newRequest("POST", path, nil, jsonBody, &retVal)
	return &retVal, err
}

func (c *Client) MagicLinksEmailRevokeInvite(body *MagicLinksEmailRevokeInvite) (
	*MagicLinksEmailRevokeInviteResponse, error) {
	path := "/magic_links/email/revoke_invite"

	var jsonBody []byte
	var err error
	if body != nil {
		jsonBody, err = json.Marshal(body)
		if err != nil {
			return nil, newInternalServerError("Oops, something seems to have gone wrong " +
				"marshalling the MagicLinksEmailRevokeInvite request body")
		}
	}

	var retVal MagicLinksEmailRevokeInviteResponse
	err = c.newRequest("POST", path, nil, jsonBody, &retVal)
	return &retVal, err
}

func (c *Client) MagicLinksAuthenticate(
	body *MagicLinksAuthenticate,
) (*MagicLinksAuthenticateResponse, error) {
	path := "/magic_links/authenticate"

	var jsonBody []byte
	var err error
	if body != nil {
		jsonBody, err = json.Marshal(body)
		if err != nil {
			return nil, newInternalServerError("Oops, something seems to have gone wrong " +
				"marshalling the MagicLinksAuthenticate request body")
		}
	}

	var retVal MagicLinksAuthenticateResponse
	err = c.newRequest("POST", path, nil, jsonBody, &retVal)
	return &retVal, err
}
