package models

import "testing"

// TestUser function
func TestUser(t *testing.T) *User {
	return &User{
		ID: "1",
		Email: "user@example.com",
		Password: "password",
	}
}
