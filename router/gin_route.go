package router

import (
	"github.com/codestates/WBABEProject-05/contorller"
	"github.com/codestates/WBABEProject-05/logger"
	"github.com/gin-gonic/gin"
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

func LoadGin(mode string, ctl map[string]contorller.Controller) *GinRoute {
	if instance != nil {
		return instance
	}
	setMode(mode)
	instance = &GinRoute{
		engin:      loadEngine(),
		controller: ctl,
	}
	return instance
}

func setMode(mode string) {
	switch mode {
	case gin.DebugMode:
		gin.SetMode(gin.DebugMode)
	case gin.ReleaseMode:
		gin.SetMode(gin.ReleaseMode)
	case gin.TestMode:
		gin.SetMode(gin.TestMode)
	default:
		gin.SetMode(gin.DebugMode)
	}
}

func loadEngine() *gin.Engine {
	grt := gin.Default()
	grt.Use(logger.GinLogger())
	grt.Use(logger.GinRecovery(true))
	grt.GET("/", func(g *gin.Context) {
		g.JSON(200, "ok")
	})
	//g.Use(CORS())
	return grt
}
