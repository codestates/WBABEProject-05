package review

import "github.com/gin-gonic/gin"

var MenuReviewControl MenuReviewController

type MenuReviewController interface {
	GetMenuReviewSortedPagesByCustomerID(c *gin.Context)

	GetMenuReviewSortedPagesByMenuID(c *gin.Context)

	PostMenuReview(c *gin.Context)
}
