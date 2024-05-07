package config

import (
	"DB_Project/internal/constants"
	"os"
)

type Config struct {
	AppConfig AppConfig
	DBConfig  DBConfig
}

type AppConfig struct {
	ApiPort string
}

type DBConfig struct {
	Host string
	Port string
	User string
	Pass string
	Name string
}

func mustGetEnv(key string) string {
	v, present := os.LookupEnv(key)
	if !present {
		panic("Missing required environment variable: " + key)
	}

	return v
}

func NewConfig() *Config {
	return &Config{
		AppConfig: AppConfig{
			ApiPort: mustGetEnv(constants.API_PORT),
		},
		DBConfig: DBConfig{
			Host: mustGetEnv("DB_HOST"),
			Port: mustGetEnv("DB_PORT"),
			User: mustGetEnv("DB_USER"),
			Pass: mustGetEnv("DB_PASSWORD"),
			Name: mustGetEnv("DB_NAME"),
		},
	}
}

func (c *Config) GetApiPort() string {
	return c.AppConfig.ApiPort
}

func (c *Config) GetDBHost() string {
	return c.DBConfig.Host
}

func (c *Config) GetDBPort() string {
	return c.DBConfig.Port
}

func (c *Config) GetDBUser() string {
	return c.DBConfig.User
}

func (c *Config) GetDBPassword() string {
	return c.DBConfig.Pass
}

func (c *Config) GetDBName() string {
	return c.DBConfig.Name
}
