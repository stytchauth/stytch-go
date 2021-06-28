package user

import (
	"encoding/json"
	"strconv"

	"github.com/stytchauth/stytch-go/stytch"
)

type Client struct {
	C *stytch.Client
}

func (c *Client) Create(body *stytch.UsersCreateParams) (*stytch.UsersCreateResponse, error) {
	var jsonBody []byte
	var err error
	if body != nil {
		jsonBody, err = json.Marshal(body)
		if err != nil {
			return nil, stytch.NewClientLibraryError("Oops, something seems to have gone wrong " +
				"marshalling the create user request body")
		}
	}

	var retVal stytch.UsersCreateResponse
	err = c.C.NewRequest("POST", "/users", nil, jsonBody, &retVal)
	return &retVal, err
}

func (c *Client) Get(userID string) (*stytch.UsersGetResponse, error) {
	path := "/users/" + userID

	var retVal stytch.UsersGetResponse
	err := c.C.NewRequest("GET", path, nil, nil, &retVal)
	return &retVal, err
}

func (c *Client) GetPending(
	body *stytch.UsersGetPendingParams) (*stytch.UsersGetPendingResponse, error) {
	var queryParams map[string]string
	if body != nil {
		limitString := ""
		if body.Limit != 0 {
			limitString = strconv.Itoa(int(body.Limit))
		}

		queryParams = map[string]string{
			"limit":             limitString,
			"starting_after_id": body.StartingAfterID,
		}
	}

	var retVal stytch.UsersGetPendingResponse
	err := c.C.NewRequest("GET", "/users/pending", queryParams, nil, &retVal)
	return &retVal, err
}

func (c *Client) Update(
	userID string, body *stytch.UsersUpdateParams) (*stytch.UsersUpdateResponse, error) {
	path := "/users/" + userID

	var jsonBody []byte
	var err error
	if body != nil {
		jsonBody, err = json.Marshal(body)
		if err != nil {
			return nil, stytch.NewClientLibraryError("Oops, something seems to have gone wrong " +
				"marshalling the update users request body")
		}
	}

	var retVal stytch.UsersUpdateResponse
	err = c.C.NewRequest("PUT", path, nil, jsonBody, &retVal)
	return &retVal, err
}

func (c *Client) Delete(userID string) (*stytch.UsersDeleteResponse, error) {
	path := "/users/" + userID

	var retVal stytch.UsersDeleteResponse
	err := c.C.NewRequest("DELETE", path, nil, nil, &retVal)
	return &retVal, err
}

func (c *Client) DeleteEmail(emailID string) (*stytch.UsersDeleteEmailResponse, error) {
	path := "/users/emails/" + emailID

	var retVal stytch.UsersDeleteEmailResponse
	err := c.C.NewRequest("DELETE", path, nil, nil, &retVal)
	return &retVal, err
}

func (c *Client) DeletePhoneNumber(
	phoneID string) (*stytch.UsersDeletePhoneNumberResponse, error) {
	path := "/users/phone_numbers/" + phoneID

	var retVal stytch.UsersDeletePhoneNumberResponse
	err := c.C.NewRequest("DELETE", path, nil, nil, &retVal)
	return &retVal, err
}
