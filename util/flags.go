package util

import (
	"flag"
)

var ConfPath string

func FlagsLoad() {
	conP := flag.String("confP", "./config/config.toml", "toml file to use for configuration")
	flag.Parse()
	ConfPath = *conP
}
