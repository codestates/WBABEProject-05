package contorller

import (
	"github.com/codestates/WBABEProject-05/contorller/info"
)

var instance *ginControl

type ginControl struct {
	health info.InfoController
}

func GetInstance() *ginControl {
	if instance != nil {
		return instance
	}
	instance = &ginControl{
		health: info.GetInstance(),
	}
	return instance
}

func (g *ginControl) GetInfoControl() info.InfoController {
	if g.health != nil {
		return g.health
	}
	g.health = info.GetInstance()
	return g.health
}
