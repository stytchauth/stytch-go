package sso

import (
	"context"
	"encoding/json"

	"github.com/stytchauth/stytch-go/v9/stytch"
	"github.com/stytchauth/stytch-go/v9/stytch/b2b"
	"github.com/stytchauth/stytch-go/v9/stytch/b2b/sso/oidc"
	"github.com/stytchauth/stytch-go/v9/stytch/b2b/sso/saml"
	"github.com/stytchauth/stytch-go/v9/stytch/stytcherror"
)

type Client struct {
	C    *stytch.Client
	SAML *saml.Client
	OIDC *oidc.Client
}

func (c *Client) Get(
	ctx context.Context,
	organizationID string,
) (*b2b.SSOGetConnectionsResponse, error) {
	path := "/b2b/sso/" + organizationID

	var retVal b2b.SSOGetConnectionsResponse
	err := c.C.NewRequest(ctx, "GET", path, nil, nil, &retVal)
	return &retVal, err
}

func (c *Client) Authenticate(
	ctx context.Context,
	body *b2b.SSOAuthenticateParams,
) (*b2b.SSOAuthenticateResponse, error) {
	path := "/b2b/sso/authenticate"

	var jsonBody []byte
	var err error
	if body != nil {
		jsonBody, err = json.Marshal(body)
		if err != nil {
			return nil, stytcherror.NewClientLibraryError(
				"Oops, something seems to have gone wrong " +
					"marshalling the /b2b/sso/authenticate request body")
		}
	}

	var retVal b2b.SSOAuthenticateResponse
	err = c.C.NewRequest(ctx, "POST", path, nil, jsonBody, &retVal)
	return &retVal, err
}

func (c *Client) Delete(
	ctx context.Context,
	organizationID string,
	connectionID string,
) (*b2b.SSODeleteConnectionResponse, error) {
	path := "/b2b/sso/" + organizationID + "/connections/" + connectionID

	var retVal b2b.SSODeleteConnectionResponse
	err := c.C.NewRequest(ctx, "DELETE", path, nil, nil, &retVal)
	return &retVal, err
}
