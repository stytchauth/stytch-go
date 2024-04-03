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

	"github.com/stytchauth/stytch-go/v13/stytch"
	"github.com/stytchauth/stytch-go/v13/stytch/consumer/users"
	"github.com/stytchauth/stytch-go/v13/stytch/stytcherror"
)

type UsersClient struct {
	C stytch.Client
}

func NewUsersClient(c stytch.Client) *UsersClient {
	return &UsersClient{
		C: c,
	}
}

// Create: Add a User to Stytch. A `user_id` is returned in the response that can then be used to perform
// other operations within Stytch. An `email` or a `phone_number` is required.
func (c *UsersClient) Create(
	ctx context.Context,
	body *users.CreateParams,
) (*users.CreateResponse, error) {
	var jsonBody []byte
	var err error
	if body != nil {
		jsonBody, err = json.Marshal(body)
		if err != nil {
			return nil, stytcherror.NewClientLibraryError("error marshaling request body")
		}
	}

	headers := make(map[string][]string)

	var retVal users.CreateResponse
	err = c.C.NewRequest(
		ctx,
		"POST",
		"/v1/users",
		nil,
		jsonBody,
		&retVal,
		headers,
	)
	return &retVal, err
}

// Get information about a specific User.
func (c *UsersClient) Get(
	ctx context.Context,
	body *users.GetParams,
) (*users.GetResponse, error) {
	headers := make(map[string][]string)

	var retVal users.GetResponse
	err := c.C.NewRequest(
		ctx,
		"GET",
		fmt.Sprintf("/v1/users/%s", body.UserID),
		nil,
		nil,
		&retVal,
		headers,
	)
	return &retVal, err
}

// Search for Users within your Stytch Project. Submit an empty `query` in the request to return all Users.
func (c *UsersClient) Search(
	ctx context.Context,
	body *users.SearchParams,
) (*users.SearchResponse, error) {
	var jsonBody []byte
	var err error
	if body != nil {
		jsonBody, err = json.Marshal(body)
		if err != nil {
			return nil, stytcherror.NewClientLibraryError("error marshaling request body")
		}
	}

	headers := make(map[string][]string)

	var retVal users.SearchResponse
	err = c.C.NewRequest(
		ctx,
		"POST",
		"/v1/users/search",
		nil,
		jsonBody,
		&retVal,
		headers,
	)
	return &retVal, err
}

// Update a User's attributes.
//
// **Note:** In order to add a new email address or phone number to an existing User object, pass the new
// email address or phone number into the respective `/send` endpoint for the authentication method of your
// choice. If you specify the existing User's `user_id` while calling the `/send` endpoint, the new,
// unverified email address or phone number will be added to the existing User object. If the user
// successfully authenticates within 5 minutes of the `/send` request, the new email address or phone
// number will be marked as verified and remain permanently on the existing Stytch User. Otherwise, it will
// be removed from the User object, and any subsequent login requests using that phone number will create a
// new User. We require this process to guard against an account takeover vulnerability.
func (c *UsersClient) Update(
	ctx context.Context,
	body *users.UpdateParams,
) (*users.UpdateResponse, error) {
	var jsonBody []byte
	var err error
	if body != nil {
		jsonBody, err = json.Marshal(body)
		if err != nil {
			return nil, stytcherror.NewClientLibraryError("error marshaling request body")
		}
	}

	headers := make(map[string][]string)

	var retVal users.UpdateResponse
	err = c.C.NewRequest(
		ctx,
		"PUT",
		fmt.Sprintf("/v1/users/%s", body.UserID),
		nil,
		jsonBody,
		&retVal,
		headers,
	)
	return &retVal, err
}

// ExchangePrimaryFactor: Exchange a user's email address or phone number for another.
//
// Must pass either an `email_address` or a `phone_number`.
//
// This endpoint only works if the user has exactly one factor. You are able to exchange the type of factor
// for another as well, i.e. exchange an `email_address` for a `phone_number`.
//
// Use this endpoint with caution as it performs an admin level action.
func (c *UsersClient) ExchangePrimaryFactor(
	ctx context.Context,
	body *users.ExchangePrimaryFactorParams,
) (*users.ExchangePrimaryFactorResponse, error) {
	var jsonBody []byte
	var err error
	if body != nil {
		jsonBody, err = json.Marshal(body)
		if err != nil {
			return nil, stytcherror.NewClientLibraryError("error marshaling request body")
		}
	}

	headers := make(map[string][]string)

	var retVal users.ExchangePrimaryFactorResponse
	err = c.C.NewRequest(
		ctx,
		"PUT",
		fmt.Sprintf("/v1/users/%s/exchange_primary_factor", body.UserID),
		nil,
		jsonBody,
		&retVal,
		headers,
	)
	return &retVal, err
}

// Delete a User from Stytch.
func (c *UsersClient) Delete(
	ctx context.Context,
	body *users.DeleteParams,
) (*users.DeleteResponse, error) {
	headers := make(map[string][]string)

	var retVal users.DeleteResponse
	err := c.C.NewRequest(
		ctx,
		"DELETE",
		fmt.Sprintf("/v1/users/%s", body.UserID),
		nil,
		nil,
		&retVal,
		headers,
	)
	return &retVal, err
}

// DeleteEmail: Delete an email from a User.
func (c *UsersClient) DeleteEmail(
	ctx context.Context,
	body *users.DeleteEmailParams,
) (*users.DeleteEmailResponse, error) {
	headers := make(map[string][]string)

	var retVal users.DeleteEmailResponse
	err := c.C.NewRequest(
		ctx,
		"DELETE",
		fmt.Sprintf("/v1/users/emails/%s", body.EmailID),
		nil,
		nil,
		&retVal,
		headers,
	)
	return &retVal, err
}

// DeletePhoneNumber: Delete a phone number from a User.
func (c *UsersClient) DeletePhoneNumber(
	ctx context.Context,
	body *users.DeletePhoneNumberParams,
) (*users.DeletePhoneNumberResponse, error) {
	headers := make(map[string][]string)

	var retVal users.DeletePhoneNumberResponse
	err := c.C.NewRequest(
		ctx,
		"DELETE",
		fmt.Sprintf("/v1/users/phone_numbers/%s", body.PhoneID),
		nil,
		nil,
		&retVal,
		headers,
	)
	return &retVal, err
}

// DeleteWebAuthnRegistration: Delete a WebAuthn registration from a User.
func (c *UsersClient) DeleteWebAuthnRegistration(
	ctx context.Context,
	body *users.DeleteWebAuthnRegistrationParams,
) (*users.DeleteWebAuthnRegistrationResponse, error) {
	headers := make(map[string][]string)

	var retVal users.DeleteWebAuthnRegistrationResponse
	err := c.C.NewRequest(
		ctx,
		"DELETE",
		fmt.Sprintf("/v1/users/webauthn_registrations/%s", body.WebAuthnRegistrationID),
		nil,
		nil,
		&retVal,
		headers,
	)
	return &retVal, err
}

// DeleteBiometricRegistration: Delete a biometric registration from a User.
func (c *UsersClient) DeleteBiometricRegistration(
	ctx context.Context,
	body *users.DeleteBiometricRegistrationParams,
) (*users.DeleteBiometricRegistrationResponse, error) {
	headers := make(map[string][]string)

	var retVal users.DeleteBiometricRegistrationResponse
	err := c.C.NewRequest(
		ctx,
		"DELETE",
		fmt.Sprintf("/v1/users/biometric_registrations/%s", body.BiometricRegistrationID),
		nil,
		nil,
		&retVal,
		headers,
	)
	return &retVal, err
}

// DeleteTOTP: Delete a TOTP from a User.
func (c *UsersClient) DeleteTOTP(
	ctx context.Context,
	body *users.DeleteTOTPParams,
) (*users.DeleteTOTPResponse, error) {
	headers := make(map[string][]string)

	var retVal users.DeleteTOTPResponse
	err := c.C.NewRequest(
		ctx,
		"DELETE",
		fmt.Sprintf("/v1/users/totps/%s", body.TOTPID),
		nil,
		nil,
		&retVal,
		headers,
	)
	return &retVal, err
}

// DeleteCryptoWallet: Delete a crypto wallet from a User.
func (c *UsersClient) DeleteCryptoWallet(
	ctx context.Context,
	body *users.DeleteCryptoWalletParams,
) (*users.DeleteCryptoWalletResponse, error) {
	headers := make(map[string][]string)

	var retVal users.DeleteCryptoWalletResponse
	err := c.C.NewRequest(
		ctx,
		"DELETE",
		fmt.Sprintf("/v1/users/crypto_wallets/%s", body.CryptoWalletID),
		nil,
		nil,
		&retVal,
		headers,
	)
	return &retVal, err
}

// DeletePassword: Delete a password from a User.
func (c *UsersClient) DeletePassword(
	ctx context.Context,
	body *users.DeletePasswordParams,
) (*users.DeletePasswordResponse, error) {
	headers := make(map[string][]string)

	var retVal users.DeletePasswordResponse
	err := c.C.NewRequest(
		ctx,
		"DELETE",
		fmt.Sprintf("/v1/users/passwords/%s", body.PasswordID),
		nil,
		nil,
		&retVal,
		headers,
	)
	return &retVal, err
}

// DeleteOAuthRegistration: Delete an OAuth registration from a User.
func (c *UsersClient) DeleteOAuthRegistration(
	ctx context.Context,
	body *users.DeleteOAuthRegistrationParams,
) (*users.DeleteOAuthRegistrationResponse, error) {
	headers := make(map[string][]string)

	var retVal users.DeleteOAuthRegistrationResponse
	err := c.C.NewRequest(
		ctx,
		"DELETE",
		fmt.Sprintf("/v1/users/oauth/%s", body.OAuthUserRegistrationID),
		nil,
		nil,
		&retVal,
		headers,
	)
	return &retVal, err
}
