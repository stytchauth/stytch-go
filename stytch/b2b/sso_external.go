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

	"github.com/stytchauth/stytch-go/v16/stytch"
	"github.com/stytchauth/stytch-go/v16/stytch/b2b/sso/external"
	"github.com/stytchauth/stytch-go/v16/stytch/stytcherror"
)

type SSOExternalClient struct {
	C stytch.Client
}

func NewSSOExternalClient(c stytch.Client) *SSOExternalClient {
	return &SSOExternalClient{
		C: c,
	}
}

// CreateConnection: Create a new External SSO Connection.
func (c *SSOExternalClient) CreateConnection(
	ctx context.Context,
	body *external.CreateConnectionParams,
	methodOptions ...*external.CreateConnectionRequestOptions,
) (*external.CreateConnectionResponse, error) {
	var jsonBody []byte
	var err error
	if body != nil {
		jsonBody, err = json.Marshal(body)
		if err != nil {
			return nil, stytcherror.NewClientLibraryError("error marshaling request body")
		}
	}

	headers := make(map[string][]string)
	for _, methodOption := range methodOptions {
		headers = methodOption.AddHeaders(headers)
	}

	var retVal external.CreateConnectionResponse
	err = c.C.NewRequest(
		ctx,
		stytch.RequestParams{
			Method:      "POST",
			Path:        fmt.Sprintf("/v1/b2b/sso/external/%s", body.OrganizationID),
			QueryParams: nil,
			Body:        jsonBody,
			V:           &retVal,
			Headers:     headers,
		},
	)
	return &retVal, err
}

// UpdateConnection: Updates an existing External SSO connection.
func (c *SSOExternalClient) UpdateConnection(
	ctx context.Context,
	body *external.UpdateConnectionParams,
	methodOptions ...*external.UpdateConnectionRequestOptions,
) (*external.UpdateConnectionResponse, error) {
	var jsonBody []byte
	var err error
	if body != nil {
		jsonBody, err = json.Marshal(body)
		if err != nil {
			return nil, stytcherror.NewClientLibraryError("error marshaling request body")
		}
	}

	headers := make(map[string][]string)
	for _, methodOption := range methodOptions {
		headers = methodOption.AddHeaders(headers)
	}

	var retVal external.UpdateConnectionResponse
	err = c.C.NewRequest(
		ctx,
		stytch.RequestParams{
			Method:      "PUT",
			Path:        fmt.Sprintf("/v1/b2b/sso/external/%s/connections/%s", body.OrganizationID, body.ConnectionID),
			QueryParams: nil,
			Body:        jsonBody,
			V:           &retVal,
			Headers:     headers,
		},
	)
	return &retVal, err
}
