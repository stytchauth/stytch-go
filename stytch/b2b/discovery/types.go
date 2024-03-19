package discovery

// !!!
// WARNING: This file is autogenerated
// Only modify code within MANUAL() sections
// or your changes may be overwritten later!
// !!!

import (
	"github.com/stytchauth/stytch-go/v12/stytch/b2b/mfa"
	"github.com/stytchauth/stytch-go/v12/stytch/b2b/organizations"
	"github.com/stytchauth/stytch-go/v12/stytch/b2b/sessions"
)

// DiscoveredOrganization:
type DiscoveredOrganization struct {
	// MemberAuthenticated: Indicates whether the Member has all of the factors needed to fully authenticate to
	// this Organization. If false, the Member may need to complete an MFA step or complete a different primary
	// authentication flow. See the `primary_required` and `mfa_required` fields for more details on each.
	MemberAuthenticated bool `json:"member_authenticated,omitempty"`
	// Organization: The [Organization object](https://stytch.com/docs/b2b/api/organization-object).
	Organization *organizations.Organization `json:"organization,omitempty"`
	// Membership: Information about the membership.
	Membership *Membership `json:"membership,omitempty"`
	// PrimaryRequired: Information about the primary authentication requirements of the Organization.
	PrimaryRequired *sessions.PrimaryRequired `json:"primary_required,omitempty"`
	// MFARequired: Information about the MFA requirements of the Organization and the Member's options for
	// fulfilling MFA.
	MFARequired *mfa.MfaRequired `json:"mfa_required,omitempty"`
}

// Membership:
type Membership struct {
	// Type: Either `active_member`, `pending_member`, `invited_member`, or `eligible_to_join_by_email_domain`
	Type string `json:"type,omitempty"`
	// Details: An object containing additional metadata about the membership, if available.
	Details map[string]any `json:"details,omitempty"`
	// Member: The [Member object](https://stytch.com/docs/b2b/api/member-object) if one already exists, or
	// null if one does not.
	Member *organizations.Member `json:"member,omitempty"`
}
