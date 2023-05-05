package session

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

func (c *Client) Get(
	ctx context.Context,
	body *b2b.SessionGetParams,
) (*b2b.SessionGetResponse, error) {
	path := "/b2b/sessions"

	queryParams := make(map[string]string)
	if body != nil {
		queryParams["organization_id"] = body.OrganizationID
		queryParams["member_id"] = body.MemberID
	}

	var retVal b2b.SessionGetResponse
	err := c.C.NewRequest(ctx, "GET", path, queryParams, nil, &retVal)
	return &retVal, err
}

func (c *Client) Authenticate(
	ctx context.Context,
	body *b2b.SessionAuthenticateParams,
) (*b2b.SessionAuthenticateResponse, error) {
	path := "/b2b/sessions/authenticate"

	var jsonBody []byte
	var err error
	if body != nil {
		jsonBody, err = json.Marshal(body)
		if err != nil {
			return nil, stytcherror.NewClientLibraryError(
				"Oops, something seems to have gone wrong " +
					"marshalling the /b2b/sessions/authenticate request body")
		}
	}

	var retVal b2b.SessionAuthenticateResponse
	err = c.C.NewRequest(ctx, "POST", path, nil, jsonBody, &retVal)
	return &retVal, err
}

func (c *Client) Revoke(
	ctx context.Context,
	body *b2b.SessionRevokeParams,
) (*b2b.SessionRevokeResponse, error) {
	path := "/b2b/sessions/revoke"

	var jsonBody []byte
	var err error
	if body != nil {
		jsonBody, err = json.Marshal(body)
		if err != nil {
			return nil, stytcherror.NewClientLibraryError(
				"Oops, something seems to have gone wrong " +
					"marshalling the /b2b/sessions/revoke request body")
		}
	}

	var retVal b2b.SessionRevokeResponse
	err = c.C.NewRequest(ctx, "POST", path, nil, jsonBody, &retVal)
	return &retVal, err
}

func (c *Client) Exchange(
	ctx context.Context,
	body *b2b.SessionExchangeParams,
) (*b2b.SessionExchangeResponse, error) {
	path := "/b2b/sessions/exchange"

	var jsonBody []byte
	var err error
	if body != nil {
		jsonBody, err = json.Marshal(body)
		if err != nil {
			return nil, stytcherror.NewClientLibraryError(
				"Oops, something seems to have gone wrong " +
					"marshalling the /b2b/sessions/exchange request body")
		}
	}

	var retVal b2b.SessionExchangeResponse
	err = c.C.NewRequest(ctx, "POST", path, nil, jsonBody, &retVal)
	return &retVal, err
}

func (c *Client) GetJWKS(
	ctx context.Context, body *b2b.SessionsGetJWKSParams,
) (*b2b.SessionsGetJWKSResponse, error) {
	path := "/b2b/sessions/jwks/" + body.ProjectID

	var retVal b2b.SessionsGetJWKSResponse
	err := c.C.NewRequest(ctx, "GET", path, nil, nil, &retVal)
	return &retVal, err
}
