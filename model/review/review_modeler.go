package review

import (
	"github.com/codestates/WBABEProject-05/model/entity"
	"github.com/codestates/WBABEProject-05/protocol/page"
)

var ReviewModel ReviewModeler

type ReviewModeler interface {
	// SelectReviews 메뉴 조회 : 개별 메뉴별 평점 및 리뷰 보기, / 해당 메뉴 선택시 메뉴에 따른 평점 및 리뷰 데이터 리턴
	SelectReviews() // 메뉴 아이디로
	// InsertReview 메뉴별 평점 작성 : 과거 주문 내역 중, 평점 및 리뷰 작성, 해당 주문내역을 기준, 평점 정보, 리뷰 스트링을 입력받아 과거 주문내역 업데이트 저장, / 성공 여부 리턴
	InsertReview(review *entity.Review) (string, error)
	SelectSortLimitedReviewsByMenuID(menuID string, sort *page.Sort, skip int, limit int) ([]*entity.Review, error)
	SelectSortLimitedReviewsByUserID(userID string, sort *page.Sort, skip int, limit int) ([]*entity.Review, error)
	SelectTotalCountByMenuID(menuID string) (int, error)
	SelectTotalCountByUserID(userID string) (int, error)
}
