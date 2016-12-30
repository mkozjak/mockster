// Package services implements server mocks for testing.
package services

import (
	"log"
	// "github.com/mkozjak/mockster/services/nats"
)

func RunAll(list map[string]bool) error {
	log.Println(list)

	return nil
}
