package main

import (
	"github.com/codestates/WBABEProject-05/contorller"
	"github.com/codestates/WBABEProject-05/logger"
	"github.com/codestates/WBABEProject-05/router"
	"github.com/codestates/WBABEProject-05/util"
)

var (
	app   = util.NewApp()
	flags = []*util.FlagCategory{
		util.ConfigFlag,
	}

	controllers = map[string]contorller.Controller{}
)

func init() {
	app.LoadFlags(flags)
	app.LoadConfig()
	logger.InitLogger(app.Config)
	gin := router.GetGin(app.Config.Server.Mode, controllers)
	app.SetRouter(gin)
}

func main() {
	app.Run()
}
