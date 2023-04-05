package whatsapp

import (
	"context"
	"encoding/json"

	"github.com/stytchauth/stytch-go/v7/stytch"
	"github.com/stytchauth/stytch-go/v7/stytch/stytcherror"
)

type Client struct {
	C *stytch.Client
}

func (c *Client) Send(
	ctx context.Context,
	body *stytch.OTPsWhatsAppSendParams,
) (*stytch.OTPsWhatsAppSendResponse, error) {
	path := "/otps/whatsapp/send"

	var jsonBody []byte
	var err error
	if body != nil {
		jsonBody, err = json.Marshal(body)
		if err != nil {
			return nil, stytcherror.NewClientLibraryError(
				"Oops, something seems to have gone wrong marshalling the /otps/whatsapp/send request body")
		}
	}

	var retVal stytch.OTPsWhatsAppSendResponse
	err = c.C.NewRequest(ctx, "POST", path, nil, jsonBody, &retVal)
	return &retVal, err
}

func (c *Client) LoginOrCreate(
	ctx context.Context,
	body *stytch.OTPsWhatsAppLoginOrCreateParams,
) (*stytch.OTPsWhatsAppLoginOrCreateResponse, error) {
	path := "/otps/whatsapp/login_or_create"

	var jsonBody []byte
	var err error
	if body != nil {
		jsonBody, err = json.Marshal(body)
		if err != nil {
			return nil, stytcherror.NewClientLibraryError("Oops, something seems to have gone wrong " +
				"marshalling the /otps/whatsapp/login_or_create request body")
		}
	}

	var retVal stytch.OTPsWhatsAppLoginOrCreateResponse
	err = c.C.NewRequest(ctx, "POST", path, nil, jsonBody, &retVal)
	return &retVal, err
}
