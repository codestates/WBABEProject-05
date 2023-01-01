package contorller

import (
	"github.com/codestates/WBABEProject-05/contorller/info"
	"github.com/codestates/WBABEProject-05/contorller/order"
	"github.com/codestates/WBABEProject-05/contorller/review"
	"github.com/codestates/WBABEProject-05/contorller/store"
	"github.com/codestates/WBABEProject-05/contorller/user"
	"github.com/codestates/WBABEProject-05/service/customer"
	"github.com/codestates/WBABEProject-05/service/login"
	order2 "github.com/codestates/WBABEProject-05/service/order"
	store2 "github.com/codestates/WBABEProject-05/service/store"
)

func InjectControllerDependency() {
	info.InfoControl = info.NewInfoControl()
	order.OrderRecordControl = order.NewOrderRecordControl(order2.OrderRecordService)
	store.StoreControl = store.NewStoreControl(store2.StoreMenuService)
	user.UserControl = user.NewUserControl(login.UserService)
	review.MenuReviewControl = review.NewMenuReviewControl(customer.MenuReviewService)
}
