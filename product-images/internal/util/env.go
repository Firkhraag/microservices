package util

import "os"

// Env returns environment variable if set or a default value
func Env(varName string, defaultValue string) string {
	varValue := os.Getenv(varName)
	if varValue == "" {
		return defaultValue
	}
	return varValue
}
