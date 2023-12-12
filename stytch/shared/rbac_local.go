package shared

import (
	"github.com/stytchauth/stytch-go/v11/stytch/b2b/rbac"
	"github.com/stytchauth/stytch-go/v11/stytch/b2b/sessions"
	"github.com/stytchauth/stytch-go/v11/stytch/stytcherror"
)

func PerformAuthorizationCheck(
	policy *rbac.Policy,
	subjectRoles []string,
	subjectOrgID string,
	authorizationCheck *sessions.AuthorizationCheck,
) error {
	if authorizationCheck == nil {
		return nil
	}

	if subjectOrgID != authorizationCheck.OrganizationID {
		return stytcherror.NewSessionAuthorizationTenancyError(subjectOrgID, authorizationCheck.OrganizationID)
	}

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

func contains(stringList []string, target string) bool {
	for _, s := range stringList {
		if target == s {
			return true
		}
	}
	return false
}
