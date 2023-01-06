package store

import (
	"github.com/codestates/WBABEProject-05/model/entity"
	"github.com/codestates/WBABEProject-05/model/query"
)

var StoreModel StoreModeler

type StoreModeler interface {
	SelectStoreByID(storeId string) (*entity.Store, error)

	SelectStoreByIDAndUserID(storeID, userID string) (*entity.Store, error)

	SelectStoreByPhone(storePhone string) (*entity.Store, error)

	InsertStore(store *entity.Store) (string, error)

	UpdateStore(store *entity.Store) (int64, error)

	SelectSortLimitedStore(pageQuery *query.PageQuery) ([]*entity.Store, error)

	SelectTotalCount() (int64, error)

	UpdatePullRecommendMenu(storeID, menuID string) (int64, error)
}
