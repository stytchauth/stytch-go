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

	"github.com/stytchauth/stytch-go/v12/stytch"
	"github.com/stytchauth/stytch-go/v12/stytch/b2b/scim/connections"
	"github.com/stytchauth/stytch-go/v12/stytch/stytcherror"
)

type SCIMConnectionsClient struct {
	C stytch.Client
}

func NewSCIMConnectionsClient(c stytch.Client) *SCIMConnectionsClient {
	return &SCIMConnectionsClient{
		C: c,
	}
}

// Update a SCIM Connection. /%}
func (c *SCIMConnectionsClient) Update(
	ctx context.Context,
	body *connections.UpdateParams,
) (*connections.UpdateResponse, error) {
	var jsonBody []byte
	var err error
	if body != nil {
		jsonBody, err = json.Marshal(body)
		if err != nil {
			return nil, stytcherror.NewClientLibraryError("error marshaling request body")
		}
	}

	headers := make(map[string][]string)

	var retVal connections.UpdateResponse
	err = c.C.NewRequest(
		ctx,
		"PUT",
		fmt.Sprintf("/v1/b2b/scim/%s/connections/%s", body.OrganizationID, body.ConnectionID),
		nil,
		jsonBody,
		&retVal,
		headers,
	)
	return &retVal, err
}

// Delete: Deletes a SCIM Connection. /%}
func (c *SCIMConnectionsClient) Delete(
	ctx context.Context,
	body *connections.DeleteParams,
) (*connections.DeleteResponse, error) {
	headers := make(map[string][]string)

	var retVal connections.DeleteResponse
	err := c.C.NewRequest(
		ctx,
		"DELETE",
		fmt.Sprintf("/v1/b2b/scim/%s/connections/%s", body.OrganizationID, body.ConnectionID),
		nil,
		nil,
		&retVal,
		headers,
	)
	return &retVal, err
}

// RotateStart: Start a SCIM token rotation. /%}
func (c *SCIMConnectionsClient) RotateStart(
	ctx context.Context,
	body *connections.RotateStartParams,
) (*connections.RotateStartResponse, error) {
	var jsonBody []byte
	var err error
	if body != nil {
		jsonBody, err = json.Marshal(body)
		if err != nil {
			return nil, stytcherror.NewClientLibraryError("error marshaling request body")
		}
	}

	headers := make(map[string][]string)

	var retVal connections.RotateStartResponse
	err = c.C.NewRequest(
		ctx,
		"POST",
		fmt.Sprintf("/v1/b2b/scim/%s/connections/%s/rotate/start", body.OrganizationID, body.ConnectionID),
		nil,
		jsonBody,
		&retVal,
		headers,
	)
	return &retVal, err
}

// RotateComplete: Completes a SCIM token rotation. This will complete the current token rotation process
// and update the active token to be the new token supplied in the
// [start SCIM token rotation](https://stytch.com/docs/b2b/api/scim-rotate-token-start) response. /%}
func (c *SCIMConnectionsClient) RotateComplete(
	ctx context.Context,
	body *connections.RotateCompleteParams,
) (*connections.RotateCompleteResponse, error) {
	var jsonBody []byte
	var err error
	if body != nil {
		jsonBody, err = json.Marshal(body)
		if err != nil {
			return nil, stytcherror.NewClientLibraryError("error marshaling request body")
		}
	}

	headers := make(map[string][]string)

	var retVal connections.RotateCompleteResponse
	err = c.C.NewRequest(
		ctx,
		"POST",
		fmt.Sprintf("/v1/b2b/scim/%s/connections/%s/rotate/complete", body.OrganizationID, body.ConnectionID),
		nil,
		jsonBody,
		&retVal,
		headers,
	)
	return &retVal, err
}

// RotateCancel: Cancel a SCIM token rotation. This will cancel the current token rotation process, keeping
// the original token active. /%}
func (c *SCIMConnectionsClient) RotateCancel(
	ctx context.Context,
	body *connections.RotateCancelParams,
) (*connections.RotateCancelResponse, error) {
	var jsonBody []byte
	var err error
	if body != nil {
		jsonBody, err = json.Marshal(body)
		if err != nil {
			return nil, stytcherror.NewClientLibraryError("error marshaling request body")
		}
	}

	headers := make(map[string][]string)

	var retVal connections.RotateCancelResponse
	err = c.C.NewRequest(
		ctx,
		"POST",
		fmt.Sprintf("/v1/b2b/scim/%s/connections/%s/rotate/cancel", body.OrganizationID, body.ConnectionID),
		nil,
		jsonBody,
		&retVal,
		headers,
	)
	return &retVal, err
}

// Create a new SCIM Connection. /%}
func (c *SCIMConnectionsClient) Create(
	ctx context.Context,
	body *connections.CreateParams,
) (*connections.CreateResponse, error) {
	var jsonBody []byte
	var err error
	if body != nil {
		jsonBody, err = json.Marshal(body)
		if err != nil {
			return nil, stytcherror.NewClientLibraryError("error marshaling request body")
		}
	}

	headers := make(map[string][]string)

	var retVal connections.CreateResponse
	err = c.C.NewRequest(
		ctx,
		"POST",
		fmt.Sprintf("/v1/b2b/scim/%s/connections", body.OrganizationID),
		nil,
		jsonBody,
		&retVal,
		headers,
	)
	return &retVal, err
}

// Get SCIM Connections. /%}
func (c *SCIMConnectionsClient) Get(
	ctx context.Context,
	body *connections.GetParams,
) (*connections.GetResponse, error) {
	headers := make(map[string][]string)

	var retVal connections.GetResponse
	err := c.C.NewRequest(
		ctx,
		"GET",
		fmt.Sprintf("/v1/b2b/scim/%s/connections", body.OrganizationID),
		nil,
		nil,
		&retVal,
		headers,
	)
	return &retVal, err
}
