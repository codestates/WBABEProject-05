package main

import (
	"github.com/codestates/WBABEProject-05/contorller"
	"github.com/codestates/WBABEProject-05/logger"
	"github.com/codestates/WBABEProject-05/router"
	"github.com/codestates/WBABEProject-05/util"
)

var (
	App   = util.NewApp()
	flags = []*util.FlagCategory{
		util.ConfigFlag,
		util.LogConfigFlag,
	}

	controllers = map[string]contorller.Controller{}
)

func init() {
	// read flags
	App.ReadFlags(flags)
	App.LoadConfig()

	// setting logger
	lcfg := App.GetLogConfig()
	zapLogger := logger.InitLogger(lcfg)
	App.SetLogger(zapLogger)

	// setting http
	gin := router.GetGin(App.Config.Server.Mode, controllers)
	App.SetRouter(gin)
}

func main() {
	App.Run()
}
