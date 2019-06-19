package acast

import (
	"fmt"
	"testing"
)

var personMap = map[string]interface{}{
	"name":   "Mitchell",
	"age":    91,
	"emails": []string{"one", "two", "three"},
	"extra": map[string]string{
		"twitter":  "mitchellh",
		"twitter2": "mitchellh2",
	},
}

type person struct {
	Name   string
	Age    int
	Emails []string
	Extra  map[string]string
}

var personStruct = person{
	Name:   "asktop",
	Age:    18,
	Emails: []string{"aaa", "bbb"},
	Extra: map[string]string{
		"github":  "wendao",
		"github2": "wendao2",
	},
}

func TestMapToStruct(t *testing.T) {
	output := person{}
	err := MapToStruct(personMap, &output)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(output)
	}
}

func TestStructToMap(t *testing.T) {
	output := map[string]interface{}{}
	err := StructToMap(personStruct, &output)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(output)
	}
}

type person2 struct {
	UserName string            `json:"name"`
	UserAge  int               `json:"age"`
	Emails   []string          `json:"emails"`
	Extra    map[string]string `json:"extra"`
}

var personStruct2 = person2{
	UserName: "asktop",
	UserAge:  18,
	Emails:   []string{"aaa", "bbb"},
	Extra: map[string]string{
		"github":  "wendao",
		"github2": "wendao2",
	},
}

func TestMapToStruct2(t *testing.T) {
	output := person2{}
	err := MapToStruct(personMap, &output)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(output)
	}
}

func TestStructToMap2(t *testing.T) {
	output := map[string]interface{}{}
	err := StructToMap(personStruct2, &output)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(output)
	}
}
