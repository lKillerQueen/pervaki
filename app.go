package main

import (
	"context"
	"net/http"
	"pervaki/anilibria"
	"pervaki/database"
	"pervaki/database/titlerepo"

	"pervaki/api"
	"pervaki/config"
	"pervaki/lib/pctx"
	"pervaki/service"

	"go.uber.org/zap"
)

type App struct {
	logger   *zap.SugaredLogger
	settings config.Settings
	server   *http.Server
}

func NewApp(ctxProvider pctx.DefaultProvider, logger *zap.SugaredLogger, settings config.Settings) App {
	pgDb, err := database.NewPgx(settings.Postgres)
	if err != nil {
		panic(err)
	}

	err = database.UpMigrations(pgDb)
	if err != nil {
		panic(err)
	}

	var (
		cli = &http.Client{}

		anilibriaClient = anilibria.NewClient(logger, cli)

		anilibriaRepo = titlerepo.NewRepository(logger, pgDb)

		animalService    = service.NewAnimalService()
		anilibriaService = service.NewAnilibriaService(logger, anilibriaClient, anilibriaRepo)

		server = api.NewServer(ctxProvider, logger, settings, animalService, anilibriaService)
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
