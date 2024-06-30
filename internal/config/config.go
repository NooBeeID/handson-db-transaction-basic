package config

import (
	"os"

	"gopkg.in/yaml.v3"
)

type Config struct {
	App AppConfig `yaml:"app"`
	DB  DBConfig  `yaml:"database"`
}

type DBConfig struct {
	Host           string                 `yaml:"host"`
	Port           string                 `yaml:"port"`
	User           string                 `yaml:"user"`
	Password       string                 `yaml:"password"`
	Name           string                 `yaml:"name"`
	SSLMode        string                 `yaml:"sslmode"`
	ConnectionPool DBConnectionPoolConfig `yaml:"connection_pool"`
}

type DBConnectionPoolConfig struct {
	MaxOpenConnection int `yaml:"max_open_connection"`
	MaxIdleConnection int `yaml:"max_idle_connection"`

	// this in second
	MaxLifeTime int `yaml:"max_lifetime"`
	MaxIdleTime int `yaml:"max_idletime"`
}

type AppConfig struct {
	Name    string `yaml:"name"`
	Port    string `yaml:"port"`
	Prefork bool   `yaml:"prefork"`
}

var cfg *Config = &Config{}

// load config file with extension `yaml`
func LoadConfigFromYaml(filename string) error {
	fileByte, err := os.ReadFile(filename)
	if err != nil {
		return err
	}

	return yaml.Unmarshal(fileByte, &cfg)
}

func GetConfig() *Config {
	return cfg
}
