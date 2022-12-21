package util

import (
	"flag"
	"log"
)

var ConfPath string

func FlagsLoad() {
	conP := flag.String("confPath", "./config/config.toml", "toml file to use for configuration")
	flag.Parse()
	ConfPath = *conP
	log.Println("toml file load ::", ConfPath)
}
