package b2b

type DiscoveredOrganization struct {
	Organization        Organization `json:"organization,omitempty"`
	Membership          Membership   `json:"membership,omitempty"`
	MemberAuthenticated bool         `json:"member_authenticated,omitempty"`
}

type Membership struct {
	Type    string                 `json:"type,omitempty"`
	Details map[string]interface{} `json:"details,omitempty"`
	Member  Member                 `json:"member,omitempty"`
}
