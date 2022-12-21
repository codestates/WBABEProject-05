package main

import (
	"fmt"
	"github.com/codestates/WBABEProject-05/config"
	"github.com/codestates/WBABEProject-05/logger"
	"github.com/codestates/WBABEProject-05/util"
	"log"
)

var (
	Conf *config.Config
)

func init() {
	util.FlagsLoad()
	Conf = config.NewConfig(util.ConfPath)
	err := logger.InitLogger(Conf)
	if err != nil {
		log.Println("logger load, fail")
		panic(err)
	}
}

func main() {
	fmt.Println("Hello World")
}
