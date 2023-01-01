package store

import (
	"github.com/codestates/WBABEProject-05/protocol/page"
	"github.com/codestates/WBABEProject-05/protocol/request"
	"github.com/codestates/WBABEProject-05/protocol/response"
)

var StoreMenuService StoreMenuServicer

type StoreMenuServicer interface {
	RegisterStore(store *request.RequestPostStore) (string, error)

	ModifyStore(storeID string, store *request.RequestPutStore) (int, error)

	RegisterMenu(menu *request.RequestMenu) (string, error)

	ModifyMenu(menuId string, menu *request.RequestMenu) (int, error)

	DeleteMenuAndBackup(menuId string) (int, error)

	FindMenusSortedPage(storeID string, page *request.RequestPage) (*page.PageData[any], error)

	FindRecommendMenus(storeID string) (*response.ResponseStoreRecommendMenus, error)

	FindStore(storeId string) (*response.ResponseStore, error)

	FindStoresSortedPage(page *request.RequestPage) (*page.PageData[any], error)
}
