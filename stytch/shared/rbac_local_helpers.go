package shared

import (
	b2brbac "github.com/stytchauth/stytch-go/v17/stytch/b2b/rbac"
	b2bsessions "github.com/stytchauth/stytch-go/v17/stytch/b2b/sessions"
	consumerrbac "github.com/stytchauth/stytch-go/v17/stytch/consumer/rbac"
	consumersessions "github.com/stytchauth/stytch-go/v17/stytch/consumer/sessions"
)

// This file contains boilerplate stubs to convert B2B and Consumer RBAC policies
// and authorization check API shapes to the same internal data shape
// for use in local authorization methods

type policy struct {
	Roles     []policyRole
	Resources []policyResource
	Scopes    []policyScope
}

type policyRole struct {
	RoleID      string
	Description string
	Permissions []policyRolePermission
}

type policyRolePermission struct {
	ResourceID string
	Actions    []string
}

type policyScope struct {
	Scope       string
	Description string
	Permissions []policyScopePermission
}

type policyScopePermission struct {
	ResourceID string
	Actions    []string
}

type policyResource struct {
	ResourceID  string
	Description string
	Actions     []string
}

func policyFromB2B(projectPolicy *b2brbac.Policy, orgPolicy *b2brbac.OrgPolicy) *policy {
	var roles []policyRole
	for _, role := range projectPolicy.Roles {
		var permissions []policyRolePermission
		for _, permission := range role.Permissions {
			permissions = append(permissions, policyRolePermission{
				ResourceID: permission.ResourceID,
				Actions:    permission.Actions,
			})
		}
		roles = append(roles, policyRole{
			RoleID:      role.RoleID,
			Description: role.Description,
			Permissions: permissions,
		})
	}
	if orgPolicy != nil {
		for _, role := range orgPolicy.Roles {
			var permissions []policyRolePermission
			for _, permission := range role.Permissions {
				permissions = append(permissions, policyRolePermission{
					ResourceID: permission.ResourceID,
					Actions:    permission.Actions,
				})
			}
			roles = append(roles, policyRole{
				RoleID:      role.RoleID,
				Description: role.Description,
				Permissions: permissions,
			})
		}
	}

	var resources []policyResource
	for _, resource := range projectPolicy.Resources {
		resources = append(resources, policyResource{
			ResourceID:  resource.ResourceID,
			Description: resource.Description,
			Actions:     resource.Actions,
		})
	}

	var scopes []policyScope
	for _, scope := range projectPolicy.Scopes {
		var permissions []policyScopePermission
		for _, permission := range scope.Permissions {
			permissions = append(permissions, policyScopePermission{
				ResourceID: permission.ResourceID,
				Actions:    permission.Actions,
			})
		}
		scopes = append(scopes, policyScope{
			Scope:       scope.Scope,
			Description: scope.Description,
			Permissions: permissions,
		})
	}

	return &policy{
		Roles:     roles,
		Resources: resources,
		Scopes:    scopes,
	}
}

func policyFromConsumer(p *consumerrbac.Policy) *policy {
	var roles []policyRole
	for _, role := range p.Roles {
		var permissions []policyRolePermission
		for _, permission := range role.Permissions {
			permissions = append(permissions, policyRolePermission{
				ResourceID: permission.ResourceID,
				Actions:    permission.Actions,
			})
		}

		roles = append(roles, policyRole{
			RoleID:      role.RoleID,
			Description: role.Description,
			Permissions: permissions,
		})
	}

	var resources []policyResource
	for _, resource := range p.Resources {
		resources = append(resources, policyResource{
			ResourceID:  resource.ResourceID,
			Description: resource.Description,
			Actions:     resource.Actions,
		})
	}

	var scopes []policyScope
	for _, scope := range p.Scopes {
		var permissions []policyScopePermission
		for _, permission := range scope.Permissions {
			permissions = append(permissions, policyScopePermission{
				ResourceID: permission.ResourceID,
				Actions:    permission.Actions,
			})
		}
		scopes = append(scopes, policyScope{
			Scope:       scope.Scope,
			Description: scope.Description,
			Permissions: permissions,
		})
	}

	return &policy{
		Roles:     roles,
		Resources: resources,
		Scopes:    scopes,
	}
}

type authorizationCheck struct {
	ResourceID string
	Action     string
}

func authorizationCheckFromConsumer(c *consumersessions.AuthorizationCheck) *authorizationCheck {
	return &authorizationCheck{
		ResourceID: c.ResourceID,
		Action:     c.Action,
	}
}

func authorizationCheckFromB2B(c *b2bsessions.AuthorizationCheck) *authorizationCheck {
	return &authorizationCheck{
		ResourceID: c.ResourceID,
		Action:     c.Action,
	}
}
