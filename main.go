package main

import (
	app2 "github.com/codestates/WBABEProject-05/common/app"
	"github.com/codestates/WBABEProject-05/common/flag"
	config2 "github.com/codestates/WBABEProject-05/config"
	gin2 "github.com/codestates/WBABEProject-05/contorller"
	"github.com/codestates/WBABEProject-05/logger"
	"github.com/codestates/WBABEProject-05/model"
	"github.com/codestates/WBABEProject-05/router"
	"github.com/codestates/WBABEProject-05/service"
)

var (
	app   = app2.NewApp()
	flags = []*flag.FlagCategory{
		flag.ServerConfigFlag,
		flag.LogConfigFlag,
		flag.InformationFlag,
		flag.DatabaseFlag,
	}

	mongoCollectionNames = []string{
		model.ReceiptCollectionName,
		model.ReviewCollectionName,
		model.StoreCollectionName,
		model.UserCollectionName,
	}
)

func init() {
	// read flags
	flag.FlagsLoad(flags)
	config2.LoadConfigs(flag.Flags)

	//logger
	logger.LoadZapLogger(config2.LogConfig)
	logger.SetAppLog(logger.ZapLog)

	// model
	err := model.LoadMongoModel(config2.DBConfig.URI)
	if err != nil {
		panic(err)
		return
	}
	model.SetModeler(model.MongoModel)
	model.LoadMongoCollections(mongoCollectionNames, config2.DBConfig.DBName)
	model.CreateIndexesInModels()
	model.InjectModelsMongoDependency(model.MongoCollection)

	// service
	service.InjectServicesDependency()

	// controller
	gin2.InjectControllerDependency()

	// router
	router.SetAppRoute(
		router.NewGinRoute(config2.AppServerConfig.Mode),
	)

}

// TODO 확인할 사항 생성, 싱글톤 등 struct 의 생성에도 방식이있는데 New 또는 Get 등 각각의 경우를 확실히 정하도록하자
func main() {
	app.Run()
}
