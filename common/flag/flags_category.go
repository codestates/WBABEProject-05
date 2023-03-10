package flag

import "flag"

type FlagCategory struct {
	Name    string
	Default string
	Usage   string
}

var (
	ServerConfigFlag = &FlagCategory{
		Name:    "server",
		Default: "./config/server/config.toml",
		Usage:   "toml file to use for server configuration",
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

	DatabaseFlag = &FlagCategory{
		Name:    "db",
		Default: "./config/db/config.toml",
		Usage:   "toml file to use for database configuration",
	}
)

func (f *FlagCategory) Load() *string {
	return flag.String(f.Name, f.Default, f.Usage)
}
