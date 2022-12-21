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
	path := util.GetConfPath()
	Config = config.NewConfig(path)
}

func main() {
	fmt.Println("Hello World")
}
