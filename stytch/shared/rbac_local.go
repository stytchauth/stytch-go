package shared

import (
	b2brbac "github.com/stytchauth/stytch-go/v17/stytch/b2b/rbac"
	b2bsessions "github.com/stytchauth/stytch-go/v17/stytch/b2b/sessions"
	consumerrbac "github.com/stytchauth/stytch-go/v17/stytch/consumer/rbac"
	consumersessions "github.com/stytchauth/stytch-go/v17/stytch/consumer/sessions"
	"github.com/stytchauth/stytch-go/v17/stytch/stytcherror"
)

type PerformB2BAuthorizationCheckIn struct {
	ProjectPolicy      *b2brbac.Policy
	OrgPolicy          *b2brbac.OrgPolicy
	SubjectRoles       []string
	SubjectOrgID       string
	AuthorizationCheck *b2bsessions.AuthorizationCheck
}

func PerformB2BAuthorizationCheck(in PerformB2BAuthorizationCheckIn) error {
	if in.AuthorizationCheck == nil {
		return nil
	}

	if in.SubjectOrgID != in.AuthorizationCheck.OrganizationID {
		return stytcherror.NewSessionAuthorizationTenancyError(in.SubjectOrgID, in.AuthorizationCheck.OrganizationID)
	}

	return performRoleAuthorizationCheck(
		policyFromB2B(in.ProjectPolicy, in.OrgPolicy),
		authorizationCheckFromB2B(in.AuthorizationCheck),
		in.SubjectRoles,
	)
}

func PerformB2BScopeAuthorizationCheck(
	policy *b2brbac.Policy, // NOTE: org policy check not needed here.
	tokenScopes []string,
	subjectOrgID string,
	authorizationCheck *b2bsessions.AuthorizationCheck,
) error {
	if authorizationCheck == nil {
		return nil
	}

	if subjectOrgID != authorizationCheck.OrganizationID {
		return stytcherror.NewSessionAuthorizationTenancyError(subjectOrgID, authorizationCheck.OrganizationID)
	}
	// Custom Org Roles don't interact with scopes.
	return performScopeAuthorizationCheck(policyFromB2B(policy, nil), authorizationCheckFromB2B(authorizationCheck), tokenScopes)
}

func PerformConsumerAuthorizationCheck(
	policy *consumerrbac.Policy,
	subjectRoles []string,
	authorizationCheck *consumersessions.AuthorizationCheck,
) error {
	if authorizationCheck == nil {
		return nil
	}
	return performRoleAuthorizationCheck(policyFromConsumer(policy), authorizationCheckFromConsumer(authorizationCheck), subjectRoles)
}

func PerformConsumerScopeAuthorizationCheck(
	policy *consumerrbac.Policy,
	tokenScopes []string,
	authorizationCheck *consumersessions.AuthorizationCheck,
) error {
	if authorizationCheck == nil {
		return nil
	}
	return performScopeAuthorizationCheck(policyFromConsumer(policy), authorizationCheckFromConsumer(authorizationCheck), tokenScopes)
}

func performRoleAuthorizationCheck(policy *policy, authorizationCheck *authorizationCheck, subjectRoles []string) error {
	for _, role := range policy.Roles {
		if contains(subjectRoles, role.RoleID) {
			for _, permission := range role.Permissions {
				hasMatchingAction := contains(permission.Actions, "*") ||
					contains(permission.Actions, authorizationCheck.Action)
				hasMatchingResource := permission.ResourceID == authorizationCheck.ResourceID
				if hasMatchingAction && hasMatchingResource {
					return nil
				}
			}
		}
	}
	return stytcherror.NewPermissionError()
}

func performScopeAuthorizationCheck(policy *policy, authorizationCheck *authorizationCheck, tokenScopes []string) error {
	for _, scope := range policy.Scopes {
		if contains(tokenScopes, scope.Scope) {
			for _, permission := range scope.Permissions {
				hasMatchingAction := contains(permission.Actions, "*") ||
					contains(permission.Actions, authorizationCheck.Action)
				hasMatchingResource := permission.ResourceID == authorizationCheck.ResourceID
				if hasMatchingAction && hasMatchingResource {
					return nil
				}
			}
		}
	}

	return stytcherror.NewPermissionError()
}

func contains(stringList []string, target string) bool {
	for _, s := range stringList {
		if target == s {
			return true
		}
	}
	return false
}
