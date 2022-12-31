package store

import (
	"encoding/json"
	"github.com/codestates/WBABEProject-05/common/flag"
	"github.com/codestates/WBABEProject-05/common/util"
	"github.com/codestates/WBABEProject-05/config/db"
	"github.com/codestates/WBABEProject-05/logger"
	"github.com/codestates/WBABEProject-05/model/entity"
	"github.com/codestates/WBABEProject-05/model/menu"
	"github.com/codestates/WBABEProject-05/model/store"
	error2 "github.com/codestates/WBABEProject-05/protocol/error"
	"github.com/codestates/WBABEProject-05/protocol/page"
	"github.com/codestates/WBABEProject-05/protocol/request"
	"github.com/codestates/WBABEProject-05/protocol/response"
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

	_, err = s.storeModel.SelectStoreByPhone(postStore.StorePhone)
	if err == nil {
		return "", &error2.DuplicatedDataError
	}

	savedId, err := s.storeModel.InsertStore(postStore)
	if err != nil {
		return "", err
	}
	return savedId, nil
}

func (s *storeMenuService) ModifyStore(storeID string, store *request.RequestPutStore) (int, error) {
	putStore, err := store.NewPutStore(storeID)
	if err != nil {
		return 0, err
	}
	updateStore, err := s.storeModel.UpdateStore(putStore)
	if err != nil {
		return 0, err
	}
	return updateStore, err
}

func (s *storeMenuService) RegisterMenu(menu *request.RequestMenu) (string, error) {
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
	return cnt, nil
}

func (s *storeMenuService) DeleteMenuAndBackup(menuId string) (int, error) {
	deletedM, err := s.menuModel.SelectMenuByIdsAndDelete(menuId)
	if err != nil || deletedM == nil {
		return 0, err
	}

	path := flag.Flags[flag.DatabaseFlag.Name]
	dbcfg := db.NewDbConfig(*path)
	err = db.WriteBackup(dbcfg.BackupPath, &deletedM)
	if err != nil {
		logger.AppLog.Error(err.Error())
		if m, err := json.Marshal(deletedM); err == nil {
			logger.AppLog.Error(string(m))
		}
		return 0, err
	}

	return 1, nil
}

func (s *storeMenuService) FindMenusSortedPage(storeID string, pg *request.RequestPage) (*page.PageData[any], error) {
	skip := pg.CurrentPage * pg.ContentCount
	if skip > 0 {
		skip--
	}

	menus, err := s.menuModel.SelectSortLimitedMenus(storeID, pg.Sort, skip, pg.ContentCount)
	if err != nil {
		return nil, err
	}

	totalCount, err := s.menuModel.SelectTotalCount(storeID)
	if err != nil {
		return nil, err
	}

	pgInfo := pg.NewPageInfo(totalCount)

	return page.NewPageData(menus, pgInfo), nil
}

func (s *storeMenuService) FindRecommendMenus(storeID string) (*response.ResponseStore, error) {
	foundStore, err := s.storeModel.SelectStoreByID(storeID)
	if err != nil {
		return nil, err
	}

	strMIDs := util.ConvertObjIDsToStrings(foundStore.RecommendMenus)
	menus, err := s.menuModel.SelectMenusByIDs(storeID, strMIDs)
	if err != nil {
		return nil, err
	}

	responseStore := response.NewResponseStore(foundStore, menus)
	return responseStore, nil
}

func (s *storeMenuService) FindStore(storeID string) (*entity.Store, error) {
	foundStore, err := s.storeModel.SelectStoreByID(storeID)
	if err != nil {
		return nil, err
	}
	return foundStore, nil
}
