package stytch

import (
	"encoding/json"
)

func (c *Client) CreateUser(body *CreateUser) (*CreateUserResponse, error) {
	jsonBody, err := json.Marshal(body)
	if err != nil {
		return nil, err
	}

	var retVal *CreateUserResponse
	err = c.newRequest("POST", "/users", jsonBody, &retVal)
	return retVal, err
}

func (c *Client) GetUser(userID string) (*GetUserResponse, error) {
	path := "/users/" + userID

	var retVal *GetUserResponse
	err := c.newRequest("GET", path, nil, &retVal)
	return retVal, err
}

func (c *Client) UpdateUser(userID string, body *UpdateUser) (*UpdateUserResponse, error) {
	path := "/users/" + userID

	jsonBody, err := json.Marshal(body)
	if err != nil {
		return nil, err
	}

	var retVal *UpdateUserResponse
	err = c.newRequest("PUT", path, jsonBody, &retVal)
	return retVal, err
}

func (c *Client) DeleteUser(userID string) (*DeleteUserResponse, error) {
	path := "/users/" + userID

	var retVal *DeleteUserResponse
	err := c.newRequest("DELETE", path, nil, &retVal)
	return retVal, err
}

func (c *Client) DeleteUserEmail(userID string, email string) (*DeleteUserEmailResponse, error) {
	path := "/users/" + userID + "/emails/" + email

	var retVal *DeleteUserEmailResponse
	err := c.newRequest("DELETE", path, nil, retVal)
	return retVal, err
}
