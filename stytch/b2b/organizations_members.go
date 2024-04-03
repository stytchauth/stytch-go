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

	"github.com/stytchauth/stytch-go/v13/stytch"
	"github.com/stytchauth/stytch-go/v13/stytch/b2b/organizations/members"
	"github.com/stytchauth/stytch-go/v13/stytch/stytcherror"
)

type OrganizationsMembersClient struct {
	C              stytch.Client
	OAuthProviders *OrganizationsMembersOAuthProvidersClient
}

func NewOrganizationsMembersClient(c stytch.Client) *OrganizationsMembersClient {
	return &OrganizationsMembersClient{
		C: c,

		OAuthProviders: NewOrganizationsMembersOAuthProvidersClient(c),
	}
}

// Update: Updates a Member specified by `organization_id` and `member_id`.
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
func (c *OrganizationsMembersClient) Update(
	ctx context.Context,
	body *members.UpdateParams,
	methodOptions ...*members.UpdateRequestOptions,
) (*members.UpdateResponse, error) {
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

	var retVal members.UpdateResponse
	err = c.C.NewRequest(
		ctx,
		"PUT",
		fmt.Sprintf("/v1/b2b/organizations/%s/members/%s", body.OrganizationID, body.MemberID),
		nil,
		jsonBody,
		&retVal,
		headers,
	)
	return &retVal, err
}

// Delete: Deletes a Member specified by `organization_id` and `member_id`. /%}
func (c *OrganizationsMembersClient) Delete(
	ctx context.Context,
	body *members.DeleteParams,
	methodOptions ...*members.DeleteRequestOptions,
) (*members.DeleteResponse, error) {
	headers := make(map[string][]string)
	for _, methodOption := range methodOptions {
		headers = methodOption.AddHeaders(headers)
	}

	var retVal members.DeleteResponse
	err := c.C.NewRequest(
		ctx,
		"DELETE",
		fmt.Sprintf("/v1/b2b/organizations/%s/members/%s", body.OrganizationID, body.MemberID),
		nil,
		nil,
		&retVal,
		headers,
	)
	return &retVal, err
}

// Reactivate: Reactivates a deleted Member's status and its associated email status (if applicable) to
// active, specified by `organization_id` and `member_id`. /%}
func (c *OrganizationsMembersClient) Reactivate(
	ctx context.Context,
	body *members.ReactivateParams,
	methodOptions ...*members.ReactivateRequestOptions,
) (*members.ReactivateResponse, error) {
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

	var retVal members.ReactivateResponse
	err = c.C.NewRequest(
		ctx,
		"PUT",
		fmt.Sprintf("/v1/b2b/organizations/%s/members/%s/reactivate", body.OrganizationID, body.MemberID),
		nil,
		jsonBody,
		&retVal,
		headers,
	)
	return &retVal, err
}

// DeleteMFAPhoneNumber: Delete a Member's MFA phone number.
//
// To change a Member's phone number, you must first call this endpoint to delete the existing phone number.
//
// Existing Member Sessions that include a phone number authentication factor will not be revoked if the
// phone number is deleted, and MFA will not be enforced until the Member logs in again.
// If you wish to enforce MFA immediately after a phone number is deleted, you can do so by prompting the
// Member to enter a new phone number
// and calling the [OTP SMS send](https://stytch.com/docs/b2b/api/otp-sms-send) endpoint, then calling the
// [OTP SMS Authenticate](https://stytch.com/docs/b2b/api/authenticate-otp-sms) endpoint.
//
//	/%}
func (c *OrganizationsMembersClient) DeleteMFAPhoneNumber(
	ctx context.Context,
	body *members.DeleteMFAPhoneNumberParams,
	methodOptions ...*members.DeleteMFAPhoneNumberRequestOptions,
) (*members.DeleteMFAPhoneNumberResponse, error) {
	headers := make(map[string][]string)
	for _, methodOption := range methodOptions {
		headers = methodOption.AddHeaders(headers)
	}

	var retVal members.DeleteMFAPhoneNumberResponse
	err := c.C.NewRequest(
		ctx,
		"DELETE",
		fmt.Sprintf("/v1/b2b/organizations/%s/members/mfa_phone_numbers/%s", body.OrganizationID, body.MemberID),
		nil,
		nil,
		&retVal,
		headers,
	)
	return &retVal, err
}

func (c *OrganizationsMembersClient) DeleteTOTP(
	ctx context.Context,
	body *members.DeleteTOTPParams,
	methodOptions ...*members.DeleteTOTPRequestOptions,
) (*members.DeleteTOTPResponse, error) {
	headers := make(map[string][]string)
	for _, methodOption := range methodOptions {
		headers = methodOption.AddHeaders(headers)
	}

	var retVal members.DeleteTOTPResponse
	err := c.C.NewRequest(
		ctx,
		"DELETE",
		fmt.Sprintf("/v1/b2b/organizations/%s/members/%s/totp", body.OrganizationID, body.MemberID),
		nil,
		nil,
		&retVal,
		headers,
	)
	return &retVal, err
}

// Search for Members within specified Organizations. An array with at least one `organization_id` is
// required. Submitting an empty `query` returns all non-deleted Members within the specified Organizations.
//
// *All fuzzy search filters require a minimum of three characters.
//
// Our RBAC implementation offers out-of-the-box handling of authorization checks for this endpoint. If you
// pass in
// a header containing a `session_token` or a `session_jwt` for an unexpired Member Session, we will check
// that the
// Member Session has permission to perform the `search` action on the `stytch.member` Resource. In
// addition, enforcing
// RBAC on this endpoint means that you may only search for Members within the calling Member's
// Organization, so the
// `organization_ids` argument may only contain the `organization_id` of the Member Session passed in the
// header.
//
// If the Member Session does not contain a Role that satisfies the requested permission, or if the
// `organization_ids`
// argument contains an `organization_id` that the Member Session does not belong to, a 403 error will be
// thrown.
// Otherwise, the request will proceed as normal.
//
// To learn more about our RBAC implementation, see our
// [RBAC guide](https://stytch.com/docs/b2b/guides/rbac/overview).
func (c *OrganizationsMembersClient) Search(
	ctx context.Context,
	body *members.SearchParams,
	methodOptions ...*members.SearchRequestOptions,
) (*members.SearchResponse, error) {
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

	var retVal members.SearchResponse
	err = c.C.NewRequest(
		ctx,
		"POST",
		"/v1/b2b/organizations/members/search",
		nil,
		jsonBody,
		&retVal,
		headers,
	)
	return &retVal, err
}

// DeletePassword: Delete a Member's password. /%}
func (c *OrganizationsMembersClient) DeletePassword(
	ctx context.Context,
	body *members.DeletePasswordParams,
	methodOptions ...*members.DeletePasswordRequestOptions,
) (*members.DeletePasswordResponse, error) {
	headers := make(map[string][]string)
	for _, methodOption := range methodOptions {
		headers = methodOption.AddHeaders(headers)
	}

	var retVal members.DeletePasswordResponse
	err := c.C.NewRequest(
		ctx,
		"DELETE",
		fmt.Sprintf("/v1/b2b/organizations/%s/members/passwords/%s", body.OrganizationID, body.MemberPasswordID),
		nil,
		nil,
		&retVal,
		headers,
	)
	return &retVal, err
}

// DangerouslyGet: Get a Member by `member_id`. This endpoint does not require an `organization_id`,
// enabling you to get members across organizations. This is a dangerous operation. Incorrect use may open
// you up to indirect object reference (IDOR) attacks. We recommend using the
// [Get Member](https://stytch.com/docs/b2b/api/get-member) API instead.
func (c *OrganizationsMembersClient) DangerouslyGet(
	ctx context.Context,
	body *members.DangerouslyGetParams,
) (*members.GetResponse, error) {
	headers := make(map[string][]string)

	var retVal members.GetResponse
	err := c.C.NewRequest(
		ctx,
		"GET",
		fmt.Sprintf("/v1/b2b/organizations/members/dangerously_get/%s", body.MemberID),
		nil,
		nil,
		&retVal,
		headers,
	)
	return &retVal, err
}

// Create: Creates a Member. An `organization_id` and `email_address` are required. /%}
func (c *OrganizationsMembersClient) Create(
	ctx context.Context,
	body *members.CreateParams,
	methodOptions ...*members.CreateRequestOptions,
) (*members.CreateResponse, error) {
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

	var retVal members.CreateResponse
	err = c.C.NewRequest(
		ctx,
		"POST",
		fmt.Sprintf("/v1/b2b/organizations/%s/members", body.OrganizationID),
		nil,
		jsonBody,
		&retVal,
		headers,
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

	headers := make(map[string][]string)

	var retVal members.GetResponse
	err := c.C.NewRequest(
		ctx,
		"GET",
		fmt.Sprintf("/v1/b2b/organizations/%s/member", body.OrganizationID),
		queryParams,
		nil,
		&retVal,
		headers,
	)
	return &retVal, err
}
