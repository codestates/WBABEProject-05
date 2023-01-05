package review

import (
	"github.com/codestates/WBABEProject-05/model/common/query"
	"github.com/codestates/WBABEProject-05/model/entity"
)

var ReviewModel ReviewModeler

type ReviewModeler interface {
	InsertReview(review *entity.Review) (string, error)

	SelectSortLimitedReviewsByMenuID(menuID string, pageQuery *query.PageQuery) ([]*entity.Review, error)

	SelectSortLimitedReviewsByUserID(ID, userRole string, pageQuery *query.PageQuery) ([]*entity.Review, error)

	SelectTotalCountByMenuID(menuID string) (int64, error)

	SelectTotalCountByUserID(ID, userRole string) (int64, error)
}
