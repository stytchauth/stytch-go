package scim

// !!!
// WARNING: This file is autogenerated
// Only modify code within MANUAL() sections
// or your changes may be overwritten later!
// !!!

import (
	"time"
)

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

type SCIMGroupImplicitRoleAssignments struct {
	RoleID    string `json:"role_id,omitempty"`
	GroupID   string `json:"group_id,omitempty"`
	GroupName string `json:"group_name,omitempty"`
}
