package config

import (
	"io/ioutil"

	"github.com/pelletier/go-toml/v2"
)

type ServerConfig struct {
	Addr         string
	AllowOrigins []string
}

type DbConfig struct {
	Addr       string
	Pw         string
	SelectedDb int
}

type Config struct {
	Server ServerConfig
	Db     DbConfig
}

func LoadConfigFromFile(filename string) *Config {
	configData, err := ioutil.ReadFile(filename)
	if err != nil {
		panic(err)
	}

	var config Config
	if err = toml.Unmarshal(configData, &config); err != nil {
		panic(err)
	}

	return &config
}
