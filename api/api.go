// Package api provides a simple REST API to manage mock services
package api

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/mkozjak/mockster/types"
)

// Env holds an API instance configuration.
type Env struct {
	cfg types.Api
}

// InteractionParams represents 'interactions' api parameters.
type InteractionParams struct {
	Action string
}

// New returns an initiated Env that takes an api configuration.
func New(cfg types.Api) *Env {
	a := Env{cfg}

	return &a
}

// Run starts an api server and sets up route listeners.
func (a *Env) Run() error {
	http.HandleFunc("/interactions", a.interactions)

	err := http.ListenAndServe(a.cfg.Hostname+":"+strconv.Itoa(a.cfg.Port), nil)
	if err != nil {
		return err
	}

	return nil
}

// interactions is a handler implementation for the 'interactions' api.
func (a *Env) interactions(resp http.ResponseWriter, req *http.Request) {
	switch req.Method {
	case "POST":
		decoder := json.NewDecoder(req.Body)
		var t InteractionParams

		if err := decoder.Decode(&t); err != nil {
			log.Println(err)
		}

		defer req.Body.Close()
		log.Println(t.Action)
	default:
		defer req.Body.Close()
	}
}
