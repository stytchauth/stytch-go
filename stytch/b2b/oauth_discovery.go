package b2b

// !!!
// WARNING: This file is autogenerated
// Only modify code within MANUAL() sections
// or your changes may be overwritten later!
// !!!

import (
	"context"
	"encoding/json"

	"github.com/stytchauth/stytch-go/v9/stytch"
	"github.com/stytchauth/stytch-go/v9/stytch/b2b/oauth/discovery"
	"github.com/stytchauth/stytch-go/v9/stytch/stytcherror"
)

type OAuthDiscoveryClient struct {
	C stytch.Client
}

func NewOAuthDiscoveryClient(c stytch.Client) *OAuthDiscoveryClient {
	return &OAuthDiscoveryClient{
		C: c,
	}
}

// Authenticate: Authenticates the Discovery OAuth token and exchanges it for an Intermediate Session
// Token. Intermediate Session Tokens can be used for various Discovery login flows and are valid for 10
// minutes.
func (c *OAuthDiscoveryClient) Authenticate(
	ctx context.Context,
	body *discovery.AuthenticateParams,
) (*discovery.AuthenticateResponse, error) {
	var jsonBody []byte
	var err error
	if body != nil {
		jsonBody, err = json.Marshal(body)
		if err != nil {
			return nil, stytcherror.NewClientLibraryError("error marshaling request body")
		}
	}

	var retVal discovery.AuthenticateResponse
	err = c.C.NewRequest(
		ctx,
		"POST",
		"/v1/b2b/oauth/discovery/authenticate",
		nil,
		jsonBody,
		&retVal,
	)
	return &retVal, err
}
