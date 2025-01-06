package consumer

// !!!
// WARNING: This file is autogenerated
// Only modify code within MANUAL() sections
// or your changes may be overwritten later!
// !!!

import (
	"context"
	"encoding/json"

	"github.com/stytchauth/stytch-go/v16/stytch"
	"github.com/stytchauth/stytch-go/v16/stytch/consumer/fraud/rules"
	"github.com/stytchauth/stytch-go/v16/stytch/stytcherror"
)

type FraudRulesClient struct {
	C stytch.Client
}

func NewFraudRulesClient(c stytch.Client) *FraudRulesClient {
	return &FraudRulesClient{
		C: c,
	}
}

// Set a rule for a particular `visitor_id`, `browser_id`, `visitor_fingerprint`, `browser_fingerprint`,
// `hardware_fingerprint`, or `network_fingerprint`. This is helpful in cases where you want to allow or
// block a specific user or fingerprint. You should be careful when setting rules for
// `browser_fingerprint`, `hardware_fingerprint`, or `network_fingerprint` as they can be shared across
// multiple users, and you could affect more users than intended.
//
// Rules are applied in the order specified above. For example, if an end user has an `ALLOW` rule set for
// their `visitor_id` but a `BLOCK` rule set for their `hardware_fingerprint`, they will receive an `ALLOW`
// verdict because the `visitor_id` rule takes precedence.
func (c *FraudRulesClient) Set(
	ctx context.Context,
	body *rules.SetParams,
) (*rules.SetResponse, error) {
	var jsonBody []byte
	var err error
	if body != nil {
		jsonBody, err = json.Marshal(body)
		if err != nil {
			return nil, stytcherror.NewClientLibraryError("error marshaling request body")
		}
	}

	headers := make(map[string][]string)

	var retVal rules.SetResponse
	err = c.C.NewRequest(
		ctx,
		stytch.RequestParams{
			Method:      "POST",
			Path:        "/v1/rules/set",
			QueryParams: nil,
			Body:        jsonBody,
			V:           &retVal,
			Headers:     headers,
			BaseURLType: "FRAUD",
		},
	)
	return &retVal, err
}
