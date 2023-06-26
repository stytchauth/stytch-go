package b2b

// !!!
// WARNING: This file is autogenerated
// Only modify code within MANUAL() sections
// or your changes may be overwritten later!
// !!!

import (
	"context"
	"encoding/json"

	"github.com/stytchauth/stytch-go/v8/stytch"
	"github.com/stytchauth/stytch-go/v8/stytch/b2b/discovery/intermediatesessions"
	"github.com/stytchauth/stytch-go/v8/stytch/stytcherror"
)

type DiscoveryIntermediateSessionsClient struct {
	C *stytch.Client
}

func NewDiscoveryIntermediateSessionsClient(c *stytch.Client) *DiscoveryIntermediateSessionsClient {
	return &DiscoveryIntermediateSessionsClient{
		C: c,
	}
}

// Exchange an Intermediate Session for a fully realized
// [Member Session](https://stytch.com/docs/b2b/api/session-object) in a desired
// [Organization](https://stytch.com/docs/b2b/api/organization-object).
// This operation consumes the Intermediate Session.
//
// This endpoint can be used to accept invites and create new members via domain matching.
func (c *DiscoveryIntermediateSessionsClient) Exchange(
	ctx context.Context,
	body *intermediatesessions.ExchangeParams,
) (*intermediatesessions.ExchangeResponse, error) {
	var jsonBody []byte
	var err error
	if body != nil {
		jsonBody, err = json.Marshal(body)
		if err != nil {
			return nil, stytcherror.NewClientLibraryError("error marshaling request body")
		}
	}

	var retVal intermediatesessions.ExchangeResponse
	err = c.C.NewRequest(
		ctx,
		"POST",
		"/v1/b2b/discovery/intermediate_sessions/exchange",
		nil,
		jsonBody,
		&retVal,
	)
	return &retVal, err
}
