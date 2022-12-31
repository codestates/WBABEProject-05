package store

import (
	"github.com/codestates/WBABEProject-05/model/entity"
	"github.com/codestates/WBABEProject-05/protocol/page"
)

var StoreModel StoreModeler

// TODO 디테일한 점이 많아 서비스 만들어야할듯하다
type StoreModeler interface {
	SelectStoreByID(storeId string) (*entity.Store, error)
	SelectStoreByPhone(storePhone string) (*entity.Store, error)
	InsertStore(store *entity.Store) (string, error)
	UpdateStore(store *entity.Store) (int, error)
	SelectSortLimitedStore(sort *page.Sort, skip int, limit int) ([]*entity.Store, error)
	SelectTotalCount() (int, error)
}
