package boot

import (
	"os"

	"github.com/joho/godotenv"
)

func loadEnv() bool {
	godotenv.Load(".env")
	return true

}
func envCheck() bool {
	s := true
	env := []string{
		"ENVIRONMENT",
		"IP_ADDR",
		"PORT",
		"APP_DEBUG",
		"APP_DB_DEBUG",
		"MYSQL_DB_ADDR",
		"MYSQL_DB_PORT",
		"MYSQL_DB_USERNAME",
		"MYSQL_DB_NAME",
		"JWT_SECRET_KEY",
	}

	for _, e := range env {
		if os.Getenv(e) == "" {
			s = false
		}
	}
	return s
}
