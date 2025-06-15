package config

import (
	"To_Do/pkg/logger"
	"os"
	"strconv"
	"time"
)

type ServerConfig struct {
	ReadTimeout       time.Duration
	ReadHeaderTimeout time.Duration
	WriteTimeout      time.Duration
	IdleTimeout       time.Duration
}

type Config struct {
	Port          string
	Database      string
	Env           string
	LogLevel      string
	Server        ServerConfig
	RedisAddr     string
	RedisPassword string
	RedisDB       int
	RedisTTL      time.Duration
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
		RedisAddr:     getEnv("REDIS_ADDR", "localhost:6379"),
		RedisPassword: getEnv("REDIS_PASSWORD", ""),
		RedisDB:       getEnvAsInt("REDIS_DB", 0),
		RedisTTL:      getEnvAsDuration("REDIS_CACHE_TTL", 10*time.Minute),
	}
}

func getEnv(key, defaultValue string) string {
	if value, ok := os.LookupEnv(key); ok {
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

func getEnvAsInt(name string, defaultValue int) int {
	val := os.Getenv(name)
	if val == "" {
		return defaultValue
	}
	valCnt, err := strconv.Atoi(val)
	if err != nil {
		logger.Logger.Error("Ошибка выбора таблицы для Redis. Будет выбрано значение по умолчанию")
		return defaultValue
	}
	return valCnt
}
