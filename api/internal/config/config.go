package config

import (
	"io/ioutil"

	"github.com/pelletier/go-toml"
)

type Config struct {
	Api struct {
		Addr         string
		AllowOrigins []string
	}
	Db struct {
		Addr       string
		Pw         string
		SelectedDb int
	}
}

func Load(filename string) (Config, error) {
	configData, err := ioutil.ReadFile(filename)
	if err != nil {
		return Config{}, err
	}

	var config Config
	if err = toml.Unmarshal(configData, &config); err != nil {
		panic(err)
	}

	return config, nil
}
