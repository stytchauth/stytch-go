package b2b

// !!!
// WARNING: This file is autogenerated
// Only modify code within MANUAL() sections
// or your changes may be overwritten later!
// !!!

import (
	"context"
	"encoding/json"

	"github.com/stytchauth/stytch-go/v15/stytch"
	"github.com/stytchauth/stytch-go/v15/stytch/b2b/passwords/discovery/email"
	"github.com/stytchauth/stytch-go/v15/stytch/stytcherror"
)

type PasswordsDiscoveryEmailClient struct {
	C stytch.Client
}

func NewPasswordsDiscoveryEmailClient(c stytch.Client) *PasswordsDiscoveryEmailClient {
	return &PasswordsDiscoveryEmailClient{
		C: c,
	}
}

func (c *PasswordsDiscoveryEmailClient) ResetStart(
	ctx context.Context,
	body *email.ResetStartParams,
) (*email.ResetStartResponse, error) {
	var jsonBody []byte
	var err error
	if body != nil {
		jsonBody, err = json.Marshal(body)
		if err != nil {
			return nil, stytcherror.NewClientLibraryError("error marshaling request body")
		}
	}

	headers := make(map[string][]string)

	var retVal email.ResetStartResponse
	err = c.C.NewRequest(
		ctx,
		"POST",
		"/v1/b2b/passwords/discovery/email/reset/start",
		nil,
		jsonBody,
		&retVal,
		headers,
	)
	return &retVal, err
}

func (c *PasswordsDiscoveryEmailClient) Reset(
	ctx context.Context,
	body *email.ResetParams,
) (*email.ResetResponse, error) {
	var jsonBody []byte
	var err error
	if body != nil {
		jsonBody, err = json.Marshal(body)
		if err != nil {
			return nil, stytcherror.NewClientLibraryError("error marshaling request body")
		}
	}

	headers := make(map[string][]string)

	var retVal email.ResetResponse
	err = c.C.NewRequest(
		ctx,
		"POST",
		"/v1/b2b/passwords/discovery/email/reset",
		nil,
		jsonBody,
		&retVal,
		headers,
	)
	return &retVal, err
}
