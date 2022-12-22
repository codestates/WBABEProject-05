package util

import (
	"context"
	"fmt"
	"github.com/codestates/WBABEProject-05/config"
	"github.com/codestates/WBABEProject-05/logger"
	"github.com/codestates/WBABEProject-05/router"
	"golang.org/x/sync/errgroup"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
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
	Server      *http.Server
}

func LoadApp() *App {
	instance = &App{
		Name:        Name,
		Description: Description,
		Author:      Author,
	}
	return instance
}

func GetApp() *App {
	if instance != nil {
		return instance
	}
	return LoadApp()
}

func (a *App) SetConfig(cfg *config.Config) {
	a.Config = cfg
}

func (a *App) SetRouter(rt router.Router) {
	a.Router = rt
}

func (a *App) Run() {
	var g errgroup.Group
	a.Server = &http.Server{
		Addr:           a.Config.Server.Port,
		Handler:        a.Router.Handle(),
		ReadTimeout:    5 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	g.Go(func() error {
		return a.startServer()
	})

	a.graceExit()

	g.Wait()
}

func (a *App) startServer() error {
	pt := a.Config.Server.Port
	md := a.Config.Server.Mode
	stl := fmt.Sprintf("Start Server ... mode is %s and port is %s", md, pt)
	logger.Info(stl)
	return a.Server.ListenAndServe()
}

func (a *App) graceExit() {
	quit := make(chan os.Signal)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	logger.Warn("Shutdown Server ...")

	rt := 3 * time.Second
	ctx, cancel := context.WithTimeout(context.Background(), rt)
	defer cancel()

	if err := a.Server.Shutdown(ctx); err != nil {
		logger.Error("Server Shutdown:", err)
	}
	select {
	case <-ctx.Done():
		tl := fmt.Sprintf("timeout of %s seconds.", rt.String())
		logger.Info(tl)
	}
	logger.Info("Server exiting")
}
