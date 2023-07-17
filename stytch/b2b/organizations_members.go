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

	"github.com/stytchauth/stytch-go/v9/stytch"
	"github.com/stytchauth/stytch-go/v9/stytch/b2b/organizations/members"
	"github.com/stytchauth/stytch-go/v9/stytch/stytcherror"
)

type OrganizationsMembersClient struct {
	C *stytch.Client
}

func NewOrganizationsMembersClient(c *stytch.Client) *OrganizationsMembersClient {
	return &OrganizationsMembersClient{
		C: c,
	}
}

// Update: Updates a Member specified by `organization_id` and `member_id`.
func (c *OrganizationsMembersClient) Update(
	ctx context.Context,
	body *members.UpdateParams,
) (*members.UpdateResponse, error) {
	var jsonBody []byte
	var err error
	if body != nil {
		jsonBody, err = json.Marshal(body)
		if err != nil {
			return nil, stytcherror.NewClientLibraryError("error marshaling request body")
		}
	}

	var retVal members.UpdateResponse
	err = c.C.NewRequest(
		ctx,
		"PUT",
		fmt.Sprintf("/v1/b2b/organizations/%s/members/%s", body.OrganizationID, body.MemberID),
		nil,
		jsonBody,
		&retVal,
	)
	return &retVal, err
}

// Delete: Deletes a Member specified by `organization_id` and `member_id`.
func (c *OrganizationsMembersClient) Delete(
	ctx context.Context,
	body *members.DeleteParams,
) (*members.DeleteResponse, error) {
	var retVal members.DeleteResponse
	err := c.C.NewRequest(
		ctx,
		"DELETE",
		fmt.Sprintf("/v1/b2b/organizations/%s/members/%s", body.OrganizationID, body.MemberID),
		nil,
		nil,
		&retVal,
	)
	return &retVal, err
}

// Search for Members within specified Organizations. An array with at least one `organization_id` is
// required. Submitting an empty `query` returns all Members within the specified Organizations.
//
// *All fuzzy search filters require a minimum of three characters.
func (c *OrganizationsMembersClient) Search(
	ctx context.Context,
	body *members.SearchParams,
) (*members.SearchResponse, error) {
	var jsonBody []byte
	var err error
	if body != nil {
		jsonBody, err = json.Marshal(body)
		if err != nil {
			return nil, stytcherror.NewClientLibraryError("error marshaling request body")
		}
	}

	var retVal members.SearchResponse
	err = c.C.NewRequest(
		ctx,
		"POST",
		"/v1/b2b/organizations/members/search",
		nil,
		jsonBody,
		&retVal,
	)
	return &retVal, err
}

// DeletePassword: Delete a Member's password.
func (c *OrganizationsMembersClient) DeletePassword(
	ctx context.Context,
	body *members.DeletePasswordParams,
) (*members.DeletePasswordResponse, error) {
	var retVal members.DeletePasswordResponse
	err := c.C.NewRequest(
		ctx,
		"DELETE",
		fmt.Sprintf("/v1/b2b/organizations/%s/members/passwords/%s", body.OrganizationID, body.MemberPasswordID),
		nil,
		nil,
		&retVal,
	)
	return &retVal, err
}

// Create: Creates a Member. An `organization_id` and `email_address` are required.
func (c *OrganizationsMembersClient) Create(
	ctx context.Context,
	body *members.CreateParams,
) (*members.CreateResponse, error) {
	var jsonBody []byte
	var err error
	if body != nil {
		jsonBody, err = json.Marshal(body)
		if err != nil {
			return nil, stytcherror.NewClientLibraryError("error marshaling request body")
		}
	}

	var retVal members.CreateResponse
	err = c.C.NewRequest(
		ctx,
		"POST",
		fmt.Sprintf("/v1/b2b/organizations/%s/members", body.OrganizationID),
		nil,
		jsonBody,
		&retVal,
	)
	return &retVal, err
}

// Get a Member by `member_id` or `email_address`.
func (c *OrganizationsMembersClient) Get(
	ctx context.Context,
	body *members.GetParams,
) (*members.GetResponse, error) {
	queryParams := make(map[string]string)
	if body != nil {
		queryParams["member_id"] = body.MemberID
		queryParams["email_address"] = body.EmailAddress
	}

	var retVal members.GetResponse
	err := c.C.NewRequest(
		ctx,
		"GET",
		fmt.Sprintf("/v1/b2b/organizations/%s/member", body.OrganizationID),
		queryParams,
		nil,
		&retVal,
	)
	return &retVal, err
}
