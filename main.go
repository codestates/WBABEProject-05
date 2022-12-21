package main

import (
	"fmt"
	"github.com/codestates/WBABEProject-05/config"
	"github.com/codestates/WBABEProject-05/util"
)

var (
	Config *config.Config
)

func init() {
	util.FlagsLoad()
	Config = config.NewConfig(util.ConfPath)
}

func main() {
	fmt.Println("Hello World")
}
