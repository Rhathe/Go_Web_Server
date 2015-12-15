package router

import (
	"controllers"
	"net/http"
	"logger"
)

type Route struct {
	Name       string
	Method     string
	Pattern    string
	Controller http.HandlerFunc
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

func (r Route) GetController() http.Handler {
	return logger.Logger(r.Controller, r.Name)
}
