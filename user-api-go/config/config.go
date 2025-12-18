package config

import "os"

type Config struct {
	DBURL string
	Port  string
}

func Load() Config {
	return Config{
		DBURL: getEnv("DB_URL", "postgres://user:pass@localhost:5432/usersdb?sslmode=disable"),
		Port:  getEnv("PORT", "8080"),
	}
}

func getEnv(key, fallback string) string {
	if v, ok := os.LookupEnv(key); ok {
		return v
	}
	return fallback
}
