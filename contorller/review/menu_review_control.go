package review

import (
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

func NeMenuReviewControl(svc customer.MenuReviewServicer) *menuReviewControl {
	if instance != nil {
		return instance
	}
	instance = &menuReviewControl{
		menuReviewService: svc,
	}
	return instance
}

func (m *menuReviewControl) GetMenuSortedPagesByUserID(c *gin.Context) {
	page := &request.RequestPage{}
	userID := c.Query("user-id")
	if err := c.ShouldBindQuery(page); err != nil || userID == "" {
		protocol.Fail(utilErr.BadRequestError).Response(c)
		return
	}

	reviews, err := m.menuReviewService.FindReviewSortedPageByUserID(userID, page)
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
