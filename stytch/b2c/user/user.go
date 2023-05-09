package user

import (
	"context"
	"encoding/json"

	"github.com/stytchauth/stytch-go/v9/stytch/b2c"

	"github.com/stytchauth/stytch-go/v9/stytch"
	"github.com/stytchauth/stytch-go/v9/stytch/stytcherror"
)

type Client struct {
	C *stytch.Client
}

func (c *Client) Create(
	ctx context.Context,
	body *b2c.UsersCreateParams,
) (*b2c.UsersCreateResponse, error) {
	var jsonBody []byte
	var err error
	if body != nil {
		jsonBody, err = json.Marshal(body)
		if err != nil {
			return nil, stytcherror.NewClientLibraryError("Oops, something seems to have gone wrong " +
				"marshalling the create user request body")
		}
	}

	var retVal b2c.UsersCreateResponse
	err = c.C.NewRequest(ctx, "POST", "/users", nil, jsonBody, &retVal)
	return &retVal, err
}

func (c *Client) Get(
	ctx context.Context,
	userID string,
) (*b2c.UsersGetResponse, error) {
	path := "/users/" + userID

	var retVal b2c.UsersGetResponse
	err := c.C.NewRequest(ctx, "GET", path, nil, nil, &retVal)
	return &retVal, err
}

func (c *Client) Search(
	ctx context.Context,
	body *b2c.UsersSearchParams,
) (*b2c.UsersSearchResponse, error) {
	var jsonBody []byte
	var err error
	if body != nil {
		jsonBody, err = json.Marshal(body)
		if err != nil {
			return nil, stytcherror.NewClientLibraryError("Oops, something seems to have gone wrong " +
				"marshalling the search users request body")
		}
	}

	var retVal b2c.UsersSearchResponse
	err = c.C.NewRequest(ctx, "POST", "/users/search", nil, jsonBody, &retVal)
	return &retVal, err
}

type iteratorState string

const (
	iteratorStatePending    iteratorState = "Pending"
	iteratorStateInProgress iteratorState = "In Progress"
	iteratorStateErrored    iteratorState = "Errored"
	iteratorStateComplete   iteratorState = "Complete"
)

type UserSearchIterator struct {
	c     *Client
	body  *b2c.UsersSearchParams
	state iteratorState
}

func (i *UserSearchIterator) HasNext() bool {
	return i.state == iteratorStatePending || i.state == iteratorStateInProgress
}

func (i *UserSearchIterator) Next(ctx context.Context) ([]b2c.User, error) {
	res, err := i.c.Search(ctx, i.body)
	if err != nil {
		i.state = iteratorStateErrored
		return nil, err
	}

	i.body.Cursor = res.ResultsMetadata.NextCursor
	if i.body.Cursor == "" {
		i.state = iteratorStateComplete
	}

	return res.Results, nil
}

func (c *Client) SearchAll(body *b2c.UsersSearchParams) *UserSearchIterator {
	return &UserSearchIterator{
		c:     c,
		body:  body,
		state: iteratorStatePending,
	}
}

func (c *Client) Update(
	ctx context.Context,
	userID string,
	body *b2c.UsersUpdateParams,
) (*b2c.UsersUpdateResponse, error) {
	path := "/users/" + userID

	var jsonBody []byte
	var err error
	if body != nil {
		jsonBody, err = json.Marshal(body)
		if err != nil {
			return nil, stytcherror.NewClientLibraryError("Oops, something seems to have gone wrong " +
				"marshalling the update users request body")
		}
	}

	var retVal b2c.UsersUpdateResponse
	err = c.C.NewRequest(ctx, "PUT", path, nil, jsonBody, &retVal)
	return &retVal, err
}

func (c *Client) Delete(
	ctx context.Context,
	userID string,
) (*b2c.UsersDeleteResponse, error) {
	path := "/users/" + userID

	var retVal b2c.UsersDeleteResponse
	err := c.C.NewRequest(ctx, "DELETE", path, nil, nil, &retVal)
	return &retVal, err
}

func (c *Client) DeleteEmail(
	ctx context.Context,
	emailID string,
) (*b2c.UsersDeleteEmailResponse, error) {
	path := "/users/emails/" + emailID

	var retVal b2c.UsersDeleteEmailResponse
	err := c.C.NewRequest(ctx, "DELETE", path, nil, nil, &retVal)
	return &retVal, err
}

func (c *Client) DeletePhoneNumber(
	ctx context.Context,
	phoneID string,
) (*b2c.UsersDeletePhoneNumberResponse, error) {
	path := "/users/phone_numbers/" + phoneID

	var retVal b2c.UsersDeletePhoneNumberResponse
	err := c.C.NewRequest(ctx, "DELETE", path, nil, nil, &retVal)
	return &retVal, err
}

func (c *Client) DeleteWebAuthnRegistration(
	ctx context.Context,
	webAuthnRegistration string,
) (*b2c.UsersDeleteWebAuthnRegistrationResponse, error) {
	path := "/users/webauthn_registrations/" + webAuthnRegistration

	var retVal b2c.UsersDeleteWebAuthnRegistrationResponse
	err := c.C.NewRequest(ctx, "DELETE", path, nil, nil, &retVal)
	return &retVal, err
}

func (c *Client) DeleteBiometricRegistration(
	ctx context.Context,
	biometricRegistrationID string,
) (*b2c.UsersDeleteBiometricRegistrationResponse, error) {
	path := "/users/biometric_registrations/" + biometricRegistrationID

	var retVal b2c.UsersDeleteBiometricRegistrationResponse
	err := c.C.NewRequest(ctx, "DELETE", path, nil, nil, &retVal)
	return &retVal, err
}

func (c *Client) DeleteTOTP(
	ctx context.Context,
	totpID string,
) (*b2c.UsersDeleteTOTPResponse, error) {
	path := "/users/totps/" + totpID

	var retVal b2c.UsersDeleteTOTPResponse
	err := c.C.NewRequest(ctx, "DELETE", path, nil, nil, &retVal)
	return &retVal, err
}

func (c *Client) DeleteCryptoWallet(
	ctx context.Context,
	cryptoWalletID string,
) (*b2c.UsersDeleteCryptoWalletResponse, error) {
	path := "/users/crypto_wallets/" + cryptoWalletID

	var retVal b2c.UsersDeleteCryptoWalletResponse
	err := c.C.NewRequest(ctx, "DELETE", path, nil, nil, &retVal)
	return &retVal, err
}

func (c *Client) DeletePassword(
	ctx context.Context,
	passwordID string,
) (*b2c.UsersDeletePasswordResponse, error) {
	path := "/users/passwords/" + passwordID

	var retVal b2c.UsersDeletePasswordResponse
	err := c.C.NewRequest(ctx, "DELETE", path, nil, nil, &retVal)
	return &retVal, err
}

func (c *Client) DeleteOAuthUserRegistration(
	ctx context.Context,
	oauthUserRegistrationID string,
) (*b2c.UsersDeleteOAuthRegistrationResponse, error) {
	path := "/users/oauth/" + oauthUserRegistrationID

	var retVal b2c.UsersDeleteOAuthRegistrationResponse
	err := c.C.NewRequest(ctx, "DELETE", path, nil, nil, &retVal)
	return &retVal, err
}
