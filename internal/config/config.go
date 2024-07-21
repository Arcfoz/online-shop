package config

import (
	"os"

	"gopkg.in/yaml.v3"
)

type Config struct {
	App AppConfig `yaml:"app"`
	DB  DBconfig  `yaml:"db"`
}

type AppConfig struct {
	Name      string           `yaml:"name"`
	Port      string           `yaml:"port"`
	Encrytion EncryptionConfig `yaml:"encrytion"`
}

type EncryptionConfig struct {
	Salt uint8 `yaml:"salt"`
	JWTSecret string `yaml:"jwt_secret"`
}

type DBconfig struct {
	Host           string                 `yaml:"host"`
	Port           string                 `yaml:"port"`
	User           string                 `yaml:"user"`
	Password       string                 `yaml:"password"`
	Name           string                 `yaml:"name"`
	ConnectionPool DBConnectionPoolConfig `yaml:"connection_pool"`
}

type DBConnectionPoolConfig struct {
	MaxIdleConnection     uint8 `yaml:"max_idle_connection"`
	MaxOpenConnection     uint8 `yaml:"max_open_connection"`
	MaxLifetimeConnection uint8 `yaml:"max_lifetime_connection"`
	MaxIdletimeConnection uint8 `yaml:"max_idletime_connection"`
}

var Cfg Config

func LoadConfig(filename string) (err error) {
	configByte, err := os.ReadFile(filename)
	if err != nil {
		return
	}

	err = yaml.Unmarshal(configByte, &Cfg)
	return
}
