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
	if file, err := os.Open(fPath); err != nil {
		log.Println("start app... does not exists config file in ", fPath)
		panic(err)
	} else {
		defer file.Close()
		if err := toml.NewDecoder(file).Decode(cfg); err != nil {
			log.Println("start app... toml decode, fail")
			panic(err)
		}
		return cfg
	}
}
