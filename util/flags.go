package util

import "flag"

type FlagCategory struct {
	Name    string
	Default string
	Usage   string
}

var (
	ConfigFlag = &FlagCategory{
		Name:    "config",
		Default: "./config/config.toml",
		Usage:   "toml file to use for configuration",
	}
)

func (f *FlagCategory) Load() *string {
	return flag.String(f.Name, f.Default, f.Usage)
}
