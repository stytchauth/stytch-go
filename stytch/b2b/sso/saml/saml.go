package saml

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
	body *b2b.SAMLCreateConnectionParams,
) (*b2b.SAMLCreateConnectionResponse, error) {
	path := "/b2b/sso/saml/" + organizationID

	var jsonBody []byte
	var err error
	if body != nil {
		jsonBody, err = json.Marshal(body)
		if err != nil {
			return nil, stytcherror.NewClientLibraryError(
				"Oops, something seems to have gone wrong " +
					"marshalling the SAML Create request body")
		}
	}

	var retVal b2b.SAMLCreateConnectionResponse
	err = c.C.NewRequest(ctx, "POST", path, nil, jsonBody, &retVal)
	return &retVal, err
}

func (c *Client) Update(
	ctx context.Context,
	organizationID string,
	connectionID string,
	body *b2b.SAMLUpdateConnectionParams,
) (*b2b.SAMLUpdateConnectionResponse, error) {
	path := "/b2b/sso/saml/" + organizationID + "/connections/" + connectionID

	var jsonBody []byte
	var err error
	if body != nil {
		jsonBody, err = json.Marshal(body)
		if err != nil {
			return nil, stytcherror.NewClientLibraryError(
				"Oops, something seems to have gone wrong " +
					"marshalling the SAML Update request body")
		}
	}

	var retVal b2b.SAMLUpdateConnectionResponse
	err = c.C.NewRequest(ctx, "PUT", path, nil, jsonBody, &retVal)
	return &retVal, err
}

func (c *Client) Delete(
	ctx context.Context,
	organizationID string,
	connectionID string,
	certificateID string,
) (*b2b.SAMLDeleteVerificationCertificateResponse, error) {
	path := "/b2b/sso/saml/" + organizationID + "/connections/" + connectionID + "/verification_certificates/" + certificateID

	var retVal b2b.SAMLDeleteVerificationCertificateResponse
	err := c.C.NewRequest(ctx, "DELETE", path, nil, nil, &retVal)
	return &retVal, err
}
