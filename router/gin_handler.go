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
	usersUrl.POST("/user", user2.UserControl.PostUser)
	usersUrl.GET("/user", user2.UserControl.GetUser)
	usersUrl.PUT("/user", user2.UserControl.PutUser)
	usersUrl.DELETE("/user", user2.UserControl.DeleteUser)
}

// StoresHandler ("/stores")
func StoresHandler(storesUrl *gin.RouterGroup) {
	storesUrl.GET("", store2.StoreControl.GetStoresSortedPage)                // 가게들 정보
	storesUrl.GET("/store", store2.StoreControl.GetStore)                     // 특정 가게 정보
	storesUrl.POST("/store", store2.StoreControl.PostStore)                   // 가게 등록
	storesUrl.PUT("/store", store2.StoreControl.PutSore)                      // 가게 수정
	storesUrl.GET("/store/recommends", store2.StoreControl.GetRecommendMenus) // 가게 추천메뉴
	// "/swag/store" sagger 테스트용
}

// MenusHandler ("/stores/store/menus")
func MenusHandler(menusUrl *gin.RouterGroup) {
	menusUrl.GET("", store2.StoreControl.GetMenuSortedPages) // 가게 메뉴들
	menusUrl.POST("/menu", store2.StoreControl.PostMenu)     // 가게 메뉴 등록
	menusUrl.PUT("/menu", store2.StoreControl.PutMenu)       // 가게 메뉴 수정
	menusUrl.DELETE("/menu", store2.StoreControl.DeleteMenu) // 가게 메뉴 삭제
}

// OrdersHandler ("/orders")
func OrdersHandler(ordersUrl *gin.RouterGroup) {
	ordersUrl.GET("/pages/store", order2.OrderRecordControl.GetStoreOrderRecordsSortedPage)       // 가게내 주문 기록들
	ordersUrl.GET("/pages/customer", order2.OrderRecordControl.GetCustomerOrderRecordsSortedPage) // 사용자의 주문 기록들
	ordersUrl.POST("/order", order2.OrderRecordControl.PostOrderRecord)                           // 주문 등록
	ordersUrl.GET("/order", order2.OrderRecordControl.GetOrderRecord)                             // 특정 주문 조회
	ordersUrl.PUT("/order/customer", order2.OrderRecordControl.PutOrderRecordFromCustomer)        // 사용자의 주문 수정
	ordersUrl.PUT("/order/store", order2.OrderRecordControl.PutOrderRecordFromStore)              // 가게의 주문 수정
	ordersUrl.GET("/order/price", order2.OrderRecordControl.GetSelectedMenusTotalPrice)           // 주문 메뉴들 총 가격

}

// ReviewHandler ("/reviews")
func ReviewHandler(reviewsUrl *gin.RouterGroup) {
	reviewsUrl.POST("/review", review.MenuReviewControl.PostMenuReview)                  // 리뷰 등록
	reviewsUrl.GET("/menu", review.MenuReviewControl.GetMenuReviewSortedPagesByMenuID)   // 메뉴별 리뷰 조회
	reviewsUrl.GET("/customer", review.MenuReviewControl.GetMenuSortedPagesByCustomerID) // 사용자의 리뷰 조회
}
