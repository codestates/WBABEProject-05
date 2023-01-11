package order

import "github.com/gin-gonic/gin"

var OrderRecordControl OrderRecordController

type OrderRecordController interface {
	PostOrderRecord(c *gin.Context)

	PutOrderRecordFromCustomer(c *gin.Context)

	PutOrderRecordFromStore(c *gin.Context)

	GetCustomerOrderRecordsSortedPage(c *gin.Context)

	GetStoreOrderRecordsSortedPage(c *gin.Context)

	GetOrderRecord(c *gin.Context)

	GetSelectedMenusTotalPrice(c *gin.Context)
}
