package config

import (
	"github.com/BurntSushi/toml"
)

type ConfigurationLoader struct {
}

type Config struct {
	Services []Service
}

type Service struct {
	Url        string
	Ip         string
	Host       string
	Method     string
	Type       string
	QueryParam string
}

func New() *ConfigurationLoader {
	return &ConfigurationLoader{}
}

func (c *ConfigurationLoader) Load() (config Config, err error) {
	var conf Config
	if _, err := toml.DecodeFile("config.toml", &conf); err != nil {
		return conf, err
	}
	return conf, nil
}
