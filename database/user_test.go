package database

import "testing"

func TestUser(t *testing.T) {
	cfg := createConfig()
	CreateDatabaseManager(cfg.Database.DataSourceName())
	ResetDB()

	const dummyUsername = "testuser"
	const dummyPassword = "testpasswd"

	_, err := Mgr.CreateUser(dummyUsername, dummyPassword)
	checkErrNil(t, err)

	_, err = Mgr.GetUserByID(12345678)
	if err == nil {
		t.Error("GetUserByID should have failed with an ID that doesn't exist")
	}

	_, err = Mgr.GetUserByUsername(dummyUsername)
	checkErrNil(t, err)

	_, err = Mgr.GetUserByUsername("This should not exist")
	if err == nil {
		t.Error("GetUserByUser should have failed with a username that doesn't exist")
	}
}

func TestUserPasswords(t *testing.T) {
	cfg := createConfig()
	CreateDatabaseManager(cfg.Database.DataSourceName())
	ResetDB()

	const dummyUsername = "testuser"
	const dummyPassword = "testpasswd"

	user, _ := Mgr.CreateUser(dummyUsername, dummyPassword)

	// Check password hash
	// Correct username & password
	if !CheckPasswordWithHash(dummyPassword, user.PasswordHash) {
		t.Errorf("Incorrect hash %s for password %s", user.PasswordHash, dummyPassword)
	}
}
