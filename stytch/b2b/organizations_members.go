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

	"github.com/stytchauth/stytch-go/v15/stytch"
	"github.com/stytchauth/stytch-go/v15/stytch/b2b/organizations/members"
	"github.com/stytchauth/stytch-go/v15/stytch/stytcherror"
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

// Update: Updates a specified by `organization_id` and `member_id`.
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

// Delete: Deletes a specified by `organization_id` and `member_id`.
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

// Reactivate: Reactivates a deleted's status and its associated email status (if applicable) to active,
// specified by `organization_id` and `member_id`. This endpoint will only work for Members with at least
// one verified email where their `email_address_verified` is `true`.
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

// DeleteMFAPhoneNumber: Delete a's MFA phone number.
//
// To change a Member's phone number, you must first call this endpoint to delete the existing phone number.
//
// Existing Member Sessions that include a phone number authentication factor will not be revoked if the
// phone number is deleted, and MFA will not be enforced until the Member logs in again.
// If you wish to enforce MFA immediately after a phone number is deleted, you can do so by prompting the
// Member to enter a new phone number
// and calling the [OTP SMS send](https://stytch.com/docs/b2b/api/otp-sms-send) endpoint, then calling the
// [OTP SMS Authenticate](https://stytch.com/docs/b2b/api/authenticate-otp-sms) endpoint.
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

// DeleteTOTP: Delete a Member's MFA TOTP registration.
//
// To mint a new registration for a Member, you must first call this endpoint to delete the existing
// registration.
//
// Existing Member Sessions that include the TOTP authentication factor will not be revoked if the
// registration is deleted, and MFA will not be enforced until the Member logs in again.
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

// DeletePassword: Delete a's password.
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

func (c *OrganizationsMembersClient) OIDCProviders(
	ctx context.Context,
	body *members.OIDCProviderInformationParams,
) (*members.OIDCProvidersResponse, error) {
	queryParams := make(map[string]string)
	if body != nil {
		if body.IncludeRefreshToken {
			queryParams["include_refresh_token"] = "true"
		} else {
			queryParams["include_refresh_token"] = "false"
		}
	}

	headers := make(map[string][]string)

	var retVal members.OIDCProvidersResponse
	err := c.C.NewRequest(
		ctx,
		"GET",
		fmt.Sprintf("/v1/b2b/organizations/%s/members/%s/oidc_providers", body.OrganizationID, body.MemberID),
		queryParams,
		nil,
		&retVal,
		headers,
	)
	return &retVal, err
}

// UnlinkRetiredEmail: Unlinks a retired email address from a specified by their `organization_id` and
// `member_id`. The email address
// to be retired can be identified in the request body by either its `email_id`, its `email_address`, or
// both. If using
// both identifiers they must refer to the same email.
//
// A previously active email address can be marked as retired in one of two ways:
//
// - It's replaced with a new primary email address during an explicit Member update.
// - A new email address is surfaced by an OAuth, SAML or OIDC provider. In this case the new email address
// becomes the
//
//	Member's primary email address and the old primary email address is retired.
//
// A retired email address cannot be used by other Members in the same Organization. However, unlinking
// retired email
// addresses allows them to be subsequently re-used by other Organization Members. Retired email addresses
// can be viewed
// on the [Member object](https://stytch.com/docs/b2b/api/member-object).
//
//	%}
func (c *OrganizationsMembersClient) UnlinkRetiredEmail(
	ctx context.Context,
	body *members.UnlinkRetiredEmailParams,
	methodOptions ...*members.UnlinkRetiredEmailRequestOptions,
) (*members.UnlinkRetiredEmailResponse, error) {
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

	var retVal members.UnlinkRetiredEmailResponse
	err = c.C.NewRequest(
		ctx,
		"POST",
		fmt.Sprintf("/v1/b2b/organizations/%s/members/%s/unlink_retired_email", body.OrganizationID, body.MemberID),
		nil,
		jsonBody,
		&retVal,
		headers,
	)
	return &retVal, err
}

// Create: Creates a. An `organization_id` and `email_address` are required.
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
