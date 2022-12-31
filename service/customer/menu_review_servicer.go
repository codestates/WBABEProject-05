package customer

import (
	"github.com/codestates/WBABEProject-05/protocol/page"
	"github.com/codestates/WBABEProject-05/protocol/request"
)

type MenuReviewServicer interface {
	// InsertReview 메뉴별 평점 작성 : 과거 주문 내역 중, 평점 및 리뷰 작성, 해당 주문내역을 기준, 평점 정보, 리뷰 스트링을 입력받아 과거 주문내역 업데이트 저장, / 성공 여부 리턴
	RegisterMenuReview(review *request.RequestPostReview) (string, error)
	// FindReviewsSortedPage 메뉴 조회 : 개별 메뉴별 평점 및 리뷰 보기, / 해당 메뉴 선택시 메뉴에 따른 평점 및 리뷰 데이터 리턴
	FindReviewSortedPageByMenuID(menuID string, page *request.RequestPage) (*page.PageData[any], error)
	FindReviewSortedPageByUserID(userID string, page *request.RequestPage) (*page.PageData[any], error)
}
