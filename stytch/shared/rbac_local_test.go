package shared_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	b2brbac "github.com/stytchauth/stytch-go/v17/stytch/b2b/rbac"
	b2bsessions "github.com/stytchauth/stytch-go/v17/stytch/b2b/sessions"
	consumerrbac "github.com/stytchauth/stytch-go/v17/stytch/consumer/rbac"
	consumersessions "github.com/stytchauth/stytch-go/v17/stytch/consumer/sessions"
	"github.com/stytchauth/stytch-go/v17/stytch/shared"
	"github.com/stytchauth/stytch-go/v17/stytch/stytcherror"
)

func Test_PerformB2BAuthorizationCheck(t *testing.T) {
	const orgID = "organization-1234"
	sampleProjectPolicy := &b2brbac.Policy{
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
	sampleOrgPolicy := &b2brbac.OrgPolicy{
		Roles: []b2brbac.PolicyRole{
			{
				RoleID: "editor",
				Permissions: []b2brbac.PolicyRolePermission{
					{
						ResourceID: "document",
						Actions:    []string{"*"},
					},
				},
			},
			{
				RoleID: "guest",
				Permissions: []b2brbac.PolicyRolePermission{
					{
						ResourceID: "program",
						Actions:    []string{"read"},
					},
				},
			},
			{
				RoleID: "sysadmin",
				Permissions: []b2brbac.PolicyRolePermission{
					{
						ResourceID: "program",
						Actions:    []string{"read", "execute"},
					},
				},
			},
		},
	}

	for _, testCase := range []struct {
		description string
		in          shared.PerformB2BAuthorizationCheckIn
		err         error
	}{
		{
			description: "error: tenancy mismatch",
			in: shared.PerformB2BAuthorizationCheckIn{
				ProjectPolicy: sampleProjectPolicy,
				OrgPolicy:     sampleOrgPolicy,
				SubjectRoles:  []string{"stytch_member"},
				SubjectOrgID:  orgID,
				AuthorizationCheck: &b2bsessions.AuthorizationCheck{
					OrganizationID: "different-organization-id",
					ResourceID:     "document",
					Action:         "read",
				},
			},
			err: stytcherror.NewSessionAuthorizationTenancyError(orgID, "different-organization-id"),
		},
		{
			description: "error: action exists in project policy but resource does not",
			in: shared.PerformB2BAuthorizationCheckIn{
				ProjectPolicy: sampleProjectPolicy,
				OrgPolicy:     sampleOrgPolicy,
				SubjectRoles:  []string{"stytch_member"},
				SubjectOrgID:  orgID,
				AuthorizationCheck: &b2bsessions.AuthorizationCheck{
					OrganizationID: orgID,
					ResourceID:     "resource_that_doesnt_exist",
					Action:         "read",
				},
			},
			err: stytcherror.NewPermissionError(),
		},
		{
			description: "error: resource exists in project policy but action does not",
			in: shared.PerformB2BAuthorizationCheckIn{
				ProjectPolicy: sampleProjectPolicy,
				OrgPolicy:     sampleOrgPolicy,
				SubjectRoles:  []string{"stytch_member"},
				SubjectOrgID:  orgID,
				AuthorizationCheck: &b2bsessions.AuthorizationCheck{
					OrganizationID: orgID,
					ResourceID:     "document",
					Action:         "action_that_doesnt_exist",
				},
			},
			err: stytcherror.NewPermissionError(),
		},
		{
			description: "error: member has this action but on a different resource",
			in: shared.PerformB2BAuthorizationCheckIn{
				ProjectPolicy: sampleProjectPolicy,
				OrgPolicy:     sampleOrgPolicy,
				SubjectRoles:  []string{"stytch_member"},
				SubjectOrgID:  orgID,
				AuthorizationCheck: &b2bsessions.AuthorizationCheck{
					OrganizationID: orgID,
					ResourceID:     "program",
					Action:         "write",
				},
			},
			err: stytcherror.NewPermissionError(),
		},
		{
			description: "error: another authorization check for a member with more elevated privileges",
			in: shared.PerformB2BAuthorizationCheckIn{
				ProjectPolicy: sampleProjectPolicy,
				OrgPolicy:     nil,
				SubjectRoles:  []string{"stytch_editor"},
				SubjectOrgID:  orgID,
				AuthorizationCheck: &b2bsessions.AuthorizationCheck{
					OrganizationID: orgID,
					ResourceID:     "program",
					Action:         "edit",
				},
			},
			err: stytcherror.NewPermissionError(),
		},
		{
			description: "success when the member is authorized",
			in: shared.PerformB2BAuthorizationCheckIn{
				ProjectPolicy: sampleProjectPolicy,
				OrgPolicy:     sampleOrgPolicy,
				SubjectRoles:  []string{"stytch_admin"},
				SubjectOrgID:  orgID,
				AuthorizationCheck: &b2bsessions.AuthorizationCheck{
					OrganizationID: orgID,
					ResourceID:     "document",
					Action:         "delete",
				},
			},
			err: nil,
		},
		{
			description: "success when the member is authorized for role in org policy",
			in: shared.PerformB2BAuthorizationCheckIn{
				ProjectPolicy: sampleProjectPolicy,
				OrgPolicy:     sampleOrgPolicy,
				SubjectRoles:  []string{"sysadmin"},
				SubjectOrgID:  orgID,
				AuthorizationCheck: &b2bsessions.AuthorizationCheck{
					OrganizationID: orgID,
					ResourceID:     "program",
					Action:         "execute",
				},
			},
			err: nil,
		},
		{
			description: "success when the member is authorized for role in org policy (multiple roles)",
			in: shared.PerformB2BAuthorizationCheckIn{
				ProjectPolicy: sampleProjectPolicy,
				OrgPolicy:     sampleOrgPolicy,
				SubjectRoles:  []string{"stytch_member", "sysadmin"},
				SubjectOrgID:  orgID,
				AuthorizationCheck: &b2bsessions.AuthorizationCheck{
					OrganizationID: orgID,
					ResourceID:     "program",
					Action:         "execute",
				},
			},
			err: nil,
		},
	} {
		t.Run(testCase.description, func(t *testing.T) {
			err := shared.PerformB2BAuthorizationCheck(testCase.in)
			if testCase.err != nil {
				require.Error(t, err)
				assert.ErrorContains(t, err, testCase.err.Error())
			} else {
				assert.NoError(t, err)
			}
		})
	}
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

func Test_PerformB2BScopeAuthorizationCheck(t *testing.T) {
	const orgID = "organization-123"
	policy := &b2brbac.Policy{
		Roles: []b2brbac.PolicyRole{
			{
				RoleID:      "stytch_member",
				Description: "member",
				Permissions: []b2brbac.PolicyRolePermission{
					{
						ResourceID: "documents",
						Actions:    []string{"read", "write"},
					},
					{
						ResourceID: "images",
						Actions:    []string{"create", "read", "write", "delete"},
					},
				},
			},
		},
		Resources: []b2brbac.PolicyResource{
			{
				ResourceID:  "documents",
				Description: "All documents",
				Actions:     []string{"read", "write", "delete"},
			},
			{
				ResourceID:  "images",
				Description: "All images",
				Actions:     []string{"create", "read", "write", "delete"},
			},
		},
		Scopes: []b2brbac.PolicyScope{
			{
				Scope:       "read:data",
				Description: "Read data scope",
				Permissions: []b2brbac.PolicyScopePermission{
					{
						ResourceID: "documents",
						Actions:    []string{"read"},
					},
				},
			},
			{
				Scope:       "wildcard:data",
				Description: "Wildcard data scope",
				Permissions: []b2brbac.PolicyScopePermission{
					{
						ResourceID: "documents",
						Actions:    []string{"read"},
					},
				},
			},
			{
				Scope:       "write:data",
				Description: "Write data scope",
				Permissions: []b2brbac.PolicyScopePermission{
					{
						ResourceID: "documents",
						Actions:    []string{"write"},
					},
				},
			},
			{
				Scope:       "crud:data",
				Description: "CRUD data scope",
				Permissions: []b2brbac.PolicyScopePermission{
					{
						ResourceID: "images",
						Actions:    []string{"create", "read", "write", "delete"},
					},
				},
			},
		},
	}

	t.Run("success case - exact match", func(t *testing.T) {
		err := shared.PerformB2BScopeAuthorizationCheck(
			policy,
			[]string{"read:data"},
			orgID,
			&b2bsessions.AuthorizationCheck{
				OrganizationID: orgID,
				ResourceID:     "documents",
				Action:         "read",
			},
		)
		assert.NoError(t, err)
	})

	t.Run("success case - wildcard match", func(t *testing.T) {
		err := shared.PerformB2BScopeAuthorizationCheck(
			policy,
			[]string{"wildcard:data"},
			orgID,
			&b2bsessions.AuthorizationCheck{
				OrganizationID: orgID,
				ResourceID:     "documents",
				Action:         "read",
			},
		)
		assert.NoError(t, err)
	})

	t.Run("success case - multiple matches", func(t *testing.T) {
		err := shared.PerformB2BScopeAuthorizationCheck(
			policy,
			[]string{"read:data", "write:data"},
			orgID,
			&b2bsessions.AuthorizationCheck{
				OrganizationID: orgID,
				ResourceID:     "documents",
				Action:         "read",
			},
		)
		assert.NoError(t, err)
	})

	t.Run("success case - multiple matches II", func(t *testing.T) {
		err := shared.PerformB2BScopeAuthorizationCheck(
			policy,
			[]string{"read:data", "write:data", "crud:data"},
			orgID,
			&b2bsessions.AuthorizationCheck{
				OrganizationID: orgID,
				ResourceID:     "images",
				Action:         "create",
			},
		)
		assert.NoError(t, err)
	})

	t.Run("failure case - invalid action", func(t *testing.T) {
		err := shared.PerformB2BScopeAuthorizationCheck(
			policy,
			[]string{"write:data"},
			orgID,
			&b2bsessions.AuthorizationCheck{
				OrganizationID: orgID,
				ResourceID:     "documents",
				Action:         "delete",
			},
		)
		assert.ErrorContains(t, err, stytcherror.NewPermissionError().Error())
	})

	t.Run("failure case - invalid resource", func(t *testing.T) {
		err := shared.PerformB2BScopeAuthorizationCheck(
			policy,
			[]string{"crud:data"},
			orgID,
			&b2bsessions.AuthorizationCheck{
				OrganizationID: orgID,
				ResourceID:     "spreadsheets",
				Action:         "write",
			},
		)
		assert.ErrorContains(t, err, stytcherror.NewPermissionError().Error())
	})

	t.Run("failure case - invalid tenancy check", func(t *testing.T) {
		diffOrgID := "organization-456"
		err := shared.PerformB2BScopeAuthorizationCheck(
			policy,
			[]string{"crud:data"},
			diffOrgID,
			&b2bsessions.AuthorizationCheck{
				OrganizationID: orgID,
				ResourceID:     "images",
				Action:         "write",
			},
		)
		assert.ErrorContains(t, err, stytcherror.NewSessionAuthorizationTenancyError(diffOrgID, orgID).Error())
	})
}

func Test_PerformConsumerScopeAuthorizationCheck(t *testing.T) {
	policy := &consumerrbac.Policy{
		Roles: []consumerrbac.PolicyRole{
			{
				RoleID:      "stytch_member",
				Description: "member",
				Permissions: []consumerrbac.PolicyRolePermission{
					{
						ResourceID: "documents",
						Actions:    []string{"read", "write"},
					},
					{
						ResourceID: "images",
						Actions:    []string{"create", "read", "write", "delete"},
					},
				},
			},
		},
		Resources: []consumerrbac.PolicyResource{
			{
				ResourceID:  "documents",
				Description: "All documents",
				Actions:     []string{"read", "write", "delete"},
			},
			{
				ResourceID:  "images",
				Description: "All images",
				Actions:     []string{"create", "read", "write", "delete"},
			},
		},
		Scopes: []consumerrbac.PolicyScope{
			{
				Scope:       "read:data",
				Description: "Read data scope",
				Permissions: []consumerrbac.PolicyScopePermission{
					{
						ResourceID: "documents",
						Actions:    []string{"read"},
					},
				},
			},
			{
				Scope:       "wildcard:data",
				Description: "Wildcard data scope",
				Permissions: []consumerrbac.PolicyScopePermission{
					{
						ResourceID: "documents",
						Actions:    []string{"read"},
					},
				},
			},
			{
				Scope:       "write:data",
				Description: "Write data scope",
				Permissions: []consumerrbac.PolicyScopePermission{
					{
						ResourceID: "documents",
						Actions:    []string{"write"},
					},
				},
			},
			{
				Scope:       "crud:data",
				Description: "CRUD data scope",
				Permissions: []consumerrbac.PolicyScopePermission{
					{
						ResourceID: "images",
						Actions:    []string{"create", "read", "write", "delete"},
					},
				},
			},
		},
	}

	t.Run("success case - exact match", func(t *testing.T) {
		err := shared.PerformConsumerScopeAuthorizationCheck(
			policy,
			[]string{"read:data"},
			&consumersessions.AuthorizationCheck{
				ResourceID: "documents",
				Action:     "read",
			},
		)
		assert.NoError(t, err)
	})

	t.Run("success case - wildcard match", func(t *testing.T) {
		err := shared.PerformConsumerScopeAuthorizationCheck(
			policy,
			[]string{"wildcard:data"},
			&consumersessions.AuthorizationCheck{
				ResourceID: "documents",
				Action:     "read",
			},
		)
		assert.NoError(t, err)
	})

	t.Run("success case - multiple matches", func(t *testing.T) {
		err := shared.PerformConsumerScopeAuthorizationCheck(
			policy,
			[]string{"read:data", "write:data"},
			&consumersessions.AuthorizationCheck{
				ResourceID: "documents",
				Action:     "read",
			},
		)
		assert.NoError(t, err)
	})

	t.Run("success case - multiple matches II", func(t *testing.T) {
		err := shared.PerformConsumerScopeAuthorizationCheck(
			policy,
			[]string{"read:data", "write:data", "crud:data"},
			&consumersessions.AuthorizationCheck{
				ResourceID: "images",
				Action:     "create",
			},
		)
		assert.NoError(t, err)
	})

	t.Run("failure case - invalid action", func(t *testing.T) {
		err := shared.PerformConsumerScopeAuthorizationCheck(
			policy,
			[]string{"write:data"},
			&consumersessions.AuthorizationCheck{
				ResourceID: "documents",
				Action:     "delete",
			},
		)
		assert.ErrorContains(t, err, stytcherror.NewPermissionError().Error())
	})

	t.Run("failure case - invalid resource", func(t *testing.T) {
		err := shared.PerformConsumerScopeAuthorizationCheck(
			policy,
			[]string{"crud:data"},
			&consumersessions.AuthorizationCheck{
				ResourceID: "spreadsheets",
				Action:     "write",
			},
		)
		assert.ErrorContains(t, err, stytcherror.NewPermissionError().Error())
	})
}
