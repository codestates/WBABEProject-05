package order

import (
	"github.com/codestates/WBABEProject-05/protocol"
	utilErr "github.com/codestates/WBABEProject-05/protocol/error"
	"github.com/codestates/WBABEProject-05/protocol/request"
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

// RegisterOrderRecord godoc
// @Summary call Post Order, return posted id by json.
// @Description 메뉴 주문을 할 수 있다.
// @name RegisterOrderRecord
// @Accept  json
// @Produce  json
// @Router /app/v1/orders [post]
// @Param order body protocol.RequestOrder true "RequestOrder JSON"
// @Success 200 {object} protocol.ApiResponse[any]
func (o *orderRecordControl) RegisterOrderRecord(c *gin.Context) {
	reqO := &request.RequestOrder{}
	if err := c.ShouldBindJSON(reqO); err != nil {
		protocol.Fail(utilErr.BadRequestError).Response(c)
		return
	}

	recordedID, err := o.orderService.RegisterOrderRecord(reqO)
	if err != nil {
		protocol.Fail(utilErr.NewError(err)).Response(c)
		return
	}

	protocol.SuccessData(gin.H{"posted_id": recordedID}).Response(c)
}
func (o *orderRecordControl) ModifyOrderRecord(c *gin.Context) {

}
func (o *orderRecordControl) FindOrderRecordsSortedPage(c *gin.Context) {

}
func (o *orderRecordControl) SelectReceipts(c *gin.Context) {

}
