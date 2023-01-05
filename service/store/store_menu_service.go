package store

import (
	"github.com/codestates/WBABEProject-05/model/menu"
	"github.com/codestates/WBABEProject-05/model/store"
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
