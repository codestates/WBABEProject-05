package service

import (
	"github.com/codestates/WBABEProject-05/model/menu"
	"github.com/codestates/WBABEProject-05/model/receipt"
	"github.com/codestates/WBABEProject-05/model/review"
	store2 "github.com/codestates/WBABEProject-05/model/store"
	user2 "github.com/codestates/WBABEProject-05/model/user"
	"github.com/codestates/WBABEProject-05/service/customer"
	"github.com/codestates/WBABEProject-05/service/order"
	"github.com/codestates/WBABEProject-05/service/store"
	"github.com/codestates/WBABEProject-05/service/user"
)

func InjectServicesDependency() {
	store.StoreMenuService = store.NewStoreMenuService(store2.StoreModel, menu.MenuModel)
	order.OrderRecordService = order.NewOrderRecordService(receipt.ReceiptModel, menu.MenuModel, user2.UserModel)
	customer.MenuReviewService = customer.NewMenuReviewService(review.ReviewModel, menu.MenuModel, receipt.ReceiptModel)
	user.UserService = user.NewUserService(user2.UserModel)
}
