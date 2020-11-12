package stytch

import (
	"encoding/json"
)

/**
 * Add a user to Stytch. A user_id is returned in the response that can then be used to perform other operations within Stytch.
 * @param    *models_pkg.UserCreate        body     parameter: Required
 * @return	Returns the *models_pkg.UserCreateResponse response from the API call
 */
func (c *Client) CreateUser (body *CreateUser) (*CreateUserResponse, error) {
	jsonBody, err := json.Marshal(body)
	if err != nil {
		return nil, err
	}

	var retVal *CreateUserResponse
	err = c.newRequest("POST", "/users", jsonBody, &retVal)
	return retVal, err
}

/**
* Fetch a given user to see what their various attributes are.
* @param    string        userId      parameter: Required
* @return	Returns the *models_pkg.UserGetResponse response from the API call
*/
func (c *Client) GetUser (userID string) (*GetUserResponse, error) {

	path := "/users/" + userID

	var retVal *GetUserResponse
	err := c.newRequest("GET", path, nil, &retVal)
	return retVal, err
}

/**
* Update a user's attributes. For example, you can add additional emails or change the user's primary email.
* @param    string                        userId      parameter: Required
* @param    *models_pkg.UserUpdate        body        parameter: Required
* @return	Returns the *models_pkg.UserUpdateResponse response from the API call
*/
func (c *Client) UpdateUser (userID string, body *UpdateUser) (*UpdateUserResponse, error) {
	path := "/users/" + userID

	jsonBody, err := json.Marshal(body)
	if err != nil {
		return nil, err
	}

	var retVal *UpdateUserResponse
	err = c.newRequest("PUT", path, jsonBody, &retVal)
	return retVal, err
}

func (c *Client) DeleteUser (userID string) (*DeleteUserResponse, error) {
	path := "/users/" + userID

	var retVal *DeleteUserResponse
	err := c.newRequest("DELETE", path, nil, &retVal)
	return retVal, err
}