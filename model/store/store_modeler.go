package store

import (
	"github.com/codestates/WBABEProject-05/model/entity"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// TODO 디테일한 점이 많아 서비스 만들어야할듯하다
type StoreModeler interface {
	// InsertMenu 신규 메뉴 등록 : 관련 정보를 등록(이름,주문가능여부,한정수량,가격,맵기정도 등) , / 성공 여부 리턴
	InsertMenu(store *entity.Store) (int, error)
	// DeleteMenu 메뉴 삭제 : 메뉴 삭제시 실제 데이터 백업이나 뷰플레그로 안보임 처리 , / 성공 여부 리턴
	DeleteMenu()
	// UpdateMenu 메뉴 수정 : 사업장에서 기존의 메뉴 정보 변경기능(ex. 가격변경, 원산지 변경, soldout) 금일 추천 메뉴 설정 변경, 리스트 출력, / 성공 여부 리턴
	UpdateMenu()
	// SelectMenus 메뉴 리스트 조회 : 메뉴 리스트 조회 및 정렬(추천/평점/주문수/최신) / 각 카테고리별  sort 리스트 출력(ex. order by 추천, 평점, 재주문수, 최신), 결과 5~10여개 임의 생성 출력, sorting 여부 확인
	SelectMenus()

	// SelectMenusByIds 메뉴 5개를 가져온다. (추천메뉴 최대 5개로 정했기에, 최신순서로)
	SelectMenusByIds(menuIds []primitive.ObjectID) ([]*entity.Menu, error)
	InsertStore(store *entity.Store) (string, error)
}
