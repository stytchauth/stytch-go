package stytch

import (
	"encoding/json"
	"strconv"
)

func (c *Client) CreateUser(body *CreateUser) (*CreateUserResponse, error) {
	jsonBody, err := json.Marshal(body)
	if err != nil {
		return nil, newInternalServerError("Oops, something seems to have gone wrong " +
			"marshalling the CreateUser request body")
	}

	var retVal *CreateUserResponse
	err = c.newRequest("POST", "/users", nil, jsonBody, retVal)
	return retVal, err
}

func (c *Client) GetUser(userID string) (*GetUserResponse, error) {
	path := "/users/" + userID

	var retVal *GetUserResponse
	err := c.newRequest("GET", path, nil, nil, retVal)
	return retVal, err
}

func (c *Client) GetInvitedUsers(body *GetInvitedUsers) (*GetInvitedUsersResponse, error) {
	limitString := ""
	if body.Limit != 0 {
		limitString = strconv.Itoa(int(body.Limit))
	}

	queryParams := map[string]string{
		"limit":             limitString,
		"starting_after_id": body.StartingAfterID,
	}

	var retVal *GetInvitedUsersResponse
	err := c.newRequest("GET", "/users/invites", queryParams, nil, retVal)
	return retVal, err
}

func (c *Client) UpdateUser(userID string, body *UpdateUser) (*UpdateUserResponse, error) {
	path := "/users/" + userID

	jsonBody, err := json.Marshal(body)
	if err != nil {
		return nil, newInternalServerError("Oops, something seems to have gone wrong " +
			"marshalling the UpdateUser request body")
	}

	var retVal *UpdateUserResponse
	err = c.newRequest("PUT", path, nil, jsonBody, retVal)
	return retVal, err
}

func (c *Client) DeleteUser(userID string) (*DeleteUserResponse, error) {
	path := "/users/" + userID

	var retVal *DeleteUserResponse
	err := c.newRequest("DELETE", path, nil, nil, retVal)
	return retVal, err
}

func (c *Client) DeleteUserEmail(userID string, email string) (*DeleteUserEmailResponse, error) {
	path := "/users/" + userID + "/emails/" + email

	var retVal *DeleteUserEmailResponse
	err := c.newRequest("DELETE", path, nil, nil, retVal)
	return retVal, err
}
