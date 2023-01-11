package main

import (
	app "github.com/codestates/WBABEProject-05/common/app"
	"github.com/codestates/WBABEProject-05/common/enum"
	"github.com/codestates/WBABEProject-05/common/flag"
	"github.com/codestates/WBABEProject-05/config"
	gin "github.com/codestates/WBABEProject-05/controller"
	zap "github.com/codestates/WBABEProject-05/logger"
	mongo "github.com/codestates/WBABEProject-05/model"
	"github.com/codestates/WBABEProject-05/router"
	"github.com/codestates/WBABEProject-05/service"
)

var (
	App   = app.NewApp()
	flags = []*flag.FlagCategory{
		flag.ServerConfigFlag,
		flag.LogConfigFlag,
		flag.InformationFlag,
		flag.DatabaseFlag,
	}

	mongoCollectionNames = []string{
		enum.ReceiptCollectionName,
		enum.ReviewCollectionName,
		enum.StoreCollectionName,
		enum.UserCollectionName,
		enum.MenuCollectionName,
	}
)

func init() {
	// read flags
	flag.FlagsLoad(flags)
	config.LoadConfigs(flag.Flags)

	//logger
	zap.LoadZapLogger(config.LogConfig)
	zap.SetAppLog(zap.ZapLog)

	// model
	err := mongo.LoadMongoModel(config.DBConfig.URI)
	if err != nil {
		panic(err)
		return
	}
	mongo.SetModeler(mongo.MongoModel)
	mongo.LoadMongoCollections(mongoCollectionNames, config.DBConfig.DBName)
	mongo.CreateIndexesInModels()
	mongo.InjectModelsMongoDependency(mongo.MongoCollection)

	// service
	service.InjectServicesDependency()

	// controller
	gin.InjectControllerDependency()

	// router
	router.SetAppRoute(
		router.NewGinRoute(config.AppServerConfig.Mode),
	)

}

func main() {
	App.Run()
}
