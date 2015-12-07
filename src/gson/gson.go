package gson

import (
	"encoding/json"
	"reflect"
)

type Gson map[string]interface{}

func (g *Gson) ToJson() ([]byte, error) {
	return json.Marshal(g)
}

func (g *Gson) IsEqual(obj interface{}) bool {
	j1, _ := g.ToJson()
	var m1, m2 map[string]interface{}
	json.Unmarshal(j1, &m1)

	j2, _ := json.Marshal(obj)
	json.Unmarshal(j2, &m2)

	return reflect.DeepEqual(m1, m2)
}
