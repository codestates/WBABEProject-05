package customer

import (
	"github.com/codestates/WBABEProject-05/protocol/page"
	"github.com/codestates/WBABEProject-05/protocol/request"
)

var MenuReviewService MenuReviewServicer

type MenuReviewServicer interface {
	RegisterMenuReview(review *request.RequestPostReview) (string, error)

	FindReviewSortedPageByMenuID(menuID string, page *request.RequestPage) (*page.PageData[any], error)

	FindReviewSortedPageByUserID(ID, userRole string, page *request.RequestPage) (*page.PageData[any], error)
}
