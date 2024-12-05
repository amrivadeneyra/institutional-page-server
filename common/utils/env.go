package utils

import (
	"os"
	"strconv"

	"github.com/pkg/errors"
)

// Get the environment variable.
// If panicEmpty=true, it will panic if the environment variable is not presented or empty.
func GetEnvVar(key string, panicEmpty bool) string {
	value := os.Getenv(key)
	if panicEmpty && value == "" {
		panic(errors.Errorf("Missing %v environment variable", key))
	}
	return value
}

func GetEnvVarInt(key string, panicEmpty bool) int {
	value := os.Getenv(key)
	if panicEmpty && value == "" {
		panic(errors.Errorf("Missing %v environment variable", key))
	}

	if value == "" {
		value = "0"
	}

	valueInt, err := strconv.Atoi(value)
	if err != nil {
		panic(errors.Errorf("Missing %v environment variable", key))
	}
	return valueInt
}

func GetEnvVarBool(key string, panicEmpty bool) bool {
	value := os.Getenv(key)
	if panicEmpty && value == "" {
		panic(errors.Errorf("Missing %v environment variable", key))
	}

	if value == "" {
		value = "false"
	}

	valueBool, err := strconv.ParseBool(value)
	if err != nil {
		panic(errors.Errorf("Missing %v environment variable", key))
	}
	return valueBool
}
