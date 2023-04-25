package sms

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
	body *b2c.OTPsSMSSendParams,
) (*b2c.OTPsSMSSendResponse, error) {
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

	var retVal b2c.OTPsSMSSendResponse
	err = c.C.NewRequest(ctx, "POST", path, nil, jsonBody, &retVal)
	return &retVal, err
}

func (c *Client) LoginOrCreate(
	ctx context.Context,
	body *b2c.OTPsSMSLoginOrCreateParams,
) (*b2c.OTPsSMSLoginOrCreateResponse, error) {
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

	var retVal b2c.OTPsSMSLoginOrCreateResponse
	err = c.C.NewRequest(ctx, "POST", path, nil, jsonBody, &retVal)
	return &retVal, err
}
