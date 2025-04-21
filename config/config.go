package config

import (
	"os"
	"time"
)

type ServerConfig struct {
	ReadTimeout       time.Duration
	ReadHeaderTimeout time.Duration
	WriteTimeout      time.Duration
	IdleTimeout       time.Duration
}

type Config struct {
	Port     string
	Database string
	Env      string
	LogLevel string
	Server   ServerConfig
}

func LoadConfig() *Config {
	return &Config{
		Port:     getEnv("PORT", "8080"),
		Database: getEnv("DATABASE_URL", ""),
		Env:      getEnv("ENV", "development"),
		LogLevel: getEnv("LOG_LEVEL", "INFO"),
		Server: ServerConfig{
			ReadTimeout:       getEnvAsDuration("SERVER_READ_TIMEOUT", 5*time.Second),
			ReadHeaderTimeout: getEnvAsDuration("SERVER_READ_HEADER_TIMEOUT", 2*time.Second),
			WriteTimeout:      getEnvAsDuration("SERVER_WRITE_TIMEOUT", 10*time.Second),
			IdleTimeout:       getEnvAsDuration("SERVER_IDLE_TIMEOUT", 120*time.Second),
		},
	}
}

func getEnv(key, defaultValue string) string {
	if value, exist := os.LookupEnv(key); exist {
		return value
	}
	return defaultValue
}
func getEnvAsDuration(name string, defaultValue time.Duration) time.Duration {
	valueCnf := os.Getenv(name)
	if valueCnf == "" {
		return defaultValue
	}
	value, err := time.ParseDuration(valueCnf)
	if err != nil {
		return defaultValue
	}
	return value
}
