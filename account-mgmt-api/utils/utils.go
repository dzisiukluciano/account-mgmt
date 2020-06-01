package utils

import (
	"fmt"
	"os"

	uuid "github.com/satori/go.uuid"
)

// GetEnv returns an environment variable or the given default value if it isn't defined
// All environment variables are returned as string as must be parsed to the appropiate
// type
func GetEnv(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}

	return fallback
}

// GetUUID generates a 16 bytes uuid as string
func GetUUID() string {
	uuid := uuid.NewV4()
	return fmt.Sprintf("id_%v", uuid)
}
