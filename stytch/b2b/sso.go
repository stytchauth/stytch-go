package b2b

type X509Certificate struct {
	CertificateID string `json:"certificate_id"`
	Certificate   string `json:"certificate"`
	Issuer        string `json:"issuer"`
	CreatedAt     string `json:"created_at"`
	ExpiresAt     string `json:"expires_at"`
}

type OIDCConnection struct {
	OrganizationID   string `json:"organization_id"`
	ConnectionID     string `json:"connection_id"`
	Status           string `json:"status"`
	DisplayName      string `json:"display_name"`
	RedirectURL      string `json:"redirect_url"`
	ClientID         string `json:"client_id"`
	ClientSecret     string `json:"client_secret"`
	Issuer           string `json:"issuer"`
	AuthorizationURL string `json:"authorization_url"`
	TokenURL         string `json:"token_url"`
	UserinfoURL      string `json:"userinfo_url"`
	JwksURL          string `json:"jwks_url"`
}

type SAMLConnection struct {
	OrganizationID           string            `json:"organization_id"`
	ConnectionID             string            `json:"connection_id"`
	Status                   string            `json:"status"`
	AttributeMapping         map[string]string `json:"attribute_mapping"`
	IdpEntityID              string            `json:"idp_entity_id"`
	DisplayName              string            `json:"display_name"`
	IdpSsoURL                string            `json:"idp_sso_url"`
	AcsURL                   string            `json:"acs_url"`
	AudienceURI              string            `json:"audience_uri"`
	SigningCertificates      []X509Certificate `json:"signing_certificates"`
	VerificationCertificates []X509Certificate `json:"verification_certificates"`
}

type SSOGetConnectionsResponse struct {
	RequestID       string           `json:"request_id"`
	StatusCode      int              `json:"status_code"`
	SAMLConnections []SAMLConnection `json:"saml_connections"`
	OIDCConnections []OIDCConnection `json:"oidc_connections"`
}

type SSODeleteConnectionResponse struct {
	RequestID    string `json:"request_id"`
	StatusCode   int    `json:"status_code"`
	ConnectionID string `json:"connection_id"`
}

type SSOAuthenticateParams struct {
	SSOToken               string                 `json:"sso_token,omitempty"`
	SessionToken           string                 `json:"session_token,omitempty"`
	SessionDurationMinutes int32                  `json:"session_duration_minutes,omitempty"`
	SessionJWT             string                 `json:"session_jwt,omitempty"`
	SessionCustomClaims    map[string]interface{} `json:"session_custom_claims,omitempty"`
	PKCECodeVerifier       string                 `json:"pkce_code_verifier,omitempty"`
}

type SSOAuthenticateResponse struct {
	RequestID      string        `json:"request_id,omitempty"`
	StatusCode     int           `json:"status_code,omitempty"`
	MemberID       string        `json:"member_id,omitempty"`
	MethodID       string        `json:"method_id,omitempty"`
	ResetSessions  bool          `json:"reset_sessions,omitempty"`
	OrganizationID string        `json:"organization_id,omitempty"`
	Member         Member        `json:"member,omitempty"`
	SessionToken   string        `json:"session_token,omitempty"`
	SessionJWT     string        `json:"session_jwt,omitempty"`
	MemberSession  MemberSession `json:"member_session,omitempty"`
	Organization   Organization  `json:"organization,omitempty"`
}

type OIDCCreateConnectionParams struct {
	OrganizationID string `json:"organization_id"`
	DisplayName    string `json:"display_name,omitempty"`
}

type OIDCCreateConnectionResponse struct {
	RequestID  string         `json:"request_id,omitempty"`
	StatusCode int            `json:"status_code,omitempty"`
	Connection OIDCConnection `json:"connection"`
}

type OIDCUpdateConnectionParams struct {
	OrganizationID   string `json:"organization_id"`
	ConnectionID     string `json:"connection_id"`
	DisplayName      string `json:"display_name,omitempty"`
	ClientID         string `json:"client_id,omitempty"`
	ClientSecret     string `json:"client_secret,omitempty"`
	Issuer           string `json:"issuer,omitempty"`
	AuthorizationURL string `json:"authorization_url,omitempty"`
	TokenURL         string `json:"token_url,omitempty"`
	UserinfoURL      string `json:"userinfo_url,omitempty"`
	JwksURL          string `json:"jwks_url,omitempty"`
}

type OIDCUpdateConnectionResponse struct {
	RequestID  string         `json:"request_id,omitempty"`
	StatusCode int            `json:"status_code,omitempty"`
	Connection OIDCConnection `json:"connection"`
}

type SAMLCreateConnectionParams struct {
	OrganizationID string `json:"organization_id"`
	DisplayName    string `json:"display_name,omitempty"`
}

type SAMLCreateConnectionResponse struct {
	RequestID  string         `json:"request_id,omitempty"`
	StatusCode int            `json:"status_code,omitempty"`
	Connection SAMLConnection `json:"connection"`
}

type SAMLUpdateConnectionParams struct {
	OrganizationID   string            `json:"organization_id"`
	ConnectionID     string            `json:"connection_id"`
	IdpEntityID      string            `json:"idp_entity_id,omitempty"`
	DisplayName      string            `json:"display_name,omitempty"`
	AttributeMapping map[string]string `json:"attribute_mapping,omitempty"`
	X509Certificate  string            `json:"x509_certificate,omitempty"`
	IdpSsoURL        string            `json:"idp_sso_url,omitempty"`
}

type SAMLUpdateConnectionResponse struct {
	RequestID  string         `json:"request_id,omitempty"`
	StatusCode int            `json:"status_code,omitempty"`
	Connection SAMLConnection `json:"connection"`
}

type SAMLDeleteVerificationCertificateResponse struct {
	RequestID     string `json:"request_id,omitempty"`
	StatusCode    int    `json:"status_code,omitempty"`
	CertificateID string `json:"certificate_id"`
}
