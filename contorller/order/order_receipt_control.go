package order

import (
	utilErr "github.com/codestates/WBABEProject-05/common/error"
	"github.com/codestates/WBABEProject-05/protocol"
	"github.com/codestates/WBABEProject-05/service"
	"github.com/gin-gonic/gin"
)

var instance *orderControl

type orderControl struct {
	orderService service.OrderReceiptServicer
}

func GetOrderControl(svc service.OrderReceiptServicer) *orderControl {
	if instance != nil {
		return instance
	}
	instance = &orderControl{
		orderService: svc,
	}
	return instance
}

func (o *orderControl) RegisterOrderRecord(c *gin.Context) {
	reqO := &protocol.RequestOrder{}
	err := c.ShouldBindJSON(reqO)
	if err != nil {
		protocol.Fail(utilErr.BadRequestError).Response(c)
		return
	}

	recordedId, err := o.orderService.RegisterOrderRecord(reqO)
	if err != nil {
		// TODO ERR
		return
	}
	protocol.SuccessData(gin.H{
		"posted_id": recordedId,
	}).Response(c)
}
func (o *orderControl) ModifyOrderRecord(c *gin.Context) {

}
func (o *orderControl) FindOrderRecordsSortedPage(c *gin.Context) {

}
func (o *orderControl) SelectReceipts(c *gin.Context) {

}
