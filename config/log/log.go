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
	cfg := new(Log)
	file, err := os.Open(fPath)
	defer file.Close()
	if err != nil {
		log.Println("start app... does not exists config file in ", fPath)
		panic(err)
	}

	if err := toml.NewDecoder(file).Decode(cfg); err != nil {
		log.Println("start app... toml decode, fail")
		panic(err)
	}
	return cfg
}
