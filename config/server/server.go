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
	//아래 코드를 early-return 하도록 만들 수 있을까요?
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
