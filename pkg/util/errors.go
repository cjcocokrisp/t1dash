package util

import "fmt"

func NullDBConnection() error {
	return fmt.Errorf("DB connection is null")
}
