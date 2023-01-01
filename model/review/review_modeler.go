package review

import (
	"github.com/codestates/WBABEProject-05/model/entity"
	"github.com/codestates/WBABEProject-05/protocol/page"
)

var ReviewModel ReviewModeler

type ReviewModeler interface {
	InsertReview(review *entity.Review) (string, error)

	SelectSortLimitedReviewsByMenuID(menuID string, sort *page.Sort, skip int, limit int) ([]*entity.Review, error)

	SelectSortLimitedReviewsByUserID(ID, userRole string, sort *page.Sort, skip int, limit int) ([]*entity.Review, error)

	SelectTotalCountByMenuID(menuID string) (int64, error)

	SelectTotalCountByUserID(ID, userRole string) (int64, error)
}
