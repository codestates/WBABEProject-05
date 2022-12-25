package log

import (
	"github.com/naoina/toml"
	log "log"
	"os"
)

var LogConfig *Log

type Log struct {
	Level     string
	FilePath  string
	MaxSize   int
	MaxAge    int
	MaxBackup int
}

func (l *Log) GetSettingValues() (path, level string, size, backup, age int) {
	path = l.FilePath
	level = l.Level
	size = l.MaxSize
	backup = l.MaxBackup
	age = l.MaxAge
	return
}

func NewLogConfig(fPath string) *Log {
	lcfg := new(Log)
	if file, err := os.Open(fPath); err != nil {
		log.Println("start app... does not exists config file in ", fPath)
		panic(err)
	} else {
		defer file.Close()
		if err := toml.NewDecoder(file).Decode(lcfg); err != nil {
			log.Println("start app... toml decode, fail")
			panic(err)
		}
		return lcfg
	}
}
