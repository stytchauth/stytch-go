package scim

// !!!
// WARNING: This file is autogenerated
// Only modify code within MANUAL() sections
// or your changes may be overwritten later!
// !!!

import (
	"time"
)

type Address struct {
	Formatted     string `json:"formatted,omitempty"`
	StreetAddress string `json:"street_address,omitempty"`
	Locality      string `json:"locality,omitempty"`
	Region        string `json:"region,omitempty"`
	PostalCode    string `json:"postal_code,omitempty"`
	Country       string `json:"country,omitempty"`
	Type          string `json:"type,omitempty"`
	Primary       bool   `json:"primary,omitempty"`
}

type Email struct {
	Value   string `json:"value,omitempty"`
	Type    string `json:"type,omitempty"`
	Primary bool   `json:"primary,omitempty"`
}

type EnterpriseExtension struct {
	EmployeeNumber string   `json:"employee_number,omitempty"`
	CostCenter     string   `json:"cost_center,omitempty"`
	Division       string   `json:"division,omitempty"`
	Department     string   `json:"department,omitempty"`
	Organization   string   `json:"organization,omitempty"`
	Manager        *Manager `json:"manager,omitempty"`
}

type Entitlement struct {
	Value   string `json:"value,omitempty"`
	Type    string `json:"type,omitempty"`
	Primary bool   `json:"primary,omitempty"`
}

type Group struct {
	Value   string `json:"value,omitempty"`
	Display string `json:"display,omitempty"`
}

type IMs struct {
	Value   string `json:"value,omitempty"`
	Type    string `json:"type,omitempty"`
	Primary bool   `json:"primary,omitempty"`
}

type Manager struct {
	Value       string `json:"value,omitempty"`
	Ref         string `json:"ref,omitempty"`
	DisplayName string `json:"display_name,omitempty"`
}

type Name struct {
	Formatted       string `json:"formatted,omitempty"`
	FamilyName      string `json:"family_name,omitempty"`
	GivenName       string `json:"given_name,omitempty"`
	MiddleName      string `json:"middle_name,omitempty"`
	HonorificPrefix string `json:"honorific_prefix,omitempty"`
	HonorificSuffix string `json:"honorific_suffix,omitempty"`
}

type PhoneNumber struct {
	Value   string `json:"value,omitempty"`
	Type    string `json:"type,omitempty"`
	Primary bool   `json:"primary,omitempty"`
}

type Photo struct {
	Value   string `json:"value,omitempty"`
	Type    string `json:"type,omitempty"`
	Primary bool   `json:"primary,omitempty"`
}

type Role struct {
	Value   string `json:"value,omitempty"`
	Type    string `json:"type,omitempty"`
	Primary bool   `json:"primary,omitempty"`
}

type SCIMAttributes struct {
	UserName            string               `json:"user_name,omitempty"`
	ID                  string               `json:"id,omitempty"`
	ExternalID          string               `json:"external_id,omitempty"`
	Active              bool                 `json:"active,omitempty"`
	Groups              []Group              `json:"groups,omitempty"`
	DisplayName         string               `json:"display_name,omitempty"`
	NickName            string               `json:"nick_name,omitempty"`
	ProfileURL          string               `json:"profile_url,omitempty"`
	UserType            string               `json:"user_type,omitempty"`
	Title               string               `json:"title,omitempty"`
	PreferredLanguage   string               `json:"preferred_language,omitempty"`
	Locale              string               `json:"locale,omitempty"`
	Timezone            string               `json:"timezone,omitempty"`
	Emails              []Email              `json:"emails,omitempty"`
	PhoneNumbers        []PhoneNumber        `json:"phone_numbers,omitempty"`
	Addresses           []Address            `json:"addresses,omitempty"`
	Ims                 []IMs                `json:"ims,omitempty"`
	Photos              []Photo              `json:"photos,omitempty"`
	Entitlements        []Entitlement        `json:"entitlements,omitempty"`
	Roles               []Role               `json:"roles,omitempty"`
	X509Certificates    []X509Certificate    `json:"x509certificates,omitempty"`
	Name                *Name                `json:"name,omitempty"`
	EnterpriseExtension *EnterpriseExtension `json:"enterprise_extension,omitempty"`
}

type SCIMConnection struct {
	OrganizationID                   string                             `json:"organization_id,omitempty"`
	ConnectionID                     string                             `json:"connection_id,omitempty"`
	Status                           string                             `json:"status,omitempty"`
	DisplayName                      string                             `json:"display_name,omitempty"`
	IdentityProvider                 string                             `json:"identity_provider,omitempty"`
	BaseURL                          string                             `json:"base_url,omitempty"`
	BearerTokenLastFour              string                             `json:"bearer_token_last_four,omitempty"`
	SCIMGroupImplicitRoleAssignments []SCIMGroupImplicitRoleAssignments `json:"scim_group_implicit_role_assignments,omitempty"`
	NextBearerTokenLastFour          string                             `json:"next_bearer_token_last_four,omitempty"`
	BearerTokenExpiresAt             *time.Time                         `json:"bearer_token_expires_at,omitempty"`
	NextBearerTokenExpiresAt         *time.Time                         `json:"next_bearer_token_expires_at,omitempty"`
}

type SCIMConnectionWithNextToken struct {
	OrganizationID                   string                             `json:"organization_id,omitempty"`
	ConnectionID                     string                             `json:"connection_id,omitempty"`
	Status                           string                             `json:"status,omitempty"`
	DisplayName                      string                             `json:"display_name,omitempty"`
	BaseURL                          string                             `json:"base_url,omitempty"`
	IdentityProvider                 string                             `json:"identity_provider,omitempty"`
	BearerTokenLastFour              string                             `json:"bearer_token_last_four,omitempty"`
	NextBearerToken                  string                             `json:"next_bearer_token,omitempty"`
	SCIMGroupImplicitRoleAssignments []SCIMGroupImplicitRoleAssignments `json:"scim_group_implicit_role_assignments,omitempty"`
	BearerTokenExpiresAt             *time.Time                         `json:"bearer_token_expires_at,omitempty"`
	NextBearerTokenExpiresAt         *time.Time                         `json:"next_bearer_token_expires_at,omitempty"`
}

type SCIMConnectionWithToken struct {
	OrganizationID                   string                             `json:"organization_id,omitempty"`
	ConnectionID                     string                             `json:"connection_id,omitempty"`
	Status                           string                             `json:"status,omitempty"`
	DisplayName                      string                             `json:"display_name,omitempty"`
	IdentityProvider                 string                             `json:"identity_provider,omitempty"`
	BaseURL                          string                             `json:"base_url,omitempty"`
	BearerToken                      string                             `json:"bearer_token,omitempty"`
	SCIMGroupImplicitRoleAssignments []SCIMGroupImplicitRoleAssignments `json:"scim_group_implicit_role_assignments,omitempty"`
	BearerTokenExpiresAt             *time.Time                         `json:"bearer_token_expires_at,omitempty"`
}

// SCIMGroup:
type SCIMGroup struct {
	// GroupID: Globally unique UUID that identifies a specific SCIM Group.
	GroupID string `json:"group_id,omitempty"`
	// GroupName: The name of the SCIM group.
	GroupName string `json:"group_name,omitempty"`
	// OrganizationID: Globally unique UUID that identifies a specific Organization. The organization_id is
	// critical to perform operations on an Organization, so be sure to preserve this value.
	OrganizationID string `json:"organization_id,omitempty"`
	// ConnectionID: The ID of the SCIM connection.
	ConnectionID string `json:"connection_id,omitempty"`
}

// SCIMGroupImplicitRoleAssignments:
type SCIMGroupImplicitRoleAssignments struct {
	// RoleID: The ID of the role.
	RoleID string `json:"role_id,omitempty"`
	// GroupID: The ID of the group.
	GroupID   string `json:"group_id,omitempty"`
	GroupName string `json:"group_name,omitempty"`
}

type X509Certificate struct {
	Value   string `json:"value,omitempty"`
	Type    string `json:"type,omitempty"`
	Primary bool   `json:"primary,omitempty"`
}
