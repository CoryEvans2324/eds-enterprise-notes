package database

import "testing"

func TestUser(t *testing.T) {
	cfg := createConfig()
	CreateDatabaseManager(cfg.Database.DataSourceName())

	const dummyUsername = "testuser"
	const dummyPassword = "testpasswd"

	Mgr.DropUserTable()
	Mgr.CreateUserTable()

	userId, err := Mgr.CreateUser(dummyUsername, dummyPassword)
	checkErrNil(t, err)

	_, err = Mgr.GetUserByID(userId)
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

	const dummyUsername = "testuser"
	const dummyPassword = "testpasswd"

	Mgr.DropUserTable()
	Mgr.CreateUserTable()

	userId, _ := Mgr.CreateUser(dummyUsername, dummyPassword)

	user, _ := Mgr.GetUserByID(userId)

	// Check password hash
	// Correct username & password
	hash, err := Mgr.GetPasswordHash(user.Username)
	checkErrNil(t, err)
	if !CheckPasswordWithHash(dummyPassword, hash) {
		t.Errorf("Incorrect hash %s for password %s", hash, dummyPassword)
	}

	// Username doesn't exist
	_, err = Mgr.GetPasswordHash("ThisUsernameShouldNotExist")
	if err == nil {
		t.Error("GetPasswordHash should have returned an error as no user exist")
	}
}
