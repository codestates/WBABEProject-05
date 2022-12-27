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
	//Acronym의 경우에는 Uppercase를 유지하는 것이 권장되고 있습니다.
	MongoUri   string
	DbName     string
	User       string
	Pwd        string
	BackupPath string
}

func NewDbConfig(fPath string) *DBConfig {
	dbcfg := new(DBConfig)
	if file, err := os.Open(fPath); err != nil {
		log.Println("start app... does not exists config file in ", fPath)
		panic(err)
	} else {
		defer file.Close()
		if err := toml.NewDecoder(file).Decode(dbcfg); err != nil {
			log.Println("start app... toml decode, fail")
			panic(err)
		}
		return dbcfg
	}
}

func WriteBackup(fPath string, T any) error {
	if data, err := json.MarshalIndent(T, "", "    "); err != nil {
		return err
	} else {
		fileP := fmt.Sprintf(fPath + time.Now().Format("2006-01-02") + ".txt")
		f, err := os.OpenFile(
			fileP,
			os.O_CREATE|os.O_RDWR|os.O_APPEND,
			os.FileMode(0644))
		defer f.Close()
		w := bufio.NewWriter(f)

		_, err = fmt.Fprint(w, string(data)+"\n")
		if err != nil {
			return err
		}

		err = w.Flush()
		if err != nil {
			return err
		}
		return nil
	}
}
