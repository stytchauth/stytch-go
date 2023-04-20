package oauth

// !!!
// WARNING: This file is autogenerated
// Only modify code within MANUAL() sections
// or your changes may be overwritten later!
// !!!

import (
	"context"
	"encoding/json"

	"github.com/stytchauth/stytch-go/v7/stytch"
	"github.com/stytchauth/stytch-go/v7/stytch/stytcherror"
)

type Client struct {
	C *stytch.Client
}

const subURL = "oauth"

func (c *Client) Attach(
	ctx context.Context,
	body *stytch.AttachParams,
) (*stytch.AttachResponse, error) {
	path := subURL + "//v1/oauth/attach"

	var retVal stytch.AttachResponse
	var jsonBody []byte
	var err error
	if body != nil {
		jsonBody, err = json.Marshal(body)
		if err != nil {
			return nil, stytcherror.NewClientLibraryError("Oops, something seems to have gone wrong marshalling the request body")
		}
	}

	err = c.C.NewRequest(ctx, "POST", path, nil, jsonBody, &retVal)

	return &retVal, err
}

func (c *Client) Authenticate(
	ctx context.Context,
	body *stytch.AuthenticateParams,
) (*stytch.AuthenticateResponse, error) {
	path := subURL + "//v1/oauth/authenticate"

	var retVal stytch.AuthenticateResponse
	var jsonBody []byte
	var err error
	if body != nil {
		jsonBody, err = json.Marshal(body)
		if err != nil {
			return nil, stytcherror.NewClientLibraryError("Oops, something seems to have gone wrong marshalling the request body")
		}
	}

	err = c.C.NewRequest(ctx, "POST", path, nil, jsonBody, &retVal)

	return &retVal, err
}
