package configs

import (
	"os"
	"strconv"
	"github.com/gofiber/fiber/v2/middleware/encryptcookie"
)

type Config struct {
	HTTPServerConfig *HTTPServerConfig
	PostgreSQLConfig *PostgreSQLConfig
	MigrationVersion int
}

type HTTPServerConfig struct {
	Endpoint      string
	CookieSecret  string
	AccessTokenCookieName string
	AllowedOrigin string
}

type PostgreSQLConfig struct {
	Uri string
}

var config *Config

func GetConfig() *Config {
	if config != nil {
		return config
	}
	config = &Config{
		HTTPServerConfig: &HTTPServerConfig{
			Endpoint:      getEnv("HTTP_ENDPOINT", ":3001"),
			CookieSecret:  getEnv("COOKIE_SECRET", encryptcookie.GenerateKey()),
			AllowedOrigin: getEnv("ALLOWED_ORIGIN", "http://localhost:3000"),
			AccessTokenCookieName: getEnv("ACCESS_TOKEN_COOKIE_NAME", "access_token"),
		},
		PostgreSQLConfig: &PostgreSQLConfig{
			Uri: getEnv("POSTGRES_URI", "postgresql://admin:admin@postgres:5432/dungeons-of-yapp"),
		},
		MigrationVersion: getEnvInt("MIGRATION_VERSION", 1),
	}
	return config
}

func getEnv(key, fallback string) string {
	value := os.Getenv(key)
	if len(value) == 0 {
		return fallback
	}
	return value
}

func getEnvInt(key string, fallback int) int {
	value := os.Getenv(key)
	if int_value, err := strconv.Atoi(value); err != nil {
		return fallback
	} else {
		return int_value
	}
}