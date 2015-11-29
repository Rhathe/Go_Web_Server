package router

import (
	"github.com/gorilla/mux"
	"logger"
)

func Router() *mux.Router {

	router := mux.NewRouter().StrictSlash(true)
	for _, route := range routes {
		controller := logger.Logger(route.Controller, route.Name)

		router.
			Methods(route.Method).
			Path(route.Pattern).
			Name(route.Name).
			Handler(controller)
	}

	return router
}
