package discovery

import (
	"context"
	"encoding/json"

	"github.com/stytchauth/stytch-go/v7/stytch"
	"github.com/stytchauth/stytch-go/v7/stytch/b2b"
	"github.com/stytchauth/stytch-go/v7/stytch/stytcherror"
)

type Client struct {
	C *stytch.Client
}

func (c *Client) List(
	ctx context.Context,
	body *b2b.DiscoveryListOrganizationsParams,
) (*b2b.DiscoveryListOrganizationsResponse, error) {
	path := "/b2b/discovery/organizations"

	var jsonBody []byte
	var err error
	if body != nil {
		jsonBody, err = json.Marshal(body)
		if err != nil {
			return nil, stytcherror.NewClientLibraryError(
				"Oops, something seems to have gone wrong " +
					"marshalling the /b2b/discovery/organizations request body")
		}
	}

	var retVal b2b.DiscoveryListOrganizationsResponse
	err = c.C.NewRequest(ctx, "POST", path, nil, jsonBody, &retVal)
	return &retVal, err
}

func (c *Client) IntermediateSessionExchange(
	ctx context.Context,
	body *b2b.DiscoveryIntermediateSessionExchangeParams,
) (*b2b.DiscoveryIntermediateSessionExchangeResponse, error) {
	path := "/b2b/discovery/intermediate_sessions/exchange"

	var jsonBody []byte
	var err error
	if body != nil {
		jsonBody, err = json.Marshal(body)
		if err != nil {
			return nil, stytcherror.NewClientLibraryError(
				"Oops, something seems to have gone wrong " +
					"marshalling the /b2b/discovery/intermediate_sessions/exchange request body")
		}
	}

	var retVal b2b.DiscoveryIntermediateSessionExchangeResponse
	err = c.C.NewRequest(ctx, "POST", path, nil, jsonBody, &retVal)
	return &retVal, err
}

func (c *Client) OrganizationCreate(
	ctx context.Context,
	body *b2b.DiscoveryOrganizationCreateParams,
) (*b2b.DiscoveryOrganizationCreateResponse, error) {
	path := "/b2b/discovery/organizations/create"

	var jsonBody []byte
	var err error
	if body != nil {
		jsonBody, err = json.Marshal(body)
		if err != nil {
			return nil, stytcherror.NewClientLibraryError(
				"Oops, something seems to have gone wrong " +
					"marshalling the /b2b/discovery/organizations/create request body")
		}
	}

	var retVal b2b.DiscoveryOrganizationCreateResponse
	err = c.C.NewRequest(ctx, "POST", path, nil, jsonBody, &retVal)
	return &retVal, err
}
