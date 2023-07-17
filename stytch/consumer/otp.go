package consumer

// !!!
// WARNING: This file is autogenerated
// Only modify code within MANUAL() sections
// or your changes may be overwritten later!
// !!!

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/mitchellh/mapstructure"
	"github.com/stytchauth/stytch-go/v9/stytch"
	"github.com/stytchauth/stytch-go/v9/stytch/consumer/otp"
	"github.com/stytchauth/stytch-go/v9/stytch/stytcherror"
)

type OTPsClient struct {
	C        *stytch.Client
	Sms      *OTPsSmsClient
	Whatsapp *OTPsWhatsappClient
	Email    *OTPsEmailClient
}

func NewOTPsClient(c *stytch.Client) *OTPsClient {
	return &OTPsClient{
		C:        c,
		Sms:      NewOTPsSmsClient(c),
		Whatsapp: NewOTPsWhatsappClient(c),
		Email:    NewOTPsEmailClient(c),
	}
}

// Authenticate a User given a `method_id` (the associated `email_id` or `phone_id`) and a `code`. This
// endpoint verifies that the code is valid, hasn't expired or been previously used, and any optional
// security settings such as IP match or user agent match are satisfied. A given `method_id` may only have
// a single active OTP code at any given time, if a User requests another OTP code before the first one has
// expired, the first one will be invalidated.
func (c *OTPsClient) Authenticate(
	ctx context.Context,
	body *otp.AuthenticateParams,
) (*otp.AuthenticateResponse, error) {
	var jsonBody []byte
	var err error
	if body != nil {
		jsonBody, err = json.Marshal(body)
		if err != nil {
			return nil, stytcherror.NewClientLibraryError("error marshaling request body")
		}
	}

	var retVal otp.AuthenticateResponse
	err = c.C.NewRequest(
		ctx,
		"POST",
		"/v1/otps/authenticate",
		nil,
		jsonBody,
		&retVal,
	)
	return &retVal, err
}

// AuthenticateWithClaims fills in the claims pointer with custom claims from the response.
// Pass in a map with the types of values you're expecting so that this function can marshal
// the claims from the response. See ExampleClient_AuthenticateWithClaims_map,
// ExampleClient_AuthenticateWithClaims_struct for examples
func (c *OTPsClient) AuthenticateWithClaims(
	ctx context.Context,
	body *otp.AuthenticateParams,
	claims any,
) (*otp.AuthenticateResponse, error) {
	var jsonBody []byte
	var err error
	if body != nil {
		jsonBody, err = json.Marshal(body)
		if err != nil {
			return nil, stytcherror.NewClientLibraryError("error marshaling request body")
		}
	}

	b, err := c.C.RawRequest(
		ctx,
		"POST",
		"/v1/otps/authenticate",
		nil,
		jsonBody,
	)
	if err != nil {
		return nil, err
	}

	// First extract the Stytch data.
	var retVal otp.AuthenticateResponse
	if err := json.Unmarshal(b, &retVal); err != nil {
		return nil, fmt.Errorf("unmarshal otp.AuthenticateResponse: %w", err)
	}

	if claims == nil {
		return &retVal, nil
	}

	if m, ok := claims.(*map[string]any); ok {
		*m = retVal.Session.CustomClaims
		return &retVal, nil
	}

	// This is where we need to convert claims into a claimsMap
	decoder, err := mapstructure.NewDecoder(&mapstructure.DecoderConfig{
		Result:  &claims,
		TagName: "json",
	})
	if err != nil {
		return nil, err
	}

	err = decoder.Decode(retVal.Session.CustomClaims)
	if err != nil {
		return nil, err
	}

	return &retVal, err
}
