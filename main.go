package main

import (
	gin2 "github.com/codestates/WBABEProject-05/contorller"
	"github.com/codestates/WBABEProject-05/router"
	"github.com/codestates/WBABEProject-05/util"
)

var (
	App   = util.NewApp()
	flags = []*util.FlagCategory{
		util.ConfigFlag,
		util.LogConfigFlag,
		util.InformationFlag,
	}
)

func init() {
	// read flags
	util.FlagsLoad(flags)
	App.LoadConfig()

	// setting logger
	App.LoadLogger()

	//setting http
	gin := router.GetGin(App.Config.Server.Mode, gin2.GetInstance())
	App.SetRouter(gin)
}

// TODO 확인할 사항 생성, 싱글톤 등 struct 의 생성에도 방식이있는데 New 또는 Get 등 각각의 경우를 확실히 정하도록하자
func main() {
	//fmt.Println(util.GetAppInfo())
	App.Run()
}
