package sessions

// !!!
// WARNING: This file is autogenerated
// Only modify code within MANUAL() sections
// or your changes may be overwritten later!
// !!!

import (
	"context"
	"encoding/json"

	"github.com/stytchauth/stytch-go/v7/stytch"
	"github.com/stytchauth/stytch-go/v7/stytch/stytcherror"
)

type Client struct {
	C *stytch.Client
}

const subURL = "sessions"

func (c *Client) Multitenantsessionsjwks(
	ctx context.Context,
	body *stytch.MultitenantsessionsjwksParams,
) (*stytch.MultitenantsessionsjwksResponse, error) {
	path := subURL + "//v1/b2b/sessions/jwks/{project_id}"

	var retVal stytch.MultitenantsessionsjwksResponse
	var jsonBody []byte
	var err error
	if body != nil {
		jsonBody, err = json.Marshal(body)
		if err != nil {
			return nil, stytcherror.NewClientLibraryError("Oops, something seems to have gone wrong marshalling the request body")
		}
	}

	err = c.C.NewRequest(ctx, "GET", path, nil, jsonBody, &retVal)

	return &retVal, err
}
