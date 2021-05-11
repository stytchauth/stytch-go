package stytch

import (
	"encoding/json"
)

func (c *Client) SendOTPBySMS(body *SendOTPBySMS) (*SendOTPBySMSResponse, error) {
	path := "/otp/send_by_sms"

	var jsonBody []byte
	var err error
	if body != nil {
		jsonBody, err = json.Marshal(body)
		if err != nil {
			return nil, newClientLibraryError("Oops, something seems to have gone wrong " +
				"marshalling the SendOTPBySMS request body")
		}
	}

	var retVal SendOTPBySMSResponse
	err = c.newRequest("POST", path, nil, jsonBody, &retVal)
	return &retVal, err
}

func (c *Client) LoginOrCreateUserBySMS(body *LoginOrCreateUserBySMS,
) (*LoginOrCreateUserBySMSResponse, error) {
	path := "/otp/login_or_create"

	var jsonBody []byte
	var err error
	if body != nil {
		jsonBody, err = json.Marshal(body)
		if err != nil {
			return nil, newClientLibraryError("Oops, something seems to have gone wrong " +
				"marshalling the LoginOrCreateUserBySMS request body")
		}
	}

	var retVal LoginOrCreateUserBySMSResponse
	err = c.newRequest("POST", path, nil, jsonBody, &retVal)
	return &retVal, err
}

func (c *Client) AuthenticateOTP(body *AuthenticateOTP) (*AuthenticateOTPResponse, error) {
	path := "/otp/authenticate"

	var jsonBody []byte
	var err error
	if body != nil {
		jsonBody, err = json.Marshal(body)
		if err != nil {
			return nil, newClientLibraryError("Oops, something seems to have gone wrong " +
				"marshalling the AuthenticateOTP request body")
		}
	}

	var retVal AuthenticateOTPResponse
	err = c.newRequest("POST", path, nil, jsonBody, &retVal)
	return &retVal, err
}
