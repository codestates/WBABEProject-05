package contorller

import (
	"github.com/codestates/WBABEProject-05/contorller/info"
	order2 "github.com/codestates/WBABEProject-05/contorller/order"
	store2 "github.com/codestates/WBABEProject-05/contorller/store"
	user2 "github.com/codestates/WBABEProject-05/contorller/user"
	"github.com/gin-gonic/gin"
)

// HomeHandler ("")
func HomeHandler(homeUrl *gin.RouterGroup) {
	homeUrl.GET("/", func(c *gin.Context) {
		c.JSON(200, "ok")
	})
	homeUrl.GET("/info", info.InfoControl.GetInformation)
}

// UsersHandler ("/users")
func UsersHandler(usersUrl *gin.RouterGroup) {
	usersUrl.POST("/join", user2.UserControl.PostUser)
}

// StoresHandler ("/stores")
func StoresHandler(storesUrl *gin.RouterGroup) {
	storesUrl.POST("", store2.StoreControl.PostStore)         // 가게 등록
	storesUrl.POST("/menu", store2.StoreControl.PostMenu)     // 메뉴 등록
	storesUrl.DELETE("/menu", store2.StoreControl.DeleteMenu) // 메뉴 삭제
	storesUrl.PUT("/menu", store2.StoreControl.PutMenu)
}

// OrdersHandler ("/orders")
func OrdersHandler(ordersUrl *gin.RouterGroup) {
	ordersUrl.POST("", order2.OrderRecordControl.RegisterOrderRecord)
}
