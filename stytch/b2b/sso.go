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
	"github.com/stytchauth/stytch-go/v12/stytch"
	"github.com/stytchauth/stytch-go/v12/stytch/b2b/sso"
	"github.com/stytchauth/stytch-go/v12/stytch/stytcherror"
)

type SSOClient struct {
	C    stytch.Client
	OIDC *SSOOIDCClient
	SAML *SSOSAMLClient
}

func NewSSOClient(c stytch.Client) *SSOClient {
	return &SSOClient{
		C: c,

		OIDC: NewSSOOIDCClient(c),
		SAML: NewSSOSAMLClient(c),
	}
}

// GetConnections: Get all SSO Connections owned by the organization. /%}
func (c *SSOClient) GetConnections(
	ctx context.Context,
	body *sso.GetConnectionsParams,
	methodOptions ...*sso.GetConnectionsRequestOptions,
) (*sso.GetConnectionsResponse, error) {
	headers := make(map[string][]string)
	for _, methodOption := range methodOptions {
		headers = methodOption.AddHeaders(headers)
	}

	var retVal sso.GetConnectionsResponse
	err := c.C.NewRequest(
		ctx,
		"GET",
		fmt.Sprintf("/v1/b2b/sso/%s", body.OrganizationID),
		nil,
		nil,
		&retVal,
		headers,
	)
	return &retVal, err
}

// DeleteConnection: Delete an existing SSO connection. /%}
func (c *SSOClient) DeleteConnection(
	ctx context.Context,
	body *sso.DeleteConnectionParams,
	methodOptions ...*sso.DeleteConnectionRequestOptions,
) (*sso.DeleteConnectionResponse, error) {
	headers := make(map[string][]string)
	for _, methodOption := range methodOptions {
		headers = methodOption.AddHeaders(headers)
	}

	var retVal sso.DeleteConnectionResponse
	err := c.C.NewRequest(
		ctx,
		"DELETE",
		fmt.Sprintf("/v1/b2b/sso/%s/connections/%s", body.OrganizationID, body.ConnectionID),
		nil,
		nil,
		&retVal,
		headers,
	)
	return &retVal, err
}

// Authenticate a user given a token.
// This endpoint verifies that the user completed the SSO Authentication flow by verifying that the token
// is valid and hasn't expired.
// Provide the `session_duration_minutes` parameter to set the lifetime of the session.
// If the `session_duration_minutes` parameter is not specified, a Stytch session will be created with a 60
// minute duration.
// To link this authentication event to an existing Stytch session, include either the `session_token` or
// `session_jwt` param.
//
// If the Member is required to complete MFA to log in to the Organization, the returned value of
// `member_authenticated` will be `false`, and an `intermediate_session_token` will be returned.
// The `intermediate_session_token` can be passed into the
// [OTP SMS Authenticate endpoint](https://stytch.com/docs/b2b/api/authenticate-otp-sms) to complete the
// MFA step and acquire a full member session.
// The `session_duration_minutes` and `session_custom_claims` parameters will be ignored.
//
// If a valid `session_token` or `session_jwt` is passed in, the Member will not be required to complete an
// MFA step.
func (c *SSOClient) Authenticate(
	ctx context.Context,
	body *sso.AuthenticateParams,
) (*sso.AuthenticateResponse, error) {
	var jsonBody []byte
	var err error
	if body != nil {
		jsonBody, err = json.Marshal(body)
		if err != nil {
			return nil, stytcherror.NewClientLibraryError("error marshaling request body")
		}
	}

	headers := make(map[string][]string)

	var retVal sso.AuthenticateResponse
	err = c.C.NewRequest(
		ctx,
		"POST",
		"/v1/b2b/sso/authenticate",
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
func (c *SSOClient) AuthenticateWithClaims(
	ctx context.Context,
	body *sso.AuthenticateParams,
	claims any,
) (*sso.AuthenticateResponse, error) {
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
		"/v1/b2b/sso/authenticate",
		nil,
		jsonBody,
		headers,
	)
	if err != nil {
		return nil, err
	}

	// First extract the Stytch data.
	var retVal sso.AuthenticateResponse
	if err := json.Unmarshal(b, &retVal); err != nil {
		return nil, fmt.Errorf("unmarshal sso.AuthenticateResponse: %w", err)
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
