package database

import "testing"

func TestHashPassword(t *testing.T) {
	password := "password"
	_, err := HashPassword(password)
	checkErrNil(t, err)
}

func TestCheckPasswordWithHash(t *testing.T) {
	password := "password"
	hash, _ := HashPassword(password)
	result := CheckPasswordWithHash(password, hash)

	if !result {
		t.Error("password check returned false")
	}
}

func BenchmarkHashPassword(b *testing.B) {
	for i := 0; i < b.N; i++ {
		HashPassword("this is a secure password")
	}
}
