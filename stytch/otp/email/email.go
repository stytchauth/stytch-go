package email

import (
	"encoding/json"

	"github.com/stytchauth/stytch-go/v6/stytch"
	"github.com/stytchauth/stytch-go/v6/stytch/stytcherror"
)

type Client struct {
	C *stytch.Client
}

func (c *Client) Send(body *stytch.OTPsEmailSendParams) (*stytch.OTPsEmailSendResponse, error) {
	path := "/otps/email/send"

	var jsonBody []byte
	var err error
	if body != nil {
		jsonBody, err = json.Marshal(body)
		if err != nil {
			return nil, stytcherror.NewClientLibraryError(
				"Oops, something seems to have gone wrong marshalling the /otps/email/send request body")
		}
	}

	var retVal stytch.OTPsEmailSendResponse
	err = c.C.NewRequest("POST", path, nil, jsonBody, &retVal)
	return &retVal, err
}

func (c *Client) LoginOrCreate(
	body *stytch.OTPsEmailLoginOrCreateParams,
) (*stytch.OTPsEmailLoginOrCreateResponse, error) {
	path := "/otps/email/login_or_create"

	var jsonBody []byte
	var err error
	if body != nil {
		jsonBody, err = json.Marshal(body)
		if err != nil {
			return nil, stytcherror.NewClientLibraryError("Oops, something seems to have gone wrong " +
				"marshalling the /otps/email/login_or_create request body")
		}
	}

	var retVal stytch.OTPsEmailLoginOrCreateResponse
	err = c.C.NewRequest("POST", path, nil, jsonBody, &retVal)
	return &retVal, err
}
