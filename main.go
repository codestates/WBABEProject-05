package main

import (
	"fmt"
	"github.com/codestates/WBABEProject-05/config"
	"github.com/codestates/WBABEProject-05/logger"
	"github.com/codestates/WBABEProject-05/util"
)

var app = util.LoadApp()

func init() {
	path := util.LoadConfPath()
	config := config.LoadConfig(path)
	logger.InitLogger(config)

	app.SetConfig(config)

}

func main() {
	defer fmt.Println("실행되나")
	fmt.Println("Hello World")
}
