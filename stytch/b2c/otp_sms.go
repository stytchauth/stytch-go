package b2c

// !!!
// WARNING: This file is autogenerated
// Only modify code within MANUAL() sections
// or your changes may be overwritten later!
// !!!

import (
	"context"
	"encoding/json"

	"github.com/stytchauth/stytch-go/v8/stytch"
	"github.com/stytchauth/stytch-go/v8/stytch/b2c/otp/sms"
	"github.com/stytchauth/stytch-go/v8/stytch/stytcherror"
)

type OTPsSmsClient struct {
	C *stytch.Client
}

func NewOTPsSmsClient(c *stytch.Client) *OTPsSmsClient {
	return &OTPsSmsClient{
		C: c,
	}
}

// Send a one-time passcode (OTP) to a user's phone number. If you'd like to create a user and send them a
// passcode with one request, use our [log in or
// create](https://stytch.com/docs/api/log-in-or-create-user-by-sms) endpoint.
//
// Note that sending another OTP code before the first has expired will invalidate the first code.
//
// ### Add a phone number to an existing user
//
// This endpoint also allows you to add a new phone number to an existing Stytch User. Including a
// `user_id`, `session_token`, or `session_jwt` in the request will add the phone number to the
// pre-existing Stytch User upon successful authentication.
//
// Adding a new phone number to an existing Stytch User requires the user to be present and validate the
// phone number via OTP. This requirement is in place to prevent account takeover attacks.
//
// ### Next steps
//
// Collect the OTP which was delivered to the user. Call [Authenticate
// OTP](https://stytch.com/docs/api/authenticate-otp) using the OTP `code` along with the `phone_id` found
// in the response as the `method_id`.
func (c *OTPsSmsClient) Send(
	ctx context.Context,
	body *sms.SendParams,
) (*sms.SendResponse, error) {
	var jsonBody []byte
	var err error
	if body != nil {
		jsonBody, err = json.Marshal(body)
		if err != nil {
			return nil, stytcherror.NewClientLibraryError("error marshaling request body")
		}
	}

	var retVal sms.SendResponse
	err = c.C.NewRequest(
		ctx,
		"POST",
		"/v1/otps/sms/send",
		nil,
		jsonBody,
		&retVal,
	)
	return &retVal, err
}

// LoginOrCreate: Send a one-time passcode (OTP) to a User using their phone number. If the phone number is
// not associated with a user already, a user will be created.
//
// ### Next steps
//
// Collect the OTP which was delivered to the User. Call [Authenticate
// OTP](https://stytch.com/docs/api/authenticate-otp) using the OTP `code` along with the `phone_id` found
// in the response as the `method_id`.
func (c *OTPsSmsClient) LoginOrCreate(
	ctx context.Context,
	body *sms.LoginOrCreateParams,
) (*sms.LoginOrCreateResponse, error) {
	var jsonBody []byte
	var err error
	if body != nil {
		jsonBody, err = json.Marshal(body)
		if err != nil {
			return nil, stytcherror.NewClientLibraryError("error marshaling request body")
		}
	}

	var retVal sms.LoginOrCreateResponse
	err = c.C.NewRequest(
		ctx,
		"POST",
		"/v1/otps/sms/login_or_create",
		nil,
		jsonBody,
		&retVal,
	)
	return &retVal, err
}
