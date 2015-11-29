package router

import (
	"controllers"
	"net/http"
)

type Route struct {
	Name        string
	Method      string
	Pattern     string
	Controller  http.HandlerFunc
}

type Routes []Route

var routes = Routes{
	Route{
		"Index",
		"GET",
		"/",
		controllers.Index,
	},
}
