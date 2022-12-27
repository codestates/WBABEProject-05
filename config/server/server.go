package config

import (
	"github.com/naoina/toml"
	"log"
	"os"
)

type ServerConfig struct {
	Mode string
	Port string
}

func NewSeverConfig(fPath string) *ServerConfig {
	cfg := new(ServerConfig)
	file, err := os.Open(fPath)
	defer file.Close()
	if err != nil {
		log.Println("start app... does not exists config file in ", fPath)
		panic(err)
	}

	if err := toml.NewDecoder(file).Decode(cfg); err != nil {
		log.Println("start app... toml decode, fail")
		panic(err)
	}
	return cfg
}
