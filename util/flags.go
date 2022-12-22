package util

import (
	"flag"
)

const (
	DefaultConfPath  = "./config/config.toml"
	Usage            = "toml file to use for configuration"
	ConfPathFlagName = "confP"
)

func LoadConfPath() string {
	conP := flag.String(ConfPathFlagName, DefaultConfPath, Usage)
	flag.Parse()
	return *conP
}
