package database

import (
	"testing"

	"github.com/CoryEvans2324/eds-enterprise-notes/config"
	"github.com/CoryEvans2324/eds-enterprise-notes/models"
)

func checkErrNil(t *testing.T, err error) {
	if err != nil {
		t.Error(err)
	}
}

func createConfig() config.Config {
	cfgData := []byte(`
database:
  host: 0.0.0.0
  port: 5432
  username: testing
  password: testing
  dbname: testing
server:
  address: ":8000"
  staticFolder: web/static
`)
	config.LoadConfig(cfgData)
	return *config.Get()
}

func ResetDB() {
	db.Exec("DROP TABLE IF EXISTS permissions")
	db.Exec("DROP TABLE IF EXISTS notes")
	db.Exec("DROP TABLE IF EXISTS users")

	db.AutoMigrate(&models.User{})
	db.AutoMigrate(&models.Note{})
	db.AutoMigrate(&models.Permission{})
}

func TestCreateDatabaseManager(t *testing.T) {
	cfg := createConfig()
	err := CreateDatabaseManager(cfg.Database.DataSourceName())
	checkErrNil(t, err)
	if Mgr == nil {
		t.Fatalf("Failed to create database manager")
	}

	err = CreateDatabaseManager("")
	if err == nil {
		t.Error("sql.Open(...) should have failed with incorrect datasourcename")
	}

	cfg.Database.Port = 0

	err = CreateDatabaseManager(cfg.Database.DataSourceName())
	if err == nil {
		t.Error("Should have failed with incorrect port")
	}
}
