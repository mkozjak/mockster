// Package messaging implements all service bus operations.
package messaging

import (
	"log"

	"github.com/nats-io/go-nats"
)

type Broker interface {
	Publish(string, []byte) error
}

type Store struct {
	nc *nats.Conn
}

// Run sets up a service bus connection.
// It also subscribes our app to configured topics.
// Returns a Broker instance and an error if any occured.
func Run(url string) (*Store, error) {
	nc, err := nats.Connect(url)

	if err != nil {
		log.Println("failed connecting to service bus with", err)
		return nil, err
	}

	subscription, err := nc.Subscribe("foo", MsgHandler)

	if err != nil {
		log.Println("failed subscribing to myque with", err, subscription.Subject)
		return nil, err
	}

	log.Println("connected to a service bus at", url)
	log.Println("listening for new requests...")
	nc.Flush()

	return &Store{nc}, nil
}

// MsgHandler handles messages received for a particular subject
func MsgHandler(msg *nats.Msg) {
}

func (s *Store) Publish(subject string, data []byte) error {
	err := s.nc.Publish(subject, data)

	if err != nil {
		return err
	}

	return nil
}
