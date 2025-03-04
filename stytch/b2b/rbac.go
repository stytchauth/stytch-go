package b2b

// !!!
// WARNING: This file is autogenerated
// Only modify code within MANUAL() sections
// or your changes may be overwritten later!
// !!!

import (
	"context"
	"time"

	"github.com/stytchauth/stytch-go/v16/stytch"
	"github.com/stytchauth/stytch-go/v16/stytch/b2b/rbac"
)

type RBACClient struct {
	C stytch.Client
}

func NewRBACClient(c stytch.Client) *RBACClient {
	return &RBACClient{
		C: c,
	}
}

// Policy: Get the active RBAC Policy for your current Stytch Project. An RBAC Policy is the canonical
// document that stores all defined Resources and Roles within your RBAC permissioning model.
//
// When using the backend SDKs, the RBAC Policy will be cached to allow for local evaluations, eliminating
// the need for an extra request to Stytch. The policy will be refreshed if an authorization check is
// requested and the RBAC policy was last updated more than 5 minutes ago.
//
// Resources and Roles can be created and managed within the
// [Dashboard](https://stytch.com/docs/dashboard/rbac). Additionally,
// [Role assignment](https://stytch.com/docs/b2b/guides/rbac/role-assignment) can be programmatically
// managed through certain Stytch API endpoints.
//
// Check out the [RBAC overview](https://stytch.com/docs/b2b/guides/rbac/overview) to learn more about
// Stytch's RBAC permissioning model.
func (c *RBACClient) Policy(
	ctx context.Context,
	body *rbac.PolicyParams,
) (*rbac.PolicyResponse, error) {
	headers := make(map[string][]string)

	var retVal rbac.PolicyResponse
	err := c.C.NewRequest(
		ctx,
		stytch.RequestParams{
			Method:      "GET",
			Path:        "/v1/b2b/rbac/policy",
			QueryParams: nil,
			Body:        nil,
			V:           &retVal,
			Headers:     headers,
		},
	)
	return &retVal, err
}

// MANUAL(PolicyCache)(TYPES)

type PolicyCache struct {
	rbacClient    *RBACClient
	policy        *rbac.Policy
	lastUpdatedAt time.Time
}

const refreshCadence = 5 * time.Minute

func NewPolicyCache(rbacClient *RBACClient) *PolicyCache {
	return &PolicyCache{rbacClient: rbacClient}
}

func (pc *PolicyCache) Get(ctx context.Context) (*rbac.Policy, error) {
	if pc.policy == nil || time.Since(pc.lastUpdatedAt) > refreshCadence {
		policyResp, err := pc.rbacClient.Policy(ctx, &rbac.PolicyParams{})
		if err != nil {
			return nil, err
		}

		pc.policy = policyResp.Policy
		pc.lastUpdatedAt = time.Now()
	}
	return pc.policy, nil
}

// ENDMANUAL(PolicyCache)
