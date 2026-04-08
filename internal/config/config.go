package config

import (
	"os"
	"strconv"
	"time"

	"github.com/joho/godotenv"
)

type Config struct {
	Port string

	DBHost string
	DBPort string
	DBUser string
	DBPass string
	DBName string

	RedisAddr string

	RedisTimeout time.Duration
	DBTimeout    time.Duration
}

func loadEnv() {
	env := os.Getenv("APP_ENV")
	if env == "" {
		env = "local"
	}
	// Try .env.<env> first, fall back to .env
	if err := godotenv.Load(".env." + env); err != nil {
		_ = godotenv.Load()
	}
}

func getEnv(key, defaultVal string) string {
	val := os.Getenv(key)
	if val == "" {
		return defaultVal
	}
	return val
}

func getEnvAsInt(key string, defaultVal int) int {
	valStr := os.Getenv(key)
	if val, err := strconv.Atoi(valStr); err == nil {
		return val
	}
	return defaultVal
}

func LoadConfig() *Config {
	loadEnv()

	redisTimeout := getEnvAsInt("REDIS_TIMEOUT_MS", 50)
	dbTimeout := getEnvAsInt("DB_TIMEOUT_MS", 200)

	return &Config{
		Port: getEnv("PORT", "8080"),

		DBHost: getEnv("DB_HOST", "localhost"),
		DBPort: getEnv("DB_PORT", "3306"),
		DBUser: getEnv("DB_USER", "root"),
		DBPass: getEnv("DB_PASSWORD", "root"),
		DBName: getEnv("DB_NAME", "url_shortener"),

		RedisAddr: getEnv("REDIS_ADDR", "localhost:6379"),

		RedisTimeout: time.Duration(redisTimeout) * time.Millisecond,
		DBTimeout:    time.Duration(dbTimeout) * time.Millisecond,
	}
}