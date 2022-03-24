package sms

import (
	"encoding/json"

	"github.com/stytchauth/stytch-go/v4/stytch"
	"github.com/stytchauth/stytch-go/v4/stytch/stytcherror"
)

type Client struct {
	C *stytch.Client
}

func (c *Client) Send(body *stytch.OTPsSMSSendParams) (*stytch.OTPsSMSSendResponse, error) {
	path := "/otps/sms/send"

	var jsonBody []byte
	var err error
	if body != nil {
		jsonBody, err = json.Marshal(body)
		if err != nil {
			return nil, stytcherror.NewClientLibraryError(
				"Oops, something seems to have gone wrong marshalling the /otps/sms/send request body")
		}
	}

	var retVal stytch.OTPsSMSSendResponse
	err = c.C.NewRequest("POST", path, nil, jsonBody, &retVal)
	return &retVal, err
}

func (c *Client) LoginOrCreate(
	body *stytch.OTPsSMSLoginOrCreateParams) (*stytch.OTPsSMSLoginOrCreateResponse, error) {
	path := "/otps/sms/login_or_create"

	var jsonBody []byte
	var err error
	if body != nil {
		jsonBody, err = json.Marshal(body)
		if err != nil {
			return nil, stytcherror.NewClientLibraryError("Oops, something seems to have gone wrong " +
				"marshalling the /otps/sms/login_or_create request body")
		}
	}

	var retVal stytch.OTPsSMSLoginOrCreateResponse
	err = c.C.NewRequest("POST", path, nil, jsonBody, &retVal)
	return &retVal, err
}
