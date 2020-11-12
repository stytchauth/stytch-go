package stytch

import "encoding/json"

func (c *Client) SendMagicLink (body *SendMagicLink) (*SendMagicLinkResponse, error) {
	path := "/magic_links/send"

	jsonBody, err := json.Marshal(body)
	if err != nil {
		return nil, err
	}

	var retVal *SendMagicLinkResponse
	err = c.newRequest("POST", path, jsonBody, &retVal)
	return retVal, err
}


func (c *Client) SendMagicLinkByEmail (body *SendMagicLinkByEmail) (*SendMagicLinkResponse, error) {
	path := "/magic_links/send_by_email"

	jsonBody, err := json.Marshal(body)
	if err != nil {
		return nil, err
	}

	var retVal *SendMagicLinkResponse
	err = c.newRequest("POST", path, jsonBody, &retVal)
	return retVal, err
}


func (c *Client) AuthenticateMagicLink (token string, body *AuthenticateMagicLink) (*AuthenticateMagicLinkResponse, error) {
	path := "/magic_links/" + token + "/authenticate"

	jsonBody, err := json.Marshal(body)
	if err != nil {
		return nil, err
	}

	var retVal *AuthenticateMagicLinkResponse
	err = c.newRequest("POST", path, jsonBody, &retVal)
	return retVal, err
}