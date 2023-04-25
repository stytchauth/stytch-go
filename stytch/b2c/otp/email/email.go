package email

import (
	"context"
	"encoding/json"
	"github.com/stytchauth/stytch-go/v7/stytch/b2c"

	"github.com/stytchauth/stytch-go/v7/stytch"
	"github.com/stytchauth/stytch-go/v7/stytch/stytcherror"
)

type Client struct {
	C *stytch.Client
}

func (c *Client) Send(
	ctx context.Context,
	body *b2c.OTPsEmailSendParams,
) (*b2c.OTPsEmailSendResponse, error) {
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

	var retVal b2c.OTPsEmailSendResponse
	err = c.C.NewRequest(ctx, "POST", path, nil, jsonBody, &retVal)
	return &retVal, err
}

func (c *Client) LoginOrCreate(
	ctx context.Context,
	body *b2c.OTPsEmailLoginOrCreateParams,
) (*b2c.OTPsEmailLoginOrCreateResponse, error) {
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

	var retVal b2c.OTPsEmailLoginOrCreateResponse
	err = c.C.NewRequest(ctx, "POST", path, nil, jsonBody, &retVal)
	return &retVal, err
}
