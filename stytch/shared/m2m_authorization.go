package shared

import (
	"strings"

	"github.com/stytchauth/stytch-go/v15/stytch/stytcherror"
)

func PerformM2MAuthorizationCheck(
	hasScopes []string,
	requiredScopes []string,
) error {
	clientScopes := map[string][]string{}
	for _, scope := range hasScopes {
		action := scope
		resource := "*"
		if strings.Contains(scope, ":") {
			parts := strings.SplitN(scope, ":", 2)
			action = parts[0]
			resource = parts[1]
		}
		clientScopes[action] = append(clientScopes[action], resource)
	}

	for _, requiredScope := range requiredScopes {
		requiredAction := requiredScope
		requiredResource := "*"
		if strings.Contains(requiredScope, ":") {
			parts := strings.SplitN(requiredScope, ":", 2)
			requiredAction = parts[0]
			requiredResource = parts[1]
		}
		resources, ok := clientScopes[requiredAction]
		if !ok {
			return stytcherror.NewM2MPermissionError()
		}

		found := false
		for _, resource := range resources {
			if resource == "*" || resource == requiredResource {
				found = true
				break
			}
		}

		if !found {
			return stytcherror.NewM2MPermissionError()
		}
	}
	return nil
}
