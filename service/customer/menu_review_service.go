package customer

import (
	"github.com/codestates/WBABEProject-05/model/review"
	"github.com/codestates/WBABEProject-05/protocol/request"
)

var instance *menuReviewService

type menuReviewService struct {
	reviewModel review.ReviewModeler
}

func NewMenuReviewService(mod review.ReviewModeler) *menuReviewService {
	if instance != nil {
		return instance
	}
	instance = &menuReviewService{
		reviewModel: mod,
	}
	return instance
}

// FindReviewsSortedPage 메뉴 조회 : 개별 메뉴별 평점 및 리뷰 보기, / 해당 메뉴 선택시 메뉴에 따른 평점 및 리뷰 데이터 리턴
func (m *menuReviewService) FindReviewsSortedPageByMenu() {

}

// InsertReview 메뉴별 평점 작성 : 과거 주문 내역 중, 평점 및 리뷰 작성, 해당 주문내역을 기준, 평점 정보, 리뷰 스트링을 입력받아 과거 주문내역 업데이트 저장, / 성공 여부 리턴
func (m *menuReviewService) RegisterMenuReview(review *request.RequestPostReview) (string, error) {
	r, err := review.NewReview()
	if err != nil {
		return "", err
	}

	savedID, err := m.reviewModel.InsertReview(r)
	if err != nil {
		return "", err
	}
	return savedID, nil
}
