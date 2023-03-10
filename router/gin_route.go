package router

import (
	"github.com/codestates/WBABEProject-05/docs"
	"github.com/codestates/WBABEProject-05/logger"
	"github.com/gin-gonic/gin"
	swgFiles "github.com/swaggo/files"
	ginSwg "github.com/swaggo/gin-swagger"
	"net/http"
)

var ginR *GinRoute

type GinRoute struct {
	engin *gin.Engine
}

func NewGinRoute(mode string) *GinRoute {
	if ginR != nil {
		return ginR
	}
	setMode(mode)
	ginR = &GinRoute{
		engin: newEngine(),
	}
	return ginR
}

func (r *GinRoute) Handle() http.Handler {
	gr := r.engin

	registerSwagger(gr)

	home := gr.Group("/home")
	{
		HomeHandler(home)
	}

	v1 := gr.Group("app/v1")
	{
		users := v1.Group("/users")
		{
			UsersHandler(users)
		}

		stores := v1.Group("/stores")
		{
			StoresHandler(stores)
			menus := stores.Group("/store/menus")
			MenusHandler(menus)
		}

		orders := v1.Group("/orders")
		{
			OrdersHandler(orders)
		}

		reviews := v1.Group("/reviews")
		{
			ReviewHandler(reviews)
		}
	}

	return r.engin
}

func registerSwagger(gr *gin.Engine) {
	gr.GET("/swagger/:any", ginSwg.WrapHandler(swgFiles.Handler))
	docs.SwaggerInfo.Host = "localhost:8080" //swagger 정보 등록
}

func setMode(mode string) {
	switch mode {
	case "dev":
		logger.AppLog.Info("Start gin mod", gin.DebugMode)
		gin.SetMode(gin.DebugMode)
	case "prod":
		logger.AppLog.Info("Start gin mod", gin.ReleaseMode)
		gin.SetMode(gin.ReleaseMode)
	case "test":
		logger.AppLog.Info("Start gin mod", gin.TestMode)
		gin.SetMode(gin.TestMode)
	default:
		logger.AppLog.Info("Start gin mod", gin.DebugMode)
		gin.SetMode(gin.DebugMode)
	}
}

// newEngine generate gin engin and global middleware setting
func newEngine() *gin.Engine {
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
