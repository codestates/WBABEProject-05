package review

import (
	"github.com/codestates/WBABEProject-05/model/entity"
	"github.com/codestates/WBABEProject-05/protocol"
	utilErr "github.com/codestates/WBABEProject-05/protocol/error"
	"github.com/codestates/WBABEProject-05/protocol/request"
	"github.com/codestates/WBABEProject-05/service/customer"
	"github.com/gin-gonic/gin"
	"net/http"
)

var instance *menuReviewControl

type menuReviewControl struct {
	menuReviewService customer.MenuReviewServicer
}

func NewMenuReviewControl(svc customer.MenuReviewServicer) *menuReviewControl {
	if instance != nil {
		return instance
	}
	instance = &menuReviewControl{
		menuReviewService: svc,
	}
	return instance
}

func (m *menuReviewControl) GetMenuSortedPagesByCustomerID(c *gin.Context) {
	page := &request.RequestPage{}
	customerID := c.Query("customer-id")
	if err := c.ShouldBindQuery(page); err != nil || customerID == "" {
		protocol.Fail(utilErr.BadRequestError).Response(c)
		return
	}

	reviews, err := m.menuReviewService.FindReviewSortedPageByUserID(customerID, entity.CustomerRole, page)
	if err != nil {
		protocol.Fail(utilErr.NewApiError(err)).Response(c)
		return
	}
	protocol.SuccessData(reviews).Response(c)
}

func (m *menuReviewControl) GetMenuSortedPagesByMenuID(c *gin.Context) {
	page := &request.RequestPage{}
	menuID := c.Query("menu-id")
	if err := c.ShouldBindQuery(page); err != nil || menuID == "" {
		protocol.Fail(utilErr.BadRequestError).Response(c)
		return
	}

	reviews, err := m.menuReviewService.FindReviewSortedPageByMenuID(menuID, page)
	if err != nil {
		protocol.Fail(utilErr.NewApiError(err)).Response(c)
		return
	}
	protocol.SuccessData(reviews).Response(c)
}

func (m *menuReviewControl) PostMenuReview(c *gin.Context) {
	reqR := &request.RequestPostReview{}
	if err := c.ShouldBindJSON(reqR); err != nil {
		protocol.Fail(utilErr.BadRequestError).Response(c)
		return
	}

	savedID, err := m.menuReviewService.RegisterMenuReview(reqR)
	if err != nil {
		protocol.Fail(utilErr.NewApiError(err)).Response(c)
		return
	}
	protocol.SuccessCodeAndData(
		http.StatusCreated,
		gin.H{"saved_id": savedID},
	).Response(c)
}
