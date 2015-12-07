package api

import (
	"encoding/json"
	"net/http"
)

type Response struct {
	Data  interface{} `json:"data"`
	Error `json:"error"`
	Info  `json:""`
	Page  Page `json:"page_info"`
	w     http.ResponseWriter
	r     *http.Request
}

type Info struct {
	Revision Revision `json:"revision"`
	Error    bool     `json:"error"`
	Env      string   `json:"env"`
	Status   int      `json:"status"`
}

type Revision struct {
	Hash      string `json:"hash"`
	Timestamp string `json:"timestamp"`
	Author    string `json:"author"`
}

type Error struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

type Page struct {
	ItemCount  int    `json:"item_count"`
	TotalPages int    `json:"total_pages"`
	Current    string `json:"current"`
	Previous   string `json:"previous"`
	Next       string `json:"next"`
}

func NewResponse(w http.ResponseWriter, r *http.Request) *Response {
	return &Response{
		Data: []string{},
		Info: Info{
			Status: 200,
		},
		w: w,
		r: r,
	}
}

func (r *Response) SetData(data interface{}) {
	r.Data = data
}

func (r *Response) JSON() {
	r.w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	r.w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(r.w).Encode(r); err != nil {
		panic(err)
	}
}
