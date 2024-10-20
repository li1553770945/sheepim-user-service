package config

import (
	"gopkg.in/yaml.v3"
	"os"
)

type ServerConfig struct {
	ServiceName   string `yaml:"service-name"`
	ListenAddress string `yaml:"listen-address"`
}

type TracingConfig struct {
	Endpoint string `yaml:"endpoint"`
}

type DatabaseConfig struct {
	Username string `yaml:"username"`
	Password string `yaml:"password"`
	Database string `yaml:"database"`
	Address  string `yaml:"address"`
	Port     string `yaml:"port"`
}

type Config struct {
	ServerConfig   ServerConfig   `yaml:"server"`
	TracingConfig  TracingConfig  `yaml:"tracing"`
	DatabaseConfig DatabaseConfig `yaml:"database"`
}

func InitConfig(path string) *Config {
	conf := &Config{}
	f, err := os.Open(path)
	if err != nil {
		panic(err)
	}
	err = yaml.NewDecoder(f).Decode(conf)
	if err != nil {
		panic(err)
	}

	return conf
}
