package otp

import (
	"encoding/json"

	"github.com/stytchauth/stytch-go/stytch"
	"github.com/stytchauth/stytch-go/stytch/otp/sms"
)

type Client struct {
	C   *stytch.Client
	SMS *sms.Client
}

func (c *Client) AuthenticateOTP(
	body *stytch.OTPsAuthenticateParams) (*stytch.OTPsAuthenticateResponse, error) {
	path := "/otps/authenticate"

	var jsonBody []byte
	var err error
	if body != nil {
		jsonBody, err = json.Marshal(body)
		if err != nil {
			return nil, stytch.NewInternalServerError("Oops, something seems to have gone wrong " +
				"marshalling the /otps/authenticate request body")
		}
	}

	var retVal stytch.OTPsAuthenticateResponse
	err = c.C.NewRequest("POST", path, nil, jsonBody, &retVal)
	return &retVal, err
}
