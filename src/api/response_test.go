package api

import (
	. "gson"
	"testing"
	"github.com/stretchr/testify/assert"
	"net/http/httptest"
)

var w = httptest.NewRecorder()

func TestNewResponse(t *testing.T) {
	assert := assert.New(t)
	r := NewResponse(w, nil)
	g := Gson{
		"data": []int{},
		"error": Gson{
			"code": 0,
			"message": "",
		},
		"revision": Gson{
			"hash": "",
			"timestamp": "",
			"author": "",
		},
		"env": "",
		"status": 200,
		"page_info": Gson{
			"item_count": 0,
			"total_pages": 0,
			"current": "",
			"next": "",
			"previous": "",
		},
	}

	assert.Equal(g.IsEqual(r), true, "they should be equal")
}

func TestSetData(t *testing.T) {
	assert := assert.New(t)
	r := Response{}
	r.SetData("test")
	assert.Equal(r.Data, "test", "they should be equal")
}
