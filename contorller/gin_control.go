package contorller

import (
	error2 "github.com/codestates/WBABEProject-05/common/error"
	"github.com/codestates/WBABEProject-05/contorller/info"
	"github.com/codestates/WBABEProject-05/contorller/order"
	"github.com/codestates/WBABEProject-05/contorller/store"
	"github.com/codestates/WBABEProject-05/contorller/user"
)

var instance *ginControl

type ginControl struct {
	infoControl  info.InfoController
	userControl  user.UserController
	storeControl store.StoreContoller
	orderControl order.OrderController
}

func GetInstance(
	inf info.InfoController,
	usr user.UserController,
	str store.StoreContoller,
	od order.OrderController,
) *ginControl {
	if instance != nil {
		return instance
	}
	instance = &ginControl{
		infoControl:  inf,
		userControl:  usr,
		storeControl: str,
		orderControl: od,
	}
	return instance
}

func (g *ginControl) InfoControl() (info.InfoController, error) {
	if g.infoControl != nil {
		return g.infoControl, nil
	}
	// TODO logger
	return nil, error2.NonInjectedError.Err
}

func (g *ginControl) UserControl() (user.UserController, error) {
	if g.userControl != nil {
		return g.userControl, nil
	}
	// TODO logger
	return nil, error2.NonInjectedError.Err
}
func (g *ginControl) StoreControl() (store.StoreContoller, error) {
	if g.storeControl != nil {
		return g.storeControl, nil
	}
	// TODO logger
	return nil, error2.NonInjectedError.Err
}

func (g *ginControl) OrderControl() (order.OrderController, error) {
	if g.orderControl != nil {
		return g.orderControl, nil
	}
	// TODO logger
	return nil, error2.NonInjectedError.Err
}
