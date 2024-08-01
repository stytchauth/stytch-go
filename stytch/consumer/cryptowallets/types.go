package cryptowallets

// !!!
// WARNING: This file is autogenerated
// Only modify code within MANUAL() sections
// or your changes may be overwritten later!
// !!!

import (
	"time"

	"github.com/stytchauth/stytch-go/v15/stytch/consumer/sessions"
	"github.com/stytchauth/stytch-go/v15/stytch/consumer/users"
)

// AuthenticateParams: Request type for `CryptoWallets.Authenticate`.
type AuthenticateParams struct {
	// CryptoWalletType: The type of wallet to authenticate. Currently `ethereum` and `solana` are supported.
	// Wallets for any EVM-compatible chains (such as Polygon or BSC) are also supported and are grouped under
	// the `ethereum` type.
	CryptoWalletType string `json:"crypto_wallet_type,omitempty"`
	// CryptoWalletAddress: The crypto wallet address to authenticate.
	CryptoWalletAddress string `json:"crypto_wallet_address,omitempty"`
	// Signature: The signature from the message challenge.
	Signature string `json:"signature,omitempty"`
	// SessionToken: The `session_token` associated with a User's existing Session.
	SessionToken string `json:"session_token,omitempty"`
	// SessionDurationMinutes: Set the session lifetime to be this many minutes from now. This will start a new
	// session if one doesn't already exist,
	//   returning both an opaque `session_token` and `session_jwt` for this session. Remember that the
	// `session_jwt` will have a fixed lifetime of
	//   five minutes regardless of the underlying session duration, and will need to be refreshed over time.
	//
	//   This value must be a minimum of 5 and a maximum of 527040 minutes (366 days).
	//
	//   If a `session_token` or `session_jwt` is provided then a successful authentication will continue to
	// extend the session this many minutes.
	//
	//   If the `session_duration_minutes` parameter is not specified, a Stytch session will not be created.
	SessionDurationMinutes int32 `json:"session_duration_minutes,omitempty"`
	// SessionJWT: The `session_jwt` associated with a User's existing Session.
	SessionJWT string `json:"session_jwt,omitempty"`
	// SessionCustomClaims: Add a custom claims map to the Session being authenticated. Claims are only created
	// if a Session is initialized by providing a value in `session_duration_minutes`. Claims will be included
	// on the Session object and in the JWT. To update a key in an existing Session, supply a new value. To
	// delete a key, supply a null value.
	//
	//   Custom claims made with reserved claims ("iss", "sub", "aud", "exp", "nbf", "iat", "jti") will be
	// ignored. Total custom claims size cannot exceed four kilobytes.
	SessionCustomClaims map[string]any `json:"session_custom_claims,omitempty"`
}

// AuthenticateStartParams: Request type for `CryptoWallets.AuthenticateStart`.
type AuthenticateStartParams struct {
	// CryptoWalletType: The type of wallet to authenticate. Currently `ethereum` and `solana` are supported.
	// Wallets for any EVM-compatible chains (such as Polygon or BSC) are also supported and are grouped under
	// the `ethereum` type.
	CryptoWalletType string `json:"crypto_wallet_type,omitempty"`
	// CryptoWalletAddress: The crypto wallet address to authenticate.
	CryptoWalletAddress string `json:"crypto_wallet_address,omitempty"`
	// UserID: The unique ID of a specific User.
	UserID string `json:"user_id,omitempty"`
	// SessionToken: The `session_token` associated with a User's existing Session.
	SessionToken string `json:"session_token,omitempty"`
	// SessionJWT: The `session_jwt` associated with a User's existing Session.
	SessionJWT string `json:"session_jwt,omitempty"`
	// SiweParams: The parameters for a Sign In With Ethereum (SIWE) message. May only be passed if the
	// `crypto_wallet_type` is `ethereum`.
	SiweParams *SIWEParams `json:"siwe_params,omitempty"`
}

// SIWEParams:
type SIWEParams struct {
	// Domain: Only required if `siwe_params` is passed. The domain that is requesting the crypto wallet
	// signature. Must be an RFC 3986 authority.
	Domain string `json:"domain,omitempty"`
	// URI: Only required if `siwe_params` is passed. An RFC 3986 URI referring to the resource that is the
	// subject of the signing.
	URI string `json:"uri,omitempty"`
	// Resources:  A list of information or references to information the user wishes to have resolved as part
	// of authentication. Every resource must be an RFC 3986 URI.
	Resources []string `json:"resources,omitempty"`
	// ChainID: The EIP-155 Chain ID to which the session is bound. Defaults to 1.
	ChainID uint32 `json:"chain_id,omitempty"`
	// Statement: A human-readable ASCII assertion that the user will sign.
	Statement string `json:"statement,omitempty"`
	// IssuedAt: The time when the message was generated. Defaults to the current time. All timestamps in our
	// API conform to the RFC 3339 standard and are expressed in UTC, e.g. `2021-12-29T12:33:09Z`.
	IssuedAt *time.Time `json:"issued_at,omitempty"`
	// NotBefore: The time when the signed authentication message will become valid. Defaults to the current
	// time. All timestamps in our API conform to the RFC 3339 standard and are expressed in UTC, e.g.
	// `2021-12-29T12:33:09Z`.
	NotBefore *time.Time `json:"not_before,omitempty"`
	// MessageRequestID: A system-specific identifier that may be used to uniquely refer to the sign-in request.
	MessageRequestID string `json:"message_request_id,omitempty"`
}

// AuthenticateResponse: Response type for `CryptoWallets.Authenticate`.
type AuthenticateResponse struct {
	// RequestID: Globally unique UUID that is returned with every API call. This value is important to log for
	// debugging purposes; we may ask for this value to help identify a specific API call when helping you
	// debug an issue.
	RequestID string `json:"request_id,omitempty"`
	// UserID: The unique ID of the affected User.
	UserID string `json:"user_id,omitempty"`
	// SessionToken: A secret token for a given Stytch Session.
	SessionToken string `json:"session_token,omitempty"`
	// SessionJWT: The JSON Web Token (JWT) for a given Stytch Session.
	SessionJWT string `json:"session_jwt,omitempty"`
	// User: The `user` object affected by this API call. See the
	// [Get user endpoint](https://stytch.com/docs/api/get-user) for complete response field details.
	User users.User `json:"user,omitempty"`
	// StatusCode: The HTTP status code of the response. Stytch follows standard HTTP response status code
	// patterns, e.g. 2XX values equate to success, 3XX values are redirects, 4XX are client errors, and 5XX
	// are server errors.
	StatusCode int32 `json:"status_code,omitempty"`
	// Session: If you initiate a Session, by including `session_duration_minutes` in your authenticate call,
	// you'll receive a full Session object in the response.
	//
	//   See [GET sessions](https://stytch.com/docs/api/session-get) for complete response fields.
	//
	Session *sessions.Session `json:"session,omitempty"`
	// SiweParams: The parameters of the Sign In With Ethereum (SIWE) message that was signed.
	SiweParams *SIWEParamsResponse `json:"siwe_params,omitempty"`
}

// AuthenticateStartResponse: Response type for `CryptoWallets.AuthenticateStart`.
type AuthenticateStartResponse struct {
	// RequestID: Globally unique UUID that is returned with every API call. This value is important to log for
	// debugging purposes; we may ask for this value to help identify a specific API call when helping you
	// debug an issue.
	RequestID string `json:"request_id,omitempty"`
	// UserID: The unique ID of the affected User.
	UserID string `json:"user_id,omitempty"`
	// Challenge: A challenge string to be signed by the wallet in order to prove ownership.
	Challenge string `json:"challenge,omitempty"`
	// UserCreated: In `login_or_create` endpoints, this field indicates whether or not a User was just created.
	UserCreated bool `json:"user_created,omitempty"`
	// StatusCode: The HTTP status code of the response. Stytch follows standard HTTP response status code
	// patterns, e.g. 2XX values equate to success, 3XX values are redirects, 4XX are client errors, and 5XX
	// are server errors.
	StatusCode int32 `json:"status_code,omitempty"`
}

// SIWEParamsResponse:
type SIWEParamsResponse struct {
	// Domain: The domain that requested the crypto wallet signature.
	Domain string `json:"domain,omitempty"`
	// URI: An RFC 3986 URI referring to the resource that is the subject of the signing.
	URI string `json:"uri,omitempty"`
	// ChainID: The EIP-155 Chain ID to which the session is bound.
	ChainID uint32 `json:"chain_id,omitempty"`
	// Resources:  A list of information or references to information the user wishes to have resolved as part
	// of authentication. Every resource must be an RFC 3986 URI.
	Resources  []string `json:"resources,omitempty"`
	StatusCode int32    `json:"status_code,omitempty"`
	// IssuedAt: The time when the message was generated. All timestamps in our API conform to the RFC 3339
	// standard and are expressed in UTC, e.g. `2021-12-29T12:33:09Z`.
	IssuedAt *time.Time `json:"issued_at,omitempty"`
	// MessageRequestID: A system-specific identifier that may be used to uniquely refer to the sign-in request.
	MessageRequestID string `json:"message_request_id,omitempty"`
}
