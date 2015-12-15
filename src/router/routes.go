package router

import (
	"controllers"
	"net/http"
	"github.com/gorilla/mux"
)

type Route struct {
	Name       string
	Method     string
	Pattern    string
	Controller controllers.Controller
	Router     *mux.Router
}

type Routes []Route

var routes = Routes{
	Route{
		"Index",
		"GET",
		"/",
		controllers.Index,
		nil,
	},
}

func (r Route) GetController() http.Handler {
	r.Controller.Route = r
	r.Controller.Router = r.Router
	return r.Controller.GetFunc()
}

func (r Route) GetName() string {
	return r.Name
}
