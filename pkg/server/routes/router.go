package router

import (
	"net/http"

	"github.com/gorilla/mux"
)

type Route struct {
	Method  string
	Path    string
	Handler http.HandlerFunc
}

func SetupRouter(r *mux.Router) *mux.Router {
	var routes = [][]Route{UserRoutes}
	for _, router := range routes {
		for _, route := range router {
			r.HandleFunc(route.Path, route.Handler).Methods(route.Method)
		}
	}

	return r
}

func GetRouter() *mux.Router {
	r := mux.NewRouter()
	return SetupRouter(r)
}
