package flag

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

	LogConfigFlag = &FlagCategory{
		Name:    "log",
		Default: "./config/log/config.toml",
		Usage:   "toml file to use for log configuration",
	}

	InformationFlag = &FlagCategory{
		Name:    "info",
		Default: "./config/info/config.toml",
		Usage:   "toml file to use for information configuration",
	}
)

func (f *FlagCategory) Load() *string {
	return flag.String(f.Name, f.Default, f.Usage)
}
