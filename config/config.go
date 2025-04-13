package config

import "os"

type Config struct {
	Port     string
	Database string
	Env      string
	LogLevel string
}

func LoadConfig() *Config {
	return &Config{
		Port:     getEnv("PORT", "8080"),
		Database: getEnv("DATABASE_URL", ""),
		Env:      getEnv("ENV", "development"),
		LogLevel: getEnv("LOG_LEVEL", "INFO"),
	}
}

func getEnv(key, defaultValue string) string {
	if value, exist := os.LookupEnv(key); exist {
		return value
	}
	return defaultValue
}
