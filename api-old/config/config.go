package api_config

import (
	"io/ioutil"
	"os"
	"strings"

	"github.com/pelletier/go-toml/v2"
)

// The config file contains configurations for both frontend and backend.
// Therefore it's not reasonable to specify all the configs here
var (
	Api ApiConfig
	Db  DbConfig
)

type ApiConfig struct {
	Addr         string
	AllowOrigins []string
}

type DbConfig struct {
	Addr       string
	Pw         string
	SelectedDb int
}

type Config struct {
	Api ApiConfig
	Db  DbConfig
}

func init() {
	// When running "go test" the args might look something like this:
	// /var/folders/{path-to-build-folder}/api.test -test.paniconexit0 -test.timeout=10m0s
	if strings.HasSuffix(os.Args[0], ".test") {
		// assume we are running unit tests
		// no configurations init
		return
	}

	configData, err := ioutil.ReadFile("./config/config.toml")
	if err != nil {
		panic(err)
	}

	var config Config
	if err = toml.Unmarshal(configData, &config); err != nil {
		panic(err)
	}

	Api = config.Api
	Db = config.Db
}
