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

	"github.com/stytchauth/stytch-go/v11/stytch"
	"github.com/stytchauth/stytch-go/v11/stytch/b2b/organizations"
	"github.com/stytchauth/stytch-go/v11/stytch/stytcherror"
)

type OrganizationsClient struct {
	C       stytch.Client
	Members *OrganizationsMembersClient
}

func NewOrganizationsClient(c stytch.Client) *OrganizationsClient {
	return &OrganizationsClient{
		C: c,

		Members: NewOrganizationsMembersClient(c),
	}
}

// Create: Creates an Organization. An `organization_name` and a unique `organization_slug` are required.
//
// By default, `email_invites` and `sso_jit_provisioning` will be set to `ALL_ALLOWED`, and `mfa_policy`
// will be set to `OPTIONAL` if no Organization authentication settings are explicitly defined in the
// request.
//
// *See the [Organization authentication settings](https://stytch.com/docs/b2b/api/org-auth-settings)
// resource to learn more about fields like `email_jit_provisioning`, `email_invites`,
// `sso_jit_provisioning`, etc., and their behaviors.
func (c *OrganizationsClient) Create(
	ctx context.Context,
	body *organizations.CreateParams,
) (*organizations.CreateResponse, error) {
	var jsonBody []byte
	var err error
	if body != nil {
		jsonBody, err = json.Marshal(body)
		if err != nil {
			return nil, stytcherror.NewClientLibraryError("error marshaling request body")
		}
	}

	headers := make(map[string][]string)

	var retVal organizations.CreateResponse
	err = c.C.NewRequest(
		ctx,
		"POST",
		"/v1/b2b/organizations",
		nil,
		jsonBody,
		&retVal,
		headers,
	)
	return &retVal, err
}

// Get: Returns an Organization specified by `organization_id`.
func (c *OrganizationsClient) Get(
	ctx context.Context,
	body *organizations.GetParams,
) (*organizations.GetResponse, error) {
	headers := make(map[string][]string)

	var retVal organizations.GetResponse
	err := c.C.NewRequest(
		ctx,
		"GET",
		fmt.Sprintf("/v1/b2b/organizations/%s", body.OrganizationID),
		nil,
		nil,
		&retVal,
		headers,
	)
	return &retVal, err
}

// Update: Updates an Organization specified by `organization_id`. An Organization must always have at
// least one auth setting set to either `RESTRICTED` or `ALL_ALLOWED` in order to provision new Members.
//
// *See the [Organization authentication settings](https://stytch.com/docs/b2b/api/org-auth-settings)
// resource to learn more about fields like `email_jit_provisioning`, `email_invites`,
// `sso_jit_provisioning`, etc., and their behaviors.
//
// Our RBAC implementation offers out-of-the-box handling of authorization checks for this endpoint. If you
// pass in
// a header containing a `session_token` or a `session_jwt` for an unexpired Member Session, we will check
// that the
// Member Session has the necessary permissions. The specific permissions needed depend on which of the
// optional fields
// are passed in the request. For example, if the `organization_name` argument is provided, the Member
// Session must have
// permission to perform the `update.info.name` action on the `stytch.organization` Resource.
//
// If the Member Session does not contain a Role that satisfies the requested permissions, or if the
// Member's Organization
// does not match the `organization_id` passed in the request, a 403 error will be thrown. Otherwise, the
// request will
// proceed as normal.
//
// To learn more about our RBAC implementation, see our
// [RBAC guide](https://stytch.com/docs/b2b/guides/rbac/overview).
func (c *OrganizationsClient) Update(
	ctx context.Context,
	body *organizations.UpdateParams,
	methodOptions ...*organizations.UpdateRequestOptions,
) (*organizations.UpdateResponse, error) {
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

	var retVal organizations.UpdateResponse
	err = c.C.NewRequest(
		ctx,
		"PUT",
		fmt.Sprintf("/v1/b2b/organizations/%s", body.OrganizationID),
		nil,
		jsonBody,
		&retVal,
		headers,
	)
	return &retVal, err
}

// Delete: Deletes an Organization specified by `organization_id`. All Members of the Organization will
// also be deleted. /%}
func (c *OrganizationsClient) Delete(
	ctx context.Context,
	body *organizations.DeleteParams,
	methodOptions ...*organizations.DeleteRequestOptions,
) (*organizations.DeleteResponse, error) {
	headers := make(map[string][]string)
	for _, methodOption := range methodOptions {
		headers = methodOption.AddHeaders(headers)
	}

	var retVal organizations.DeleteResponse
	err := c.C.NewRequest(
		ctx,
		"DELETE",
		fmt.Sprintf("/v1/b2b/organizations/%s", body.OrganizationID),
		nil,
		nil,
		&retVal,
		headers,
	)
	return &retVal, err
}

// Search for Organizations. If you send a request with no body params, no filtering will be applied and
// the endpoint will return all Organizations. All fuzzy search filters require a minimum of three
// characters.
func (c *OrganizationsClient) Search(
	ctx context.Context,
	body *organizations.SearchParams,
) (*organizations.SearchResponse, error) {
	var jsonBody []byte
	var err error
	if body != nil {
		jsonBody, err = json.Marshal(body)
		if err != nil {
			return nil, stytcherror.NewClientLibraryError("error marshaling request body")
		}
	}

	headers := make(map[string][]string)

	var retVal organizations.SearchResponse
	err = c.C.NewRequest(
		ctx,
		"POST",
		"/v1/b2b/organizations/search",
		nil,
		jsonBody,
		&retVal,
		headers,
	)
	return &retVal, err
}
