package config

import (
	"context"
	"os"
)

var (
	CTX     = context.Background()
	PodName = getEnv("POD_NAME", "undefined")
	Version = getEnv("VERSION", "undefined")

	dbUser     = getEnv("POSTGRES_USER", "moffy")
	dbPassword = getEnv("POSTGRES_PASSWORD", "moffy0223")
	dbIp       = getEnv("POSTGRES_IP", "localhost")
	dbPort     = getEnv("POSTGRES_PORT", "5432")
	dbName     = getEnv("POSTGRES_DB_NAME", "chatuser")
	DbSource   = "postgresql://" + dbUser + ":" + dbPassword + "@" + dbIp + ":" + dbPort + "/" + dbName + "?sslmode=disable"

	FrontAppSource = getEnv("FRONT_APP_SOURCE", "http://localhost:3000")
)

func getEnv(key, defaultValue string) string {
	value := os.Getenv(key)
	if value == "" {
		return defaultValue
	}
	return value
}
