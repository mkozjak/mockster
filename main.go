package main

import (
	"fmt"
	"log"
	"os"
	"runtime"

	srv "github.com/mkozjak/mockster/services"
	"github.com/mkozjak/mockster/types"
	"github.com/spf13/pflag"
	"github.com/spf13/viper"
)

var (
	mockNats bool
	natsHost string
	natsPort int
)

type Env struct {
	conf *viper.Viper
}

func init() {
	// bind all configuration options from cli, config and env
	pflag.BoolVar(&mockNats, "nats", false, "mock nats")
	pflag.StringVar(&natsHost, "nats-hostname", "localhost", "nats hostname")
	pflag.IntVar(&natsPort, "nats-port", 4222, "nats port")

	viper.BindEnv("services.nats.hostname", "NATS_HOSTNAME")
	viper.BindEnv("services.nats.port", "NATS_PORT")

	viper.BindPFlag("services.nats.hostname", pflag.Lookup("nats-hostname"))
	viper.BindPFlag("services.nats.port", pflag.Lookup("nats-port"))
	viper.BindPFlag("services.nats.enabled", pflag.Lookup("nats"))
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

	var cfg types.Config

	if err := viper.Unmarshal(&cfg); err != nil {
		log.Println("unable to decode config:", err)
		os.Exit(1)
	}

	// set up and run mock services
	s, err := srv.New(cfg.Services)

	if err != nil {
		log.Println("unable to decode config:", err)
		os.Exit(1)
	}

	if err := s.RunAll(); err != nil {
		log.Println("failed running services:", err)
		os.Exit(1)
	}

	runtime.Goexit()
}
