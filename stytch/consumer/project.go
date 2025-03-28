package consumer

// !!!
// WARNING: This file is autogenerated
// Only modify code within MANUAL() sections
// or your changes may be overwritten later!
// !!!

import (
	"context"

	"github.com/stytchauth/stytch-go/v16/stytch"
	"github.com/stytchauth/stytch-go/v16/stytch/consumer/project"
)

type ProjectClient struct {
	C stytch.Client
}

func NewProjectClient(c stytch.Client) *ProjectClient {
	return &ProjectClient{
		C: c,
	}
}

func (c *ProjectClient) Metrics(
	ctx context.Context,
	body *project.MetricsParams,
) (*project.MetricsResponse, error) {
	headers := make(map[string][]string)

	var retVal project.MetricsResponse
	err := c.C.NewRequest(
		ctx,
		stytch.RequestParams{
			Method:      "GET",
			Path:        "/v1/projects/metrics",
			QueryParams: nil,
			Body:        nil,
			V:           &retVal,
			Headers:     headers,
		},
	)
	return &retVal, err
}
