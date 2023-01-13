package config

import (
	"io/ioutil"

	"github.com/pelletier/go-toml/v2"
)

var (
	Title string
)

type Config struct {
	Title string
}

func init() {
	configData, err := ioutil.ReadFile("config.toml")
	if err != nil {
		panic(err)
	}

	var config Config
	if err = toml.Unmarshal(configData, &config); err != nil {
		panic(err)
	}

	Title = config.Title
}
