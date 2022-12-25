package order

import (
	utilErr "github.com/codestates/WBABEProject-05/common/error"
	"github.com/codestates/WBABEProject-05/protocol"
	"github.com/codestates/WBABEProject-05/service/order"
	"github.com/gin-gonic/gin"
)

var instance *orderRecordControl

type orderRecordControl struct {
	orderService order.OrderRecordServicer
}

func NewOrderRecordControl(svc order.OrderRecordServicer) *orderRecordControl {
	if instance != nil {
		return instance
	}
	instance = &orderRecordControl{
		orderService: svc,
	}
	return instance
}

func (o *orderRecordControl) RegisterOrderRecord(c *gin.Context) {
	reqO := &protocol.RequestOrder{}
	err := c.ShouldBindJSON(reqO)
	if err != nil {
		protocol.Fail(utilErr.BadRequestError).Response(c)
		return
	}

	recordedId, err := o.orderService.RegisterOrderRecord(reqO)
	if err != nil {
		protocol.Fail(utilErr.NewError(err)).Response(c)
		return
	}
	protocol.SuccessData(gin.H{
		"posted_id": recordedId,
	}).Response(c)
}
func (o *orderRecordControl) ModifyOrderRecord(c *gin.Context) {

}
func (o *orderRecordControl) FindOrderRecordsSortedPage(c *gin.Context) {

}
func (o *orderRecordControl) SelectReceipts(c *gin.Context) {

}
