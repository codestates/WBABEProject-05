package validator

import (
	"github.com/codestates/WBABEProject-05/common/enum"
	error2 "github.com/codestates/WBABEProject-05/common/error"
	"github.com/codestates/WBABEProject-05/model/menu"
	store2 "github.com/codestates/WBABEProject-05/model/store"
	"github.com/codestates/WBABEProject-05/model/user"
)

// CheckRoleIsStore 사용자가 존재할경우 역할이 store 인지 확인
func CheckRoleIsStore(userID string) error {
	selectUser, err := user.UserModel.SelectUser(userID)
	if err != nil {
		return error2.UserNotFoundErr
	}
	if selectUser.Role != enum.StoreRole {
		return error2.UnauthorizedError
	}
	return nil
}

// CheckRoleIsCustomer 사용자가 존재할경우 역할이 store 인지 확인
func CheckRoleIsCustomer(userID string) error {
	selectUser, err := user.UserModel.SelectUser(userID)
	if err != nil {
		return error2.UserNotFoundErr
	}
	if selectUser.Role != enum.CustomerRole {
		return error2.UnauthorizedError
	}
	return nil
}

// CheckExistsStore 가게가 존재하지 않으면 DataNotFoundError
func CheckExistsStore(storeID string) error {
	if _, err := store2.StoreModel.SelectStoreByID(storeID); err != nil {
		return error2.DataNotFoundError
	}
	return nil
}

// CheckAlreadyExistsMenuByName 메뉴가 이미 존재하면 DuplicatedDataError
func CheckAlreadyExistsMenuByName(storeID, menuName string) error {
	if found, _ := menu.MenuModel.SelectMenuByStoreIDAndName(storeID, menuName); found != nil {
		return error2.DuplicatedDataError
	}
	return nil
}

// CheckExistsMenu 메뉴가 존재하지 않으면 DataNotFoundError
func CheckExistsMenu(storeID, menuID string) error {
	if err := CheckExistsMenus(storeID, []string{menuID}); err != nil {
		return err
	}
	return nil
}

// CheckExistsMenus 입력받은 메뉴들중 하나라도 존재하지 않으면 DataNotFoundError
func CheckExistsMenus(storeID string, menuIDs []string) error {
	if foundMenus, err := menu.MenuModel.SelectMenusByIDs(storeID, menuIDs); err != nil || len(foundMenus) != len(menuIDs) {
		return error2.DataNotFoundError
	}
	return nil
}

// CheckStoreUser 가게의 user 즉 주인이 맞지 않으면 DataNotFoundError
func CheckStoreUser(storeID, userID string) error {
	if _, err := store2.StoreModel.SelectStoreByIDAndUserID(storeID, userID); err != nil {
		return error2.DataNotFoundError
	}
	return nil
}
