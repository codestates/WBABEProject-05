package store

import (
	"github.com/codestates/WBABEProject-05/common/enum"
	error2 "github.com/codestates/WBABEProject-05/common/error"
	"github.com/codestates/WBABEProject-05/logger"
	"github.com/codestates/WBABEProject-05/model/common"
	"github.com/codestates/WBABEProject-05/protocol/request"
	"github.com/codestates/WBABEProject-05/service/validator"
)

func (s *storeMenuService) RegisterStore(store *request.RequestPostStore) (string, error) {
	if err := validator.CheckRoleIsStore(store.UserID); err != nil {
		return enum.BlankSTR, err
	}

	postStore, err := store.ToPostStore()
	if err != nil {
		return enum.BlankSTR, err
	}

	if _, err = s.storeModel.SelectStoreByPhone(postStore.StorePhone); err == nil {
		logger.AppLog.Error(err.Error())
		return enum.BlankSTR, error2.DuplicatedDataError
	}

	savedId, err := s.storeModel.InsertStore(postStore)
	if err != nil {
		return enum.BlankSTR, err
	}

	return savedId, nil
}

func (s *storeMenuService) ModifyStore(storeID string, store *request.RequestPutStore) (int, error) {
	if err := validator.CheckRoleIsStore(store.UserID); err != nil {
		return 0, err
	}

	if err := validator.CheckExistsMenus(storeID, store.RecommendMenus); err != nil {
		return 0, err
	}

	putStore, err := store.ToPutStore(storeID)
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
	if err := validator.CheckRoleIsStore(menu.UserID); err != nil {
		return enum.BlankSTR, err
	}

	if err := validator.CheckAlreadyExistsMenuByName(menu.StoreID, menu.Name); err != nil {
		return enum.BlankSTR, err
	}

	newM, err := menu.ToPostMenu()
	if err != nil {
		return enum.BlankSTR, err
	}

	savedID, err := s.menuModel.InsertMenu(newM)
	if err != nil {
		return enum.BlankSTR, err
	}

	return savedID, nil
}

func (s *storeMenuService) ModifyMenu(menuID string, menu *request.RequestMenu) (int, error) {
	if err := validator.CheckRoleIsStore(menu.UserID); err != nil {
		return 0, err
	}

	if err := validator.CheckExistsMenu(menu.StoreID, menuID); err != nil {
		return 0, err
	}

	updateMenu, err := menu.ToPutMenu(menuID)
	if err != nil {
		return 0, err
	}

	cnt, err := s.menuModel.UpdateMenu(updateMenu)
	if err != nil {
		return 0, err
	}

	return int(cnt), nil
}

func (s *storeMenuService) DeleteMenuAndBackup(menu *request.RequestDeleteMenu) (int, error) {
	if err := validator.CheckRoleIsStore(menu.UserID); err != nil {
		return 0, err
	}

	foundM, err := s.menuModel.SelectMenuByID(menu.MenuID)
	if err != nil {
		return 0, err
	}

	if common.ConvertOBJIDToString(foundM.StoreID) != menu.StoreID {
		return 0, error2.UnauthorizedError.New()
	}

	deletedM, err := s.menuModel.SelectMenuByIDsAndDelete(menu.MenuID)
	if err != nil || deletedM == nil {
		return 0, err
	}

	go s.saveDeletedMenuBackupData(deletedM)

	return 1, nil
}
