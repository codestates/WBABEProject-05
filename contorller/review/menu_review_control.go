package review

import (
	utilErr "github.com/codestates/WBABEProject-05/common/error"
	"github.com/codestates/WBABEProject-05/model/enum"
	"github.com/codestates/WBABEProject-05/protocol"
	"github.com/codestates/WBABEProject-05/protocol/request"
	"github.com/codestates/WBABEProject-05/service/customer"
	"github.com/codestates/WBABEProject-05/service/validator"
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

// GetMenuSortedPagesByCustomerID godoc
// @Tags 메뉴리뷰
// @Summary call Get sorted page menu reviews, return sorted page menu reviews by json.
// @Description 특정 사용자의 리뷰들을 볼 수 있다.
// @name GetMenuSortedPagesByCustomerID
// @Accept  json
// @Produce  json
// @Router /app/v1/reviews/customer [get]
// @Param customer-id query string true "customer-id"
// @Param RequestPage query request.RequestPage true "RequestPage"
// @Param Sort query page.Sort true "Sort"
// @Success 200 {object} protocol.ApiResponse[any]
func (m *menuReviewControl) GetMenuReviewSortedPagesByCustomerID(c *gin.Context) {
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

	reviews, err := m.menuReviewService.FindReviewSortedPageByUserID(customerID, enum.CustomerRole, page)
	if err != nil {
		protocol.Fail(utilErr.NewAppError(err)).Response(c)
		return
	}
	protocol.SuccessData(reviews).Response(c)
}

// GetMenuReviewSortedPagesByMenuID godoc
// @Tags 메뉴리뷰
// @Summary call Get sorted page menu reviews, return sorted page menu reviews by json.
// @Description 특정 메뉴의 리뷰들을 볼 수 있다.
// @name GetMenuReviewSortedPagesByMenuID
// @Accept  json
// @Produce  json
// @Router /app/v1/reviews/menu [get]
// @Param menu-id query string true "menu-id"
// @Param RequestPage query request.RequestPage true "RequestPage"
// @Param Sort query page.Sort true "Sort"
// @Success 200 {object} protocol.ApiResponse[any]
func (m *menuReviewControl) GetMenuReviewSortedPagesByMenuID(c *gin.Context) {
	page := &request.RequestPage{}
	if err := c.ShouldBindQuery(page); err != nil {
		protocol.Fail(utilErr.BadRequestError).Response(c)
		return
	}

	menuID := c.Query("menu-id")
	if err := validator.IsBlank(menuID); err != nil {
		protocol.Fail(utilErr.NewAppError(err)).Response(c)
		return
	}

	reviews, err := m.menuReviewService.FindReviewSortedPageByMenuID(menuID, page)
	if err != nil {
		protocol.Fail(utilErr.NewAppError(err)).Response(c)
		return
	}
	protocol.SuccessData(reviews).Response(c)
}

// PostMenuReview godoc
// @Tags 메뉴리뷰
// @Summary call Post menu review, return saved id by json.
// @Description 메뉴 리뷰를 작성 할 수 있다.
// @name PostMenuReview
// @Accept  json
// @Produce  json
// @Router /app/v1/reviews/review [post]
// @Param RequestPostReview body request.RequestPostReview true "RequestPostReview JSON"
// @Success 201 {object} protocol.ApiResponse[any]
func (m *menuReviewControl) PostMenuReview(c *gin.Context) {
	reqR := &request.RequestPostReview{}
	if err := c.ShouldBindJSON(reqR); err != nil {
		protocol.Fail(utilErr.BadRequestError).Response(c)
		return
	}

	savedID, err := m.menuReviewService.RegisterMenuReview(reqR)
	if err != nil {
		protocol.Fail(utilErr.NewAppError(err)).Response(c)
		return
	}
	protocol.SuccessCodeAndData(
		http.StatusCreated,
		gin.H{"saved_id": savedID},
	).Response(c)
}
