package b2b

// !!!
// WARNING: This file is autogenerated
// Only modify code within MANUAL() sections
// or your changes may be overwritten later!
// !!!

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/stytchauth/stytch-go/v16/stytch"
	"github.com/stytchauth/stytch-go/v16/stytch/b2b/sso/saml"
	"github.com/stytchauth/stytch-go/v16/stytch/stytcherror"
)

type SSOSAMLClient struct {
	C stytch.Client
}

func NewSSOSAMLClient(c stytch.Client) *SSOSAMLClient {
	return &SSOSAMLClient{
		C: c,
	}
}

// CreateConnection: Create a new SAML Connection.
func (c *SSOSAMLClient) CreateConnection(
	ctx context.Context,
	body *saml.CreateConnectionParams,
	methodOptions ...*saml.CreateConnectionRequestOptions,
) (*saml.CreateConnectionResponse, error) {
	var jsonBody []byte
	var err error
	if body != nil {
		jsonBody, err = json.Marshal(body)
		if err != nil {
			return nil, stytcherror.NewClientLibraryError("error marshaling request body")
		}
	}

	headers := make(map[string][]string)
	for _, methodOption := range methodOptions {
		headers = methodOption.AddHeaders(headers)
	}

	var retVal saml.CreateConnectionResponse
	err = c.C.NewRequest(
		ctx,
		stytch.RequestParams{
			Method:      "POST",
			Path:        fmt.Sprintf("/v1/b2b/sso/saml/%s", body.OrganizationID),
			QueryParams: nil,
			Body:        jsonBody,
			V:           &retVal,
			Headers:     headers,
		},
	)
	return &retVal, err
}

// UpdateConnection: Updates an existing SAML connection.
//
// Note that a newly created connection will not become active until all of the following are provided:
// * `idp_sso_url`
// * `attribute_mapping`
// * `idp_entity_id`
// * `x509_certificate`
func (c *SSOSAMLClient) UpdateConnection(
	ctx context.Context,
	body *saml.UpdateConnectionParams,
	methodOptions ...*saml.UpdateConnectionRequestOptions,
) (*saml.UpdateConnectionResponse, error) {
	var jsonBody []byte
	var err error
	if body != nil {
		jsonBody, err = json.Marshal(body)
		if err != nil {
			return nil, stytcherror.NewClientLibraryError("error marshaling request body")
		}
	}

	headers := make(map[string][]string)
	for _, methodOption := range methodOptions {
		headers = methodOption.AddHeaders(headers)
	}

	var retVal saml.UpdateConnectionResponse
	err = c.C.NewRequest(
		ctx,
		stytch.RequestParams{
			Method:      "PUT",
			Path:        fmt.Sprintf("/v1/b2b/sso/saml/%s/connections/%s", body.OrganizationID, body.ConnectionID),
			QueryParams: nil,
			Body:        jsonBody,
			V:           &retVal,
			Headers:     headers,
		},
	)
	return &retVal, err
}

// UpdateByURL: Used to update an existing SAML connection using an IDP metadata URL.
//
// A newly created connection will not become active until all the following are provided:
// * `idp_sso_url`
// * `idp_entity_id`
// * `x509_certificate`
// * `attribute_mapping` (must be supplied using [Update SAML Connection](update-saml-connection))
func (c *SSOSAMLClient) UpdateByURL(
	ctx context.Context,
	body *saml.UpdateByURLParams,
	methodOptions ...*saml.UpdateByURLRequestOptions,
) (*saml.UpdateByURLResponse, error) {
	var jsonBody []byte
	var err error
	if body != nil {
		jsonBody, err = json.Marshal(body)
		if err != nil {
			return nil, stytcherror.NewClientLibraryError("error marshaling request body")
		}
	}

	headers := make(map[string][]string)
	for _, methodOption := range methodOptions {
		headers = methodOption.AddHeaders(headers)
	}

	var retVal saml.UpdateByURLResponse
	err = c.C.NewRequest(
		ctx,
		stytch.RequestParams{
			Method:      "PUT",
			Path:        fmt.Sprintf("/v1/b2b/sso/saml/%s/connections/%s/url", body.OrganizationID, body.ConnectionID),
			QueryParams: nil,
			Body:        jsonBody,
			V:           &retVal,
			Headers:     headers,
		},
	)
	return &retVal, err
}

// DeleteVerificationCertificate: Delete a SAML verification certificate.
//
// You may need to do this when rotating certificates from your IdP, since Stytch allows a maximum of 5
// certificates per connection. There must always be at least one certificate per active connection.
func (c *SSOSAMLClient) DeleteVerificationCertificate(
	ctx context.Context,
	body *saml.DeleteVerificationCertificateParams,
	methodOptions ...*saml.DeleteVerificationCertificateRequestOptions,
) (*saml.DeleteVerificationCertificateResponse, error) {
	headers := make(map[string][]string)
	for _, methodOption := range methodOptions {
		headers = methodOption.AddHeaders(headers)
	}

	var retVal saml.DeleteVerificationCertificateResponse
	err := c.C.NewRequest(
		ctx,
		stytch.RequestParams{
			Method:      "DELETE",
			Path:        fmt.Sprintf("/v1/b2b/sso/saml/%s/connections/%s/verification_certificates/%s", body.OrganizationID, body.ConnectionID, body.CertificateID),
			QueryParams: nil,
			Body:        nil,
			V:           &retVal,
			Headers:     headers,
		},
	)
	return &retVal, err
}
