package utils

import (
	"os"
)

// GetEnv retrieves the value of the environment variable specified by `key`,
//
// If the environment variable is set, the function returns its value.
// If the variable is not set, it returns the provided `fallback` value instead.
//
// Parameters:
//
//	-key: The name of the environment variable to look up.
//	-fallback: The value to return if the environment variable is not set.
//
// Returns:
//
//	The value of the environment variable if it exists, or the `fallback` value if not.
func GetEnv(key, fallback string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}

	return fallback
}
