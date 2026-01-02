package env

import (
	"fmt"
	"os"
	"strconv"
	"time"
)

// Helper function to be able to parse numbers from the environment
// On any failures returns the default value that you selected
func ParseNum(env string, defaultValue int, min int, max int) int {
	str := os.Getenv(env)
	if str == "" {
		return defaultValue
	}

	num, err := strconv.Atoi(str)
	if err != nil {
		fmt.Printf("Error occured on int conversion: %v\n", err)
		return defaultValue
	}

	if num > max || num < min {
		fmt.Println("Number out of range")
		return defaultValue
	}

	return num
}

// Helper function to be able to parse strings from the environment
// On any failures returns the default value selected
func ParseString(env string, defaultValue string) string {
	str := os.Getenv(env)
	if str == "" {
		return defaultValue
	}
	return str
}

// Helper function to be able to parse durations from the environment
// On any failures returns the default value selected
func ParseDuration(env string, defaultValue time.Duration) time.Duration {
	str := os.Getenv(env)
	if str == "" {
		return defaultValue
	}

	dur, err := time.ParseDuration(str)
	if err != nil {
		return defaultValue
	}
	return dur
}
