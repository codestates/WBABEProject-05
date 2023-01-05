package store

import (
	"github.com/codestates/WBABEProject-05/model/common/query"
	"github.com/codestates/WBABEProject-05/model/entity"
)

var StoreModel StoreModeler

// TODO 디테일한 점이 많아 서비스 만들어야할듯하다
type StoreModeler interface {
	SelectStoreByID(storeId string) (*entity.Store, error)

	SelectStoreByIDAndUserID(storeID, userID string) (*entity.Store, error)

	SelectStoreByPhone(storePhone string) (*entity.Store, error)

	InsertStore(store *entity.Store) (string, error)

	UpdateStore(store *entity.Store) (int64, error)

	SelectSortLimitedStore(pageQuery *query.PageQuery) ([]*entity.Store, error)

	SelectTotalCount() (int64, error)
}
