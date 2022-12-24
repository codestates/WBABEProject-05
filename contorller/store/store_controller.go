package store

import "github.com/gin-gonic/gin"

type StoreContoller interface {
	// PostMenu 신규 메뉴 등록 : 관련 정보를 등록(이름,주문가능여부,한정수량,가격,맵기정도 등) , / 성공 여부 리턴
	PostMenu(g *gin.Context)
	// DeleteMenu 메뉴 삭제 : 메뉴 삭제시 실제 데이터 백업이나 뷰플레그로 안보임 처리 , / 성공 여부 리턴
	DeleteMenu(g *gin.Context)
	// 사업장 정보 수정와 추천메뉴 수정
	PutSoreAndRecommendMenu(g *gin.Context)
	// PutMenu 메뉴 수정 : 사업장에서 기존의 메뉴 정보 변경기능(ex. 가격변경, 원산지 변경, soldout)
	PutMenu(g *gin.Context)
	// 금일 추천 메뉴 설정 변경, 리스트 출력, / 성공 여부 리턴
	GetRecommendMenusSortedTimeDesc(g *gin.Context)
	// GetMenuSortedPages 메뉴 리스트 조회 : 메뉴 리스트 조회 및 정렬(추천/평점/주문수/최신) / 각 카테고리별  sort 리스트 출력(ex. order by 추천, 평점, 재주문수, 최신), 결과 5~10여개 임의 생성 출력, sorting 여부 확인
	GetMenuSortedPages(g *gin.Context)
}
