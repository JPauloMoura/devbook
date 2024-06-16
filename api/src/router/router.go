package router

import (
	"api/src/router/routers"

	"github.com/gorilla/mux"
)

func New() *mux.Router {
	return routers.Config(mux.NewRouter())
}
