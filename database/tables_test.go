package database

import "testing"

func TestCreatDatabaseManager(t *testing.T) {

	cfg := createConfig()
	CreateDatabaseManager(cfg.Database.DataSourceName())

	if Mgr == nil {
		t.Fatal("Database creation failed")
	}

	Mgr.Close()
}

func TestRecreateTables(t *testing.T) {
	cfg := createConfig()
	CreateDatabaseManager(cfg.Database.DataSourceName())

	err := Mgr.DropUserTable()
	checkErrNil(t, err)

	err = Mgr.CreateUserTable()
	checkErrNil(t, err)
}
