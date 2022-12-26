package main

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

var instance *App

type App struct {
	Server *http.Server
}

func NewApp() *App {
	instance = &App{}
	return instance
}

func (a *App) Run() {
	var g errgroup.Group
	a.Server = &http.Server{
		Addr:           config.AppServerConfig.Port,
		Handler:        router.Route.Handle(),
		ReadTimeout:    5 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	g.Go(func() error {
		if err := a.startServer(); err != nil {
			logger.AppLog.Error("Start Server fail,", err.Error())
			return err
		}
		return nil
	})
	a.graceExit()
	g.Wait()
}

func (a *App) startServer() error {
	pt := config.AppServerConfig.Port
	md := config.AppServerConfig.Mode
	stl := fmt.Sprintf("Start Server ... mode is %s and port is %s", md, pt)
	logger.AppLog.Info(stl)
	return a.Server.ListenAndServe()
}

func (a *App) graceExit() {
	quit := make(chan os.Signal)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	logger.AppLog.Warn("Shutdown Server ...")

	rt := 3 * time.Second
	ctx, cancel := context.WithTimeout(context.Background(), rt)
	defer cancel()

	if err := a.Server.Shutdown(ctx); err != nil {
		logger.AppLog.Error("Server Shutdown:", err)
	}
	select {
	case <-ctx.Done():
		tl := fmt.Sprintf("timeout of %s seconds.", rt.String())
		logger.AppLog.Info(tl)
	}
	logger.AppLog.Info("Server exiting")
}
