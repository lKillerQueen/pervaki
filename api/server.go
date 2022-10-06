package api

import (
	"context"
	"fmt"
	"github.com/gorilla/mux"
	"go.uber.org/zap"
	"net"
	"net/http"
	"pervaki/api/handler"
	"pervaki/config"
	"pervaki/lib/pctx"
	"pervaki/service"
)

func NewServer(
	ctxProvider pctx.DefaultProvider,
	logger *zap.SugaredLogger,
	settings config.Settings,
	animalService service.AnimalService,
) *http.Server {
	router := mux.NewRouter()

	router.HandleFunc("/ping", handler.Ping(logger)).Methods(http.MethodPost)
	router.HandleFunc("/animal/sound", handler.AnimalSound(logger, animalService)).Methods(http.MethodPost)

	return &http.Server{
		Addr: fmt.Sprintf(":%d", settings.Port),
		BaseContext: func(listener net.Listener) context.Context {
			return ctxProvider()
		},
		Handler: router,
	}
}
