package store

import "github.com/gin-gonic/gin"

var StoreControl StoreContoller

type StoreContoller interface {
	PostStore(c *gin.Context)

	PutSore(c *gin.Context)

	PostMenu(c *gin.Context)

	PutMenu(c *gin.Context)

	DeleteMenu(c *gin.Context)

	GetMenuSortedPages(c *gin.Context)

	GetRecommendMenus(c *gin.Context)

	// GetStoreInSwagForTest : swagger 테스트를 위한 store 조회
	GetStoreInSwagForTest(c *gin.Context)
}
