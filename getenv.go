package getenv

import (
	"fmt"
	"os"
	"strconv"
	"time"
)

// GetIntEnv converts an environment variable to string
func GetStringEnv(key, defaultValue string) string {
	e := os.Getenv(key)
	if e == "" {
		return defaultValue
	}
	return e
}

// GetIntEnv converts an environment variable to int
func GetIntEnv(key string, defaultValue int) int {
	e := os.Getenv(key)
	if e == "" {
		return defaultValue
	}
	res, err := strconv.Atoi(e)
	if err != nil {
		panic(fmt.Sprintf("getenv: cannot convert %s to int", e))
	}
	return res
}

// GetDurationEnv converts an integer environment variable to a duration in seconds
func GetDurationEnv(key string, defaultValue time.Duration) time.Duration {
	e := os.Getenv(key)
	if e == "" {
		return defaultValue
	}
	res, err := strconv.Atoi(e)
	if err != nil {
		panic(fmt.Sprintf("getenv: cannot convert %vs to time.Duration", res))
	}
	return time.Duration(res) * time.Second
}

// GetIntEnv converts an environment variable to bool
func GetBoolEnv(key string, defaultValue bool) bool {
	e := os.Getenv(key)
	if e == "" {
		return defaultValue
	}
	res, err := strconv.ParseBool(e)
	if err != nil {
		panic(fmt.Sprintf("getenv: cannot convert %v to bool", res))
	}
	return res
}
