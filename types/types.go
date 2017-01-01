// Package types implements custom mockster types.
package types

type Config struct {
	Services
}

type Services struct {
	Nats Service
	Amqp Service
}

type Service struct {
	Hostname string
	Port     int
	Enabled  bool
}
