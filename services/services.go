// Package services implements server mocks for testing.
package services

import (
	"log"
	// "github.com/mkozjak/mockster/services/nats"
	"github.com/mkozjak/mockster/types"
)

type Env struct {
	list interface{}
}

func New(cfg types.Services) (*Env, error) {
	log.Println("OK", cfg.Nats)
	s := new(Env)

	// s.list = list

	return s, nil
}

func (s *Env) RunAll() error {
	return nil
}
