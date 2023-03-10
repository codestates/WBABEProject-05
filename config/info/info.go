package info

import (
	"github.com/naoina/toml"
	"log"
	"os"
)

var AppInfo *Info

type Info struct {
	Version     string   `json:"version"`
	Name        string   `json:"name"`
	Description string   `json:"description"`
	Author      string   `json:"author"`
	Spec        []string `json:"spec"`
	Blog        string   `json:"blog"`
}

func NewInfo(fPath string) *Info {
	info := new(Info)
	file, err := os.Open(fPath)
	defer file.Close()
	if err != nil {
		log.Println("start app... does not exists config file in ", fPath)
		panic(err)
	}

	if err := toml.NewDecoder(file).Decode(info); err != nil {
		log.Println("start app... toml decode, fail")
		panic(err)
	}
	return info
}
