package router

import (
	"github.com/codestates/WBABEProject-05/contorller"
	"github.com/codestates/WBABEProject-05/logger"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"net/http"
)

var instance *GinRoute

type GinRoute struct {
	engin      *gin.Engine
	controller map[string]contorller.Controller
}

func (r *GinRoute) Handle() http.Handler {
	return r.engin
}

func GetGin(mode string, ctl map[string]contorller.Controller) *GinRoute {
	if instance != nil {
		return instance
	}
	setMode(mode)
	instance = &GinRoute{
		engin:      NewEngine(),
		controller: ctl,
	}
	return instance
}

func setMode(mode string) {
	switch mode {
	case "dev":
		zap.L().Info("start gin server set mod dev")
		gin.SetMode(gin.DebugMode)
	case "prod":
		zap.L().Info("start gin server set mod prod")
		gin.SetMode(gin.ReleaseMode)
	case "test":
		zap.L().Info("start gin server set mod test")
		gin.SetMode(gin.TestMode)
	default:
		zap.L().Info("start gin server set mod dev")
		gin.SetMode(gin.DebugMode)
	}
}

func NewEngine() *gin.Engine {
	grt := gin.Default()
	grt.Use(logger.GinLogger())
	grt.Use(logger.GinRecovery(true))
	grt.GET("/", func(g *gin.Context) {
		g.JSON(200, "ok")
	})
	//g.Use(CORS())
	return grt
}
