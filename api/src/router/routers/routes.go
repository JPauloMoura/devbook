package routers

import (
	"net/http"

	"github.com/gorilla/mux"
)

// Router representa a estutura de uma rota rest
type Router struct {
	Uri          string
	Method       string
	Function     func(http.ResponseWriter, *http.Request)
	AuthRequired bool
}

func Config(r *mux.Router) *mux.Router {
	for _, router := range userRouters {
		r.HandleFunc(router.Uri, router.Function).Methods(router.Method)
	}

	return r
}
