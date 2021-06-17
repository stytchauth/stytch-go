package stytch

import (
	"encoding/json"
)

func (c *Client) OTPsSMSSend(body *OTPsSMSSend) (*OTPsSMSSendResponse, error) {
	path := "/otps/sms/send"

	var jsonBody []byte
	var err error
	if body != nil {
		jsonBody, err = json.Marshal(body)
		if err != nil {
			return nil, newInternalServerError("Oops, something seems to have gone wrong " +
				"marshalling the OTPsSMSSend request body")
		}
	}

	var retVal OTPsSMSSendResponse
	err = c.newRequest("POST", path, nil, jsonBody, &retVal)
	return &retVal, err
}

func (c *Client) OTPsSMSLoginOrCreate(body *OTPsSMSLoginOrCreate,
) (*OTPsSMSLoginOrCreateResponse, error) {
	path := "/otps/sms/login_or_create"

	var jsonBody []byte
	var err error
	if body != nil {
		jsonBody, err = json.Marshal(body)
		if err != nil {
			return nil, newInternalServerError("Oops, something seems to have gone wrong " +
				"marshalling the OTPsSMSLoginOrCreate request body")
		}
	}

	var retVal OTPsSMSLoginOrCreateResponse
	err = c.newRequest("POST", path, nil, jsonBody, &retVal)
	return &retVal, err
}

func (c *Client) OTPsAuthenticate(body *OTPsAuthenticate) (*OTPsAuthenticateResponse, error) {
	path := "/otps/authenticate"

	var jsonBody []byte
	var err error
	if body != nil {
		jsonBody, err = json.Marshal(body)
		if err != nil {
			return nil, newInternalServerError("Oops, something seems to have gone wrong " +
				"marshalling the OTPsAuthenticate request body")
		}
	}

	var retVal OTPsAuthenticateResponse
	err = c.newRequest("POST", path, nil, jsonBody, &retVal)
	return &retVal, err
}
