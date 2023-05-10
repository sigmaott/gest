package config

type HostPort struct {
	Host     string `yaml:"host"  mapstructure:"host"`
	Port     int    `yaml:"port"  mapstructure:"port"`
	BasePath string `yaml:"base_path"  mapstructure:"base_path"`
}
