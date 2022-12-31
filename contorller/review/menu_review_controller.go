package review

import "github.com/gin-gonic/gin"

var MenuReviewControl MenuReviewController

type MenuReviewController interface {
	// FindReviewsSortedPage 메뉴 조회 : 개별 메뉴별 평점 및 리뷰 보기, / 해당 메뉴 선택시 메뉴에 따른 평점 및 리뷰 데이터 리턴
	GetMenuSortedPagesByUserID(c *gin.Context)
	GetMenuSortedPagesByMenuID(c *gin.Context)

	// InsertReview 메뉴별 평점 작성 : 과거 주문 내역 중, 평점 및 리뷰 작성, 해당 주문내역을 기준, 평점 정보, 리뷰 스트링을 입력받아 과거 주문내역 업데이트 저장, / 성공 여부 리턴
	PostMenuReview(c *gin.Context)
}
