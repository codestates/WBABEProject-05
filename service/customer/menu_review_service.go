package customer

import (
	"errors"
	"fmt"
	"github.com/codestates/WBABEProject-05/logger"
	"github.com/codestates/WBABEProject-05/model/menu"
	"github.com/codestates/WBABEProject-05/model/review"
	"github.com/codestates/WBABEProject-05/protocol/page"
	"github.com/codestates/WBABEProject-05/protocol/request"
	"math"
)

var instance *menuReviewService

type menuReviewService struct {
	reviewModel review.ReviewModeler
	menuModel   menu.MenuModeler
}

func NewMenuReviewService(rMod review.ReviewModeler, mMod menu.MenuModeler) *menuReviewService {
	if instance != nil {
		return instance
	}
	instance = &menuReviewService{
		reviewModel: rMod,
		menuModel:   mMod,
	}
	return instance
}

// FindReviewsSortedPage 메뉴 조회 : 개별 메뉴별 평점 및 리뷰 보기, / 해당 메뉴 선택시 메뉴에 따른 평점 및 리뷰 데이터 리턴
func (m *menuReviewService) FindReviewSortedPageByMenuID(menuID string, pg *request.RequestPage) (*page.PageData[any], error) {
	skip := pg.CurrentPage * pg.ContentCount

	reviews, err := m.reviewModel.SelectSortLimitedReviewsByMenuID(menuID, pg.Sort, skip, pg.ContentCount)
	if err != nil {
		return nil, err
	}

	totalCount, err := m.reviewModel.SelectTotalCountByMenuID(menuID)
	if err != nil {
		return nil, err
	}

	pgInfo := pg.NewPageInfo(totalCount)

	return page.NewPageData(reviews, pgInfo), nil
}
func (m *menuReviewService) FindReviewSortedPageByUserID(ID, userRole string, pg *request.RequestPage) (*page.PageData[any], error) {
	skip := pg.CurrentPage * pg.ContentCount

	reviews, err := m.reviewModel.SelectSortLimitedReviewsByUserID(ID, userRole, pg.Sort, skip, pg.ContentCount)
	if err != nil {
		return nil, err
	}

	totalCount, err := m.reviewModel.SelectTotalCountByUserID(ID, userRole)
	if err != nil {
		return nil, err
	}

	pgInfo := pg.NewPageInfo(totalCount)

	return page.NewPageData(reviews, pgInfo), nil
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

	menu, err := m.menuModel.SelectMenuByID(review.Menu)
	if err != nil {
		return "", err
	}

	menu.TotalReviewScore += review.Rating
	menu.ReviewCount++
	menu.Rating = math.Round((float64(menu.TotalReviewScore)/float64(menu.ReviewCount))*10) / 10

	// Rating 은 비즈니스상 중요하지않아 채널을 활용
	rating, err := m.menuModel.UpdateAboutRating(menu)
	if err != nil || rating == 0 {
		msg := fmt.Sprintf("does not update rating Menu ID %v", menu.ID)
		logger.AppLog.Error(errors.New(msg))
	}

	return savedID, nil
}
