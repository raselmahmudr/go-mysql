package router

import (
	"github.com/gorilla/mux"
	"hello/src/router/routes"
)

func New() *mux.Router {
	r := mux.NewRouter().StrictSlash(true)
	return routes.SetupRoutes(r)
}