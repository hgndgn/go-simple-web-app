package config

import (
	"os"
)

func Get(key string, defaultValue string) string {
	value := os.Getenv(key)

	if value == "" {
		value = defaultValue
	}

	return value
}
