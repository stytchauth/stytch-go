package consumer

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
	"github.com/stytchauth/stytch-go/v15/stytch"
	"github.com/stytchauth/stytch-go/v15/stytch/consumer/passwords"
	"github.com/stytchauth/stytch-go/v15/stytch/stytcherror"
)

type PasswordsClient struct {
	C                stytch.Client
	Email            *PasswordsEmailClient
	ExistingPassword *PasswordsExistingPasswordClient
	Sessions         *PasswordsSessionsClient
}

func NewPasswordsClient(c stytch.Client) *PasswordsClient {
	return &PasswordsClient{
		C: c,

		Email:            NewPasswordsEmailClient(c),
		ExistingPassword: NewPasswordsExistingPasswordClient(c),
		Sessions:         NewPasswordsSessionsClient(c),
	}
}

// Create a new user with a password. If `session_duration_minutes` is specified, a new session will be
// started as well.
//
// If a user with this email already exists in your Stytch project, this endpoint will return a
// `duplicate_email` error. To add a password to an existing passwordless user, you'll need to either call
// the [Migrate password endpoint](https://stytch.com/docs/api/password-migrate) or prompt the user to
// complete one of our password reset flows.
//
// This endpoint will return an error if the password provided does not meet our strength requirements,
// which you can check beforehand via the
// [Password strength check endpoint](https://stytch.com/docs/api/password-strength-check).
//
// When creating new Passwords users, it's good practice to enforce an email verification flow. We'd
// recommend checking out our
// [Email verification guide](https://stytch.com/docs/guides/passwords/email-verification/overview) for
// more information.
func (c *PasswordsClient) Create(
	ctx context.Context,
	body *passwords.CreateParams,
) (*passwords.CreateResponse, error) {
	var jsonBody []byte
	var err error
	if body != nil {
		jsonBody, err = json.Marshal(body)
		if err != nil {
			return nil, stytcherror.NewClientLibraryError("error marshaling request body")
		}
	}

	headers := make(map[string][]string)

	var retVal passwords.CreateResponse
	err = c.C.NewRequest(
		ctx,
		stytch.RequestParams{
			Method:      "POST",
			Path:        "/v1/passwords",
			QueryParams: nil,
			Body:        jsonBody,
			V:           &retVal,
			Headers:     headers,
		},
	)
	return &retVal, err
}

// Authenticate a user with their email address and password. This endpoint verifies that the user has a
// password currently set, and that the entered password is correct. There are two instances where the
// endpoint will return a `reset_password` error even if they enter their previous password:
//
// **One:** The user’s credentials appeared in the HaveIBeenPwned dataset. We force a password reset to
// ensure that the user is the legitimate owner of the email address, and not a malicious actor abusing the
// compromised credentials.
//
// **Two:** A user that has previously authenticated with email/password uses a passwordless authentication
// method tied to the same email address (e.g. Magic Links, Google OAuth) for the first time. Any
// subsequent email/password authentication attempt will result in this error. We force a password reset in
// this instance in order to safely deduplicate the account by email address, without introducing the risk
// of a pre-hijack account takeover attack.
//
// Imagine a bad actor creates many accounts using passwords and the known email addresses of their
// victims. If a victim comes to the site and logs in for the first time with an email-based passwordless
// authentication method then both the victim and the bad actor have credentials to access to the same
// account. To prevent this, any further email/password login attempts first require a password reset which
// can only be accomplished by someone with access to the underlying email address.
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

	headers := make(map[string][]string)

	var retVal passwords.AuthenticateResponse
	err = c.C.NewRequest(
		ctx,
		stytch.RequestParams{
			Method:      "POST",
			Path:        "/v1/passwords/authenticate",
			QueryParams: nil,
			Body:        jsonBody,
			V:           &retVal,
			Headers:     headers,
		},
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

	headers := make(map[string][]string)

	b, err := c.C.RawRequest(
		ctx,
		stytch.RequestParams{
			Method:      "POST",
			Path:        "/v1/passwords/authenticate",
			QueryParams: nil,
			Body:        jsonBody,
			Headers:     headers,
		},
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
		*m = retVal.Session.CustomClaims
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

	err = decoder.Decode(retVal.Session.CustomClaims)
	if err != nil {
		return nil, err
	}

	return &retVal, err
}

// StrengthCheck: This API allows you to check whether or not the user’s provided password is valid, and to
// provide feedback to the user on how to increase the strength of their password.
//
// This endpoint adapts to your Project's password strength configuration. If you're using
// [zxcvbn](https://stytch.com/docs/guides/passwords/strength-policy), the default, your passwords are
// considered valid if the strength score is >= 3. If you're using
// [LUDS](https://stytch.com/docs/guides/passwords/strength-policy), your passwords are considered valid if
// they meet the requirements that you've set with Stytch. You may update your password strength
// configuration in the [stytch dashboard](https://stytch.com/dashboard/password-strength-config).
//
// ### Password feedback
//
// The `feedback` object contains relevant fields for you to relay feedback to users that failed to create
// a strong enough password.
//
// If you're using zxcvbn, the `feedback` object will contain `warning` and `suggestions` for any password
// that does not meet the zxcvbn strength requirements. You can return these strings directly to the user
// to help them craft a strong password.
//
// If you're using LUDS, the `feedback` object will contain an object named `luds_requirements` which
// contain a collection of fields that the user failed or passed. You'll want to prompt the user to create
// a password that meets all of the requirements that they failed.
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

	headers := make(map[string][]string)

	var retVal passwords.StrengthCheckResponse
	err = c.C.NewRequest(
		ctx,
		stytch.RequestParams{
			Method:      "POST",
			Path:        "/v1/passwords/strength_check",
			QueryParams: nil,
			Body:        jsonBody,
			V:           &retVal,
			Headers:     headers,
		},
	)
	return &retVal, err
}

// Migrate: Adds an existing password to a User's email that doesn't have a password yet. We support
// migrating users from passwords stored with `bcrypt`, `scrypt`, `argon2`, `MD-5`, `SHA-1`, or `PBKDF2`.
// This endpoint has a rate limit of 100 requests per second.
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

	headers := make(map[string][]string)

	var retVal passwords.MigrateResponse
	err = c.C.NewRequest(
		ctx,
		stytch.RequestParams{
			Method:      "POST",
			Path:        "/v1/passwords/migrate",
			QueryParams: nil,
			Body:        jsonBody,
			V:           &retVal,
			Headers:     headers,
		},
	)
	return &retVal, err
}
