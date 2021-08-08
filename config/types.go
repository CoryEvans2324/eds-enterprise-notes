package config

import "fmt"

type Config struct {
	Database Database `yaml:"database"`
	Server   Server   `yaml:"server"`
}

type Database struct {
	Host     string `yaml:"host"`
	Port     int    `yaml:"port"`
	Username string `yaml:"username"`
	Password string `yaml:"password"`
	DBName   string `yaml:"dbname"`
}

type Server struct {
	Address      string `yaml:"address"`
	StaticFolder string `yaml:"staticFolder"`
}

func (d *Database) DataSourceName() string {
	return fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", d.Host, d.Port, d.Username, d.Password, d.DBName)
}
