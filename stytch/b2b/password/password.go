package password

import (
	"context"
	"encoding/json"

	"github.com/stytchauth/stytch-go/v9/stytch/b2b"
	"github.com/stytchauth/stytch-go/v9/stytch/b2b/password/email"
	"github.com/stytchauth/stytch-go/v9/stytch/b2b/password/existingpassword"
	"github.com/stytchauth/stytch-go/v9/stytch/b2b/password/session"

	"github.com/stytchauth/stytch-go/v9/stytch"
	"github.com/stytchauth/stytch-go/v9/stytch/stytcherror"
)

type Client struct {
	C                *stytch.Client
	Email            *email.Client
	ExistingPassword *existingpassword.Client
	Session          *session.Client
}

func (c *Client) Authenticate(
	ctx context.Context,
	body *b2b.PasswordsAuthenticateParams,
) (*b2b.PasswordsAuthenticateResponse, error) {
	path := "/b2b/passwords/authenticate"

	var jsonBody []byte
	var err error
	if body != nil {
		jsonBody, err = json.Marshal(body)
		if err != nil {
			return nil, stytcherror.NewClientLibraryError(
				"Oops, something seems to have gone wrong " +
					"marshalling the /b2b/passwords/authenticate request body")
		}
	}

	var retVal b2b.PasswordsAuthenticateResponse
	err = c.C.NewRequest(ctx, "POST", path, nil, jsonBody, &retVal)
	return &retVal, err
}

func (c *Client) StrengthCheck(
	ctx context.Context,
	body *b2b.PasswordsStrengthCheckParams,
) (*b2b.PasswordsStrengthCheckResponse, error) {
	path := "/b2b/passwords/strength_check"

	var jsonBody []byte
	var err error
	if body != nil {
		jsonBody, err = json.Marshal(body)
		if err != nil {
			return nil, stytcherror.NewClientLibraryError(
				"Oops, something seems to have gone wrong " +
					"marshalling the /b2b/passwords/strength_check request body")
		}
	}

	var retVal b2b.PasswordsStrengthCheckResponse
	err = c.C.NewRequest(ctx, "POST", path, nil, jsonBody, &retVal)
	return &retVal, err
}

func (c *Client) Migrate(
	ctx context.Context,
	body *b2b.PasswordsMigrateParams,
) (*b2b.PasswordsMigrateResponse, error) {
	path := "/b2b/passwords/migrate"

	var jsonBody []byte
	var err error
	if body != nil {
		jsonBody, err = json.Marshal(body)
		if err != nil {
			return nil, stytcherror.NewClientLibraryError(
				"Oops, something seems to have gone wrong " +
					"marshalling the /b2b/passwords/migrate request body")
		}
	}

	var retVal b2b.PasswordsMigrateResponse
	err = c.C.NewRequest(ctx, "POST", path, nil, jsonBody, &retVal)
	return &retVal, err
}

func (c *Client) Delete(
	ctx context.Context,
	organizationID string,
	memberPasswordID string,
) (*b2b.PasswordsDeleteResponse, error) {
	path := "/b2b/organizations/" + organizationID + "/members/passwords/" + memberPasswordID

	var retVal b2b.PasswordsDeleteResponse
	err := c.C.NewRequest(ctx, "POST", path, nil, nil, &retVal)
	return &retVal, err
}
