package info

import (
	"fmt"
	"github.com/naoina/toml"
	"log"
	"os"
)

type Info struct {
	Version     string
	Name        string
	Description string
	Author      string
}

func NewInfo(fPath string) *Info {
	fmt.Println(fPath)
	info := new(Info)
	if file, err := os.Open(fPath); err != nil {
		log.Println("start app... does not exists config file in ", fPath)
		panic(err)
	} else {
		defer file.Close()
		if err := toml.NewDecoder(file).Decode(info); err != nil {
			log.Println("start app... toml decode, fail")
			panic(err)
		}
		return info
	}
}
