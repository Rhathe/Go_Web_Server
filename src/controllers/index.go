package controllers

import (
	. "gson"
	"api"
	"fmt"
	"net/http"
)

func IndexFunc(w http.ResponseWriter, r *http.Request) {
	resp := api.NewResponse(w, r)
	resp.SetData(Gson{
		"users": fmt.Sprintf("%s://%s/users", resp.URL().Scheme, resp.URL().Host),
	})
	resp.JSON()
}

var Index = Controller{
	IndexFunc,
	nil,
	nil,
}
