package util

import (
	"flag"
)

const (
	DefaultConfigPath = "./config/config.toml"
	Usage             = "toml file to use for configuration"
	ConfPathFlagName  = "confP"
)

var ConfigPath string

func LoadConfigFilePath() string {
	cfgP := flag.String(ConfPathFlagName, DefaultConfigPath, Usage)
	flag.Parse()
	ConfigPath = *cfgP
	return ConfigPath
}
