package shared_test

import (
	consumerrbac "github.com/stytchauth/stytch-go/v16/stytch/consumer/rbac"
	consumersessions "github.com/stytchauth/stytch-go/v16/stytch/consumer/sessions"
	"testing"

	"github.com/stretchr/testify/assert"
	b2brbac "github.com/stytchauth/stytch-go/v16/stytch/b2b/rbac"
	b2bsessions "github.com/stytchauth/stytch-go/v16/stytch/b2b/sessions"
	"github.com/stytchauth/stytch-go/v16/stytch/shared"
	"github.com/stytchauth/stytch-go/v16/stytch/stytcherror"
)

func Test_PerformB2BAuthorizationCheck(t *testing.T) {
	const orgID = "organization-1234"
	policy := &b2brbac.Policy{
		Roles: []b2brbac.PolicyRole{
			{
				RoleID:      "stytch_member",
				Description: "member",
				Permissions: []b2brbac.PolicyRolePermission{
					{
						ResourceID: "document",
						Actions:    []string{"read", "write"},
					},
					{
						ResourceID: "program",
						Actions:    []string{"read"},
					},
				},
			},
			{
				RoleID:      "stytch_editor",
				Description: "member",
				Permissions: []b2brbac.PolicyRolePermission{
					{
						ResourceID: "document",
						Actions:    []string{"read", "write"},
					},
					{
						ResourceID: "program",
						Actions:    []string{"read", "execute"},
					},
				},
			},
			{
				RoleID:      "stytch_admin",
				Description: "admin",
				Permissions: []b2brbac.PolicyRolePermission{
					{
						ResourceID: "document",
						Actions:    []string{"read", "write", "delete"},
					},
					{
						ResourceID: "program",
						Actions:    []string{"read", "edit", "execute"},
					},
				},
			},
		},
		Resources: []b2brbac.PolicyResource{
			{
				ResourceID:  "document",
				Description: "All documents",
				Actions:     []string{"read", "write", "delete"},
			},
			{
				ResourceID:  "program",
				Description: "An executable program",
				Actions:     []string{"read", "write", "execute"},
			},
		},
	}

	t.Run("tenancy mismatch", func(t *testing.T) {
		diffOrgID := "different-organization-id"
		err := shared.PerformB2BAuthorizationCheck(
			policy,
			[]string{"stytch_member"},
			orgID,
			&b2bsessions.AuthorizationCheck{
				OrganizationID: diffOrgID,
				ResourceID:     "document",
				Action:         "read",
			},
		)
		assert.ErrorContains(t, err, stytcherror.NewSessionAuthorizationTenancyError(orgID, diffOrgID).Error())
	})
	t.Run("action exists but resource does not", func(t *testing.T) {
		err := shared.PerformB2BAuthorizationCheck(
			policy,
			[]string{"stytch_member"},
			orgID,
			&b2bsessions.AuthorizationCheck{
				OrganizationID: orgID,
				ResourceID:     "resource_that_doesnt_exist",
				Action:         "read",
			},
		)
		assert.ErrorContains(t, err, stytcherror.NewPermissionError().Error())
	})
	t.Run("resource exists but action does not", func(t *testing.T) {
		err := shared.PerformB2BAuthorizationCheck(
			policy,
			[]string{"stytch_member"},
			orgID,
			&b2bsessions.AuthorizationCheck{
				OrganizationID: orgID,
				ResourceID:     "document",
				Action:         "action_that_doesnt_exist",
			},
		)
		assert.ErrorContains(t, err, stytcherror.NewPermissionError().Error())
	})
	t.Run("member has this action but on a different resource", func(t *testing.T) {
		err := shared.PerformB2BAuthorizationCheck(
			policy,
			[]string{"stytch_member"},
			orgID,
			&b2bsessions.AuthorizationCheck{
				OrganizationID: orgID,
				ResourceID:     "program",
				Action:         "write",
			},
		)
		assert.ErrorContains(t, err, stytcherror.NewPermissionError().Error())
	})
	t.Run("another authorization check for a member with more elevated privileges", func(t *testing.T) {
		err := shared.PerformB2BAuthorizationCheck(
			policy,
			[]string{"stytch_editor"},
			orgID,
			&b2bsessions.AuthorizationCheck{
				OrganizationID: orgID,
				ResourceID:     "program",
				Action:         "edit",
			},
		)
		assert.ErrorContains(t, err, stytcherror.NewPermissionError().Error())
	})
	t.Run("no error when the member is authorized", func(t *testing.T) {
		err := shared.PerformB2BAuthorizationCheck(
			policy,
			[]string{"stytch_admin"},
			orgID,
			&b2bsessions.AuthorizationCheck{
				OrganizationID: orgID,
				ResourceID:     "document",
				Action:         "delete",
			},
		)
		assert.NoError(t, err)
	})
}

func Test_PerformConsumerAuthorizationCheck(t *testing.T) {
	policy := &consumerrbac.Policy{
		Roles: []consumerrbac.PolicyRole{
			{
				RoleID:      "stytch_member",
				Description: "member",
				Permissions: []consumerrbac.PolicyRolePermission{
					{
						ResourceID: "document",
						Actions:    []string{"read", "write"},
					},
					{
						ResourceID: "program",
						Actions:    []string{"read"},
					},
				},
			},
			{
				RoleID:      "stytch_editor",
				Description: "member",
				Permissions: []consumerrbac.PolicyRolePermission{
					{
						ResourceID: "document",
						Actions:    []string{"read", "write"},
					},
					{
						ResourceID: "program",
						Actions:    []string{"read", "execute"},
					},
				},
			},
			{
				RoleID:      "stytch_admin",
				Description: "admin",
				Permissions: []consumerrbac.PolicyRolePermission{
					{
						ResourceID: "document",
						Actions:    []string{"read", "write", "delete"},
					},
					{
						ResourceID: "program",
						Actions:    []string{"read", "edit", "execute"},
					},
				},
			},
		},
		Resources: []consumerrbac.PolicyResource{
			{
				ResourceID:  "document",
				Description: "All documents",
				Actions:     []string{"read", "write", "delete"},
			},
			{
				ResourceID:  "program",
				Description: "An executable program",
				Actions:     []string{"read", "write", "execute"},
			},
		},
	}

	t.Run("action exists but resource does not", func(t *testing.T) {
		err := shared.PerformConsumerAuthorizationCheck(
			policy,
			[]string{"stytch_member"},
			&consumersessions.AuthorizationCheck{
				ResourceID: "resource_that_doesnt_exist",
				Action:     "read",
			},
		)
		assert.ErrorContains(t, err, stytcherror.NewPermissionError().Error())
	})
	t.Run("resource exists but action does not", func(t *testing.T) {
		err := shared.PerformConsumerAuthorizationCheck(
			policy,
			[]string{"stytch_member"},
			&consumersessions.AuthorizationCheck{
				ResourceID: "document",
				Action:     "action_that_doesnt_exist",
			},
		)
		assert.ErrorContains(t, err, stytcherror.NewPermissionError().Error())
	})
	t.Run("member has this action but on a different resource", func(t *testing.T) {
		err := shared.PerformConsumerAuthorizationCheck(
			policy,
			[]string{"stytch_member"},
			&consumersessions.AuthorizationCheck{
				ResourceID: "program",
				Action:     "write",
			},
		)
		assert.ErrorContains(t, err, stytcherror.NewPermissionError().Error())
	})
	t.Run("another authorization check for a member with more elevated privileges", func(t *testing.T) {
		err := shared.PerformConsumerAuthorizationCheck(
			policy,
			[]string{"stytch_editor"},
			&consumersessions.AuthorizationCheck{
				ResourceID: "program",
				Action:     "edit",
			},
		)
		assert.ErrorContains(t, err, stytcherror.NewPermissionError().Error())
	})
	t.Run("no error when the member is authorized", func(t *testing.T) {
		err := shared.PerformConsumerAuthorizationCheck(
			policy,
			[]string{"stytch_admin"},
			&consumersessions.AuthorizationCheck{
				ResourceID: "document",
				Action:     "delete",
			},
		)
		assert.NoError(t, err)
	})
}
