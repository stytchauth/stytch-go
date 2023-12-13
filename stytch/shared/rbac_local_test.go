package shared_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stytchauth/stytch-go/v11/stytch/b2b/rbac"
	"github.com/stytchauth/stytch-go/v11/stytch/b2b/sessions"
	"github.com/stytchauth/stytch-go/v11/stytch/shared"
	"github.com/stytchauth/stytch-go/v11/stytch/stytcherror"
)

const orgID = "organization-1234"

func TestPerformAuthorizationCheck(t *testing.T) {
	policy := &rbac.Policy{
		Roles: []rbac.PolicyRole{
			{
				RoleID:      "stytch_member",
				Description: "member",
				Permissions: []rbac.PolicyRolePermission{
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
				Permissions: []rbac.PolicyRolePermission{
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
				Permissions: []rbac.PolicyRolePermission{
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
		Resources: []rbac.PolicyResource{
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
		err := shared.PerformAuthorizationCheck(
			policy,
			[]string{"stytch_member"},
			orgID,
			&sessions.AuthorizationCheck{
				OrganizationID: diffOrgID,
				ResourceID:     "document",
				Action:         "read",
			},
		)
		assert.ErrorIs(t, err, stytcherror.NewSessionAuthorizationTenancyError(orgID, diffOrgID))
	})
	t.Run("action exists but resource does not", func(t *testing.T) {
		err := shared.PerformAuthorizationCheck(
			policy,
			[]string{"stytch_member"},
			orgID,
			&sessions.AuthorizationCheck{
				OrganizationID: orgID,
				ResourceID:     "resource_that_doesnt_exist",
				Action:         "read",
			},
		)
		assert.ErrorIs(t, err, stytcherror.NewPermissionError())
	})
	t.Run("resource exists but action does not", func(t *testing.T) {
		err := shared.PerformAuthorizationCheck(
			policy,
			[]string{"stytch_member"},
			orgID,
			&sessions.AuthorizationCheck{
				OrganizationID: orgID,
				ResourceID:     "document",
				Action:         "action_that_doesnt_exist",
			},
		)
		assert.ErrorIs(t, err, stytcherror.NewPermissionError())
	})
	t.Run("member has this action but on a different resource", func(t *testing.T) {
		err := shared.PerformAuthorizationCheck(
			policy,
			[]string{"stytch_member"},
			orgID,
			&sessions.AuthorizationCheck{
				OrganizationID: orgID,
				ResourceID:     "program",
				Action:         "write",
			},
		)
		assert.ErrorIs(t, err, stytcherror.NewPermissionError())
	})
	t.Run("another authorization check for a member with more elevated privileges", func(t *testing.T) {
		err := shared.PerformAuthorizationCheck(
			policy,
			[]string{"stytch_editor"},
			orgID,
			&sessions.AuthorizationCheck{
				OrganizationID: orgID,
				ResourceID:     "program",
				Action:         "edit",
			},
		)
		assert.ErrorIs(t, err, stytcherror.NewPermissionError())
	})
	t.Run("no error when the member is authorized", func(t *testing.T) {
		err := shared.PerformAuthorizationCheck(
			policy,
			[]string{"stytch_admin"},
			orgID,
			&sessions.AuthorizationCheck{
				OrganizationID: orgID,
				ResourceID:     "document",
				Action:         "delete",
			},
		)
		assert.NoError(t, err)
	})
}
