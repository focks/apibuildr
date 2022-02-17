package apibuildr

import (
	"github.com/gorilla/mux"
	"net/http"
)

type ApiHandler struct {
	Name       string
	Path       string
	Method     string
	HandleFunc http.HandlerFunc
}

func (api ApiHandler) RegisterToRouter(router *mux.Router) {
	router.Methods(api.Method).
		Path(api.Path).
		Name(api.Name).
		Handler(api.HandleFunc)
}
