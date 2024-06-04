package common

import "syscall"

func EnvString(key, defaultValue string) string {
	if value, ok := syscall.Getenv(key); ok {
		return value
	}
	return defaultValue
}
