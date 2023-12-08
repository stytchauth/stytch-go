package rbac

// !!!
// WARNING: This file is autogenerated
// Only modify code within MANUAL() sections
// or your changes may be overwritten later!
// !!!

type Policy struct {
	Roles     []PolicyRole     `json:"roles,omitempty"`
	Resources []PolicyResource `json:"resources,omitempty"`
}
type (
	PolicyParams   struct{}
	PolicyResource struct {
		ResourceID  string   `json:"resource_id,omitempty"`
		Description string   `json:"description,omitempty"`
		Actions     []string `json:"actions,omitempty"`
	}
)

type PolicyRole struct {
	RoleID      string                 `json:"role_id,omitempty"`
	Description string                 `json:"description,omitempty"`
	Permissions []PolicyRolePermission `json:"permissions,omitempty"`
}
type PolicyRolePermission struct {
	ResourceID string   `json:"resource_id,omitempty"`
	Actions    []string `json:"actions,omitempty"`
}
type PolicyResponse struct {
	RequestID  string  `json:"request_id,omitempty"`
	StatusCode int32   `json:"status_code,omitempty"`
	Policy     *Policy `json:"policy,omitempty"`
}
