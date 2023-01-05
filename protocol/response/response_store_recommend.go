package response

import (
	"github.com/codestates/WBABEProject-05/model/entity"
)

type ResponseStoreRecommendMenus struct {
	StoreID        string           `json:"store_id"`
	UserID         string           `json:"user_id"`
	Name           string           `json:"name"`
	Address        *ResponseAddress `json:"address"`
	StorePhone     string           `json:"store_phone"`
	RecommendMenus []*ResponseMenu  `json:"recommend_menus"`
}

func FromStoreAndMenus(store *entity.Store, menus []*entity.Menu) *ResponseStoreRecommendMenus {
	return &ResponseStoreRecommendMenus{
		StoreID:        store.ID.Hex(),
		UserID:         store.UserID.Hex(),
		Name:           store.Name,
		Address:        FromAddr(store.Address),
		StorePhone:     store.StorePhone,
		RecommendMenus: FromMenus(menus),
	}
}
