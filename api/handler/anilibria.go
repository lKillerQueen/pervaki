package handler

import (
	"io"
	"net/http"
	"pervaki/service"

	"github.com/gorilla/mux"
	"go.uber.org/zap"

	"encoding/json"
)

func AnilibriaTitle(logger *zap.SugaredLogger, service service.AnilibriaService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		title, err := service.GetTitleName(r.Context(), mux.Vars(r)["titleCode"])
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

func AnilibriaAll(logger *zap.SugaredLogger, service service.AnilibriaService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		rows, err := service.GetAll(r.Context())
		if err != nil {
			logger.Error(err)
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		JsonBytes, err := json.Marshal(rows)

		if _, err := io.WriteString(w, string(JsonBytes)); err != nil {
			logger.Errorf("error to write response: %s", err)
		}
	}
}
