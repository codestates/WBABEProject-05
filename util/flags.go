package util

import (
	"flag"
	"log"
)

var ConfPath string

func GetConfPath() string {
	ConfPath := flag.String("confPath", "./config/config.toml", "toml file to use for configuration")
	flag.Parse()
	log.Println("toml file load ::", ConfPath)
	return *ConfPath
}
