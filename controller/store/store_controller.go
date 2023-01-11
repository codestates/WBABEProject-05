package store

import "github.com/gin-gonic/gin"

var StoreControl StoreContoller

type StoreContoller interface {
	PostStore(c *gin.Context)

	PutStore(c *gin.Context)

	PostMenu(c *gin.Context)

	PutMenu(c *gin.Context)

	DeleteMenu(c *gin.Context)

	GetMenuSortedPagesByStoreID(c *gin.Context)

	GetMenuSortedPagesByName(c *gin.Context)

	GetRecommendMenus(c *gin.Context)

	GetStoresSortedPage(c *gin.Context)

	GetStore(c *gin.Context)
}
