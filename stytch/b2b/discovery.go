package b2b

// !!!
// WARNING: This file is autogenerated
// Only modify code within MANUAL() sections
// or your changes may be overwritten later!
// !!!

import (
	"github.com/stytchauth/stytch-go/v9/stytch"
)

type DiscoveryClient struct {
	C                    *stytch.Client
	IntermediateSessions *DiscoveryIntermediateSessionsClient
	Organizations        *DiscoveryOrganizationsClient
}

func NewDiscoveryClient(c *stytch.Client) *DiscoveryClient {
	return &DiscoveryClient{
		C:                    c,
		IntermediateSessions: NewDiscoveryIntermediateSessionsClient(c),
		Organizations:        NewDiscoveryOrganizationsClient(c),
	}
}
