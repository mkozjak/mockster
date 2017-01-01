// Package services implements server mocks for testing.
package services

import (
	_ "log"

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
	// kinda hacky
	for _, cfgSrv := range structs.Fields(s.cfgList) {
		for _, srvName := range s.registered {
			if srvName != cfgSrv.Name() {
				continue
			}

			// run service
			switch srvName {
			case "Nats":
				if err := nats.Start(
					s.cfgList.Nats.Port,
					s.cfgList.Nats.Hostname); err != nil {

					return err
				}
			}

			break
		}
	}

	return nil
}
