package contorller

import (
	"github.com/codestates/WBABEProject-05/contorller/info"
	"github.com/codestates/WBABEProject-05/contorller/order"
	"github.com/codestates/WBABEProject-05/contorller/store"
	"github.com/codestates/WBABEProject-05/contorller/user"
)

type Controller interface {
	InfoControl() (info.InfoController, error)
	UserControl() (user.UserController, error)
	StoreControl() (store.StoreContoller, error)
	OrderControl() (order.OrderController, error)
}
