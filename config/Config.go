package config

import (
	"github.com/naoina/toml"
	"log"
	"os"
)

var instance *Config

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

	Log struct {
		Level   string
		Fpath   string
		Msize   int
		Mage    int
		Mbackup int
	}
}

func LoadConfig(fPath string) *Config {
	conf := new(Config)
	if file, err := os.Open(fPath); err != nil {
		log.Println("does not exists file in ", fPath)
		panic(err)
	} else {
		defer file.Close()
		if err := toml.NewDecoder(file).Decode(conf); err != nil {
			log.Println("toml decode, fail")
			panic(err)
		}
		instance = conf
	}
	return instance
}
