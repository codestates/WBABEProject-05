package db

import (
	"bufio"
	"encoding/json"
	"fmt"
	"github.com/naoina/toml"
	"log"
	"os"
	"time"
)

type DBConfig struct {
	MongoUri   string
	DbName     string
	User       string
	Pwd        string
	BackupPath string
}

func NewDbConfig(fPath string) *DBConfig {
	cfg := new(DBConfig)

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

func WriteBackup(fPath string, T any) error {
	data, err := json.MarshalIndent(T, "", "    ")
	if err != nil {
		return err
	}

	path := fPath + time.Now().Format("2006-01-02") + ".txt"
	file := fmt.Sprintf(path)
	f, err := os.OpenFile(
		file, os.O_CREATE|os.O_RDWR|os.O_APPEND, os.FileMode(0644),
	)
	defer f.Close()

	w := bufio.NewWriter(f)
	if _, err = fmt.Fprint(w, string(data)+"\n"); err != nil {
		return err
	}

	if err = w.Flush(); err != nil {
		return err
	}

	return nil
}
