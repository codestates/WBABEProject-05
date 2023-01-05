package store

import (
	"encoding/json"
	"github.com/codestates/WBABEProject-05/common/flag"
	"github.com/codestates/WBABEProject-05/config/db"
	"github.com/codestates/WBABEProject-05/logger"
	"github.com/codestates/WBABEProject-05/model/common"
	"github.com/codestates/WBABEProject-05/model/common/query"
	"github.com/codestates/WBABEProject-05/model/entity"
	"github.com/codestates/WBABEProject-05/protocol/page"
	"github.com/codestates/WBABEProject-05/protocol/request"
	"github.com/codestates/WBABEProject-05/protocol/response"
	util2 "github.com/codestates/WBABEProject-05/service/common"
)

func (s *storeMenuService) FindMenusSortedPage(storeID string, pg *request.RequestPage) (*page.PageData[any], error) {
	skip := util2.NewSkipNumber(pg.CurrentPage, pg.ContentCount)

	pageQuery := query.NewPageQuery(pg.Sort.Name, pg.Sort.Direction, skip, pg.ContentCount)

	menus, err := s.menuModel.SelectSortLimitedMenus(storeID, pageQuery)
	if err != nil {
		return nil, err
	}

	totalCount, err := s.menuModel.SelectTotalCount(storeID)
	if err != nil {
		return nil, err
	}

	pgInfo := pg.NewPageInfo(int(totalCount))

	return page.NewPageData(menus, pgInfo), nil
}

func (s *storeMenuService) FindMenusSortedPageByName(name string, pg *request.RequestPage) (*page.PageData[any], error) {
	skip := util2.NewSkipNumber(pg.CurrentPage, pg.ContentCount)

	pageQuery := query.NewPageQuery(pg.Sort.Name, pg.Sort.Direction, skip, pg.ContentCount)

	menus, err := s.menuModel.SelectSortLimitedMenusByName(name, pageQuery)
	if err != nil {
		return nil, err
	}

	totalCount, err := s.menuModel.SelectTotalCountByName(name)
	if err != nil {
		return nil, err
	}

	pgInfo := pg.NewPageInfo(int(totalCount))

	return page.NewPageData(menus, pgInfo), nil
}

func (s *storeMenuService) FindRecommendMenus(storeID string) (*response.ResponseStoreRecommendMenus, error) {
	foundStore, err := s.storeModel.SelectStoreByID(storeID)
	if err != nil {
		return nil, err
	}

	strMIDs := common.ConvertOBJIDsToStrings(foundStore.RecommendMenus)
	menus, err := s.menuModel.SelectMenusByIDs(storeID, strMIDs)
	if err != nil {
		return nil, err
	}

	return response.NewResponseStoreAndMenus(foundStore, menus), nil
}

func (s *storeMenuService) FindStore(storeID string) (*response.ResponseStore, error) {
	foundStore, err := s.storeModel.SelectStoreByID(storeID)
	if err != nil {
		return nil, err
	}

	return response.NewResponseStore(foundStore), nil
}

func (s *storeMenuService) FindStoresSortedPage(pg *request.RequestPage) (*page.PageData[any], error) {
	skip := util2.NewSkipNumber(pg.CurrentPage, pg.ContentCount)

	pageQuery := query.NewPageQuery(pg.Sort.Name, pg.Sort.Direction, skip, pg.ContentCount)

	receipts, err := s.storeModel.SelectSortLimitedStore(pageQuery)
	if err != nil {
		return nil, err
	}

	totalCount, err := s.storeModel.SelectTotalCount()
	if err != nil {
		return nil, err
	}

	pgInfo := pg.NewPageInfo(int(totalCount))

	return page.NewPageData(receipts, pgInfo), nil
}

// saveDeletedMenuBackupData 파일에 데이터를 쓰는데, 에러발생시 데이터를 로그로 남긴다
func (s *storeMenuService) saveDeletedMenuBackupData(deletedM *entity.Menu) {
	path := flag.Flags[flag.DatabaseFlag.Name]
	dbcfg := db.NewDBConfig(*path)
	err := db.WriteBackup(dbcfg.BackupPath, &deletedM)
	if err != nil {
		logger.AppLog.Error(err.Error())
		if m, err := json.Marshal(deletedM); err == nil {
			logger.AppLog.Error(string(m))
		}
	}
}
