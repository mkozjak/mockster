package main

import (
	"fmt"
	"log"
	"os"
	"runtime"

	"github.com/mkozjak/mockster/services"
	"github.com/spf13/pflag"
	"github.com/spf13/viper"
)

var (
	mockNats bool
	natsHost string
	natsPort int
)

type Env struct {
	// broker messaging.Broker
}

func init() {
	// bind all configuration options from cli, config and env
	pflag.BoolVar(&mockNats, "nats", false, "mock nats")
	pflag.StringVar(&natsHost, "nats-hostname", "localhost", "nats hostname")
	pflag.IntVar(&natsPort, "nats-port", 4222, "nats port")

	viper.BindEnv("nats.hostname", "NATS_HOSTNAME")
	viper.BindEnv("nats.port", "NATS_PORT")

	viper.BindPFlag("nats.hostname", pflag.Lookup("nats-hostname"))
	viper.BindPFlag("nats.port", pflag.Lookup("nats-port"))
	viper.BindPFlag("nats.enabled", pflag.Lookup("nats"))
}

func cliUsage() {
	fmt.Println("Usage: mockster [options]\n")
	fmt.Println("Options:")
	pflag.PrintDefaults()
}

func main() {
	// parse configuration
	pflag.Usage = cliUsage
	pflag.Parse()

	viper.AddConfigPath(".")
	viper.SetConfigName("config")

	if err := viper.ReadInConfig(); err != nil {
		log.Println("failed parsing config file:", err)
		os.Exit(1)
	}

	if err := services.RunAll(); err != nil {
		log.Println("failed running services:", err)
		os.Exit(1)
	}

	runtime.Goexit()
}
