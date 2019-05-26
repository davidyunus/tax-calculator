package config

import "os"

// Config ...
type Config struct {
	DBConnectionString string
}

const (
	dbConnectionString = "DB_CONNECTION_STRING"
)

var config *Config

func getEnvOrDefault(env string, defaultVal string) string {
	e := os.Getenv(env)
	if e == "" {
		return defaultVal
	}
	return e
}

// GetConfiguration ...
func GetConfiguration() (*Config, error) {
	if config != nil {
		return config, nil
	}

	config := &Config{
		DBConnectionString: getEnvOrDefault(dbConnectionString, "postgres://postgres@localhost:5432/user?sslmode=disable"),
	}

	return config, nil
}
