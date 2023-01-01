package response

import (
	"github.com/codestates/WBABEProject-05/model/entity"
	"github.com/codestates/WBABEProject-05/protocol/request"
)

type ResponseStore struct {
	StoreID        string                  `json:"store_id"`
	UserID         string                  `json:"user_id"`
	Name           string                  `json:"name"`
	Address        *request.RequestAddress `json:"address"`
	StorePhone     string                  `json:"store_phone"`
	RecommendMenus []*ResponseMenu         `bson:"recommend_menus"`
}

func NewResponseStore(store *entity.Store, menus []*entity.Menu) *ResponseStore {
	return &ResponseStore{
		StoreID:        store.ID.Hex(),
		UserID:         store.UserID.Hex(),
		Name:           store.Name,
		Address:        request.FromAddress(store.Address),
		StorePhone:     store.StorePhone,
		RecommendMenus: FromMenus(menus),
	}
}
