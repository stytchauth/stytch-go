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

	"github.com/mitchellh/mapstructure"
	"github.com/stytchauth/stytch-go/v11/stytch"
	"github.com/stytchauth/stytch-go/v11/stytch/b2b/passwords"
	"github.com/stytchauth/stytch-go/v11/stytch/stytcherror"
)

type PasswordsClient struct {
	C                stytch.Client
	Email            *PasswordsEmailClient
	Sessions         *PasswordsSessionsClient
	ExistingPassword *PasswordsExistingPasswordClient
}

func NewPasswordsClient(c stytch.Client) *PasswordsClient {
	return &PasswordsClient{
		C:                c,
		Email:            NewPasswordsEmailClient(c),
		Sessions:         NewPasswordsSessionsClient(c),
		ExistingPassword: NewPasswordsExistingPasswordClient(c),
	}
}

// StrengthCheck: This API allows you to check whether the user’s provided password is valid, and to
// provide feedback to the user on how to increase the strength of their password.
//
// This endpoint adapts to your Project's password strength configuration. If you're using
// [zxcvbn](https://stytch.com/docs/guides/passwords/strength-policy), the default, your passwords are
// considered valid if the strength score is >= 3. If you're using
// [LUDS](https://stytch.com/docs/guides/passwords/strength-policy), your passwords are considered valid if
// they meet the requirements that you've set with Stytch. You may update your password strength
// configuration in the [stytch dashboard](https://stytch.com/dashboard/password-strength-config).
//
// ## Password feedback
// The zxcvbn_feedback and luds_feedback objects contains relevant fields for you to relay feedback to
// users that failed to create a strong enough password.
//
// If you're using [zxcvbn](https://stytch.com/docs/guides/passwords/strength-policy), the feedback object
// will contain warning and suggestions for any password that does not meet the
// [zxcvbn](https://stytch.com/docs/guides/passwords/strength-policy) strength requirements. You can return
// these strings directly to the user to help them craft a strong password.
//
// If you're using [LUDS](https://stytch.com/docs/guides/passwords/strength-policy), the feedback object
// will contain a collection of fields that the user failed or passed. You'll want to prompt the user to
// create a password that meets all requirements that they failed.
func (c *PasswordsClient) StrengthCheck(
	ctx context.Context,
	body *passwords.StrengthCheckParams,
) (*passwords.StrengthCheckResponse, error) {
	var jsonBody []byte
	var err error
	if body != nil {
		jsonBody, err = json.Marshal(body)
		if err != nil {
			return nil, stytcherror.NewClientLibraryError("error marshaling request body")
		}
	}

	var retVal passwords.StrengthCheckResponse
	err = c.C.NewRequest(
		ctx,
		"POST",
		"/v1/b2b/passwords/strength_check",
		nil,
		jsonBody,
		&retVal,
	)
	return &retVal, err
}

// Migrate: Adds an existing password to a member's email that doesn't have a password yet. We support
// migrating members from passwords stored with bcrypt, scrypt, argon2, MD-5, SHA-1, and PBKDF2. This
// endpoint has a rate limit of 100 requests per second.
func (c *PasswordsClient) Migrate(
	ctx context.Context,
	body *passwords.MigrateParams,
) (*passwords.MigrateResponse, error) {
	var jsonBody []byte
	var err error
	if body != nil {
		jsonBody, err = json.Marshal(body)
		if err != nil {
			return nil, stytcherror.NewClientLibraryError("error marshaling request body")
		}
	}

	var retVal passwords.MigrateResponse
	err = c.C.NewRequest(
		ctx,
		"POST",
		"/v1/b2b/passwords/migrate",
		nil,
		jsonBody,
		&retVal,
	)
	return &retVal, err
}

// Authenticate a member with their email address and password. This endpoint verifies that the member has
// a password currently set, and that the entered password is correct.
//
// If you have breach detection during authentication enabled in your
// [password strength policy](https://stytch.com/docs/b2b/guides/passwords/strength-policies) and the
// member's credentials have appeared in the HaveIBeenPwned dataset, this endpoint will return a
// `member_reset_password` error even if the member enters a correct password. We force a password reset in
// this case to ensure that the member is the legitimate owner of the email address and not a malicious
// actor abusing the compromised credentials.
//
// If the Member is required to complete MFA to log in to the Organization, the returned value of
// `member_authenticated` will be `false`, and an `intermediate_session_token` will be returned.
// The `intermediate_session_token` can be passed into the
// [OTP SMS Authenticate endpoint](https://stytch.com/docs/b2b/api/authenticate-otp-sms) to complete the
// MFA step and acquire a full member session.
// The `session_duration_minutes` and `session_custom_claims` parameters will be ignored.
//
// If a valid `session_token` or `session_jwt` is passed in, the Member will not be required to complete an
// MFA step.
func (c *PasswordsClient) Authenticate(
	ctx context.Context,
	body *passwords.AuthenticateParams,
) (*passwords.AuthenticateResponse, error) {
	var jsonBody []byte
	var err error
	if body != nil {
		jsonBody, err = json.Marshal(body)
		if err != nil {
			return nil, stytcherror.NewClientLibraryError("error marshaling request body")
		}
	}

	var retVal passwords.AuthenticateResponse
	err = c.C.NewRequest(
		ctx,
		"POST",
		"/v1/b2b/passwords/authenticate",
		nil,
		jsonBody,
		&retVal,
	)
	return &retVal, err
}

// AuthenticateWithClaims fills in the claims pointer with custom claims from the response.
// Pass in a map with the types of values you're expecting so that this function can marshal
// the claims from the response. See ExampleClient_AuthenticateWithClaims_map,
// ExampleClient_AuthenticateWithClaims_struct for examples
func (c *PasswordsClient) AuthenticateWithClaims(
	ctx context.Context,
	body *passwords.AuthenticateParams,
	claims any,
) (*passwords.AuthenticateResponse, error) {
	var jsonBody []byte
	var err error
	if body != nil {
		jsonBody, err = json.Marshal(body)
		if err != nil {
			return nil, stytcherror.NewClientLibraryError("error marshaling request body")
		}
	}

	b, err := c.C.RawRequest(
		ctx,
		"POST",
		"/v1/b2b/passwords/authenticate",
		nil,
		jsonBody,
	)
	if err != nil {
		return nil, err
	}

	// First extract the Stytch data.
	var retVal passwords.AuthenticateResponse
	if err := json.Unmarshal(b, &retVal); err != nil {
		return nil, fmt.Errorf("unmarshal passwords.AuthenticateResponse: %w", err)
	}

	if claims == nil {
		return &retVal, nil
	}

	if m, ok := claims.(*map[string]any); ok {
		*m = retVal.MemberSession.CustomClaims
		return &retVal, nil
	}

	// This is where we need to convert claims into a claimsMap
	decoder, err := mapstructure.NewDecoder(&mapstructure.DecoderConfig{
		Result:  &claims,
		TagName: "json",
	})
	if err != nil {
		return nil, err
	}

	err = decoder.Decode(retVal.MemberSession.CustomClaims)
	if err != nil {
		return nil, err
	}

	return &retVal, err
}
