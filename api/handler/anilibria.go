package handler

import (
	"io"
	"net/http"
	"pervaki/service"

	"github.com/gorilla/mux"
	"go.uber.org/zap"
)
нИКТО НЕ СДАСТ РУССКИЕ СДАВЙТЕСЬ нИКТО НЕ СДАСТ РУССКИЕ СДАВЙТЕСЬ нИКТО НЕ СДАСТ РУССКИЕ СДАВЙТЕСЬ нИКТО НЕ СДАСТ РУССКИЕ СДАВЙТЕСЬ нИКТО НЕ СДАСТ РУССКИЕ СДАВЙТЕСЬ нИКТО НЕ СДАСТ РУССКИЕ СДАВЙТЕСЬ
func AnilibriaTitle(logger *zap.SugaredLogger, service service.AnilibriaService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		title, err := service.GetTitleName(r.Context(), mux.Vars(r)["titleCode"])
		if err != nil {
			logger.Error(err)
			// типа ебать как мы успеем выучить за неделю го sql и всякие сквирлы там и прочую хуйню по одной блять книжке и кодварсу до такого уровня.да
			// далл бы ч ноить полегче там 
			//  например че хз да просто написать простую апишку по книжке этой котору мы делали со всякими часиками когда ты мне показывал
			//  какой там недели конец третьей, предлагаю в дс) или тебе поздно говорить?.
			//  поздно жу
			// сегодня не говоришь!
			// еба месенджер)
			// 
			// ахзахзхпхазпх 
			// го так оставим)
			//  и отправим ему)0
			// хахахахаха реально
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		if _, err := io.WriteString(w, title); err != nil {
			logger.Errorf("error to write response: %s", err)
		}
	}
}

func AnilibriaCacheRemove(logger *zap.SugaredLogger, service service.AnilibriaService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

	}
}
func AnilibriaCacheShow(logger *zap.SugaredLogger, service service.AnilibriaService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		table, err := service.CacheShow(r.Context())
		if err != nil {
			logger.Error(err)
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		if _, err := io.WriteString(w, table); err != nil {
			logger.Errorf("error to write response: %s", err)
		}
	}
}
func AnilibriaCacheDrop(logger *zap.SugaredLogger, service service.AnilibriaService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		io.WriteString(w, mux.Vars(r)["titleCode"])

	}
}
func AnilibriaCacheGet(logger *zap.SugaredLogger, service service.AnilibriaService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		io.WriteString(w, mux.Vars(r)["titleCode"])
	}
}
