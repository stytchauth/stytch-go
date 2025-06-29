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
	"github.com/stytchauth/stytch-go/v16/stytch/consumer/fraud/verdictreasons"
	"github.com/stytchauth/stytch-go/v16/stytch/stytcherror"
)

type FraudVerdictReasonsClient struct {
	C stytch.Client
}

func NewFraudVerdictReasonsClient(c stytch.Client) *FraudVerdictReasonsClient {
	return &FraudVerdictReasonsClient{
		C: c,
	}
}

// Override: Use this endpoint to override the action returned for a specific verdict reason during a
// fingerprint lookup. For example, Stytch Device Fingerprinting returns a `CHALLENGE` verdict action by
// default for the verdict reason `VIRTUAL_MACHINE`. You can use this endpoint to override that reason to
// return an `ALLOW` verdict instead if you expect many legitimate users to be using a browser that runs in
// a virtual machine.
func (c *FraudVerdictReasonsClient) Override(
	ctx context.Context,
	body *verdictreasons.OverrideParams,
) (*verdictreasons.OverrideResponse, error) {
	var jsonBody []byte
	var err error
	if body != nil {
		jsonBody, err = json.Marshal(body)
		if err != nil {
			return nil, stytcherror.NewClientLibraryError("error marshaling request body")
		}
	}

	headers := make(map[string][]string)

	var retVal verdictreasons.OverrideResponse
	err = c.C.NewRequest(
		ctx,
		stytch.RequestParams{
			Method:      "POST",
			Path:        "/v1/verdict_reasons/override",
			QueryParams: nil,
			Body:        jsonBody,
			V:           &retVal,
			Headers:     headers,
			BaseURLType: "FRAUD",
		},
	)
	return &retVal, err
}

// List: Get the list of verdict reasons returned by the Stytch Device Fingerprinting product along with
// their default actions and any overrides you may have defined. This is not an exhaustive list of verdict
// reasons, but it contains all verdict reasons that you may set an override on.
//
// For a full list of possible verdict reasons, see
// [Warning Flags (Verdict Reasons)](https://stytch.com/docs/docs/fraud/guides/device-fingerprinting/reference/warning-flags-verdict-reasons).
func (c *FraudVerdictReasonsClient) List(
	ctx context.Context,
	body *verdictreasons.ListParams,
) (*verdictreasons.ListResponse, error) {
	var jsonBody []byte
	var err error
	if body != nil {
		jsonBody, err = json.Marshal(body)
		if err != nil {
			return nil, stytcherror.NewClientLibraryError("error marshaling request body")
		}
	}

	headers := make(map[string][]string)

	var retVal verdictreasons.ListResponse
	err = c.C.NewRequest(
		ctx,
		stytch.RequestParams{
			Method:      "POST",
			Path:        "/v1/verdict_reasons/list",
			QueryParams: nil,
			Body:        jsonBody,
			V:           &retVal,
			Headers:     headers,
			BaseURLType: "FRAUD",
		},
	)
	return &retVal, err
}
