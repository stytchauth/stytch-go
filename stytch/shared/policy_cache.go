package shared

import (
	"context"
	"time"

	"github.com/stytchauth/stytch-go/v11/stytch/b2b"
	"github.com/stytchauth/stytch-go/v11/stytch/b2b/rbac"
)

const refreshCadence = 300 * time.Second

type PolicyCache struct {
	rbacClient    *b2b.B2BRBACClient
	policy        *rbac.Policy
	lastUpdatedAt time.Time
}

func NewPolicyCache(rbacClient *b2b.B2BRBACClient) *PolicyCache {
	return &PolicyCache{rbacClient: rbacClient}
}

func (pc *PolicyCache) shouldRefreshPolicy() bool {
	return time.Since(pc.lastUpdatedAt) > refreshCadence
}

func (pc *PolicyCache) Get(ctx context.Context) (*rbac.Policy, error) {
	if pc.policy == nil || pc.shouldRefreshPolicy() {
		policyResp, err := pc.rbacClient.Policy(ctx, &rbac.PolicyParams{})
		if err != nil {
			return nil, err
		}

		pc.policy = policyResp.Policy
		pc.lastUpdatedAt = time.Now()
	}
	return pc.policy, nil
}
