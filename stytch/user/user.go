package user

import (
	"encoding/json"
	"strconv"

	"github.com/stytchauth/stytch-go/v5/stytch"
	"github.com/stytchauth/stytch-go/v5/stytch/stytcherror"
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
			return nil, stytcherror.NewClientLibraryError("Oops, something seems to have gone wrong " +
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

func (c *Client) Search(
	body *stytch.UsersSearchParams) (*stytch.UsersSearchResponse, error) {
	var jsonBody []byte
	var err error
	if body != nil {
		jsonBody, err = json.Marshal(body)
		if err != nil {
			return nil, stytcherror.NewClientLibraryError("Oops, something seems to have gone wrong " +
				"marshalling the search users request body")
		}
	}

	var retVal stytch.UsersSearchResponse
	err = c.C.NewRequest("POST", "/users/search", nil, jsonBody, &retVal)
	return &retVal, err
}

type iteratorState string

const (
	iteratorStatePending    iteratorState = "Pending"
	iteratorStateInProgress iteratorState = "In Progress"
	iteratorStateErrored    iteratorState = "Errored"
	iteratorStateComplete   iteratorState = "Complete"
)

type UserSearchIterator struct {
	c     *Client
	body  *stytch.UsersSearchParams
	state iteratorState
}

func (i *UserSearchIterator) HasNext() bool {
	return i.state == iteratorStatePending || i.state == iteratorStateInProgress
}

func (i *UserSearchIterator) Next() ([]stytch.User, error) {
	res, err := i.c.Search(i.body)
	if err != nil {
		i.state = iteratorStateErrored
		return nil, err
	}

	i.body.Cursor = res.ResultsMetadata.NextCursor
	if i.body.Cursor == "" {
		i.state = iteratorStateComplete
	}

	return res.Results, nil
}

func (c *Client) SearchAll(body *stytch.UsersSearchParams) *UserSearchIterator {
	return &UserSearchIterator{
		c:     c,
		body:  body,
		state: iteratorStatePending,
	}
}

func (c *Client) Update(
	userID string, body *stytch.UsersUpdateParams) (*stytch.UsersUpdateResponse, error) {
	path := "/users/" + userID

	var jsonBody []byte
	var err error
	if body != nil {
		jsonBody, err = json.Marshal(body)
		if err != nil {
			return nil, stytcherror.NewClientLibraryError("Oops, something seems to have gone wrong " +
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

func (c *Client) DeleteWebAuthnRegistration(
	webAuthnRegistration string) (*stytch.UsersDeleteWebAuthnRegistrationResponse, error) {
	path := "/users/webauthn_registrations/" + webAuthnRegistration

	var retVal stytch.UsersDeleteWebAuthnRegistrationResponse
	err := c.C.NewRequest("DELETE", path, nil, nil, &retVal)
	return &retVal, err
}

func (c *Client) DeleteTOTP(
	totpID string) (*stytch.UsersDeleteTOTPResponse, error) {
	path := "/users/totps/" + totpID

	var retVal stytch.UsersDeleteTOTPResponse
	err := c.C.NewRequest("DELETE", path, nil, nil, &retVal)
	return &retVal, err
}

func (c *Client) DeleteCryptoWallet(
	cryptoWalletID string) (*stytch.UsersDeleteCryptoWalletResponse, error) {
	path := "/users/crypto_wallets/" + cryptoWalletID

	var retVal stytch.UsersDeleteCryptoWalletResponse
	err := c.C.NewRequest("DELETE", path, nil, nil, &retVal)
	return &retVal, err
}
