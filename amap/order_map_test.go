package amap

import (
	"encoding/json"
	"testing"
)

func TestNewFrom(t *testing.T) {
	data := map[string]string{}
	data["k1"] = "v1"
	data["k3"] = "v3"
	data["k2"] = "v2"
	odata := NewOrderMapFrom(data)
	t.Log(odata)
	odata.SortKey()
	t.Log(odata)
}

func TestOrderMapRange(t *testing.T) {
	o := NewOrderMap()
	o.Set("number", 3)
	o.Set("string", "x")
	o.Set("c", "c")
	o.Set("a", "a")
	o.Set("b", "b")
	for _, k := range o.Keys() {
		v := o.Get(k)
		t.Log(k, v)
	}
}

func TestOrderMap(t *testing.T) {
	o := NewOrderMap()
	// number
	o.Set("number", 3)
	// string
	o.Set("string", "x")
	// string slice
	o.Set("strings", []string{
		"t",
		"u",
	})
	// slice
	o.Set("slice", []interface{}{
		"1",
		1,
	})
	// orderedmap
	n := NewOrderMap()
	n.Set("e", 1)
	n.Set("a", 2)
	o.Set("orderedmap", n)

	// Keys method
	t.Log(o)
	for _, key := range o.Keys() {
		v, _ := o.Search(key)
		t.Log(key, ":", v)
	}

	// delete
	o.Remove("strings")
	t.Log(o)
}

func TestOrderedMap_MarshalJSON(t *testing.T) {
	o := NewOrderMap()
	// number
	o.Set("number", 3)
	// string
	o.Set("string", "x")
	// string slice
	o.Set("strings", []string{
		"t",
		"u",
	})
	// slice
	o.Set("slice", []interface{}{
		"1",
		1,
	})
	// orderedmap
	n := NewOrderMap()
	n.Set("e", 1)
	n.Set("a", 2)
	o.Set("orderedmap", n)

	t.Log(o)
	// convert to json
	b, err := json.Marshal(o)
	if err != nil {
		t.Error("Marshalling json", err)
	}
	s := string(b)
	t.Log(s)
}

func TestOrderedMap_UnmarshalJSON(t *testing.T) {
	s := `{"number":3,"string":"x","strings":["t","u"],"slice":["1",1],"orderedmap":{"e":1,"a":2}}`
	o := NewOrderMap()
	err := json.Unmarshal([]byte(s), &o)
	if err != nil {
		t.Error("JSON Unmarshal error", err)
	}

	t.Log(o)
	for _, key := range o.Keys() {
		v, _ := o.Search(key)
		t.Log(key, ":", v)
	}
}

func TestOrderedMap_SortKey(t *testing.T) {
	o := NewOrderMap()
	o.Set("b", "b")
	o.Set("d", "3")
	o.Set("a", "a")
	o.Set("c", "c")
	o.Set("f", "1")
	o.Set("e", 2)
	t.Log(o)
	for _, key := range o.Keys() {
		v, _ := o.Search(key)
		t.Log(key, ":", v)
	}

	o.SortKey()

	t.Log(o)
	for _, key := range o.Keys() {
		v, _ := o.Search(key)
		t.Log(key, ":", v)
	}

	o.RSortKey()

	t.Log(o)
	for _, key := range o.Keys() {
		v, _ := o.Search(key)
		t.Log(key, ":", v)
	}
}

func TestOrderedMap_SortValue(t *testing.T) {
	o := NewOrderMap()
	o.Set("b", "b")
	o.Set("d", "3")
	o.Set("a", "a")
	o.Set("c", "c")
	o.Set("f", "1")
	o.Set("e", 2)
	t.Log(o)
	for _, key := range o.Keys() {
		v, _ := o.Search(key)
		t.Log(key, ":", v)
	}

	o.SortValue()

	t.Log(o)
	for _, key := range o.Keys() {
		v, _ := o.Search(key)
		t.Log(key, ":", v)
	}

	o.RSortValue()

	t.Log(o)
	for _, key := range o.Keys() {
		v, _ := o.Search(key)
		t.Log(key, ":", v)
	}
}
