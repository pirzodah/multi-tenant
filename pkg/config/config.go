package config

import (
	"fmt"
	"os"
)

type Config struct {
	DBUrl string
}

func LoadConfig() *Config {
	host := getEnv("DB_HOST", "localhost")
	port := getEnv("DB_PORT", "5432")
	user := getEnv("DB_USER", "marketplace")
	pass := getEnv("DB_PASSWORD", "marketplace")
	dbname := getEnv("DB_NAME", "marketplace")

	return &Config{
		DBUrl: fmt.Sprintf("postgres://%s:%s@%s:%s/%s",
			user, pass, host, port, dbname),
	}
}

func getEnv(key, fallback string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return fallback
}
