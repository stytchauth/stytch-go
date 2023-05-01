package b2b

type X509Certificate struct {
	CertificateID string `json:"certificate_id,omitempty"`
	Certificate   string `json:"certificate,omitempty"`
	Issuer        string `json:"issuer,omitempty"`
	CreatedAt     string `json:"created_at,omitempty"`
	ExpiresAt     string `json:"expires_at,omitempty"`
}

type OIDCConnection struct {
	OrganizationID   string `json:"organization_id,omitempty"`
	ConnectionID     string `json:"connection_id,omitempty"`
	Status           string `json:"status,omitempty"`
	DisplayName      string `json:"display_name,omitempty"`
	RedirectURL      string `json:"redirect_url,omitempty"`
	ClientID         string `json:"client_id,omitempty"`
	ClientSecret     string `json:"client_secret,omitempty"`
	Issuer           string `json:"issuer,omitempty"`
	AuthorizationURL string `json:"authorization_url,omitempty"`
	TokenURL         string `json:"token_url,omitempty"`
	UserinfoURL      string `json:"userinfo_url,omitempty"`
	JwksURL          string `json:"jwks_url,omitempty"`
}

type SAMLConnection struct {
	OrganizationID           string            `json:"organization_id,omitempty"`
	ConnectionID             string            `json:"connection_id,omitempty"`
	Status                   string            `json:"status,omitempty"`
	AttributeMapping         map[string]string `json:"attribute_mapping,omitempty"`
	IdpEntityID              string            `json:"idp_entity_id,omitempty"`
	DisplayName              string            `json:"display_name,omitempty"`
	IdpSsoURL                string            `json:"idp_sso_url,omitempty"`
	AcsURL                   string            `json:"acs_url,omitempty"`
	AudienceURI              string            `json:"audience_uri,omitempty"`
	SigningCertificates      []X509Certificate `json:"signing_certificates,omitempty"`
	VerificationCertificates []X509Certificate `json:"verification_certificates,omitempty"`
}

type SSOGetConnectionsResponse struct {
	RequestID       string           `json:"request_id,omitempty"`
	StatusCode      int              `json:"status_code,omitempty"`
	SAMLConnections []SAMLConnection `json:"saml_connections,omitempty"`
	OIDCConnections []OIDCConnection `json:"oidc_connections,omitempty"`
}

type SSODeleteConnectionResponse struct {
	RequestID    string `json:"request_id,omitempty"`
	StatusCode   int    `json:"status_code,omitempty"`
	ConnectionID string `json:"connection_id,omitempty"`
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
	DisplayName string `json:"display_name,omitempty"`
}

type OIDCCreateConnectionResponse struct {
	RequestID  string         `json:"request_id,omitempty"`
	StatusCode int            `json:"status_code,omitempty"`
	Connection OIDCConnection `json:"connection,omitempty"`
}

type OIDCUpdateConnectionParams struct {
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
	Connection OIDCConnection `json:"connection,omitempty"`
}

type SAMLCreateConnectionParams struct {
	DisplayName string `json:"display_name,omitempty"`
}

type SAMLCreateConnectionResponse struct {
	RequestID  string         `json:"request_id,omitempty"`
	StatusCode int            `json:"status_code,omitempty"`
	Connection SAMLConnection `json:"connection,omitempty"`
}

type SAMLUpdateConnectionParams struct {
	IdpEntityID      string            `json:"idp_entity_id,omitempty"`
	DisplayName      string            `json:"display_name,omitempty"`
	AttributeMapping map[string]string `json:"attribute_mapping,omitempty"`
	X509Certificate  string            `json:"x509_certificate,omitempty"`
	IdpSsoURL        string            `json:"idp_sso_url,omitempty"`
}

type SAMLUpdateConnectionResponse struct {
	RequestID  string         `json:"request_id,omitempty"`
	StatusCode int            `json:"status_code,omitempty"`
	Connection SAMLConnection `json:"connection,omitempty"`
}

type SAMLDeleteVerificationCertificateResponse struct {
	RequestID     string `json:"request_id,omitempty"`
	StatusCode    int    `json:"status_code,omitempty"`
	CertificateID string `json:"certificate_id,omitempty"`
}
