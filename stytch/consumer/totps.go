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
	"github.com/stytchauth/stytch-go/v16/stytch"
	"github.com/stytchauth/stytch-go/v16/stytch/consumer/totps"
	"github.com/stytchauth/stytch-go/v16/stytch/stytcherror"
)

type TOTPsClient struct {
	C stytch.Client
}

func NewTOTPsClient(c stytch.Client) *TOTPsClient {
	return &TOTPsClient{
		C: c,
	}
}

// Create a new TOTP instance for a user. The user can use the authenticator application of their choice to
// scan the QR code or enter the secret.
func (c *TOTPsClient) Create(
	ctx context.Context,
	body *totps.CreateParams,
) (*totps.CreateResponse, error) {
	var jsonBody []byte
	var err error
	if body != nil {
		jsonBody, err = json.Marshal(body)
		if err != nil {
			return nil, stytcherror.NewClientLibraryError("error marshaling request body")
		}
	}

	headers := make(map[string][]string)

	var retVal totps.CreateResponse
	err = c.C.NewRequest(
		ctx,
		stytch.RequestParams{
			Method:      "POST",
			Path:        "/v1/totps",
			QueryParams: nil,
			Body:        jsonBody,
			V:           &retVal,
			Headers:     headers,
		},
	)
	return &retVal, err
}

// Authenticate a TOTP code entered by a user.
func (c *TOTPsClient) Authenticate(
	ctx context.Context,
	body *totps.AuthenticateParams,
) (*totps.AuthenticateResponse, error) {
	var jsonBody []byte
	var err error
	if body != nil {
		jsonBody, err = json.Marshal(body)
		if err != nil {
			return nil, stytcherror.NewClientLibraryError("error marshaling request body")
		}
	}

	headers := make(map[string][]string)

	var retVal totps.AuthenticateResponse
	err = c.C.NewRequest(
		ctx,
		stytch.RequestParams{
			Method:      "POST",
			Path:        "/v1/totps/authenticate",
			QueryParams: nil,
			Body:        jsonBody,
			V:           &retVal,
			Headers:     headers,
		},
	)
	return &retVal, err
}

// AuthenticateWithClaims fills in the claims pointer with custom claims from the response.
// Pass in a map with the types of values you're expecting so that this function can marshal
// the claims from the response. See ExampleClient_AuthenticateWithClaims_map,
// ExampleClient_AuthenticateWithClaims_struct for examples
func (c *TOTPsClient) AuthenticateWithClaims(
	ctx context.Context,
	body *totps.AuthenticateParams,
	claims any,
) (*totps.AuthenticateResponse, error) {
	var jsonBody []byte
	var err error
	if body != nil {
		jsonBody, err = json.Marshal(body)
		if err != nil {
			return nil, stytcherror.NewClientLibraryError("error marshaling request body")
		}
	}

	headers := make(map[string][]string)

	b, err := c.C.RawRequest(
		ctx,
		stytch.RequestParams{
			Method:      "POST",
			Path:        "/v1/totps/authenticate",
			QueryParams: nil,
			Body:        jsonBody,
			Headers:     headers,
		},
	)
	if err != nil {
		return nil, err
	}

	// First extract the Stytch data.
	var retVal totps.AuthenticateResponse
	if err := json.Unmarshal(b, &retVal); err != nil {
		return nil, fmt.Errorf("unmarshal totps.AuthenticateResponse: %w", err)
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

// RecoveryCodes: Retrieve the recovery codes for a TOTP instance tied to a User.
func (c *TOTPsClient) RecoveryCodes(
	ctx context.Context,
	body *totps.RecoveryCodesParams,
) (*totps.RecoveryCodesResponse, error) {
	var jsonBody []byte
	var err error
	if body != nil {
		jsonBody, err = json.Marshal(body)
		if err != nil {
			return nil, stytcherror.NewClientLibraryError("error marshaling request body")
		}
	}

	headers := make(map[string][]string)

	var retVal totps.RecoveryCodesResponse
	err = c.C.NewRequest(
		ctx,
		stytch.RequestParams{
			Method:      "POST",
			Path:        "/v1/totps/recovery_codes",
			QueryParams: nil,
			Body:        jsonBody,
			V:           &retVal,
			Headers:     headers,
		},
	)
	return &retVal, err
}

// Recover: Authenticate a recovery code for a TOTP instance.
func (c *TOTPsClient) Recover(
	ctx context.Context,
	body *totps.RecoverParams,
) (*totps.RecoverResponse, error) {
	var jsonBody []byte
	var err error
	if body != nil {
		jsonBody, err = json.Marshal(body)
		if err != nil {
			return nil, stytcherror.NewClientLibraryError("error marshaling request body")
		}
	}

	headers := make(map[string][]string)

	var retVal totps.RecoverResponse
	err = c.C.NewRequest(
		ctx,
		stytch.RequestParams{
			Method:      "POST",
			Path:        "/v1/totps/recover",
			QueryParams: nil,
			Body:        jsonBody,
			V:           &retVal,
			Headers:     headers,
		},
	)
	return &retVal, err
}
