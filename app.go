package main

import (
	"context"
	"go.uber.org/zap"
	"net/http"
	"pervaki/api"
	"pervaki/config"
	"pervaki/lib/pctx"
	"pervaki/service"
)

type App struct {
	logger   *zap.SugaredLogger
	settings config.Settings
	server   *http.Server
}

func NewApp(ctxProvider pctx.DefaultProvider, logger *zap.SugaredLogger, settings config.Settings) App {
	var (
		cli = http.Client{}

		animalService    = service.NewAnimalService()
		anilibriaService = service.NewAnilibriaService(logger, &cli)
		server           = api.NewServer(ctxProvider, logger, settings, animalService, anilibriaService)
	)

	return App{
		logger:   logger,
		settings: settings,
		server:   server,
	}
}

func (a App) Run() {
	go func() {
		_ = a.server.ListenAndServe()
	}()
	a.logger.Debugf("HTTP server started on %d", a.settings.Port)
}

func (a App) Stop(ctx context.Context) {
	_ = a.server.Shutdown(ctx)
	a.logger.Debugf("HTTP server stopped")
}
