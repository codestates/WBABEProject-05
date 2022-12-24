package contorller

import (
	"github.com/codestates/WBABEProject-05/contorller/info"
	"github.com/codestates/WBABEProject-05/contorller/store"
	"github.com/codestates/WBABEProject-05/contorller/user"
	error2 "github.com/codestates/WBABEProject-05/protocol/error"
)

var instance *ginControl

type ginControl struct {
	infoControl  info.InfoController
	userControl  user.UserController
	storeControl store.StoreContoller
}

func GetInstance(
	inf info.InfoController,
	usr user.UserController,
	str store.StoreContoller,
) *ginControl {
	if instance != nil {
		return instance
	}
	instance = &ginControl{
		infoControl:  inf,
		userControl:  usr,
		storeControl: str,
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
