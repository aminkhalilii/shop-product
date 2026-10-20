package config

import (
	"fmt"
	"os"
	"strconv"
	"time"
)

func getString(key string) string {
	return os.Getenv(key)
}

func getInt(key string) int {
	value := os.Getenv(key)

	if value == "" {
		return 0
	}

	i, err := strconv.Atoi(value)
	if err != nil {
		panic(fmt.Sprintf("%s must be an integer", key))
	}

	return i
}

func getBool(key string) bool {
	value := os.Getenv(key)

	if value == "" {
		return false
	}

	b, err := strconv.ParseBool(value)
	if err != nil {
		panic(fmt.Sprintf("%s must be a boolean", key))
	}

	return b
}

func getDuration(key string) time.Duration {
	value := os.Getenv(key)

	if value == "" {
		return 0
	}

	d, err := time.ParseDuration(value)
	if err != nil {
		panic(fmt.Sprintf("%s has invalid duration", key))
	}

	return d
}
