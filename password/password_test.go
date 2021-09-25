package password

import (
	"testing"
)

func TestPasswordsMatch(t *testing.T) {
	password := "secret"
	hash, _ := HashPassword(password, 14) // ignore error for the sake of simplicity


	match, err := CheckPasswordHash(password, hash)

	if err != nil {
		t.Fatalf("err returned from CheckPasswordHash should be nil")
	}

	if !match {
		t.Fatalf("hashed passwords do not match")
	}
}
