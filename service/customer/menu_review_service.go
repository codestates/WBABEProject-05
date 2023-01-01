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

	pgInfo := pg.NewPageInfo(int(totalCount))

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

	pgInfo := pg.NewPageInfo(int(totalCount))

	return page.NewPageData(reviews, pgInfo), nil
}

func (m *menuReviewService) RegisterMenuReview(review *request.RequestPostReview) (string, error) {
	r, err := review.NewReview()
	if err != nil {
		return "", err
	}

	savedID, err := m.reviewModel.InsertReview(r)
	if err != nil {
		return "", err
	}

	foundM, err := m.menuModel.SelectMenuByID(review.MenuID)
	if err != nil {
		return "", err
	}

	foundM.TotalReviewScore += review.Rating
	foundM.ReviewCount++
	foundM.Rating = math.Round((float64(foundM.TotalReviewScore)/float64(foundM.ReviewCount))*10) / 10

	// Rating 은 비즈니스상 중요하지않아 고루틴 활용
	go func() {
		rating, err := m.menuModel.UpdateAboutRating(foundM)
		if err != nil || rating == 0 {
			msg := fmt.Sprintf("does not update rating Menu ID %v", foundM.ID)
			logger.AppLog.Error(errors.New(msg))
		}
	}()

	return savedID, nil
}
