package oidc

import (
	"context"
	"encoding/json"

	"github.com/stytchauth/stytch-go/v8/stytch"
	"github.com/stytchauth/stytch-go/v8/stytch/b2b"
	"github.com/stytchauth/stytch-go/v8/stytch/stytcherror"
)

type Client struct {
	C *stytch.Client
}

func (c *Client) Create(
	ctx context.Context,
	organizationID string,
	body *b2b.OIDCCreateConnectionParams,
) (*b2b.OIDCCreateConnectionResponse, error) {
	path := "/b2b/sso/oidc/" + organizationID

	var jsonBody []byte
	var err error
	if body != nil {
		jsonBody, err = json.Marshal(body)
		if err != nil {
			return nil, stytcherror.NewClientLibraryError(
				"Oops, something seems to have gone wrong " +
					"marshalling the OIDC Create request body")
		}
	}

	var retVal b2b.OIDCCreateConnectionResponse
	err = c.C.NewRequest(ctx, "POST", path, nil, jsonBody, &retVal)
	return &retVal, err
}

func (c *Client) Update(
	ctx context.Context,
	organizationID string,
	connectionID string,
	body *b2b.OIDCUpdateConnectionParams,
) (*b2b.OIDCUpdateConnectionResponse, error) {
	path := "/b2b/sso/oidc/" + organizationID + "/connections/" + connectionID

	var jsonBody []byte
	var err error
	if body != nil {
		jsonBody, err = json.Marshal(body)
		if err != nil {
			return nil, stytcherror.NewClientLibraryError(
				"Oops, something seems to have gone wrong " +
					"marshalling the OIDC Update request body")
		}
	}

	var retVal b2b.OIDCUpdateConnectionResponse
	err = c.C.NewRequest(ctx, "PUT", path, nil, jsonBody, &retVal)
	return &retVal, err
}
