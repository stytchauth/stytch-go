package rbac

// !!!
// WARNING: This file is autogenerated
// Only modify code within MANUAL() sections
// or your changes may be overwritten later!
// !!!

// Policy:
type Policy struct {
	// Roles: An array of [Role objects](https://stytch.com/docs/b2b/api/rbac-role-object).
	Roles []PolicyRole `json:"roles,omitempty"`
	// Resources: An array of [Resource objects](https://stytch.com/docs/b2b/api/rbac-resource-object).
	Resources []PolicyResource `json:"resources,omitempty"`
	Scopes    []PolicyScope    `json:"scopes,omitempty"`
}

// PolicyParams: Request type for `RBAC.Policy`.
type PolicyParams struct{}

// PolicyResource:
type PolicyResource struct {
	// ResourceID: A unique identifier of the RBAC Resource, provided by the developer and intended to be
	// human-readable.
	//
	//   A `resource_id` is not allowed to start with `stytch`, which is a special prefix used for Stytch
	// default Resources with reserved  `resource_id`s. These include:
	//
	//   * `stytch.organization`
	//   * `stytch.member`
	//   * `stytch.sso`
	//   * `stytch.self`
	//
	//   Check out the
	// [guide on Stytch default Resources](https://stytch.com/docs/b2b/guides/rbac/stytch-default) for a more
	// detailed explanation.
	//
	//
	ResourceID string `json:"resource_id,omitempty"`
	// Description: The description of the RBAC Resource.
	Description string `json:"description,omitempty"`
	// Actions: A list of all possible actions for a provided Resource.
	//
	//   Reserved `actions` that are predefined by Stytch include:
	//
	//   * `*`
	//   * For the `stytch.organization` Resource:
	//     * `update.info.name`
	//     * `update.info.slug`
	//     * `update.info.untrusted_metadata`
	//     * `update.info.email_jit_provisioning`
	//     * `update.info.logo_url`
	//     * `update.info.email_invites`
	//     * `update.info.allowed_domains`
	//     * `update.info.default_sso_connection`
	//     * `update.info.sso_jit_provisioning`
	//     * `update.info.mfa_policy`
	//     * `update.info.implicit_roles`
	//     * `delete`
	//   * For the `stytch.member` Resource:
	//     * `create`
	//     * `update.info.name`
	//     * `update.info.untrusted_metadata`
	//     * `update.info.mfa-phone`
	//     * `update.info.delete.mfa-phone`
	//     * `update.settings.is-breakglass`
	//     * `update.settings.mfa_enrolled`
	//     * `update.settings.roles`
	//     * `search`
	//     * `delete`
	//   * For the `stytch.sso` Resource:
	//     * `create`
	//     * `update`
	//     * `delete`
	//   * For the `stytch.self` Resource:
	//     * `update.info.name`
	//     * `update.info.untrusted_metadata`
	//     * `update.info.mfa-phone`
	//     * `update.info.delete.mfa-phone`
	//     * `update.info.delete.password`
	//     * `update.settings.mfa_enrolled`
	//     * `delete`
	//
	Actions []string `json:"actions,omitempty"`
}

// PolicyRole:
type PolicyRole struct {
	// RoleID: The unique identifier of the RBAC Role, provided by the developer and intended to be
	// human-readable.
	//
	//   Reserved `role_id`s that are predefined by Stytch include:
	//
	//   * `stytch_member`
	//   * `stytch_admin`
	//
	//   Check out the [guide on Stytch default Roles](https://stytch.com/docs/b2b/guides/rbac/stytch-default)
	// for a more detailed explanation.
	//
	//
	RoleID string `json:"role_id,omitempty"`
	// Description: The description of the RBAC Role.
	Description string `json:"description,omitempty"`
	// Permissions: A list of permissions that link a
	// [Resource](https://stytch.com/docs/b2b/api/rbac-resource-object) to a list of actions.
	Permissions []PolicyRolePermission `json:"permissions,omitempty"`
}

// PolicyRolePermission:
type PolicyRolePermission struct {
	// ResourceID: A unique identifier of the RBAC Resource, provided by the developer and intended to be
	// human-readable.
	//
	//   A `resource_id` is not allowed to start with `stytch`, which is a special prefix used for Stytch
	// default Resources with reserved  `resource_id`s. These include:
	//
	//   * `stytch.organization`
	//   * `stytch.member`
	//   * `stytch.sso`
	//   * `stytch.self`
	//
	//   Check out the
	// [guide on Stytch default Resources](https://stytch.com/docs/b2b/guides/rbac/stytch-default) for a more
	// detailed explanation.
	//
	//
	ResourceID string `json:"resource_id,omitempty"`
	// Actions: A list of permitted actions the Role is authorized to take with the provided Resource. You can
	// use `*` as a wildcard to grant a Role permission to use all possible actions related to the Resource.
	Actions []string `json:"actions,omitempty"`
}

type PolicyScope struct {
	Scope       string                  `json:"scope,omitempty"`
	Description string                  `json:"description,omitempty"`
	Permissions []PolicyScopePermission `json:"permissions,omitempty"`
}

type PolicyScopePermission struct {
	ResourceID string   `json:"resource_id,omitempty"`
	Actions    []string `json:"actions,omitempty"`
}

// PolicyResponse: Response type for `RBAC.Policy`.
type PolicyResponse struct {
	// RequestID: Globally unique UUID that is returned with every API call. This value is important to log for
	// debugging purposes; we may ask for this value to help identify a specific API call when helping you
	// debug an issue.
	RequestID string `json:"request_id,omitempty"`
	// StatusCode: The HTTP status code of the response. Stytch follows standard HTTP response status code
	// patterns, e.g. 2XX values equate to success, 3XX values are redirects, 4XX are client errors, and 5XX
	// are server errors.
	StatusCode int32 `json:"status_code,omitempty"`
	// Policy: The RBAC Policy document that contains all defined Roles and Resources – which are managed in
	// the [Dashboard](https://stytch.com/docs/dashboard/rbac). Read more about these entities and how they
	// work in our [RBAC overview](https://stytch.com/docs/b2b/guides/rbac/overview).
	Policy *Policy `json:"policy,omitempty"`
}
