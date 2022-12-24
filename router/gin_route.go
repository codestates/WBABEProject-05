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
	controller contorller.Controller
}

func (r *GinRoute) Handle() http.Handler {
	gr := r.engin
	gr.GET("/", func(c *gin.Context) {
		c.JSON(200, "ok")
	})
	infCtl, _ := r.controller.InfoControl()
	//TODO error
	gr.GET("/info", infCtl.GetInformation)

	user := gr.Group("/users")
	{
		usrCtl, _ := r.controller.UserControl()
		user.POST("/join", usrCtl.PostUser)
	}

	store := gr.Group("/stores")
	{
		strCtl, _ := r.controller.StoreControl()
		store.POST("", strCtl.PostStore)
		store.POST("/menu", strCtl.PostMenu)
	}

	return r.engin
}

func GetGin(mode string, ctl contorller.Controller) *GinRoute {
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
		gin.SetMode(gin.DebugMode)
	case "prod":
		gin.SetMode(gin.ReleaseMode)
	case "test":
		gin.SetMode(gin.TestMode)
	default:
		gin.SetMode(gin.DebugMode)
	}
}

// NewEngine global middleware setting
func NewEngine() *gin.Engine {
	grt := gin.Default()
	grt.Use(logger.GinLogger())
	grt.Use(logger.GinRecovery(true))
	grt.Use(CORS())
	return grt
}

func CORS() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, X-Forwarded-For, Authorization, accept, origin, Cache-Control, X-Requested-With")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT, DELETE")
		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}
		c.Next()
	}
}
