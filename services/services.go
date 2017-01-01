// Package services implements server mocks for testing.
package services

import (
	"log"

	"github.com/fatih/structs"
	"github.com/mkozjak/mockster/services/nats"
	"github.com/mkozjak/mockster/types"
)

type Env struct {
	cfgList    types.Services
	registered [1]string
}

func New(cfg types.Services) (*Env, error) {
	s := new(Env)

	s.registered[0] = "Nats"
	s.cfgList = cfg

	return s, nil
}

func (s *Env) RunAll() error {
	for i, cfgSrv := range structs.Fields(s.cfgList) {
		for _, srvName := range s.registered {
			if srvName == cfgSrv.Name() {
				// run service
				switch srvName {
				case "Nats":
					nats.Start()
				}

				break
			}
		}
	}

	return nil
}
