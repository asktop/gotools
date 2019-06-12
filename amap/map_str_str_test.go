package amap

import (
	"encoding/json"
	"fmt"
	"testing"
)

func TestStrStrMap_Iterator(t *testing.T) {
	data := map[string]string{}
	data["k1"] = "v1"
	data["k2"] = "v2"
	data["k3"] = "v3"
	data["k4"] = "v4"
	m := NewStrStrMapFrom(data)
	fmt.Println(m)
	m.Iterator(func(k string, v string) bool {
		if k == "k2" {
			return true
		}
		fmt.Println(k, ":", v)
		return true
	})
	rs, err := json.Marshal(m)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(string(rs))
		m2 := NewStrStrMap()
		err = json.Unmarshal(rs, m2)
		if err != nil {
			fmt.Println(err)
		} else {
			fmt.Println(m2)
		}
	}
}
