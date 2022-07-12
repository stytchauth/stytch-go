package password

import (
	"encoding/json"
	"fmt"

	"github.com/stytchauth/stytch-go/v5/stytch"
	"github.com/stytchauth/stytch-go/v5/stytch/password/email"
	"github.com/stytchauth/stytch-go/v5/stytch/stytcherror"
)

type Client struct {
	C     *stytch.Client
	Email *email.Client
}

func (c *Client) Create(
	body *stytch.PasswordsCreateParams,
) (*stytch.PasswordsCreateResponse, error) {
	path := "/passwords"

	var jsonBody []byte
	var err error
	if body != nil {
		jsonBody, err = json.Marshal(body)
		if err != nil {
			return nil, stytcherror.NewClientLibraryError(
				"Oops, something seems to have gone wrong " +
					"marshalling the /passwords request body")
		}
	}

	var retVal stytch.PasswordsCreateResponse
	err = c.C.NewRequest("POST", path, nil, jsonBody, &retVal)
	return &retVal, err
}

// CreateWithClaims fills in the claims pointer with custom claims from the response.
// Pass in a map with the types of values you're expecting so that this function can marshal
// the claims from the response. See ExampleClient_AuthenticateWithClaims_map, 
// ExampleClient_AuthenticateWithClaims_struct for examples
func (c *Client) CreateWithClaims(
	body *stytch.PasswordsCreateParams,
	claims interface{},
) (*stytch.PasswordsCreateResponse, error) {
	path := "/passwords"

	var jsonBody []byte
	var err error
	if body != nil {
		jsonBody, err = json.Marshal(body)
		if err != nil {
			return nil, stytcherror.NewClientLibraryError(
				"Oops, something seems to have gone wrong " +
					"marshalling the /passwords request body")
		}
	}

	b, err := c.C.RawRequest("POST", path, nil, jsonBody)
	if err != nil {
		return nil, err
	}

	// First extract the Stytch data.
	var retVal stytch.PasswordsCreateResponse
	if err := json.Unmarshal(b, &retVal); err != nil {
		return nil, fmt.Errorf("unmarshal PasswordsCreateResponse: %w", err)
	}

	// Then extract the custom claims. Build a claims wrapper using the caller's `claims` value so
	// the unmarshal fills it.
	wrapper := stytch.SessionWrapper{
		Session: stytch.ClaimsWrapper{
			Claims: claims,
		},
	}
	if err := json.Unmarshal(b, &wrapper); err != nil {
		return nil, fmt.Errorf("unmarshal custom claims: %w", err)
	}
	retVal.Session.CustomClaims = wrapper.Session.Claims
	return &retVal, err
}

func (c *Client) Authenticate(
	body *stytch.PasswordsAuthenticateParams,
) (*stytch.PasswordsAuthenticateResponse, error) {
	path := "/passwords/authenticate"

	var jsonBody []byte
	var err error
	if body != nil {
		jsonBody, err = json.Marshal(body)
		if err != nil {
			return nil, stytcherror.NewClientLibraryError(
				"Oops, something seems to have gone wrong " +
					"marshalling the /passwords/authenticate request body")
		}
	}

	var retVal stytch.PasswordsAuthenticateResponse
	err = c.C.NewRequest("POST", path, nil, jsonBody, &retVal)
	return &retVal, err
}

// AuthenticateWithClaims fills in the claims pointer with custom claims from the response.
// Pass in a map with the types of values you're expecting so that this function can marshal
// the claims from the response. See ExampleClient_AuthenticateWithClaims_map, 
// ExampleClient_AuthenticateWithClaims_struct for examples
func (c *Client) AuthenticateWithClaims(
	body *stytch.PasswordsAuthenticateParams,
	claims interface{},
) (*stytch.PasswordsAuthenticateResponse, error) {
	path := "/passwords/authenticate"

	var jsonBody []byte
	var err error
	if body != nil {
		jsonBody, err = json.Marshal(body)
		if err != nil {
			return nil, stytcherror.NewClientLibraryError(
				"Oops, something seems to have gone wrong " +
					"marshalling the /passwords/authenticate request body")
		}
	}

	b, err := c.C.RawRequest("POST", path, nil, jsonBody)
	if err != nil {
		return nil, err
	}

	// First extract the Stytch data.
	var retVal stytch.PasswordsAuthenticateResponse
	if err := json.Unmarshal(b, &retVal); err != nil {
		return nil, fmt.Errorf("unmarshal PasswordsAuthenticateResponse: %w", err)
	}

	// Then extract the custom claims. Build a claims wrapper using the caller's `claims` value so
	// the unmarshal fills it.
	wrapper := stytch.SessionWrapper{
		Session: stytch.ClaimsWrapper{
			Claims: claims,
		},
	}
	if err := json.Unmarshal(b, &wrapper); err != nil {
		return nil, fmt.Errorf("unmarshal custom claims: %w", err)
	}
	retVal.Session.CustomClaims = wrapper.Session.Claims
	return &retVal, err
}

func (c *Client) StrengthCheck(
	body *stytch.PasswordsStrengthCheckParams,
) (*stytch.PasswordsStrengthCheckResponse, error) {
	path := "/passwords/strength_check"

	var jsonBody []byte
	var err error
	if body != nil {
		jsonBody, err = json.Marshal(body)
		if err != nil {
			return nil, stytcherror.NewClientLibraryError(
				"Oops, something seems to have gone wrong " +
					"marshalling the /passwords/strength_check request body")
		}
	}

	var retVal stytch.PasswordsStrengthCheckResponse
	err = c.C.NewRequest("POST", path, nil, jsonBody, &retVal)
	return &retVal, err
}

func (c *Client) Migrate(
	body *stytch.PasswordsMigrateParams,
) (*stytch.PasswordsMigrateResponse, error) {
	path := "/passwords/migrate"

	var jsonBody []byte
	var err error
	if body != nil {
		jsonBody, err = json.Marshal(body)
		if err != nil {
			return nil, stytcherror.NewClientLibraryError(
				"Oops, something seems to have gone wrong " +
					"marshalling the /passwords/migrate request body")
		}
	}

	var retVal stytch.PasswordsMigrateResponse
	err = c.C.NewRequest("POST", path, nil, jsonBody, &retVal)
	return &retVal, err
}
