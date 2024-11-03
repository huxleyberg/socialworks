package env

import (
	"os"
	"strconv"
)

func GetString(key, fallback string) string {
	val, ok := os.LookupEnv(key)
	if !ok {
		return fallback
	}
	return val
}

func GetInt(key string, fallback int) (int, error) {
	val, ok := os.LookupEnv(key)
	if !ok {
		return fallback, nil
	}
	conv, err := strconv.Atoi(val)
	if err != nil {
		return conv, err
	}
	return conv, nil
}
