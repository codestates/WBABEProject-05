package config

import (
	"github.com/naoina/toml"
	"log"
	"os"
)

type Config struct {
	Server struct {
		Mode string
		Port string
	}

	Db struct {
		MongoUri string
		User     string
		Pwd      string
	}
}

func NewConfig(fPath string) *Config {
	cfg := new(Config)
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
