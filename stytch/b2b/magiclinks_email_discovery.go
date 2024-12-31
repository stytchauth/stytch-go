package b2b

// !!!
// WARNING: This file is autogenerated
// Only modify code within MANUAL() sections
// or your changes may be overwritten later!
// !!!

import (
	"context"
	"encoding/json"

	"github.com/stytchauth/stytch-go/v15/stytch"
	"github.com/stytchauth/stytch-go/v15/stytch/b2b/magiclinks/email/discovery"
	"github.com/stytchauth/stytch-go/v15/stytch/stytcherror"
)

type MagicLinksEmailDiscoveryClient struct {
	C stytch.Client
}

func NewMagicLinksEmailDiscoveryClient(c stytch.Client) *MagicLinksEmailDiscoveryClient {
	return &MagicLinksEmailDiscoveryClient{
		C: c,
	}
}

// Send a discovery magic link to an email address. The magic link is valid for 60 minutes.
func (c *MagicLinksEmailDiscoveryClient) Send(
	ctx context.Context,
	body *discovery.SendParams,
) (*discovery.SendResponse, error) {
	var jsonBody []byte
	var err error
	if body != nil {
		jsonBody, err = json.Marshal(body)
		if err != nil {
			return nil, stytcherror.NewClientLibraryError("error marshaling request body")
		}
	}

	headers := make(map[string][]string)

	var retVal discovery.SendResponse
	err = c.C.NewRequest(
		ctx,
		stytch.RequestParams{
			Method:      "POST",
			Path:        "/v1/b2b/magic_links/email/discovery/send",
			QueryParams: nil,
			Body:        jsonBody,
			V:           &retVal,
			Headers:     headers,
		},
	)
	return &retVal, err
}
