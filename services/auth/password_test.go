package auth

import "testing"



func TestHashPassword(t *testing.T) {
	hash, err := HashPassword("password")
	if err != nil {
		t.Errorf("error hashing password: %v", err)
	}

	if hash == "" {
		t.Error("expected hash to be not empty")
	}

	if hash == "password" {
		t.Error("expected hash to be different from password")
	}
}

func TestComparePasswords(t *testing.T) {
	hash, err := HashPassword("password")
	if err != nil {
		t.Errorf("error hashing password: %v", err)
	}

	err=VerifyPassword("password",hash);
	if err!=nil {
		t.Errorf("expected password to match hash: %v",err);
	}

	err=VerifyPassword("notpassword",hash);
	if err!=nil {
		t.Errorf("expected password to not match hash")
	}
}