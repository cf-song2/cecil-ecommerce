package util

import "testing"

func TestPasswordHashing(t *testing.T) {
	plain := "secret123"

	hash, err := HashPassword(plain)
	if err != nil {
		t.Fatalf("Failed to hash password: %v", err)
	}

	if !CheckPasswordHash(plain, hash) {
		t.Errorf("Password check failed: expected true")
	}

	if CheckPasswordHash("wrongpass", hash) {
		t.Errorf("Password check should fail with wrong password")
	}
}
