package util

import (
	"context"
	"fmt"
	"github.com/codestates/WBABEProject-05/config"
	"github.com/codestates/WBABEProject-05/config/log"
	"github.com/codestates/WBABEProject-05/logger"
	"github.com/codestates/WBABEProject-05/router"
	"golang.org/x/sync/errgroup"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

var instance *App

// TODO App 의 구조를 및 설정을 한번에 파악할 수 있는 App struct 로 주석 추가해 주도록하자
type App struct {
	Config *config.Config `json:"_"`
	Logger logger.Logger
	Router router.Router
	Server *http.Server
}

func NewApp() *App {
	instance = &App{}
	return instance
}

func (a *App) LoadConfig() {
	path := Flags[ConfigFlag.Name]
	a.Config = config.NewConfig(*path)
}

func (a *App) LoadLogger() {
	path := Flags[LogConfigFlag.Name]
	lcfg := log.NewLogConfig(*path)
	Logger := logger.InitLogger(lcfg)
	a.Logger = Logger
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
		if err := a.startServer(); err != nil {
			a.Logger.Error("Start Server fail,", err.Error())
			return err
		}
		return nil
	})

	a.graceExit()

	g.Wait()
}

func (a *App) startServer() error {
	pt := a.Config.Server.Port
	md := a.Config.Server.Mode
	stl := fmt.Sprintf("Start Server ... mode is %s and port is %s", md, pt)
	a.Logger.Info(stl)
	return a.Server.ListenAndServe()
}

func (a *App) graceExit() {
	quit := make(chan os.Signal)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	a.Logger.Warn("Shutdown Server ...")

	rt := 3 * time.Second
	ctx, cancel := context.WithTimeout(context.Background(), rt)
	defer cancel()

	if err := a.Server.Shutdown(ctx); err != nil {
		a.Logger.Error("Server Shutdown:", err)
	}
	select {
	case <-ctx.Done():
		tl := fmt.Sprintf("timeout of %s seconds.", rt.String())
		a.Logger.Info(tl)
	}
	a.Logger.Info("Server exiting")
}
