package config

import (
	"flag"
	"fmt"
	"github.com/gestgo/gest/package/common/config"
	"log"
	"os"
)

var configuration *Configuration

type Configuration struct {
	Http  config.HostPort
	Mongo MongoConfig
}
type MongoConfig struct {
	Uri      string `mapstructure:"uri"`
	Database string `mapstructure:"database"`
}

func init() {
	configPath := flag.String("c", "./payment/config/config.yaml", "")

	flag.Usage = func() {
		fmt.Println("Usage: server -e {mode}")
		os.Exit(1)
	}
	flag.Parse()

	var err error
	configuration, err = config.LoadConfigYaml[Configuration](*configPath)
	if err != nil {
		log.Fatal(err)
	}
}
func GetConfiguration() *Configuration {
	return configuration
}
