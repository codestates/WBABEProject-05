package store

import (
	"encoding/json"
	error2 "github.com/codestates/WBABEProject-05/common/error"
	"github.com/codestates/WBABEProject-05/common/flag"
	"github.com/codestates/WBABEProject-05/config/db"
	"github.com/codestates/WBABEProject-05/logger"
	"github.com/codestates/WBABEProject-05/model/entity"
	"github.com/codestates/WBABEProject-05/model/menu"
	"github.com/codestates/WBABEProject-05/model/store"
	"github.com/codestates/WBABEProject-05/model/util"
	"github.com/codestates/WBABEProject-05/protocol/page"
	"github.com/codestates/WBABEProject-05/protocol/request"
	"github.com/codestates/WBABEProject-05/protocol/response"
	util2 "github.com/codestates/WBABEProject-05/service/util"
)

type storeMenuService struct {
	storeModel store.StoreModeler
	menuModel  menu.MenuModeler
}

var instance *storeMenuService

func NewStoreMenuService(
	sd store.StoreModeler,
	md menu.MenuModeler,
) *storeMenuService {
	if instance != nil {
		return instance
	}
	instance = &storeMenuService{
		storeModel: sd,
		menuModel:  md,
	}
	return instance
}

func (s *storeMenuService) RegisterStore(store *request.RequestPostStore) (string, error) {
	postStore, err := store.NewPostStore()
	if err != nil {
		return "", err
	}

	if _, err = s.storeModel.SelectStoreByPhone(postStore.StorePhone); err == nil {
		return "", &error2.DuplicatedDataError
	}

	savedId, err := s.storeModel.InsertStore(postStore)
	if err != nil {
		return "", err
	}

	return savedId, nil
}

func (s *storeMenuService) ModifyStore(storeID string, store *request.RequestPutStore) (int, error) {
	foundMenus, err := s.menuModel.SelectMenusByIDs(storeID, store.RecommendMenus)
	if err != nil || len(foundMenus) == len(store.RecommendMenus) {
		return 0, error2.AddNotRecommendMenusError.New()
	}

	putStore, err := store.NewPutStore(storeID)
	if err != nil {
		return 0, err
	}

	updateStore, err := s.storeModel.UpdateStore(putStore)
	if err != nil {
		return 0, err
	}

	return int(updateStore), err
}

func (s *storeMenuService) RegisterMenu(menu *request.RequestMenu) (string, error) {
	if foundMenu, _ := s.menuModel.SelectMenuByStoreIDAndName(menu.StoreID, menu.Name); foundMenu != nil {
		return "", error2.DuplicatedDataError.New()
	}

	newM, err := menu.NewMenu()
	if err != nil {
		return "", err
	}

	savedID, err := s.menuModel.InsertMenu(newM)
	if err != nil {
		return "", err
	}

	return savedID, nil
}

func (s *storeMenuService) ModifyMenu(menuId string, menu *request.RequestMenu) (int, error) {
	updateMenu, err := menu.NewUpdateMenu(menuId)
	if err != nil {
		return 0, err
	}

	cnt, err := s.menuModel.UpdateMenu(updateMenu)
	if err != nil {
		return 0, err
	}

	return int(cnt), nil
}

func (s *storeMenuService) DeleteMenuAndBackup(menuId string) (int, error) {
	deletedM, err := s.menuModel.SelectMenuByIDsAndDelete(menuId)
	if err != nil || deletedM == nil {
		return 0, err
	}

	go s.saveDeletedMenuBackupData(deletedM)

	return 1, nil
}

func (s *storeMenuService) FindMenusSortedPage(storeID string, pg *request.RequestPage) (*page.PageData[any], error) {
	skip := util2.NewSkipNumber(pg.CurrentPage, pg.ContentCount)

	menus, err := s.menuModel.SelectSortLimitedMenus(storeID, pg.Sort, skip, pg.ContentCount)
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

	menus, err := s.menuModel.SelectSortLimitedMenusByName(name, pg.Sort, skip, pg.ContentCount)
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

	strMIDs := util.ConvertOBJIDsToStrings(foundStore.RecommendMenus)
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

	receipts, err := s.storeModel.SelectSortLimitedStore(pg.Sort, skip, pg.ContentCount)
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
