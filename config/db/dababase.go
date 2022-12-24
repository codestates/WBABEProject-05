package db

import (
	"github.com/naoina/toml"
	"log"
	"os"
)

type Database struct {
	MongoUri string
	DbName   string
	User     string
	Pwd      string
}

func NewDbConfig(fPath string) *Database {
	dbcfg := new(Database)
	if file, err := os.Open(fPath); err != nil {
		log.Println("start app... does not exists config file in ", fPath)
		panic(err)
	} else {
		defer file.Close()
		if err := toml.NewDecoder(file).Decode(dbcfg); err != nil {
			log.Println("start app... toml decode, fail")
			panic(err)
		}
		return dbcfg
	}
}
