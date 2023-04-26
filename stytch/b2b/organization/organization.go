package organization

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

func (c *Client) Create(
	ctx context.Context,
	body *b2b.OrganizationCreateParams,
) (*b2b.OrganizationCreateResponse, error) {
	path := "/b2b/organizations"

	var jsonBody []byte
	var err error
	if body != nil {
		jsonBody, err = json.Marshal(body)
		if err != nil {
			return nil, stytcherror.NewClientLibraryError(
				"Oops, something seems to have gone wrong " +
					"marshalling the /b2b/organizations request body")
		}
	}

	var retVal b2b.OrganizationCreateResponse
	err = c.C.NewRequest(ctx, "POST", path, nil, jsonBody, &retVal)
	return &retVal, err
}

func (c *Client) Get(
	ctx context.Context,
	organizationID string,
) (*b2b.OrganizationGetResponse, error) {
	path := "/b2b/organizations/" + organizationID

	var retVal b2b.OrganizationGetResponse
	err := c.C.NewRequest(ctx, "GET", path, nil, nil, &retVal)
	return &retVal, err
}

func (c *Client) Update(
	ctx context.Context,
	organizationID string,
	body *b2b.OrganizationUpdateParams,
) (*b2b.OrganizationUpdateResponse, error) {
	path := "/b2b/organizations" + organizationID

	var jsonBody []byte
	var err error
	if body != nil {
		jsonBody, err = json.Marshal(body)
		if err != nil {
			return nil, stytcherror.NewClientLibraryError(
				"Oops, something seems to have gone wrong " +
					"marshalling the update organization request body")
		}
	}

	var retVal b2b.OrganizationUpdateResponse
	err = c.C.NewRequest(ctx, "PUT", path, nil, jsonBody, &retVal)
	return &retVal, err
}

func (c *Client) Delete(
	ctx context.Context,
	organizationID string,
) (*b2b.OrganizationDeleteResponse, error) {
	path := "/b2b/organizations/" + organizationID

	var retVal b2b.OrganizationDeleteResponse
	err := c.C.NewRequest(ctx, "DELETE", path, nil, nil, &retVal)
	return &retVal, err
}

func (c *Client) Search(
	ctx context.Context,
	body *b2b.OrganizationSearchParams,
) (*b2b.OrganizationSearchResponse, error) {
	path := "/b2b/organizations/search"

	var jsonBody []byte
	var err error
	if body != nil {
		jsonBody, err = json.Marshal(body)
		if err != nil {
			return nil, stytcherror.NewClientLibraryError(
				"Oops, something seems to have gone wrong " +
					"marshalling the /b2b/organizations/search request body")
		}
	}

	var retVal b2b.OrganizationSearchResponse
	err = c.C.NewRequest(ctx, "POST", path, nil, jsonBody, &retVal)
	return &retVal, err
}
