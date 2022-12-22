package util

import (
	"github.com/codestates/WBABEProject-05/config"
	"github.com/codestates/WBABEProject-05/contorller"
	"github.com/codestates/WBABEProject-05/router"
)

const (
	Name        = "WBA-띵동주문이요"
	Description = "온라인 주문 시스템"
	Author      = "Hooneats"
)

var instance *App

type App struct {
	Name        string
	Description string
	Author      string
	Config      *config.Config
	Router      router.Router
	Controller  contorller.Controller
}

func LoadApp() *App {
	instance = &App{
		Name:        Name,
		Description: Description,
		Author:      Author,
	}
	return instance
}

func (a *App) SetConfig(config *config.Config) {
	a.Config = config
}

func (a *App) SetRouter(router router.Router) {
	a.Router = router
}

func (a *App) SetController(controller contorller.Controller) {
	a.Controller = controller
}
