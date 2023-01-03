package order

import (
	utilErr "github.com/codestates/WBABEProject-05/common/error"
	"github.com/codestates/WBABEProject-05/model/enum"
	"github.com/codestates/WBABEProject-05/protocol"
	"github.com/codestates/WBABEProject-05/protocol/request"
	"github.com/codestates/WBABEProject-05/service/order"
	"github.com/codestates/WBABEProject-05/service/validator"
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

// PostOrderRecord godoc
// @Tags 주문기록
// @Summary call Post Order, return posted id by json.
// @Description 메뉴 주문을 할 수 있다.
// @name PostOrderRecord
// @Accept  json
// @Produce  json
// @Router /app/v1/orders/order [post]
// @Param order body request.RequestOrder true "RequestOrder JSON"
// @Success 200 {object} protocol.ApiResponse[any]
func (o *orderRecordControl) PostOrderRecord(c *gin.Context) {
	reqO := &request.RequestOrder{}
	if err := c.ShouldBindJSON(reqO); err != nil {
		protocol.Fail(utilErr.BadRequestError).Response(c)
		return
	}

	recordedID, err := o.orderService.RegisterOrderRecord(reqO)
	if err != nil {
		protocol.Fail(utilErr.NewAppError(err)).Response(c)
		return
	}

	protocol.SuccessCodeAndData(
		http.StatusCreated,
		gin.H{"posted_id": recordedID},
	).Response(c)
}

// PutOrderRecordFromCustomer godoc
// @Tags 주문기록
// @Summary call Put order records in customer, return updated count by json.
// @Description 사용자가 주문을 변경 할 수 있다.
// @name PutOrderRecordFromCustomer
// @Accept  json
// @Produce  json
// @Router /app/v1/orders/order/customer [put]
// @Param RequestPutCustomerOrder body request.RequestPutCustomerOrder true "RequestPutCustomerOrder"
// @Success 200 {object} protocol.ApiResponse[any]
func (o *orderRecordControl) PutOrderRecordFromCustomer(c *gin.Context) {
	reqO := &request.RequestPutCustomerOrder{}
	if err := c.ShouldBindJSON(reqO); err != nil {
		protocol.Fail(utilErr.BadRequestError).Response(c)
		return
	}

	newOrderID, err := o.orderService.ModifyOrderRecordFromCustomer(reqO)
	if err != nil {
		protocol.Fail(utilErr.NewAppError(err)).Response(c)
		return
	}
	protocol.SuccessData(gin.H{
		"new_order_id": newOrderID,
	}).Response(c)
}

// PutOrderRecordFromStore godoc
// @Tags 주문기록
// @Summary call Put order records in store, return updated count by json.
// @Description 가게에서 주문 상태를 변경 할 수 있다.
// @name PutOrderRecordFromStore
// @Accept  json
// @Produce  json
// @Router /app/v1/orders/order/store [put]
// @Param RequestPutStoreOrder body request.RequestPutStoreOrder true "RequestPutStoreOrder"
// @Success 200 {object} protocol.ApiResponse[any]
func (o *orderRecordControl) PutOrderRecordFromStore(c *gin.Context) {
	reqO := &request.RequestPutStoreOrder{}
	if err := c.ShouldBindJSON(reqO); err != nil {
		protocol.Fail(utilErr.BadRequestError).Response(c)
		return
	}

	updatedCnt, err := o.orderService.ModifyOrderRecordFromStore(reqO)
	if err != nil {
		protocol.Fail(utilErr.NewAppError(err)).Response(c)
		return
	}
	protocol.SuccessData(gin.H{
		"updated_count": updatedCnt,
	}).Response(c)
}

// GetCustomerOrderRecordsSortedPage godoc
// @Tags 주문기록
// @Summary call Get sorted pages customer order records, return order records by json.
// @Description 특정 사용자의 주문기록들을 볼 수 있다.
// @name GetCustomerOrderRecordsSortedPage
// @Accept  json
// @Produce  json
// @Router /app/v1/orders/pages/customer [get]
// @Param customer-id query string true "customer-id"
// @Param RequestPage query request.RequestPage true "RequestPage"
// @Param Sort query page.Sort true "Sort"
// @Success 200 {object} protocol.ApiResponse[any]
func (o *orderRecordControl) GetCustomerOrderRecordsSortedPage(c *gin.Context) {
	page := &request.RequestPage{}
	if err := c.ShouldBindQuery(page); err != nil {
		protocol.Fail(utilErr.BadRequestError).Response(c)
		return
	}

	customerID := c.Query("customer-id")
	if err := validator.IsBlank(customerID); err != nil {
		protocol.Fail(utilErr.NewAppError(err)).Response(c)
		return
	}

	receipts, err := o.orderService.FindOrderRecordsSortedPage(customerID, enum.CustomerRole, page)
	if err != nil {
		protocol.Fail(utilErr.NewAppError(err)).Response(c)
		return
	}
	protocol.SuccessData(receipts).Response(c)
}

// GetStoreOrderRecordsSortedPage godoc
// @Tags 주문기록
// @Summary call Get sorted pages store order records, return order records by json.
// @Description 특정 가게의 주문기록들을 볼 수 있다.
// @name GetStoreOrderRecordsSortedPage
// @Accept  json
// @Produce  json
// @Router /app/v1/orders/pages/store [get]
// @Param store-id query string true "store-id"
// @Param RequestPage query request.RequestPage true "RequestPage"
// @Param Sort query page.Sort true "Sort"
// @Success 200 {object} protocol.ApiResponse[any]
func (o *orderRecordControl) GetStoreOrderRecordsSortedPage(c *gin.Context) {
	page := &request.RequestPage{}
	if err := c.ShouldBindQuery(page); err != nil {
		protocol.Fail(utilErr.BadRequestError).Response(c)
		return
	}

	storeID := c.Query("store-id")
	if err := validator.IsBlank(storeID); err != nil {
		protocol.Fail(utilErr.NewAppError(err)).Response(c)
		return
	}

	receipts, err := o.orderService.FindOrderRecordsSortedPage(storeID, enum.StoreRole, page)
	if err != nil {
		protocol.Fail(utilErr.NewAppError(err)).Response(c)
		return
	}
	protocol.SuccessData(receipts).Response(c)
}

// GetOrderRecord godoc
// @Tags 주문기록
// @Summary call Get order-record, return order-record by json.
// @Description 특정 주문기록을 볼 수 있다.
// @name GetOrderRecord
// @Accept  json
// @Produce  json
// @Router /app/v1/orders/order [get]
// @Param order-id query string true "order-id"
// @Success 200 {object} protocol.ApiResponse[any]
func (o *orderRecordControl) GetOrderRecord(c *gin.Context) {
	orderID := c.Query("order-id")
	if err := validator.IsBlank(orderID); err != nil {
		protocol.Fail(utilErr.NewAppError(err)).Response(c)
		return
	}

	resOrder, err := o.orderService.FindOrderRecord(orderID)
	if err != nil {
		protocol.Fail(utilErr.NewAppError(err)).Response(c)
		return
	}

	protocol.SuccessData(resOrder).Response(c)
}

// GetSelectedMenusTotalPrice godoc
// @Tags 주문기록
// @Summary call Get selected menus total price, return total price by json.
// @Description 선택한 메뉴들의 총 가격을 알 수 있다.
// @name GetSelectedMenusTotalPrice
// @Accept  json
// @Produce  json
// @Router /app/v1/orders/order/price [get]
// @Param RequestCheckPrice query request.RequestCheckPrice true "RequestCheckPrice"
// @Success 200 {object} protocol.ApiResponse[any]
func (o *orderRecordControl) GetSelectedMenusTotalPrice(c *gin.Context) {
	reqCheckP := &request.RequestCheckPrice{}
	if err := c.ShouldBindQuery(reqCheckP); err != nil || reqCheckP.Menus == nil {
		protocol.Fail(utilErr.BadRequestError).Response(c)
		return
	}
	reqCheckP.Menus = strings.Split(reqCheckP.Menus[0], ",")

	resCheckP, err := o.orderService.FiendSelectedMenusTotalPrice(reqCheckP.StoreID, reqCheckP.Menus)
	if err != nil {
		protocol.Fail(utilErr.NewAppError(err)).Response(c)
		return
	}
	protocol.SuccessData(resCheckP).Response(c)
}
