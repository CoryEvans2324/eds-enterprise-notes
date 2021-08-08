package config

import (
	"gopkg.in/yaml.v2"
)

var configInstance Config

func LoadConfig(data []byte) error {
	return yaml.Unmarshal(data, &configInstance)
}

func Get() *Config {
	return &configInstance
}
