// Package services implements server mocks for testing.
package services

import (
	_ "log"

	"github.com/fatih/structs"
	"github.com/mkozjak/mockster/services/nats"
	"github.com/mkozjak/mockster/types"
)

// Env holds a services instance configuration.
type Env struct {
	cfgList    types.Services
	registered [1]string
}

// New returns an initiated Env that takes services' configuration.
func New(cfg types.Services) *Env {
	s := new(Env)

	s.registered[0] = "Nats"
	s.cfgList = cfg

	return s
}

// RunAll starts all defined mock services.
func (s *Env) RunAll() error {
	// kinda hacky
	for _, cfgSrv := range structs.Fields(s.cfgList) {
		for _, srvName := range s.registered {
			if srvName != cfgSrv.Name() {
				continue
			}

			// run service
			switch srvName {
			case "Nats":
				// TODO: handle errors
				go nats.Start(s.cfgList.Nats.Port, s.cfgList.Nats.Hostname)
			}

			break
		}
	}

	return nil
}
