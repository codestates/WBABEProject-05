package order

import (
	"github.com/codestates/WBABEProject-05/model/entity"
	"github.com/codestates/WBABEProject-05/protocol"
	utilErr "github.com/codestates/WBABEProject-05/protocol/error"
	"github.com/codestates/WBABEProject-05/protocol/request"
	"github.com/codestates/WBABEProject-05/service/order"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
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
func (o *orderRecordControl) PostOrderRecord(c *gin.Context) {
	reqO := &request.RequestOrder{}
	if err := c.ShouldBindJSON(reqO); err != nil {
		protocol.Fail(utilErr.BadRequestError).Response(c)
		return
	}

	recordedID, err := o.orderService.RegisterOrderRecord(reqO)
	if err != nil {
		protocol.Fail(utilErr.NewApiError(err)).Response(c)
		return
	}

	protocol.SuccessCodeAndData(
		http.StatusCreated,
		gin.H{"posted_id": recordedID},
	).Response(c)
}

func (o *orderRecordControl) PutOrderRecordFromCustomer(c *gin.Context) {
	reqO := &request.RequestPutCustomerOrder{}
	err := c.ShouldBindJSON(reqO)
	if err != nil {
		protocol.Fail(utilErr.BadRequestError).Response(c)
		return
	}

	newOrderID, err := o.orderService.ModifyOrderRecordFromCustomer(reqO)
	if err != nil {
		protocol.Fail(utilErr.NewApiError(err)).Response(c)
		return
	}
	protocol.SuccessData(gin.H{
		"new_order_id": newOrderID,
	}).Response(c)
}

func (o *orderRecordControl) PutOrderRecordFromStore(c *gin.Context) {
	reqO := &request.RequestPutStoreOrder{}
	err := c.ShouldBindJSON(reqO)
	if err != nil {
		protocol.Fail(utilErr.BadRequestError).Response(c)
		return
	}

	updatedCnt, err := o.orderService.ModifyOrderRecordFromStore(reqO)
	if err != nil {
		protocol.Fail(utilErr.NewApiError(err)).Response(c)
		return
	}
	protocol.SuccessData(gin.H{
		"updated_count": updatedCnt,
	}).Response(c)
}

func (o *orderRecordControl) GetCustomerOrderRecordsSortedPage(c *gin.Context) {
	page := &request.RequestPage{}
	customerID := c.Query("customer-id")
	if err := c.ShouldBindQuery(page); err != nil || customerID == "" {
		protocol.Fail(utilErr.BadRequestError).Response(c)
		return
	}

	receipts, err := o.orderService.FindOrderRecordsSortedPage(customerID, entity.CustomerRole, page)
	if err != nil {
		protocol.Fail(utilErr.NewApiError(err)).Response(c)
		return
	}
	protocol.SuccessData(receipts).Response(c)
}

func (o *orderRecordControl) GetStoreOrderRecordsSortedPage(c *gin.Context) {
	page := &request.RequestPage{}
	storeID := c.Query("store-id")
	if err := c.ShouldBindQuery(page); err != nil || storeID == "" {
		protocol.Fail(utilErr.BadRequestError).Response(c)
		return
	}

	receipts, err := o.orderService.FindOrderRecordsSortedPage(storeID, entity.StoreRole, page)
	if err != nil {
		protocol.Fail(utilErr.NewApiError(err)).Response(c)
		return
	}
	protocol.SuccessData(receipts).Response(c)
}

func (o *orderRecordControl) GetOrderRecord(c *gin.Context) {
	ordrID := c.Query("order-id")
	resOrder, err := o.orderService.FindOrderRecord(ordrID)
	if err != nil {
		protocol.Fail(utilErr.NewApiError(err)).Response(c)
		return
	}
	protocol.SuccessData(resOrder).Response(c)
}

func (o *orderRecordControl) GetSelectedMenusTotalPrice(c *gin.Context) {
	reqCheckP := &request.RequestCheckPrice{}
	if err := c.ShouldBindQuery(reqCheckP); err != nil || reqCheckP.Menus == nil {
		protocol.Fail(utilErr.BadRequestError).Response(c)
		return
	}
	reqCheckP.Menus = strings.Split(reqCheckP.Menus[0], ",")

	resCheckP, err := o.orderService.FiendSelectedMenusTotalPrice(reqCheckP.StoreID, reqCheckP.Menus)
	if err != nil {
		protocol.Fail(utilErr.NewApiError(err)).Response(c)
		return
	}
	protocol.SuccessData(resCheckP).Response(c)
}
