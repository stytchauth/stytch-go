package consumer

// !!!
// WARNING: This file is autogenerated
// Only modify code within MANUAL() sections
// or your changes may be overwritten later!
// !!!

import (
	"context"
	"encoding/json"

	"github.com/stytchauth/stytch-go/v8/stytch"
	"github.com/stytchauth/stytch-go/v8/stytch/consumer/magiclinks"
	"github.com/stytchauth/stytch-go/v8/stytch/stytcherror"
)

type MagicLinksClient struct {
	C     *stytch.Client
	Email *MagicLinksEmailClient
}

func NewMagicLinksClient(c *stytch.Client) *MagicLinksClient {
	return &MagicLinksClient{
		C:     c,
		Email: NewMagicLinksEmailClient(c),
	}
}

// Authenticate a User given a Magic Link. This endpoint verifies that the Magic Link token is valid,
// hasn't expired or been previously used, and any optional security settings such as IP match or user
// agent match are satisfied.
func (c *MagicLinksClient) Authenticate(
	ctx context.Context,
	body *magiclinks.AuthenticateParams,
) (*magiclinks.AuthenticateResponse, error) {
	var jsonBody []byte
	var err error
	if body != nil {
		jsonBody, err = json.Marshal(body)
		if err != nil {
			return nil, stytcherror.NewClientLibraryError("error marshaling request body")
		}
	}

	var retVal magiclinks.AuthenticateResponse
	err = c.C.NewRequest(
		ctx,
		"POST",
		"/v1/magic_links/authenticate",
		nil,
		jsonBody,
		&retVal,
	)
	return &retVal, err
}

// Create an embeddable Magic Link token for a User. Access to this endpoint is restricted. To enable it,
// please send us a note at support@stytch.com.
//
// ### Next steps
// Send the returned `token` value to the end user in a link which directs to your application. When the
// end user follows your link, collect the token, and call [Authenticate Magic
// Link](https://stytch.com/docs/api/authenticate-magic-link) to complete authentication.
func (c *MagicLinksClient) Create(
	ctx context.Context,
	body *magiclinks.CreateParams,
) (*magiclinks.CreateResponse, error) {
	var jsonBody []byte
	var err error
	if body != nil {
		jsonBody, err = json.Marshal(body)
		if err != nil {
			return nil, stytcherror.NewClientLibraryError("error marshaling request body")
		}
	}

	var retVal magiclinks.CreateResponse
	err = c.C.NewRequest(
		ctx,
		"POST",
		"/v1/magic_links",
		nil,
		jsonBody,
		&retVal,
	)
	return &retVal, err
}