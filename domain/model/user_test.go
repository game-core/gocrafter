package model

import (
	"testing"
)

func TestUser_EmptyUser(t *testing.T) {
	tests := []struct {
		name     string
		user     *User
		expected bool
	}{
		{
			name:     "正常系: モデルが空の場合",
			user:     EmptyUser(),
			expected: true,
		},
		{
			name: "正常系: モデルが空でない場合",
			user: &User{
				ID:       1,
				UserKey:  "test_key",
				UserName: "test_name",
				Email:    "test@example.com",
				Password: "password123",
				Token:    "token123",
			},
			expected: false,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			if empty := test.user.IsEmpty(); empty != test.expected {
				t.Errorf("Expected IsEmpty() to return %v, but got %v", test.expected, empty)
			}
		})
	}
}
