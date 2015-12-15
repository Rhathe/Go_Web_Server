package controllers

import (
	"net/http"
	"logger"
	"github.com/gorilla/mux"
)

type Route interface {
	GetName() string
}

type Controller struct {
	Func   http.HandlerFunc
	Route  Route
	Router *mux.Router
}

func (c Controller) GetFunc() http.Handler {
	return logger.Logger(c.Func, c.Route.GetName())
}
