package b2b

// !!!
// WARNING: This file is autogenerated
// Only modify code within MANUAL() sections
// or your changes may be overwritten later!
// !!!

import (
	"context"
	"encoding/json"

	"github.com/stytchauth/stytch-go/v9/stytch"
	"github.com/stytchauth/stytch-go/v9/stytch/b2b/magiclinks/email/discovery"
	"github.com/stytchauth/stytch-go/v9/stytch/stytcherror"
)

type MagicLinksEmailDiscoveryClient struct {
	C *stytch.Client
}

func NewMagicLinksEmailDiscoveryClient(c *stytch.Client) *MagicLinksEmailDiscoveryClient {
	return &MagicLinksEmailDiscoveryClient{
		C: c,
	}
}

// Send a discovery magic link to an email address.
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

	var retVal discovery.SendResponse
	err = c.C.NewRequest(
		ctx,
		"POST",
		"/v1/b2b/magic_links/email/discovery/send",
		nil,
		jsonBody,
		&retVal,
	)
	return &retVal, err
}
