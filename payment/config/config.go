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
	Lago  LagoConfig
	Kafka KafkaConfig `mapstructure:"kafka"`
}

type MongoConfig struct {
	Uri      string `mapstructure:"uri"`
	Database string `mapstructure:"database"`
}

type LagoConfig struct {
	Host           string         `mapstructure:"host"`
	Port           string         `mapstructure:"port"`
	ApiKey         string         `mapstructure:"api_key"`
	BillableMetric BillableMetric `mapstructure:"billable_metric"`
}
type BillableMetric struct {
	SSAIInsertAdsCode string `mapstructure:"ssai_insert_ads_code"`
}
type KafkaConfig struct {
	Urls    []string `yaml:"urls"  mapstructure:"urls"`
	GroupId string   `mapstructure:"group_id"`
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
