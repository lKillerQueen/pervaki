package handler

import (
	"encoding/json"
	"go.uber.org/zap"
	"net/http"
	"pervaki/model"
	"pervaki/service"
)

func AnimalSound(logger *zap.SugaredLogger, animalService service.AnimalService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var animal model.Animal
		if err := json.NewDecoder(r.Body).Decode(&animal); err != nil {
			logger.Error(err)
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		sound, err := animalService.SoundByZoo(animal.Name)
		if err != nil {
			logger.Error(err)
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		err = json.NewEncoder(w).Encode(sound)
		if err != nil {
			logger.Error(err)
			w.WriteHeader(http.StatusBadRequest)
			return
		}
	}
}
