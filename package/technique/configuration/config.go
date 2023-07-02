package configuration

import "github.com/spf13/viper"

func LoadConfigYaml[T any](path string) (*T, error) {
	config := new(T)
	configParser := viper.New()
	configParser.SetConfigType("yaml")
	configParser.SetConfigFile(path)
	err := configParser.ReadInConfig()
	if err != nil {
		return nil, err
	}
	err = configParser.Unmarshal(config)
	if err != nil {
		return nil, err
	}
	return config, nil
}
