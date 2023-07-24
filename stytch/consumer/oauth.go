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
	"github.com/stytchauth/stytch-go/v10/stytch"
	"github.com/stytchauth/stytch-go/v10/stytch/consumer/oauth"
	"github.com/stytchauth/stytch-go/v10/stytch/stytcherror"
)

type OAuthClient struct {
	C stytch.Client
}

func NewOAuthClient(c stytch.Client) *OAuthClient {
	return &OAuthClient{
		C: c,
	}
}

// Attach: Generate an OAuth Attach Token to pre-associate an OAuth flow with an existing Stytch User. Pass
// the returned `oauth_attach_token` to the same provider's OAuth Start endpoint to treat this OAuth flow
// as a login for that user instead of a signup for a new user.
//
// Exactly one of `user_id`, `session_token`, or `session_jwt` must be provided to identify the target
// Stytch User.
//
// This is an optional step in the OAuth flow. Stytch can often determine whether to create a new user or
// log in an existing one based on verified identity provider information. This endpoint is useful for
// cases where we can't, such as missing or unverified provider information.
func (c *OAuthClient) Attach(
	ctx context.Context,
	body *oauth.AttachParams,
) (*oauth.AttachResponse, error) {
	var jsonBody []byte
	var err error
	if body != nil {
		jsonBody, err = json.Marshal(body)
		if err != nil {
			return nil, stytcherror.NewClientLibraryError("error marshaling request body")
		}
	}

	var retVal oauth.AttachResponse
	err = c.C.NewRequest(
		ctx,
		"POST",
		"/v1/oauth/attach",
		nil,
		jsonBody,
		&retVal,
	)
	return &retVal, err
}

// Authenticate a User given a `token`. This endpoint verifies that the user completed the OAuth flow by
// verifying that the token is valid and hasn't expired. To initiate a Stytch session for the user while
// authenticating their OAuth token, include `session_duration_minutes`; a session with the identity
// provider, e.g. Google or Facebook, will always be initiated upon successful authentication.
func (c *OAuthClient) Authenticate(
	ctx context.Context,
	body *oauth.AuthenticateParams,
) (*oauth.AuthenticateResponse, error) {
	var jsonBody []byte
	var err error
	if body != nil {
		jsonBody, err = json.Marshal(body)
		if err != nil {
			return nil, stytcherror.NewClientLibraryError("error marshaling request body")
		}
	}

	var retVal oauth.AuthenticateResponse
	err = c.C.NewRequest(
		ctx,
		"POST",
		"/v1/oauth/authenticate",
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
func (c *OAuthClient) AuthenticateWithClaims(
	ctx context.Context,
	body *oauth.AuthenticateParams,
	claims any,
) (*oauth.AuthenticateResponse, error) {
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
		"/v1/oauth/authenticate",
		nil,
		jsonBody,
	)
	if err != nil {
		return nil, err
	}

	// First extract the Stytch data.
	var retVal oauth.AuthenticateResponse
	if err := json.Unmarshal(b, &retVal); err != nil {
		return nil, fmt.Errorf("unmarshal oauth.AuthenticateResponse: %w", err)
	}

	if claims == nil {
		return &retVal, nil
	}

	if m, ok := claims.(*map[string]any); ok {
		*m = retVal.UserSession.CustomClaims
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

	err = decoder.Decode(retVal.UserSession.CustomClaims)
	if err != nil {
		return nil, err
	}

	return &retVal, err
}
