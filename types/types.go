// Package types implements custom mockster types.
package types

// Config consists of multiple keys defined throughout the configuration.
type Config struct {
	Services
	Api
}

// Services consists of currently supported service types (mocking servers).
type Services struct {
	Nats Service
	Amqp Service
}

// Service represents services-based configuration.
type Service struct {
	Hostname string
	Port     int
	Enabled  bool
}

// Api represents a rest api configuration.
type Api struct {
	Hostname string
	Port     int
}
