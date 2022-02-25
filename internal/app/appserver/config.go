package appserver

import (
	"flag"
	"github.com/BurntSushi/toml"
	"log"
)

type TomlConfig struct {
	BindAddr    string `toml:"bind_addr"`
	LogLevel    string `toml:"log_level"`
	DatabaseUrl string `toml:"database_url"`
}

var configPath string

func init() {
	flag.StringVar(&configPath, "config-path", "config/config.toml", "path to config file")
}

func NewConfig() *TomlConfig {
	config := &TomlConfig{}
	_, err := toml.DecodeFile(configPath, config)
	if err != nil {
		log.Fatal(err)
	}

	return config
}
