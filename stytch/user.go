package stytch

import (
	"encoding/json"
	"time"
)

type Name struct {
	FirstName  string `json:"first_name,omitempty"`
	MiddleName string `json:"middle_name,omitempty"`
	LastName   string `json:"last_name,omitempty"`
}

type Email struct {
	EmailID  string `json:"email_id,omitempty"`
	Email    string `json:"email,omitempty"`
	Verified bool   `json:"verified,omitempty"`
}

type EmailString struct {
	Email string `json:"email,omitempty"`
}

type PhoneNumber struct {
	PhoneID     string `json:"phone_id,omitempty"`
	PhoneNumber string `json:"phone_number,omitempty"`
	Verified    bool   `json:"verified,omitempty"`
}

type PhoneNumberString struct {
	PhoneNumber string `json:"phone_number,omitempty"`
}

type WebAuthnRegistration struct {
	WebAuthnRegistrationID string `json:"webauthn_registration_id,omitempty"`
	Domain                 string `json:"domain,omitempty"`
	UserAgent              string `json:"user_agent,omitempty"`
	Verified               bool   `json:"verified,omitempty"`
}

type OAuthProvider struct {
	ProviderType    string `json:"provider_type,omitempty"`
	ProviderSubject string `json:"provider_subject,omitempty"`
}

type UsersCreateParams struct {
	Email               string     `json:"email,omitempty"`
	PhoneNumber         string     `json:"phone_number,omitempty"`
	Name                Name       `json:"name,omitempty"`
	CreateUserAsPending bool       `json:"create_user_as_pending,omitempty"`
	Attributes          Attributes `json:"attributes,omitempty"`
}

type UsersCreateResponse struct {
	RequestID  string `json:"request_id,omitempty"`
	StatusCode int    `json:"status_code,omitempty"`
	UserID     string `json:"user_id,omitempty"`
	EmailID    string `json:"email_id,omitempty"`
	PhoneID    string `json:"phone_id,omitempty"`
	Status     string `json:"status,omitempty"`
}

type UsersGetResponse struct {
	RequestID             string                 `json:"request_id,omitempty"`
	StatusCode            int                    `json:"status_code,omitempty"`
	UserID                string                 `json:"user_id,omitempty"`
	Name                  Name                   `json:"name,omitempty"`
	Emails                []Email                `json:"emails,omitempty"`
	PhoneNumbers          []PhoneNumber          `json:"phone_numbers,omitempty"`
	WebAuthnRegistrations []WebAuthnRegistration `json:"webauthn_registrations,omitempty"`
	OAuthProviders        []OAuthProvider        `json:"providers,omitempty"`
	Status                string                 `json:"status,omitempty"`
	CreatedAt             time.Time              `json:"created_at,omitempty"`
}

type UsersUpdateParams struct {
	Name         Name                `json:"name,omitempty"`
	Emails       []EmailString       `json:"emails,omitempty"`
	PhoneNumbers []PhoneNumberString `json:"phone_numbers,omitempty"`
	Attributes   Attributes          `json:"attributes,omitempty"`
}

type UsersUpdateResponse struct {
	RequestID    string        `json:"request_id,omitempty"`
	StatusCode   int           `json:"status_code,omitempty"`
	UserID       string        `json:"user_id,omitempty"`
	Emails       []Email       `json:"emails,omitempty"`
	PhoneNumbers []PhoneNumber `json:"phone_numbers,omitempty"`
}

type UsersDeleteResponse struct {
	RequestID  string `json:"request_id,omitempty"`
	StatusCode int    `json:"status_code,omitempty"`
	UserID     string `json:"user_id,omitempty"`
}

type UsersDeleteEmailResponse struct {
	RequestID  string `json:"request_id,omitempty"`
	StatusCode int    `json:"status_code,omitempty"`
	UserID     string `json:"user_id,omitempty"`
}

type UsersDeletePhoneNumberResponse struct {
	RequestID  string `json:"request_id,omitempty"`
	StatusCode int    `json:"status_code,omitempty"`
	UserID     string `json:"user_id,omitempty"`
}

type UsersDeleteWebAuthnRegistrationResponse struct {
	RequestID  string `json:"request_id,omitempty"`
	StatusCode int    `json:"status_code,omitempty"`
	UserID     string `json:"user_id,omitempty"`
}

/* User Search */
type UsersSearchParams struct {
	Limit  int32            `json:"limit,omitempty"`
	Query  UsersSearchQuery `json:"query,omitempty"`
	Cursor string           `json:"cursor,omitempty"`
}

type UsersSearchOperator string

const (
	UserSearchOperatorOR  UsersSearchOperator = "OR"
	UserSearchOperatorAND UsersSearchOperator = "AND"
)

type UsersSearchQuery struct {
	Operator UsersSearchOperator `json:"operator,omitempty"`
	Operands []json.Marshaler    `json:"operands,omitempty"`
}

/* Created At Filters */
type UsersSearchQueryCreatedAtGreaterThanFilter struct {
	CreatedAtGreaterThan time.Time
}

func (q UsersSearchQueryCreatedAtGreaterThanFilter) MarshalJSON() ([]byte, error) {
	return json.Marshal(map[string]interface{}{
		"filter":                  "created_at_greater_than",
		"created_at_greater_than": q.CreatedAtGreaterThan,
	})
}

type UsersSearchQueryCreatedAtLessThanFilter struct {
	CreatedAtLessThan time.Time
}

func (q UsersSearchQueryCreatedAtLessThanFilter) MarshalJSON() ([]byte, error) {
	return json.Marshal(map[string]interface{}{
		"filter":               "created_at_less_than",
		"created_at_less_than": q.CreatedAtLessThan,
	})
}

type UsersSearchQueryCreatedAtBetweenFilter struct {
	CreatedAtGreaterThan time.Time
	CreatedAtLessThan    time.Time
}

func (q UsersSearchQueryCreatedAtBetweenFilter) MarshalJSON() ([]byte, error) {
	return json.Marshal(map[string]interface{}{
		"filter":                  "created_at_between",
		"created_at_greater_than": q.CreatedAtGreaterThan,
		"created_at_less_than":    q.CreatedAtLessThan,
	})
}

/* User Filters */
type UsersSearchQueryStatusFilter struct {
	Status string
}

func (q UsersSearchQueryStatusFilter) MarshalJSON() ([]byte, error) {
	return json.Marshal(map[string]interface{}{
		"filter": "status",
		"status": q.Status,
	})
}

type UsersSearchQueryUserIDFilter struct {
	UserIDs []string
}

func (q UsersSearchQueryUserIDFilter) MarshalJSON() ([]byte, error) {
	return json.Marshal(map[string]interface{}{
		"filter":   "user_id",
		"user_ids": q.UserIDs,
	})
}

type UsersSearchQueryFullNameFuzzyFilter struct {
	FullNameFuzzy string
}

func (q UsersSearchQueryFullNameFuzzyFilter) MarshalJSON() ([]byte, error) {
	return json.Marshal(map[string]interface{}{
		"filter":          "full_name_fuzzy",
		"full_name_fuzzy": q.FullNameFuzzy,
	})
}

/* Phone Number Filters */

type UsersSearchQueryPhoneNumberFilter struct {
	PhoneNumbers []string
}

func (q UsersSearchQueryPhoneNumberFilter) MarshalJSON() ([]byte, error) {
	return json.Marshal(map[string]interface{}{
		"filter":        "phone_number",
		"phone_numbers": q.PhoneNumbers,
	})
}

type UsersSearchQueryPhoneIDFilter struct {
	PhoneIDs []string
}

func (q UsersSearchQueryPhoneIDFilter) MarshalJSON() ([]byte, error) {
	return json.Marshal(map[string]interface{}{
		"filter":    "phone_id",
		"phone_ids": q.PhoneIDs,
	})
}

type UsersSearchQueryPhoneVerifiedFilter struct {
	PhoneVerified bool
}

func (q UsersSearchQueryPhoneVerifiedFilter) MarshalJSON() ([]byte, error) {
	return json.Marshal(map[string]interface{}{
		"filter":         "phone_verified",
		"phone_verified": q.PhoneVerified,
	})
}

type UsersSearchQueryPhoneNumberFuzzyFilter struct {
	PhoneNumberFuzzy string
}

func (q UsersSearchQueryPhoneNumberFuzzyFilter) MarshalJSON() ([]byte, error) {
	return json.Marshal(map[string]interface{}{
		"filter":             "phone_number_fuzzy",
		"phone_number_fuzzy": q.PhoneNumberFuzzy,
	})
}

/* Email Filters */

type UsersSearchQueryEmailAddressFilter struct {
	EmailAddresses []string
}

func (q UsersSearchQueryEmailAddressFilter) MarshalJSON() ([]byte, error) {
	return json.Marshal(map[string]interface{}{
		"filter":          "email_address",
		"email_addresses": q.EmailAddresses,
	})
}

type UsersSearchQueryEmailIDFilter struct {
	EmailIDs []string
}

func (q UsersSearchQueryEmailIDFilter) MarshalJSON() ([]byte, error) {
	return json.Marshal(map[string]interface{}{
		"filter":    "email_id",
		"email_ids": q.EmailIDs,
	})
}

type UsersSearchQueryEmailVerifiedFilter struct {
	EmailVerified bool
}

func (q UsersSearchQueryEmailVerifiedFilter) MarshalJSON() ([]byte, error) {
	return json.Marshal(map[string]interface{}{
		"filter":         "email_verified",
		"email_verified": q.EmailVerified,
	})
}

type UsersSearchQueryEmailAddressFuzzyFilter struct {
	EmailAddressFuzzy string
}

func (q UsersSearchQueryEmailAddressFuzzyFilter) MarshalJSON() ([]byte, error) {
	return json.Marshal(map[string]interface{}{
		"filter":              "email_address_fuzzy",
		"email_address_fuzzy": q.EmailAddressFuzzy,
	})
}

/* WebAuthn Filters */

type UsersSearchQueryWebAuthnRegistrationVerifiedFilter struct {
	WebAuthnRegistrationVerified bool
}

func (q UsersSearchQueryWebAuthnRegistrationVerifiedFilter) MarshalJSON() ([]byte, error) {
	return json.Marshal(map[string]interface{}{
		"filter":                         "webauthn_registration_verified",
		"webauthn_registration_verified": q.WebAuthnRegistrationVerified,
	})
}

/* OAuth Filters */

type UsersSearchQueryOAuthProviderFilter struct {
	OAuthProviders []string
}

func (q UsersSearchQueryOAuthProviderFilter) MarshalJSON() ([]byte, error) {
	return json.Marshal(map[string]interface{}{
		"filter":          "oauth_provider",
		"oauth_providers": q.OAuthProviders,
	})
}

type User struct {
	UserID                string                 `json:"user_id,omitempty"`
	Name                  Name                   `json:"name,omitempty"`
	Emails                []Email                `json:"emails,omitempty"`
	PhoneNumbers          []PhoneNumber          `json:"phone_numbers,omitempty"`
	WebAuthnRegistrations []WebAuthnRegistration `json:"webauthn_registrations,omitempty"`
	OAuthProviders        []OAuthProvider        `json:"providers,omitempty"`
	Status                string                 `json:"status,omitempty"`
	CreatedAt             time.Time              `json:"created_at,omitempty"`
}

type UsersSearchResponse struct {
	RequestID       string `json:"request_id,omitempty"`
	StatusCode      int    `json:"status_code,omitempty"`
	Results         []User `json:"results,omitempty"`
	ResultsMetadata struct {
		NextCursor string `json:"next_cursor,omitempty"`
		Total      int    `json:"total,omitempty"`
	} `json:"results_metadata,omitempty"`
}

/* End User Search */

type PendingUsers struct {
	UserID       string        `json:"user_id,omitempty"`
	Name         Name          `json:"name,omitempty"`
	Emails       []Email       `json:"emails,omitempty"`
	PhoneNumbers []PhoneNumber `json:"phone_numbers,omitempty"`
	Status       string        `json:"status,omitempty"`
	InvitedAt    string        `json:"invited_at,omitempty"`
}

type UsersGetPendingParams struct {
	Limit           int32  `json:"limit,omitempty"`
	StartingAfterID string `json:"starting_after_id,omitempty"`
}

type UsersGetPendingResponse struct {
	RequestID       string         `json:"request_id,omitempty"`
	StatusCode      int            `json:"status_code,omitempty"`
	Users           []PendingUsers `json:"users,omitempty"`
	HasMore         bool           `json:"has_more,omitempty"`
	StartingAfterID string         `json:"starting_after_id,omitempty"`
	Total           int            `json:"total,omitempty"`
}
