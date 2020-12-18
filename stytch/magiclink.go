package stytch

import "encoding/json"

func (c *Client) SendMagicLink(body *SendMagicLink) (*SendMagicLinkResponse, error) {
	path := "/magic_links/send"

	jsonBody, err := json.Marshal(body)
	if err != nil {
		return nil, err
	}

	var retVal *SendMagicLinkResponse
	err = c.newRequest("POST", path, jsonBody, &retVal)
	return retVal, err
}

func (c *Client) SendMagicLinkByEmail(body *SendMagicLinkByEmail) (*SendMagicLinkResponse, error) {
	path := "/magic_links/send_by_email"

	jsonBody, err := json.Marshal(body)
	if err != nil {
		return nil, err
	}

	var retVal *SendMagicLinkResponse
	err = c.newRequest("POST", path, jsonBody, &retVal)
	return retVal, err
}

func (c *Client) LoginOrCreateUser(body *LoginOrCreateUser) (*LoginOrCreateResponse, error) {
	path := "/magic_links/login_or_create"

	jsonBody, err := json.Marshal(body)
	if err != nil {
		return nil, err
	}

	var retVal *LoginOrCreateResponse
	err = c.newRequest("POST", path, jsonBody, &retVal)
	return retVal, err
}

func (c *Client) LoginOrInviteByEmail(body *LoginOrInviteByEmail) (*LoginOrCreateResponse, error) {
	path := "/magic_links/login_or_invite"

	jsonBody, err := json.Marshal(body)
	if err != nil {
		return nil, err
	}

	var retVal *LoginOrCreateResponse
	err = c.newRequest("POST", path, jsonBody, &retVal)
	return retVal, err
}

func (c *Client) InviteByEmail(body *InviteByEmail) (*InviteByEmailResponse, error) {
	path := "/magic_links/invite_by_email"

	jsonBody, err := json.Marshal(body)
	if err != nil {
		return nil, err
	}

	var retVal *InviteByEmailResponse
	err = c.newRequest("POST", path, jsonBody, &retVal)
	return retVal, err
}

func (c *Client) RevokeInviteByEmail(body *RevokeInviteByEmail) (*RevokeInviteByEmailResponse, error) {
	path := "/magic_links/revoke_invite"

	jsonBody, err := json.Marshal(body)
	if err != nil {
		return nil, err
	}

	var retVal *RevokeInviteByEmailResponse
	err = c.newRequest("POST", path, jsonBody, &retVal)
	return retVal, err
}

func (c *Client) AuthenticateMagicLink(
	token string,
	body *AuthenticateMagicLink,
) (*AuthenticateMagicLinkResponse, error) {
	path := "/magic_links/" + token + "/authenticate"

	jsonBody, err := json.Marshal(body)
	if err != nil {
		return nil, err
	}

	var retVal *AuthenticateMagicLinkResponse
	err = c.newRequest("POST", path, jsonBody, &retVal)
	return retVal, err
}
