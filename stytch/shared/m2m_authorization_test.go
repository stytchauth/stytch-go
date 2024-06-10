package shared_test

import (
	"testing"

	"github.com/stytchauth/stytch-go/v15/stytch/consumer/m2m"
	"github.com/stytchauth/stytch-go/v15/stytch/shared"
)

func TestPerformM2MAuthorizationCheck(t *testing.T) {
	tests := []struct {
		name        string
		has         []string
		needs       []string
		expectError bool
	}{
		{
			name:        "basic",
			has:         []string{"read:users", "write:users"},
			needs:       []string{"read:users"},
			expectError: false,
		},
		{
			name:        "multiple required scopes",
			has:         []string{"read:users", "write:users", "read:books"},
			needs:       []string{"read:users", "read:books"},
			expectError: false,
		},
		{
			name:        "simple scopes",
			has:         []string{"read_users", "write_users"},
			needs:       []string{"read_users"},
			expectError: false,
		},
		{
			name:        "wildcard resource",
			has:         []string{"read:*", "write:*"},
			needs:       []string{"read:users"},
			expectError: false,
		},
		{
			name:        "missing required scope",
			has:         []string{"read:users"},
			needs:       []string{"write:users"},
			expectError: true,
		},
		{
			name:        "missing required scope with wildcard",
			has:         []string{"read:users", "write:*"},
			needs:       []string{"delete:books"},
			expectError: true,
		},
		{
			name:        "has simple scope and wants specific scope",
			has:         []string{"read"},
			needs:       []string{"read:users"},
			expectError: true,
		},
		{
			name:        "has specific scope and wants simple scope",
			has:         []string{"read:users"},
			needs:       []string{"read"},
			expectError: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := shared.PerformM2MAuthorizationCheck(m2m.ScopeAuthorizationFuncParams{
				HasScopes:      tt.has,
				RequiredScopes: tt.needs,
			})
			if (err != nil) != tt.expectError {
				t.Errorf("PerformM2MAuthorizationCheck(%v, %v) error = %v, expectError %v", tt.has, tt.needs, err, tt.expectError)
			}
		})
	}
}
