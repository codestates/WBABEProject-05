package customer

import (
	"fmt"
	"github.com/codestates/WBABEProject-05/common/enum"
	"github.com/codestates/WBABEProject-05/logger"
	"github.com/codestates/WBABEProject-05/model/common/query"
	"github.com/codestates/WBABEProject-05/model/menu"
	"github.com/codestates/WBABEProject-05/model/review"
	"github.com/codestates/WBABEProject-05/protocol/page"
	"github.com/codestates/WBABEProject-05/protocol/request"
	util2 "github.com/codestates/WBABEProject-05/service/common"
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

func (m *menuReviewService) FindReviewSortedPageByMenuID(menuID string, pg *request.RequestPage) (*page.PageData[any], error) {
	skip := util2.NewSkipNumber(pg.CurrentPage, pg.ContentCount)

	pageQuery := query.NewPageQuery(pg.Sort.Name, pg.Sort.Direction, skip, pg.ContentCount)

	reviews, err := m.reviewModel.SelectSortLimitedReviewsByMenuID(menuID, pageQuery)
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
	skip := util2.NewSkipNumber(pg.CurrentPage, pg.ContentCount)

	pageQuery := query.NewPageQuery(pg.Sort.Name, pg.Sort.Direction, skip, pg.ContentCount)

	reviews, err := m.reviewModel.SelectSortLimitedReviewsByUserID(ID, userRole, pageQuery)
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
	newR, err := review.NewReview()
	if err != nil {
		return enum.BlankSTR, err
	}

	savedID, err := m.reviewModel.InsertReview(newR)
	if err != nil {
		return enum.BlankSTR, err
	}

	// Rating 은 비즈니스상 중요하지않아보여 따로 컨틀롤하지 않는 고루틴 활용
	go m.updateMenuScores(review)

	return savedID, nil
}

func (m *menuReviewService) updateMenuScores(review *request.RequestPostReview) {
	foundM, err := m.menuModel.SelectMenuByID(review.MenuID)
	if err != nil {
		logger.AppLog.Error(err.Error())
	}

	foundM.TotalReviewScore += review.Rating
	foundM.ReviewCount++
	foundM.Rating = math.Round((float64(foundM.TotalReviewScore)/float64(foundM.ReviewCount))*10) / 10

	rating, err := m.menuModel.UpdateMenuRating(foundM)
	if err != nil || rating == 0 {
		MSG := fmt.Sprintf("does not update rating Menu ID %v", foundM.ID)
		logger.AppLog.Error(MSG)
	}
}
