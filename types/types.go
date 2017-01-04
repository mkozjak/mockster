// Package types implements custom mockster types.
package types

type Config struct {
	Services
	Api
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

type Api struct {
	Hostname string
	Port     int
}
