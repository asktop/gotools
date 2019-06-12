package omap

import (
	"encoding/json"
	"fmt"
	"testing"
)

func TestSetMap(t *testing.T) {
	data := map[string]string{}
	data["k1"] = "v1"
	data["k3"] = "v3"
	data["k2"] = "v2"
	odata := Adds(data)
	fmt.Println(odata)
	odata.SortKey()
	fmt.Println(odata)
}

func TestOrderMapRange(t *testing.T) {
	o := New()
	o.Set("number", 3)
	o.Set("string", "x")
	o.Set("c", "c")
	o.Set("a", "a")
	o.Set("b", "b")
	for _, k := range o.Keys() {
		v := o.Get(k)
		fmt.Println(k, v)
	}
}

func TestOrderMap(t *testing.T) {
	o := New()
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
	n := New()
	n.Set("e", 1)
	n.Set("a", 2)
	o.Set("orderedmap", n)

	// Keys method
	fmt.Println(o)
	for _, key := range o.Keys() {
		v, _ := o.Search(key)
		fmt.Println(key, ":", v)
	}

	// delete
	o.Remove("strings")
	fmt.Println(o)
}

func TestOrderedMap_MarshalJSON(t *testing.T) {
	o := New()
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
	n := New()
	n.Set("e", 1)
	n.Set("a", 2)
	o.Set("orderedmap", n)

	fmt.Println(o)
	// convert to json
	b, err := json.Marshal(o)
	if err != nil {
		t.Error("Marshalling json", err)
	}
	s := string(b)
	fmt.Println(s)
}

func TestOrderedMap_UnmarshalJSON(t *testing.T) {
	s := `{"number":3,"string":"x","strings":["t","u"],"slice":["1",1],"orderedmap":{"e":1,"a":2}}`
	o := New()
	err := json.Unmarshal([]byte(s), &o)
	if err != nil {
		t.Error("JSON Unmarshal error", err)
	}

	fmt.Println(o)
	for _, key := range o.Keys() {
		v, _ := o.Search(key)
		fmt.Println(key, ":", v)
	}
}

func TestOrderedMap_SortKey(t *testing.T) {
	o := New()
	o.Set("b", "b")
	o.Set("d", "3")
	o.Set("a", "a")
	o.Set("c", "c")
	o.Set("f", "1")
	o.Set("e", 2)
	fmt.Println(o)
	for _, key := range o.Keys() {
		v, _ := o.Search(key)
		fmt.Println(key, ":", v)
	}

	o.SortKey()

	fmt.Println(o)
	for _, key := range o.Keys() {
		v, _ := o.Search(key)
		fmt.Println(key, ":", v)
	}

	o.RSortKey()

	fmt.Println(o)
	for _, key := range o.Keys() {
		v, _ := o.Search(key)
		fmt.Println(key, ":", v)
	}
}

func TestOrderedMap_SortValue(t *testing.T) {
	o := New()
	o.Set("b", "b")
	o.Set("d", "3")
	o.Set("a", "a")
	o.Set("c", "c")
	o.Set("f", "1")
	o.Set("e", 2)
	fmt.Println(o)
	for _, key := range o.Keys() {
		v, _ := o.Search(key)
		fmt.Println(key, ":", v)
	}

	o.SortValue()

	fmt.Println(o)
	for _, key := range o.Keys() {
		v, _ := o.Search(key)
		fmt.Println(key, ":", v)
	}

	o.RSortValue()

	fmt.Println(o)
	for _, key := range o.Keys() {
		v, _ := o.Search(key)
		fmt.Println(key, ":", v)
	}
}
