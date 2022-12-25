package config

import (
	"github.com/codestates/WBABEProject-05/common/flag"
	"github.com/codestates/WBABEProject-05/config/db"
	"github.com/codestates/WBABEProject-05/config/info"
	"github.com/codestates/WBABEProject-05/config/log"
	server "github.com/codestates/WBABEProject-05/config/server"
)

var AppInfo *info.Info
var DBConfig *db.DBConfig
var LogConfig *log.Log
var AppServerConfig *server.ServerConfig

func LoadConfigs(pathMap map[string]*string) {
	ipath := pathMap[flag.InformationFlag.Name]
	AppInfo = info.NewInfo(*ipath)

	dpath := pathMap[flag.DatabaseFlag.Name]
	DBConfig = db.NewDbConfig(*dpath)

	lpath := pathMap[flag.LogConfigFlag.Name]
	LogConfig = log.NewLogConfig(*lpath)

	spath := pathMap[flag.ServerConfigFlag.Name]
	AppServerConfig = server.NewSeverConfig(*spath)

}
