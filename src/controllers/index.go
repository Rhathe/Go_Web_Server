package controllers

import (
	"api"
	"net/http"
)

func Index(w http.ResponseWriter, r *http.Request) {
	resp := api.NewResponse(w, r)
	resp.JSON()
}
