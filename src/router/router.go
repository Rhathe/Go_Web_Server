package router

import (
	"github.com/gorilla/mux"
)

func Router() *mux.Router {

	router := mux.NewRouter().StrictSlash(true)
	for _, route := range routes {
		router.
			Methods(route.Method).
			Path(route.Pattern).
			Name(route.Name).
			Handler(route.GetController())
	}

	return router
}
