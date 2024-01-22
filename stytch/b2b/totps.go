package b2b

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
	"github.com/stytchauth/stytch-go/v11/stytch"
	"github.com/stytchauth/stytch-go/v11/stytch/b2b/totps"
	"github.com/stytchauth/stytch-go/v11/stytch/stytcherror"
)

type TOTPsClient struct {
	C stytch.Client
}

func NewTOTPsClient(c stytch.Client) *TOTPsClient {
	return &TOTPsClient{
		C: c,
	}
}

// Create a new TOTP instance for a Member. The Member can use the authenticator application of their
// choice to scan the QR code or enter the secret.
//
// Passing an intermediate session token, session token, or session JWT is not required, but if passed must
// match the Member ID passed.
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
		"POST",
		"/v1/b2b/totp",
		nil,
		jsonBody,
		&retVal,
		headers,
	)
	return &retVal, err
}

// Authenticate a Member provided TOTP.
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
		"POST",
		"/v1/b2b/totp/authenticate",
		nil,
		jsonBody,
		&retVal,
		headers,
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
		"POST",
		"/v1/b2b/totp/authenticate",
		nil,
		jsonBody,
		headers,
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
		*m = retVal.MemberSession.CustomClaims
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

	err = decoder.Decode(retVal.MemberSession.CustomClaims)
	if err != nil {
		return nil, err
	}

	return &retVal, err
}

// Migrate an existing TOTP instance for a Member. Recovery codes are not required and will be minted for
// the Member if not provided.
func (c *TOTPsClient) Migrate(
	ctx context.Context,
	body *totps.MigrateParams,
) (*totps.MigrateResponse, error) {
	var jsonBody []byte
	var err error
	if body != nil {
		jsonBody, err = json.Marshal(body)
		if err != nil {
			return nil, stytcherror.NewClientLibraryError("error marshaling request body")
		}
	}

	headers := make(map[string][]string)

	var retVal totps.MigrateResponse
	err = c.C.NewRequest(
		ctx,
		"POST",
		"/v1/b2b/totp/migrate",
		nil,
		jsonBody,
		&retVal,
		headers,
	)
	return &retVal, err
}
