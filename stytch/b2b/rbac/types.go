package rbac

type Policy struct {
	Roles []PolicyRole `json:"roles"`
}

type PolicyRole struct {
	RoleID      string       `json:"role_id"`
	Description string       `json:"description"`
	Permissions []Permission `json:"permissions"`
}

type Permission struct {
	ResourceID string   `json:"resource_id"`
	Actions    []string `json:"actions"`
}
