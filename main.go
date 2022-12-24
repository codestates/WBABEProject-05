package main

import (
	"github.com/codestates/WBABEProject-05/common/flag"
	gin2 "github.com/codestates/WBABEProject-05/contorller"
	"github.com/codestates/WBABEProject-05/contorller/info"
	"github.com/codestates/WBABEProject-05/contorller/store"
	"github.com/codestates/WBABEProject-05/contorller/user"
	"github.com/codestates/WBABEProject-05/model"
	store2 "github.com/codestates/WBABEProject-05/model/store"
	"github.com/codestates/WBABEProject-05/router"
	"github.com/codestates/WBABEProject-05/service"
)

var (
	app   = NewApp()
	flags = []*flag.FlagCategory{
		flag.ConfigFlag,
		flag.LogConfigFlag,
		flag.InformationFlag,
		flag.DatabaseFlag,
	}
)

func init() {
	// read flags
	flag.FlagsLoad(flags)
	model.LoadModel()
	app.LoadConfig()

	// setting logger
	app.LoadLogger()

	//setting http
	config := model.GetDbConfig()
	mod, _ := model.GetModel(config)
	storCol := mod.GetCollection("store", config.DbName)
	storeModel := store2.GetStoreModel(storCol)
	menuService := service.GetStoreMenuService(storeModel)
	ginControl := gin2.GetInstance(
		info.GetInfoControl(),
		user.GetUserControl(),
		store.GetStoreControl(
			menuService,
		),
	)
	gin := router.GetGin(app.Config.Server.Mode, ginControl)
	app.SetRouter(gin)
}

// TODO 확인할 사항 생성, 싱글톤 등 struct 의 생성에도 방식이있는데 New 또는 Get 등 각각의 경우를 확실히 정하도록하자
func main() {
	//fmt.Println(util.GetAppInfo())
	app.Run()
}
