package stytch

// !!!
// WARNING: This file is autogenerated
// Only modify code within MANUAL() sections
// or your changes may be overwritten later!
// !!!

type MultiTenantMagicLinksMultitenantmagiclinksemailloginorsignupParams struct {
	OrganizationID    string `json:"organization_id,omitempty"`
	EmailAddress      string `json:"email_address,omitempty"`
	LoginRedirectUrl  string `json:"login_redirect_url,omitempty"`
	SignupRedirectUrl string `json:"signup_redirect_url,omitempty"`
	PkceCodeChallenge string `json:"pkce_code_challenge,omitempty"`
	LoginTemplateID   string `json:"login_template_id,omitempty"`
	SignupTemplateID  string `json:"signup_template_id,omitempty"`
	Locale            string `json:"locale,omitempty"`
}
type MultiTenantMagicLinksMultitenantmagiclinksemailinviteParams struct {
	OrganizationID    string         `json:"organization_id,omitempty"`
	EmailAddress      string         `json:"email_address,omitempty"`
	InviteRedirectUrl string         `json:"invite_redirect_url,omitempty"`
	InvitedByMemberID string         `json:"invited_by_member_id,omitempty"`
	Name              string         `json:"name,omitempty"`
	TrustedMetadata   map[string]any `json:"trusted_metadata,omitempty"`
	UntrustedMetadata map[string]any `json:"untrusted_metadata,omitempty"`
	InviteTemplateID  string         `json:"invite_template_id,omitempty"`
	Locale            string         `json:"locale,omitempty"`
}
type MultiTenantMagicLinksMultitenantmagiclinksauthenticateParams struct {
	MagicLinksToken        string         `json:"magic_links_token,omitempty"`
	PkceCodeVerifier       string         `json:"pkce_code_verifier,omitempty"`
	SessionToken           string         `json:"session_token,omitempty"`
	SessionJwt             string         `json:"session_jwt,omitempty"`
	SessionDurationMinutes int32          `json:"session_duration_minutes,omitempty"`
	SessionCustomClaims    map[string]any `json:"session_custom_claims,omitempty"`
}
type MultiTenantMagicLinksB2BmagiclinksemaildiscoverysendParams struct {
	EmailAddress         string `json:"email_address,omitempty"`
	DiscoveryRedirectUrl string `json:"discovery_redirect_url,omitempty"`
	PkceCodeChallenge    string `json:"pkce_code_challenge,omitempty"`
	LoginTemplateID      string `json:"login_template_id,omitempty"`
	Locale               string `json:"locale,omitempty"`
}
type MultiTenantMagicLinksB2BmagiclinksdiscoveryauthenticateParams struct {
	DiscoveryMagicLinksToken string `json:"discovery_magic_links_token,omitempty"`
	PkceCodeVerifier         string `json:"pkce_code_verifier,omitempty"`
}

type MultiTenantMagicLinksMultitenantmagiclinksemailloginorsignupResponse struct {
	RequestID     string       `json:"request_id,omitempty"`
	MemberID      string       `json:"member_id,omitempty"`
	MemberCreated bool         `json:"member_created,omitempty"`
	Member        Member       `json:"member,omitempty"`
	Organization  Organization `json:"organization,omitempty"`
}
type MultiTenantMagicLinksMultitenantmagiclinksemailinviteResponse struct {
	RequestID    string       `json:"request_id,omitempty"`
	MemberID     string       `json:"member_id,omitempty"`
	Member       Member       `json:"member,omitempty"`
	Organization Organization `json:"organization,omitempty"`
}
type MultiTenantMagicLinksMultitenantmagiclinksauthenticateResponse struct {
	RequestID      string        `json:"request_id,omitempty"`
	MemberID       string        `json:"member_id,omitempty"`
	MethodID       string        `json:"method_id,omitempty"`
	ResetSessions  bool          `json:"reset_sessions,omitempty"`
	OrganizationID string        `json:"organization_id,omitempty"`
	Member         Member        `json:"member,omitempty"`
	SessionToken   string        `json:"session_token,omitempty"`
	SessionJwt     string        `json:"session_jwt,omitempty"`
	MemberSession  MemberSession `json:"member_session,omitempty"`
	Organization   Organization  `json:"organization,omitempty"`
}
type MultiTenantMagicLinksB2BmagiclinksemaildiscoverysendResponse struct {
	RequestID string `json:"request_id,omitempty"`
}
type MultiTenantMagicLinksB2BmagiclinksdiscoveryauthenticateResponse struct {
	RequestID                string                   `json:"request_id,omitempty"`
	IntermediateSessionToken string                   `json:"intermediate_session_token,omitempty"`
	EmailAddress             string                   `json:"email_address,omitempty"`
	DiscoveredOrganizations  []DiscoveredOrganization `json:"discovered_organizations,omitempty"`
}
