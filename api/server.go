package api

import (
	"context"
	"fmt"
	"net"
	"net/http"
	"pervaki/api/handler"
	"pervaki/config"
	"pervaki/lib/pctx"
	"pervaki/service"

	"github.com/gorilla/mux"
	"go.uber.org/zap"
)

func NewServer(
	ctxProvider pctx.DefaultProvider,
	logger *zap.SugaredLogger,
	settings config.Settings,
	animalService service.AnimalService,
	anilibriaService service.AnilibriaService,
) *http.Server {
	router := mux.NewRouter()

	router.HandleFunc("/ping", handler.Ping(logger)).Methods(http.MethodPost)
	router.HandleFunc("/animal/sound", handler.AnimalSound(logger, animalService)).Methods(http.MethodPost)
	router.HandleFunc("/anilibria/title/{titleCode}/name", handler.AnilibriaTitle(logger, anilibriaService)).Methods(http.MethodGet)
	router.HandleFunc("/anilibria/title/{titleCode}/remove", handler.AnilibriaCacheRemove(logger, anilibriaService)).Methods(http.MethodGet)
	router.HandleFunc("/anilibria/title/cache/drop", handler.AnilibriaCacheDrop(logger, anilibriaService)).Methods(http.MethodGet)
	router.HandleFunc("/anilibria/title/cache/show", handler.AnilibriaCacheShow(logger, anilibriaService)).Methods(http.MethodGet)
	router.HandleFunc("/anilibria/title/{titleCode}/cache", handler.AnilibriaCacheGet(logger, anilibriaService)).Methods(http.MethodGet)

	return &http.Server{
		Addr: fmt.Sprintf(":%d", settings.Port),
		BaseContext: func(listener net.Listener) context.Context {
			return ctxProvider()
		},
		Handler: router,
	}
}
