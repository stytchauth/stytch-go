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
	"github.com/stytchauth/stytch-go/v15/stytch"
	"github.com/stytchauth/stytch-go/v15/stytch/consumer/webauthn"
	"github.com/stytchauth/stytch-go/v15/stytch/stytcherror"
)

type WebAuthnClient struct {
	C stytch.Client
}

func NewWebAuthnClient(c stytch.Client) *WebAuthnClient {
	return &WebAuthnClient{
		C: c,
	}
}

// RegisterStart: Initiate the process of creating a new Passkey or WebAuthn registration.
//
// To optimize for Passkeys, set the `return_passkey_credential_options` field to `true`.
//
// After calling this endpoint, the browser will need to call
// [navigator.credentials.create()](https://www.w3.org/TR/webauthn-2/#sctn-createCredential) with the data
// from
// [public_key_credential_creation_options](https://w3c.github.io/webauthn/#dictionary-makecredentialoptions)
// passed to the [navigator.credentials.create()](https://www.w3.org/TR/webauthn-2/#sctn-createCredential)
// request via the public key argument. We recommend using the `create()` wrapper provided by the
// webauthn-json library.
//
// If you are not using the [webauthn-json](https://github.com/github/webauthn-json) library, the
// `public_key_credential_creation_options` will need to be converted to a suitable public key by
// unmarshalling the JSON, base64 decoding the user ID field, and converting user ID and the challenge
// fields into an array buffer.
func (c *WebAuthnClient) RegisterStart(
	ctx context.Context,
	body *webauthn.RegisterStartParams,
) (*webauthn.RegisterStartResponse, error) {
	var jsonBody []byte
	var err error
	if body != nil {
		jsonBody, err = json.Marshal(body)
		if err != nil {
			return nil, stytcherror.NewClientLibraryError("error marshaling request body")
		}
	}

	headers := make(map[string][]string)

	var retVal webauthn.RegisterStartResponse
	err = c.C.NewRequest(
		ctx,
		"POST",
		"/v1/webauthn/register/start",
		nil,
		jsonBody,
		&retVal,
		headers,
	)
	return &retVal, err
}

// Register: Complete the creation of a WebAuthn registration by passing the response from the
// [navigator.credentials.create()](https://www.w3.org/TR/webauthn-2/#sctn-createCredential) request to
// this endpoint as the `public_key_credential` parameter.
//
// If the [webauthn-json](https://github.com/github/webauthn-json) library's `create()` method was used,
// the response can be passed directly to the
// [register endpoint](https://stytch.com/docs/api/webauthn-register). If not, some fields (the client data
// and the attestation object) from the
// [navigator.credentials.create()](https://www.w3.org/TR/webauthn-2/#sctn-createCredential) response will
// need to be converted from array buffers to strings and marshalled into JSON.
func (c *WebAuthnClient) Register(
	ctx context.Context,
	body *webauthn.RegisterParams,
) (*webauthn.RegisterResponse, error) {
	var jsonBody []byte
	var err error
	if body != nil {
		jsonBody, err = json.Marshal(body)
		if err != nil {
			return nil, stytcherror.NewClientLibraryError("error marshaling request body")
		}
	}

	headers := make(map[string][]string)

	var retVal webauthn.RegisterResponse
	err = c.C.NewRequest(
		ctx,
		"POST",
		"/v1/webauthn/register",
		nil,
		jsonBody,
		&retVal,
		headers,
	)
	return &retVal, err
}

// AuthenticateStart: Initiate the authentication of a Passkey or WebAuthn registration.
//
// To optimize for Passkeys, set the `return_passkey_credential_options` field to `true`.
//
// After calling this endpoint, the browser will need to call
// [navigator.credentials.get()](https://www.w3.org/TR/webauthn-2/#sctn-getAssertion) with the data from
// `public_key_credential_request_options` passed to the
// [navigator.credentials.get()](https://www.w3.org/TR/webauthn-2/#sctn-getAssertion) request via the
// public key argument. We recommend using the `get()` wrapper provided by the webauthn-json library.
//
// If you are not using the [webauthn-json](https://github.com/github/webauthn-json) library, `the
// public_key_credential_request_options` will need to be converted to a suitable public key by
// unmarshalling the JSON and converting some the fields to array buffers.
func (c *WebAuthnClient) AuthenticateStart(
	ctx context.Context,
	body *webauthn.AuthenticateStartParams,
) (*webauthn.AuthenticateStartResponse, error) {
	var jsonBody []byte
	var err error
	if body != nil {
		jsonBody, err = json.Marshal(body)
		if err != nil {
			return nil, stytcherror.NewClientLibraryError("error marshaling request body")
		}
	}

	headers := make(map[string][]string)

	var retVal webauthn.AuthenticateStartResponse
	err = c.C.NewRequest(
		ctx,
		"POST",
		"/v1/webauthn/authenticate/start",
		nil,
		jsonBody,
		&retVal,
		headers,
	)
	return &retVal, err
}

// Authenticate: Complete the authentication of a Passkey or WebAuthn registration by passing the response
// from the [navigator.credentials.get()](https://www.w3.org/TR/webauthn-2/#sctn-getAssertion) request to
// the authenticate endpoint.
//
// If the [webauthn-json](https://github.com/github/webauthn-json) library's `get()` method was used, the
// response can be passed directly to the
// [authenticate endpoint](https://stytch.com/docs/api/webauthn-authenticate). If not some fields from the
// [navigator.credentials.get()](https://www.w3.org/TR/webauthn-2/#sctn-getAssertion) response will need to
// be converted from array buffers to strings and marshalled into JSON.
func (c *WebAuthnClient) Authenticate(
	ctx context.Context,
	body *webauthn.AuthenticateParams,
) (*webauthn.AuthenticateResponse, error) {
	var jsonBody []byte
	var err error
	if body != nil {
		jsonBody, err = json.Marshal(body)
		if err != nil {
			return nil, stytcherror.NewClientLibraryError("error marshaling request body")
		}
	}

	headers := make(map[string][]string)

	var retVal webauthn.AuthenticateResponse
	err = c.C.NewRequest(
		ctx,
		"POST",
		"/v1/webauthn/authenticate",
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
func (c *WebAuthnClient) AuthenticateWithClaims(
	ctx context.Context,
	body *webauthn.AuthenticateParams,
	claims any,
) (*webauthn.AuthenticateResponse, error) {
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
		"/v1/webauthn/authenticate",
		nil,
		jsonBody,
		headers,
	)
	if err != nil {
		return nil, err
	}

	// First extract the Stytch data.
	var retVal webauthn.AuthenticateResponse
	if err := json.Unmarshal(b, &retVal); err != nil {
		return nil, fmt.Errorf("unmarshal webauthn.AuthenticateResponse: %w", err)
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

// Update: Updates a Passkey or WebAuthn registration.
func (c *WebAuthnClient) Update(
	ctx context.Context,
	body *webauthn.UpdateParams,
) (*webauthn.UpdateResponse, error) {
	var jsonBody []byte
	var err error
	if body != nil {
		jsonBody, err = json.Marshal(body)
		if err != nil {
			return nil, stytcherror.NewClientLibraryError("error marshaling request body")
		}
	}

	headers := make(map[string][]string)

	var retVal webauthn.UpdateResponse
	err = c.C.NewRequest(
		ctx,
		"PUT",
		fmt.Sprintf("/v1/webauthn/%s", body.WebAuthnRegistrationID),
		nil,
		jsonBody,
		&retVal,
		headers,
	)
	return &retVal, err
}
