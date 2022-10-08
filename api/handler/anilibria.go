package handler

import (
	"github.com/gorilla/mux"
	"go.uber.org/zap"
	"io"
	"net/http"
	"pervaki/model"
	"pervaki/service"
)

func AnilibriaTitle(logger *zap.SugaredLogger, service service.AnilibriaService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		titleCode := mux.Vars(r)["titleCode"]

		title, err := service.GetTitle(r.Context(), model.GetTitleFilter{Code: titleCode})
		if err != nil {
			logger.Error(err)
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		if _, err := io.WriteString(w, title); err != nil {
			logger.Errorf("error to write response: %s", err)
		}
	}
}
