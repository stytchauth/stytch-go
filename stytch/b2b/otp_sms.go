package b2b

// !!!
// WARNING: This file is autogenerated
// Only modify code within MANUAL() sections
// or your changes may be overwritten later!
// !!!

import (
	"context"
	"encoding/json"

	"github.com/stytchauth/stytch-go/v10/stytch"
	"github.com/stytchauth/stytch-go/v10/stytch/b2b/otp/sms"
	"github.com/stytchauth/stytch-go/v10/stytch/stytcherror"
)

type OTPsSmsClient struct {
	C stytch.Client
}

func NewOTPsSmsClient(c stytch.Client) *OTPsSmsClient {
	return &OTPsSmsClient{
		C: c,
	}
}

// Send a one-time passcode (OTP) to a Member's phone number.
//
// If the Member already has a phone number, the `mfa_phone_number` field is not needed; the endpoint will
// send an OTP to the number associated with the Member.
// If the Member does not have a phone number, the endpoint will send an OTP to the `mfa_phone_number`
// provided and link the `mfa_phone_number` with the Member.
//
// An error will be thrown if the Member already has a phone number and the provided `mfa_phone_number`
// does not match the existing one.
//
// Note that sending another OTP code before the first has expired will invalidate the first code.
//
// If a Member has a phone number and is enrolled in MFA, then after a successful primary authentication
// event (e.g. [email magic link](https://stytch.com/docs/b2b/api/authenticate-magic-link) or
// [SSO](https://stytch.com/docs/b2b/api/sso-authenticate) login is complete), an SMS OTP will
// automatically be sent to their phone number. In that case, this endpoint should only be used for
// subsequent authentication events, such as prompting a Member for an OTP again after a period of
// inactivity.
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
		"/v1/b2b/otps/sms/send",
		nil,
		jsonBody,
		&retVal,
	)
	return &retVal, err
}

// Authenticate: SMS OTPs may not be used as a primary authentication mechanism. They can be used to
// complete an MFA requirement, or they can be used as a step-up factor to be added to an existing session.
//
// This endpoint verifies that the one-time passcode (OTP) is valid and hasn't expired or been previously
// used. A given Member may only have a single active OTP code at any given time. If a Member requests
// another OTP code before the first one has expired, the first one will be invalidated.
//
// Exactly one of `intermediate_session_token`, `session_token`, or `session_jwt` must be provided in the
// request.
// If an intermediate session token is provided, this operation will consume it.
//
// Intermediate session tokens are generated upon successful calls to primary authenticate methods in the
// case where MFA is required,
// such as [email magic link authenticate](https://stytch.com/docs/b2b/api/authenticate-magic-link),
// or upon successful calls to discovery authenticate methods, such as
// [email magic link discovery authenticate](https://stytch.com/docs/b2b/api/authenticate-discovery-magic-link).
//
// If the Organization's MFA policy is `REQUIRED_FOR_ALL`, a successful OTP authentication will change the
// Member's `mfa_enrolled` status to `true` if it is not already `true`.
// If the Organization's MFA policy is `OPTIONAL`, the Member's MFA enrollment can be toggled by passing in
// a value for the `set_mfa_enrollment` field.
// The Member's MFA enrollment can also be toggled through the
// [Update Member](https://stytch.com/docs/b2b/api/update-member) endpoint.
//
// Provide the `session_duration_minutes` parameter to set the lifetime of the session. If the
// `session_duration_minutes` parameter is not specified, a Stytch session will be created with a duration
// of 60 minutes.
func (c *OTPsSmsClient) Authenticate(
	ctx context.Context,
	body *sms.AuthenticateParams,
) (*sms.AuthenticateResponse, error) {
	var jsonBody []byte
	var err error
	if body != nil {
		jsonBody, err = json.Marshal(body)
		if err != nil {
			return nil, stytcherror.NewClientLibraryError("error marshaling request body")
		}
	}

	var retVal sms.AuthenticateResponse
	err = c.C.NewRequest(
		ctx,
		"POST",
		"/v1/b2b/otps/sms/authenticate",
		nil,
		jsonBody,
		&retVal,
	)
	return &retVal, err
}
