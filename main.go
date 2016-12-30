package main

import (
	"fmt"
	"log"
	"os"
	"runtime"

	"github.com/spf13/pflag"
	"github.com/spf13/viper"
	"uniqcast.com/packetize/messaging"
)

var (
	mockNats bool
	natsHost string
	natsPort int
)

type Env struct {
	broker messaging.Broker
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
	fmt.Println("Usage: packetize [options]\n")
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
		log.Println("failed parsing config file with", err)
		os.Exit(1)
	}

	// connect to service bus
	if viper.GetBool("nats.enabled") == true {
		_, err := messaging.Run("nats://" +
			viper.GetString("nats.hostname") + ":" +
			viper.GetString("nats.port"))

		if err != nil {
			log.Println("failed connecting to service bus with", err)
			os.Exit(1)
		}
	}

	// env := &Env{m}

	runtime.Goexit()
}
