package stytch

func (c *Client) DeleteEmail(emailID string, userID string) (*DeleteEmailResponse, error) {
	path := "emails/" + emailID + "/users/" + userID

	var retVal *DeleteEmailResponse
	err := c.newRequest("DELETE", path, nil, &retVal)
	return retVal, err
}
