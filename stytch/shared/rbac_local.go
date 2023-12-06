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
		return stytcherror.NewClientLibraryError("Subject organization ID does not match ID from request")
	}

	for _, role := range policy.Roles {
		if doesListOfStringsContainTarget(subjectRoles, role.RoleID) {
			for _, permission := range role.Permissions {
				hasMatchingAction := doesListOfStringsContainTarget(permission.Actions, "*") ||
					doesListOfStringsContainTarget(permission.Actions, authorizationCheck.Action)
				hasMatchingResource := permission.ResourceID == authorizationCheck.ResourceID
				if hasMatchingAction && hasMatchingResource {
					return nil
				}
			}
		}
	}

	return stytcherror.NewClientLibraryError("Member is not authorized")
}

func doesListOfStringsContainTarget(stringList []string, target string) bool {
	for _, s := range stringList {
		if target == s {
			return true
		}
	}
	return false
}
