package main

import (
	"github.com/codestates/WBABEProject-05/config"
	"github.com/codestates/WBABEProject-05/contorller"
	"github.com/codestates/WBABEProject-05/logger"
	"github.com/codestates/WBABEProject-05/router"
	"github.com/codestates/WBABEProject-05/util"
)

var (
	app         = util.LoadApp()
	controllers = map[string]contorller.Controller{}
)

func init() {
	path := util.LoadConfigFilePath()
	cfg := config.LoadConfig(path)
	logger.InitLogger(cfg)
	app.SetConfig(cfg)
	rt := router.LoadGin("", nil)
	app.SetRouter(rt)
}

func main() {
	app.Run()
}
