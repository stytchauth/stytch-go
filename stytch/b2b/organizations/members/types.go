package members

// !!!
// WARNING: This file is autogenerated
// Only modify code within MANUAL() sections
// or your changes may be overwritten later!
// !!!

import (
	"github.com/stytchauth/stytch-go/v12/stytch/b2b/organizations"
	"github.com/stytchauth/stytch-go/v12/stytch/methodoptions"
)

// CreateParams: Request type for `Members.Create`.
type CreateParams struct {
	// OrganizationID: Globally unique UUID that identifies a specific Organization. The `organization_id` is
	// critical to perform operations on an Organization, so be sure to preserve this value.
	OrganizationID string `json:"organization_id,omitempty"`
	// EmailAddress: The email address of the Member.
	EmailAddress string `json:"email_address,omitempty"`
	// Name: The name of the Member.
	Name string `json:"name,omitempty"`
	// TrustedMetadata: An arbitrary JSON object for storing application-specific data or
	// identity-provider-specific data.
	TrustedMetadata map[string]any `json:"trusted_metadata,omitempty"`
	// UntrustedMetadata: An arbitrary JSON object of application-specific data. These fields can be edited
	// directly by the
	//   frontend SDK, and should not be used to store critical information. See the
	// [Metadata resource](https://stytch.com/docs/b2b/api/metadata)
	//   for complete field behavior details.
	UntrustedMetadata map[string]any `json:"untrusted_metadata,omitempty"`
	// CreateMemberAsPending: Flag for whether or not to save a Member as `pending` or `active` in Stytch. It
	// defaults to false. If true, new Members will be created with status `pending` in Stytch's backend. Their
	// status will remain `pending` and they will continue to receive signup email templates for every Email
	// Magic Link until that Member authenticates and becomes `active`. If false, new Members will be created
	// with status `active`.
	CreateMemberAsPending bool `json:"create_member_as_pending,omitempty"`
	// IsBreakglass: Identifies the Member as a break glass user - someone who has permissions to authenticate
	// into an Organization by bypassing the Organization's settings. A break glass account is typically used
	// for emergency purposes to gain access outside of normal authentication procedures. Refer to the
	// [Organization object](organization-object) and its `auth_methods` and `allowed_auth_methods` fields for
	// more details.
	IsBreakglass bool `json:"is_breakglass,omitempty"`
	// MFAPhoneNumber: The Member's phone number. A Member may only have one phone number.
	MFAPhoneNumber string `json:"mfa_phone_number,omitempty"`
	// MFAEnrolled: Sets whether the Member is enrolled in MFA. If true, the Member must complete an MFA step
	// whenever they wish to log in to their Organization. If false, the Member only needs to complete an MFA
	// step if the Organization's MFA policy is set to `REQUIRED_FOR_ALL`.
	MFAEnrolled bool `json:"mfa_enrolled,omitempty"`
	// Roles to explicitly assign to this Member. See the
	// [RBAC guide](https://stytch.com/docs/b2b/guides/rbac/role-assignment)
	//    for more information about role assignment.
	Roles []string `json:"roles,omitempty"`
}

// DangerouslyGetParams: Request type for `Members.DangerouslyGet`.
type DangerouslyGetParams struct {
	// MemberID: Globally unique UUID that identifies a specific Member. The `member_id` is critical to perform
	// operations on a Member, so be sure to preserve this value.
	MemberID string `json:"member_id,omitempty"`
}

// DeleteMFAPhoneNumberParams: Request type for `Members.DeleteMFAPhoneNumber`.
type DeleteMFAPhoneNumberParams struct {
	// OrganizationID: Globally unique UUID that identifies a specific Organization. The `organization_id` is
	// critical to perform operations on an Organization, so be sure to preserve this value.
	OrganizationID string `json:"organization_id,omitempty"`
	// MemberID: Globally unique UUID that identifies a specific Member. The `member_id` is critical to perform
	// operations on a Member, so be sure to preserve this value.
	MemberID string `json:"member_id,omitempty"`
}

// DeleteParams: Request type for `Members.Delete`.
type DeleteParams struct {
	// OrganizationID: Globally unique UUID that identifies a specific Organization. The `organization_id` is
	// critical to perform operations on an Organization, so be sure to preserve this value.
	OrganizationID string `json:"organization_id,omitempty"`
	// MemberID: Globally unique UUID that identifies a specific Member. The `member_id` is critical to perform
	// operations on a Member, so be sure to preserve this value.
	MemberID string `json:"member_id,omitempty"`
}

// DeletePasswordParams: Request type for `Members.DeletePassword`.
type DeletePasswordParams struct {
	// OrganizationID: Globally unique UUID that identifies a specific Organization. The `organization_id` is
	// critical to perform operations on an Organization, so be sure to preserve this value.
	OrganizationID string `json:"organization_id,omitempty"`
	// MemberPasswordID: Globally unique UUID that identifies a Member's password.
	MemberPasswordID string `json:"member_password_id,omitempty"`
}

type DeleteTOTPParams struct {
	OrganizationID string `json:"organization_id,omitempty"`
	MemberID       string `json:"member_id,omitempty"`
}

// GetParams: Request type for `Members.Get`.
type GetParams struct {
	// OrganizationID: Globally unique UUID that identifies a specific Organization. The `organization_id` is
	// critical to perform operations on an Organization, so be sure to preserve this value.
	OrganizationID string `json:"organization_id,omitempty"`
	// MemberID: Globally unique UUID that identifies a specific Member. The `member_id` is critical to perform
	// operations on a Member, so be sure to preserve this value.
	MemberID string `json:"member_id,omitempty"`
	// EmailAddress: The email address of the Member.
	EmailAddress string `json:"email_address,omitempty"`
}

// ReactivateParams: Request type for `Members.Reactivate`.
type ReactivateParams struct {
	// OrganizationID: Globally unique UUID that identifies a specific Organization. The `organization_id` is
	// critical to perform operations on an Organization, so be sure to preserve this value.
	OrganizationID string `json:"organization_id,omitempty"`
	// MemberID: Globally unique UUID that identifies a specific Member. The `member_id` is critical to perform
	// operations on a Member, so be sure to preserve this value.
	MemberID string `json:"member_id,omitempty"`
}

// SearchParams: Request type for `Members.Search`.
type SearchParams struct {
	// OrganizationIds: An array of organization_ids. At least one value is required.
	OrganizationIds []string `json:"organization_ids,omitempty"`
	// Cursor: The `cursor` field allows you to paginate through your results. Each result array is limited to
	// 1000 results. If your query returns more than 1000 results, you will need to paginate the responses
	// using the `cursor`. If you receive a response that includes a non-null `next_cursor` in the
	// `results_metadata` object, repeat the search call with the `next_cursor` value set to the `cursor` field
	// to retrieve the next page of results. Continue to make search calls until the `next_cursor` in the
	// response is null.
	Cursor string `json:"cursor,omitempty"`
	// Limit: The number of search results to return per page. The default limit is 100. A maximum of 1000
	// results can be returned by a single search request. If the total size of your result set is greater than
	// one page size, you must paginate the response. See the `cursor` field.
	Limit uint32 `json:"limit,omitempty"`
	// Query: The optional query object contains the operator, i.e. `AND` or `OR`, and the operands that will
	// filter your results. Only an operator is required. If you include no operands, no filtering will be
	// applied. If you include no query object, it will return all Members with no filtering applied.
	Query *organizations.SearchQuery `json:"query,omitempty"`
}

// UpdateParams: Request type for `Members.Update`.
type UpdateParams struct {
	// OrganizationID: Globally unique UUID that identifies a specific Organization. The `organization_id` is
	// critical to perform operations on an Organization, so be sure to preserve this value.
	OrganizationID string `json:"organization_id,omitempty"`
	// MemberID: Globally unique UUID that identifies a specific Member. The `member_id` is critical to perform
	// operations on a Member, so be sure to preserve this value.
	MemberID string `json:"member_id,omitempty"`
	// Name: The name of the Member.
	//
	// If this field is provided and a session header is passed into the request, the Member Session must have
	// permission to perform the `update.info.name` action on the `stytch.member` Resource.
	//   Alternatively, if the Member Session matches the Member associated with the `member_id` passed in the
	// request, the authorization check will also allow a Member Session that has permission to perform the
	// `update.info.name` action on the `stytch.self` Resource.
	Name string `json:"name,omitempty"`
	// TrustedMetadata: An arbitrary JSON object for storing application-specific data or
	// identity-provider-specific data.
	//           If a session header is passed into the request, this field may **not** be passed into the
	// request. You cannot
	//           update trusted metadata when acting as a Member.
	TrustedMetadata map[string]any `json:"trusted_metadata,omitempty"`
	// UntrustedMetadata: An arbitrary JSON object of application-specific data. These fields can be edited
	// directly by the
	//   frontend SDK, and should not be used to store critical information. See the
	// [Metadata resource](https://stytch.com/docs/b2b/api/metadata)
	//   for complete field behavior details.
	//
	// If this field is provided and a session header is passed into the request, the Member Session must have
	// permission to perform the `update.info.untrusted-metadata` action on the `stytch.member` Resource.
	//   Alternatively, if the Member Session matches the Member associated with the `member_id` passed in the
	// request, the authorization check will also allow a Member Session that has permission to perform the
	// `update.info.untrusted-metadata` action on the `stytch.self` Resource.
	UntrustedMetadata map[string]any `json:"untrusted_metadata,omitempty"`
	// IsBreakglass: Identifies the Member as a break glass user - someone who has permissions to authenticate
	// into an Organization by bypassing the Organization's settings. A break glass account is typically used
	// for emergency purposes to gain access outside of normal authentication procedures. Refer to the
	// [Organization object](organization-object) and its `auth_methods` and `allowed_auth_methods` fields for
	// more details.
	//
	// If this field is provided and a session header is passed into the request, the Member Session must have
	// permission to perform the `update.settings.is-breakglass` action on the `stytch.member` Resource.
	IsBreakglass bool `json:"is_breakglass,omitempty"`
	// MFAPhoneNumber: Sets the Member's phone number. Throws an error if the Member already has a phone
	// number. To change the Member's phone number, use the
	// [Delete member phone number endpoint](https://stytch.com/docs/b2b/api/delete-member-mfa-phone-number) to
	// delete the Member's existing phone number first.
	//
	// If this field is provided and a session header is passed into the request, the Member Session must have
	// permission to perform the `update.info.mfa-phone` action on the `stytch.member` Resource.
	//   Alternatively, if the Member Session matches the Member associated with the `member_id` passed in the
	// request, the authorization check will also allow a Member Session that has permission to perform the
	// `update.info.mfa-phone` action on the `stytch.self` Resource.
	MFAPhoneNumber string `json:"mfa_phone_number,omitempty"`
	// MFAEnrolled: Sets whether the Member is enrolled in MFA. If true, the Member must complete an MFA step
	// whenever they wish to log in to their Organization. If false, the Member only needs to complete an MFA
	// step if the Organization's MFA policy is set to `REQUIRED_FOR_ALL`.
	//
	// If this field is provided and a session header is passed into the request, the Member Session must have
	// permission to perform the `update.settings.mfa-enrolled` action on the `stytch.member` Resource.
	//   Alternatively, if the Member Session matches the Member associated with the `member_id` passed in the
	// request, the authorization check will also allow a Member Session that has permission to perform the
	// `update.settings.mfa-enrolled` action on the `stytch.self` Resource.
	MFAEnrolled bool `json:"mfa_enrolled,omitempty"`
	// Roles to explicitly assign to this Member.
	//  Will completely replace any existing explicitly assigned roles. See the
	//  [RBAC guide](https://stytch.com/docs/b2b/guides/rbac/role-assignment) for more information about role
	// assignment.
	//
	//    If a Role is removed from a Member, and the Member is also implicitly assigned this Role from an SSO
	// connection
	//    or an SSO group, we will by default revoke any existing sessions for the Member that contain any SSO
	//    authentication factors with the affected connection ID. You can preserve these sessions by passing in
	// the
	//    `preserve_existing_sessions` parameter with a value of `true`.
	//
	// If this field is provided and a session header is passed into the request, the Member Session must have
	// permission to perform the `update.settings.roles` action on the `stytch.member` Resource.
	Roles []string `json:"roles,omitempty"`
	// PreserveExistingSessions: Whether to preserve existing sessions when explicit Roles that are revoked are
	// also implicitly assigned
	//   by SSO connection or SSO group. Defaults to `false` - that is, existing Member Sessions that contain
	// SSO
	//   authentication factors with the affected SSO connection IDs will be revoked.
	PreserveExistingSessions bool `json:"preserve_existing_sessions,omitempty"`
	// DefaultMFAMethod: Sets whether the Member is enrolled in MFA. If true, the Member must complete an MFA
	// step whenever they wish to log in to their Organization. If false, the Member only needs to complete an
	// MFA step if the Organization's MFA policy is set to `REQUIRED_FOR_ALL`.
	//
	// If this field is provided and a session header is passed into the request, the Member Session must have
	// permission to perform the `update.settings.default-mfa-method` action on the `stytch.member` Resource.
	//   Alternatively, if the Member Session matches the Member associated with the `member_id` passed in the
	// request, the authorization check will also allow a Member Session that has permission to perform the
	// `update.settings.default-mfa-method` action on the `stytch.self` Resource.
	DefaultMFAMethod string `json:"default_mfa_method,omitempty"`
	// EmailAddress: Updates the Member's `email_address`, if provided.
	EmailAddress string `json:"email_address,omitempty"`
}

// CreateRequestOptions:
type CreateRequestOptions struct {
	// Authorization: Optional authorization object.
	// Pass in an active Stytch Member session token or session JWT and the request
	// will be run using that member's permissions.
	Authorization methodoptions.Authorization `json:"authorization,omitempty"`
}

func (o *CreateRequestOptions) AddHeaders(headers map[string][]string) map[string][]string {
	headers = o.Authorization.AddHeaders(headers)
	return headers
}

// DeleteMFAPhoneNumberRequestOptions:
type DeleteMFAPhoneNumberRequestOptions struct {
	// Authorization: Optional authorization object.
	// Pass in an active Stytch Member session token or session JWT and the request
	// will be run using that member's permissions.
	Authorization methodoptions.Authorization `json:"authorization,omitempty"`
}

func (o *DeleteMFAPhoneNumberRequestOptions) AddHeaders(headers map[string][]string) map[string][]string {
	headers = o.Authorization.AddHeaders(headers)
	return headers
}

// DeletePasswordRequestOptions:
type DeletePasswordRequestOptions struct {
	// Authorization: Optional authorization object.
	// Pass in an active Stytch Member session token or session JWT and the request
	// will be run using that member's permissions.
	Authorization methodoptions.Authorization `json:"authorization,omitempty"`
}

func (o *DeletePasswordRequestOptions) AddHeaders(headers map[string][]string) map[string][]string {
	headers = o.Authorization.AddHeaders(headers)
	return headers
}

// DeleteRequestOptions:
type DeleteRequestOptions struct {
	// Authorization: Optional authorization object.
	// Pass in an active Stytch Member session token or session JWT and the request
	// will be run using that member's permissions.
	Authorization methodoptions.Authorization `json:"authorization,omitempty"`
}

func (o *DeleteRequestOptions) AddHeaders(headers map[string][]string) map[string][]string {
	headers = o.Authorization.AddHeaders(headers)
	return headers
}

// DeleteTOTPRequestOptions:
type DeleteTOTPRequestOptions struct {
	// Authorization: Optional authorization object.
	// Pass in an active Stytch Member session token or session JWT and the request
	// will be run using that member's permissions.
	Authorization methodoptions.Authorization `json:"authorization,omitempty"`
}

func (o *DeleteTOTPRequestOptions) AddHeaders(headers map[string][]string) map[string][]string {
	headers = o.Authorization.AddHeaders(headers)
	return headers
}

// ReactivateRequestOptions:
type ReactivateRequestOptions struct {
	// Authorization: Optional authorization object.
	// Pass in an active Stytch Member session token or session JWT and the request
	// will be run using that member's permissions.
	Authorization methodoptions.Authorization `json:"authorization,omitempty"`
}

func (o *ReactivateRequestOptions) AddHeaders(headers map[string][]string) map[string][]string {
	headers = o.Authorization.AddHeaders(headers)
	return headers
}

// SearchRequestOptions:
type SearchRequestOptions struct {
	// Authorization: Optional authorization object.
	// Pass in an active Stytch Member session token or session JWT and the request
	// will be run using that member's permissions.
	Authorization methodoptions.Authorization `json:"authorization,omitempty"`
}

func (o *SearchRequestOptions) AddHeaders(headers map[string][]string) map[string][]string {
	headers = o.Authorization.AddHeaders(headers)
	return headers
}

// UpdateRequestOptions:
type UpdateRequestOptions struct {
	// Authorization: Optional authorization object.
	// Pass in an active Stytch Member session token or session JWT and the request
	// will be run using that member's permissions.
	Authorization methodoptions.Authorization `json:"authorization,omitempty"`
}

func (o *UpdateRequestOptions) AddHeaders(headers map[string][]string) map[string][]string {
	headers = o.Authorization.AddHeaders(headers)
	return headers
}

// CreateResponse: Response type for `Members.Create`.
type CreateResponse struct {
	// RequestID: Globally unique UUID that is returned with every API call. This value is important to log for
	// debugging purposes; we may ask for this value to help identify a specific API call when helping you
	// debug an issue.
	RequestID string `json:"request_id,omitempty"`
	// MemberID: Globally unique UUID that identifies a specific Member.
	MemberID string `json:"member_id,omitempty"`
	// Member: The [Member object](https://stytch.com/docs/b2b/api/member-object)
	Member organizations.Member `json:"member,omitempty"`
	// Organization: The [Organization object](https://stytch.com/docs/b2b/api/organization-object).
	Organization organizations.Organization `json:"organization,omitempty"`
	// StatusCode: The HTTP status code of the response. Stytch follows standard HTTP response status code
	// patterns, e.g. 2XX values equate to success, 3XX values are redirects, 4XX are client errors, and 5XX
	// are server errors.
	StatusCode int32 `json:"status_code,omitempty"`
}

// DeleteMFAPhoneNumberResponse: Response type for `Members.DeleteMFAPhoneNumber`.
type DeleteMFAPhoneNumberResponse struct {
	// RequestID: Globally unique UUID that is returned with every API call. This value is important to log for
	// debugging purposes; we may ask for this value to help identify a specific API call when helping you
	// debug an issue.
	RequestID string `json:"request_id,omitempty"`
	// MemberID: Globally unique UUID that identifies a specific Member.
	MemberID string `json:"member_id,omitempty"`
	// Member: The [Member object](https://stytch.com/docs/b2b/api/member-object)
	Member organizations.Member `json:"member,omitempty"`
	// Organization: The [Organization object](https://stytch.com/docs/b2b/api/organization-object).
	Organization organizations.Organization `json:"organization,omitempty"`
	// StatusCode: The HTTP status code of the response. Stytch follows standard HTTP response status code
	// patterns, e.g. 2XX values equate to success, 3XX values are redirects, 4XX are client errors, and 5XX
	// are server errors.
	StatusCode int32 `json:"status_code,omitempty"`
}

// DeletePasswordResponse: Response type for `Members.DeletePassword`.
type DeletePasswordResponse struct {
	// RequestID: Globally unique UUID that is returned with every API call. This value is important to log for
	// debugging purposes; we may ask for this value to help identify a specific API call when helping you
	// debug an issue.
	RequestID string `json:"request_id,omitempty"`
	// MemberID: Globally unique UUID that identifies a specific Member.
	MemberID string `json:"member_id,omitempty"`
	// Member: The [Member object](https://stytch.com/docs/b2b/api/member-object)
	Member organizations.Member `json:"member,omitempty"`
	// Organization: The [Organization object](https://stytch.com/docs/b2b/api/organization-object).
	Organization organizations.Organization `json:"organization,omitempty"`
	// StatusCode: The HTTP status code of the response. Stytch follows standard HTTP response status code
	// patterns, e.g. 2XX values equate to success, 3XX values are redirects, 4XX are client errors, and 5XX
	// are server errors.
	StatusCode int32 `json:"status_code,omitempty"`
}

// DeleteResponse: Response type for `Members.Delete`.
type DeleteResponse struct {
	// RequestID: Globally unique UUID that is returned with every API call. This value is important to log for
	// debugging purposes; we may ask for this value to help identify a specific API call when helping you
	// debug an issue.
	RequestID string `json:"request_id,omitempty"`
	// MemberID: Globally unique UUID that identifies a specific Member.
	MemberID string `json:"member_id,omitempty"`
	// StatusCode: The HTTP status code of the response. Stytch follows standard HTTP response status code
	// patterns, e.g. 2XX values equate to success, 3XX values are redirects, 4XX are client errors, and 5XX
	// are server errors.
	StatusCode int32 `json:"status_code,omitempty"`
}

type DeleteTOTPResponse struct {
	RequestID    string                     `json:"request_id,omitempty"`
	MemberID     string                     `json:"member_id,omitempty"`
	Member       organizations.Member       `json:"member,omitempty"`
	Organization organizations.Organization `json:"organization,omitempty"`
	StatusCode   int32                      `json:"status_code,omitempty"`
}

// GetResponse: Response type for `Members.DangerouslyGet`, `Members.Get`.
type GetResponse struct {
	// RequestID: Globally unique UUID that is returned with every API call. This value is important to log for
	// debugging purposes; we may ask for this value to help identify a specific API call when helping you
	// debug an issue.
	RequestID string `json:"request_id,omitempty"`
	// MemberID: Globally unique UUID that identifies a specific Member.
	MemberID string `json:"member_id,omitempty"`
	// Member: The [Member object](https://stytch.com/docs/b2b/api/member-object)
	Member organizations.Member `json:"member,omitempty"`
	// Organization: The [Organization object](https://stytch.com/docs/b2b/api/organization-object).
	Organization organizations.Organization `json:"organization,omitempty"`
	// StatusCode: The HTTP status code of the response. Stytch follows standard HTTP response status code
	// patterns, e.g. 2XX values equate to success, 3XX values are redirects, 4XX are client errors, and 5XX
	// are server errors.
	StatusCode int32 `json:"status_code,omitempty"`
}

// ReactivateResponse: Response type for `Members.Reactivate`.
type ReactivateResponse struct {
	// RequestID: Globally unique UUID that is returned with every API call. This value is important to log for
	// debugging purposes; we may ask for this value to help identify a specific API call when helping you
	// debug an issue.
	RequestID string `json:"request_id,omitempty"`
	// MemberID: Globally unique UUID that identifies a specific Member.
	MemberID string `json:"member_id,omitempty"`
	// Member: The [Member object](https://stytch.com/docs/b2b/api/member-object)
	Member organizations.Member `json:"member,omitempty"`
	// Organization: The [Organization object](https://stytch.com/docs/b2b/api/organization-object).
	Organization organizations.Organization `json:"organization,omitempty"`
	// StatusCode: The HTTP status code of the response. Stytch follows standard HTTP response status code
	// patterns, e.g. 2XX values equate to success, 3XX values are redirects, 4XX are client errors, and 5XX
	// are server errors.
	StatusCode int32 `json:"status_code,omitempty"`
}

// SearchResponse: Response type for `Members.Search`.
type SearchResponse struct {
	// RequestID: Globally unique UUID that is returned with every API call. This value is important to log for
	// debugging purposes; we may ask for this value to help identify a specific API call when helping you
	// debug an issue.
	RequestID string `json:"request_id,omitempty"`
	// Members: An array of [Member objects](member-object).
	Members []organizations.Member `json:"members,omitempty"`
	// ResultsMetadata: The search `results_metadata` object contains metadata relevant to your specific query
	// like `total` and `next_cursor`.
	ResultsMetadata organizations.ResultsMetadata `json:"results_metadata,omitempty"`
	// Organizations: A map from `organization_id` to
	// [Organization object](https://stytch.com/docs/b2b/api/organization-object). The map only contains the
	// Organizations that the Members belongs to.
	Organizations map[string]organizations.Organization `json:"organizations,omitempty"`
	// StatusCode: The HTTP status code of the response. Stytch follows standard HTTP response status code
	// patterns, e.g. 2XX values equate to success, 3XX values are redirects, 4XX are client errors, and 5XX
	// are server errors.
	StatusCode int32 `json:"status_code,omitempty"`
}

// UpdateResponse: Response type for `Members.Update`.
type UpdateResponse struct {
	// RequestID: Globally unique UUID that is returned with every API call. This value is important to log for
	// debugging purposes; we may ask for this value to help identify a specific API call when helping you
	// debug an issue.
	RequestID string `json:"request_id,omitempty"`
	// MemberID: Globally unique UUID that identifies a specific Member.
	MemberID string `json:"member_id,omitempty"`
	// Member: The [Member object](https://stytch.com/docs/b2b/api/member-object)
	Member organizations.Member `json:"member,omitempty"`
	// Organization: The [Organization object](https://stytch.com/docs/b2b/api/organization-object).
	Organization organizations.Organization `json:"organization,omitempty"`
	// StatusCode: The HTTP status code of the response. Stytch follows standard HTTP response status code
	// patterns, e.g. 2XX values equate to success, 3XX values are redirects, 4XX are client errors, and 5XX
	// are server errors.
	StatusCode int32 `json:"status_code,omitempty"`
}
