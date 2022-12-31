package router

import (
	"github.com/codestates/WBABEProject-05/contorller/info"
	order2 "github.com/codestates/WBABEProject-05/contorller/order"
	"github.com/codestates/WBABEProject-05/contorller/review"
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
	usersUrl.GET("/id", user2.UserControl.GetUser)
}

// StoresHandler ("/stores")
func StoresHandler(storesUrl *gin.RouterGroup) {
	storesUrl.POST("", store2.StoreControl.PostStore)           // 가게 등록
	storesUrl.PUT("/id", store2.StoreControl.PutSore)           // 가게 등록
	storesUrl.GET("/id", store2.StoreControl.GetRecommendMenus) // 가게 등록
	storesUrl.POST("/menu", store2.StoreControl.PostMenu)       // 메뉴 등록
	storesUrl.PUT("/menu", store2.StoreControl.PutMenu)
	storesUrl.DELETE("/menu", store2.StoreControl.DeleteMenu) // 메뉴 삭제
	// "/swag/store" sagger 테스트용
	storesUrl.GET("/swag/store", store2.StoreControl.GetStoreInSwagForTest)
	storesUrl.GET("/menus/pages/id", store2.StoreControl.GetMenuSortedPages)
	storesUrl.GET("", store2.StoreControl.GetStoresSortedPage)
}

// OrdersHandler ("/orders")
func OrdersHandler(ordersUrl *gin.RouterGroup) {
	ordersUrl.POST("", order2.OrderRecordControl.PostOrderRecord)
	ordersUrl.PUT("/customer", order2.OrderRecordControl.PutOrderRecordFromCustomer)
	ordersUrl.PUT("/store", order2.OrderRecordControl.PutOrderRecordFromStore)
	ordersUrl.GET("/price", order2.OrderRecordControl.GetSelectedMenusTotalPrice)
	ordersUrl.GET("/id", order2.OrderRecordControl.GetOrderRecord)
	ordersUrl.GET("/pages/user", order2.OrderRecordControl.GetOrderRecordsSortedPage)

}

// ReviewHandler ("/reviews")
func ReviewHandler(reviewsUrl *gin.RouterGroup) {
	reviewsUrl.POST("", review.MenuReviewControl.PostMenuReview)
	reviewsUrl.GET("/menu/id", review.MenuReviewControl.GetMenuSortedPagesByMenuID)
	reviewsUrl.GET("/user/id", review.MenuReviewControl.GetMenuSortedPagesByUserID)
}
