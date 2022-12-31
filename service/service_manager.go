package service

import (
	"github.com/codestates/WBABEProject-05/model/menu"
	"github.com/codestates/WBABEProject-05/model/receipt"
	store2 "github.com/codestates/WBABEProject-05/model/store"
	user2 "github.com/codestates/WBABEProject-05/model/user"
	"github.com/codestates/WBABEProject-05/service/login"
	"github.com/codestates/WBABEProject-05/service/order"
	"github.com/codestates/WBABEProject-05/service/store"
)

func InjectServicesDependency() {
	store.StoreMenuService = store.NewStoreMenuService(store2.StoreModel, menu.MenuModel)
	order.OrderRecordService = order.NewOrderRecordService(receipt.ReceiptModel)
	login.UserService = login.NewUserService(user2.UserModel)
}

func SetServicer() {
	Service = NewService()
}
