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
	log.Println("path is ", fPath)
	conf := new(Config)
	if file, err := os.Open(fPath); err != nil {
		panic(err)
	} else {
		defer file.Close()
		log.Println("File is ", file)
		if err := toml.NewDecoder(file).Decode(conf); err != nil {
			panic(err)
		}
		log.Println("Loading database uri is :: ", conf)
		return conf
	}

}
