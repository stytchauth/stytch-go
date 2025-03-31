package shared

import (
	"strings"

	"github.com/stytchauth/stytch-go/v17/stytch/consumer/m2m"
	"github.com/stytchauth/stytch-go/v17/stytch/stytcherror"
)

func PerformM2MAuthorizationCheck(params m2m.ScopeAuthorizationFuncParams) error {
	clientScopes := map[string][]string{}
	for _, scope := range params.HasScopes {
		action := scope
		resource := "-"
		if strings.Contains(scope, ":") {
			parts := strings.SplitN(scope, ":", 2)
			action = parts[0]
			resource = parts[1]
		}
		clientScopes[action] = append(clientScopes[action], resource)
	}

	for _, requiredScope := range params.RequiredScopes {
		requiredAction := requiredScope
		requiredResource := "-"
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
