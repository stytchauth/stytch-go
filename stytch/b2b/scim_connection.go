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
	"strconv"

	"github.com/stytchauth/stytch-go/v16/stytch"
	"github.com/stytchauth/stytch-go/v16/stytch/b2b/scim/connection"
	"github.com/stytchauth/stytch-go/v16/stytch/stytcherror"
)

type SCIMConnectionClient struct {
	C stytch.Client
}

func NewSCIMConnectionClient(c stytch.Client) *SCIMConnectionClient {
	return &SCIMConnectionClient{
		C: c,
	}
}

// Update a SCIM Connection.
func (c *SCIMConnectionClient) Update(
	ctx context.Context,
	body *connection.UpdateParams,
	methodOptions ...*connection.UpdateRequestOptions,
) (*connection.UpdateResponse, error) {
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

	var retVal connection.UpdateResponse
	err = c.C.NewRequest(
		ctx,
		stytch.RequestParams{
			Method:      "PUT",
			Path:        fmt.Sprintf("/v1/b2b/scim/%s/connection/%s", body.OrganizationID, body.ConnectionID),
			QueryParams: nil,
			Body:        jsonBody,
			V:           &retVal,
			Headers:     headers,
		},
	)
	return &retVal, err
}

// Delete: Deletes a SCIM Connection.
func (c *SCIMConnectionClient) Delete(
	ctx context.Context,
	body *connection.DeleteParams,
	methodOptions ...*connection.DeleteRequestOptions,
) (*connection.DeleteResponse, error) {
	headers := make(map[string][]string)
	for _, methodOption := range methodOptions {
		headers = methodOption.AddHeaders(headers)
	}

	var retVal connection.DeleteResponse
	err := c.C.NewRequest(
		ctx,
		stytch.RequestParams{
			Method:      "DELETE",
			Path:        fmt.Sprintf("/v1/b2b/scim/%s/connection/%s", body.OrganizationID, body.ConnectionID),
			QueryParams: nil,
			Body:        nil,
			V:           &retVal,
			Headers:     headers,
		},
	)
	return &retVal, err
}

// RotateStart: Start a SCIM token rotation.
func (c *SCIMConnectionClient) RotateStart(
	ctx context.Context,
	body *connection.RotateStartParams,
	methodOptions ...*connection.RotateStartRequestOptions,
) (*connection.RotateStartResponse, error) {
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

	var retVal connection.RotateStartResponse
	err = c.C.NewRequest(
		ctx,
		stytch.RequestParams{
			Method:      "POST",
			Path:        fmt.Sprintf("/v1/b2b/scim/%s/connection/%s/rotate/start", body.OrganizationID, body.ConnectionID),
			QueryParams: nil,
			Body:        jsonBody,
			V:           &retVal,
			Headers:     headers,
		},
	)
	return &retVal, err
}

// RotateComplete: Completes a SCIM token rotation. This will complete the current token rotation process
// and update the active token to be the new token supplied in the
// [start SCIM token rotation](https://stytch.com/docs/b2b/api/scim-rotate-token-start) response.
func (c *SCIMConnectionClient) RotateComplete(
	ctx context.Context,
	body *connection.RotateCompleteParams,
	methodOptions ...*connection.RotateCompleteRequestOptions,
) (*connection.RotateCompleteResponse, error) {
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

	var retVal connection.RotateCompleteResponse
	err = c.C.NewRequest(
		ctx,
		stytch.RequestParams{
			Method:      "POST",
			Path:        fmt.Sprintf("/v1/b2b/scim/%s/connection/%s/rotate/complete", body.OrganizationID, body.ConnectionID),
			QueryParams: nil,
			Body:        jsonBody,
			V:           &retVal,
			Headers:     headers,
		},
	)
	return &retVal, err
}

// RotateCancel: Cancel a SCIM token rotation. This will cancel the current token rotation process, keeping
// the original token active.
func (c *SCIMConnectionClient) RotateCancel(
	ctx context.Context,
	body *connection.RotateCancelParams,
	methodOptions ...*connection.RotateCancelRequestOptions,
) (*connection.RotateCancelResponse, error) {
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

	var retVal connection.RotateCancelResponse
	err = c.C.NewRequest(
		ctx,
		stytch.RequestParams{
			Method:      "POST",
			Path:        fmt.Sprintf("/v1/b2b/scim/%s/connection/%s/rotate/cancel", body.OrganizationID, body.ConnectionID),
			QueryParams: nil,
			Body:        jsonBody,
			V:           &retVal,
			Headers:     headers,
		},
	)
	return &retVal, err
}

// GetGroups: Gets a paginated list of all SCIM Groups associated with a given Connection.
func (c *SCIMConnectionClient) GetGroups(
	ctx context.Context,
	body *connection.GetGroupsParams,
	methodOptions ...*connection.GetGroupsRequestOptions,
) (*connection.GetGroupsResponse, error) {
	queryParams := make(map[string]string)
	if body != nil {
		queryParams["cursor"] = body.Cursor
		queryParams["limit"] = strconv.FormatUint(uint64(body.Limit), 10)
	}

	headers := make(map[string][]string)
	for _, methodOption := range methodOptions {
		headers = methodOption.AddHeaders(headers)
	}

	var retVal connection.GetGroupsResponse
	err := c.C.NewRequest(
		ctx,
		stytch.RequestParams{
			Method:      "GET",
			Path:        fmt.Sprintf("/v1/b2b/scim/%s/connection/%s", body.OrganizationID, body.ConnectionID),
			QueryParams: queryParams,
			Body:        nil,
			V:           &retVal,
			Headers:     headers,
		},
	)
	return &retVal, err
}

// Create a new SCIM Connection.
func (c *SCIMConnectionClient) Create(
	ctx context.Context,
	body *connection.CreateParams,
	methodOptions ...*connection.CreateRequestOptions,
) (*connection.CreateResponse, error) {
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

	var retVal connection.CreateResponse
	err = c.C.NewRequest(
		ctx,
		stytch.RequestParams{
			Method:      "POST",
			Path:        fmt.Sprintf("/v1/b2b/scim/%s/connection", body.OrganizationID),
			QueryParams: nil,
			Body:        jsonBody,
			V:           &retVal,
			Headers:     headers,
		},
	)
	return &retVal, err
}

// Get SCIM Connection.
func (c *SCIMConnectionClient) Get(
	ctx context.Context,
	body *connection.GetParams,
	methodOptions ...*connection.GetRequestOptions,
) (*connection.GetResponse, error) {
	headers := make(map[string][]string)
	for _, methodOption := range methodOptions {
		headers = methodOption.AddHeaders(headers)
	}

	var retVal connection.GetResponse
	err := c.C.NewRequest(
		ctx,
		stytch.RequestParams{
			Method:      "GET",
			Path:        fmt.Sprintf("/v1/b2b/scim/%s/connection", body.OrganizationID),
			QueryParams: nil,
			Body:        nil,
			V:           &retVal,
			Headers:     headers,
		},
	)
	return &retVal, err
}
