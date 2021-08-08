package config

import (
	"regexp"
	"testing"
)

func TestConfig(t *testing.T) {
	cfgData := []byte(`
database:
  host: 0.0.0.0
  port: 5432
  username: root
  password: rootpasswd
  dbname: golangDB
server:
  address: ":8000"
  staticFolder: web/static
`)
	err := LoadConfig(cfgData)
	if err != nil {
		t.Fatalf("Failed to load config: %v", err)
	}
	cfg := Get()
	matched, err := regexp.MatchString(`host=.+ port=\d+ user=.+ password=.+ dbname=.+ sslmode=disable`, cfg.Database.DataSourceName())
	if matched == false || err != nil {
		t.Fatalf("Failed database source string with error: %v", err)
	}
}

func TestSet(t *testing.T) {
	cfg := Config{
		Database: Database{
			Host:     "localhost",
			Port:     5432,
			Username: "testing",
			Password: "testing",
			DBName:   "testing",
		},
		Server: Server{
			Address:      ":8000",
			StaticFolder: "web/static/",
		},
	}

	Set(cfg)

	if configInstance != cfg {
		t.Error("Didn't set config")
	}
}
