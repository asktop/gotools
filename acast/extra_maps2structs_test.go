package acast

import (
	"fmt"
	"testing"
)

type user struct {
	Name  string
	Age   int
	Extra string
}

func TestMapsToStructs(t *testing.T) {
	m1 := map[string]interface{}{
		"name":  "a1",
		"age":   "11",
		"extra": "u1",
	}
	m2 := map[string]interface{}{
		"name":  "a2",
		"age":   "22",
		"extra": "u2",
	}
	ms := []map[string]interface{}{}
	ms = append(ms, m1)
	ms = append(ms, m2)

	us := []user{}

	err := MapsToStructs(ms, &us)
	if err != nil {
		fmt.Println(err)
	} else {
		for _, u := range us {
			fmt.Println(u)
		}
	}
}

type user2 struct {
	UserName string `json:"name"`
	UserAge  int    `json:"age"`
	Extra    string `json:"extra"`
}

func TestMapsToStructs2(t *testing.T) {
	m1 := map[string]interface{}{
		"name":  "a1",
		"age":   "11",
		"extra": "u1",
	}
	m2 := map[string]interface{}{
		"name":  "a2",
		"age":   "22",
		"extra": "u2",
	}
	ms := []map[string]interface{}{}
	ms = append(ms, m1)
	ms = append(ms, m2)

	us := []user2{}

	err := MapsToStructs(ms, &us)
	if err != nil {
		fmt.Println(err)
	} else {
		for _, u := range us {
			fmt.Println(u)
		}
	}
}

func TestMapsStrToStructs(t *testing.T) {
	m1 := map[string]string{
		"name":  "a1",
		"age":   "11",
		"extra": "u1",
	}
	m2 := map[string]string{
		"name":  "a2",
		"age":   "22",
		"extra": "u2",
	}
	ms := []map[string]string{}
	ms = append(ms, m1)
	ms = append(ms, m2)

	us := []user2{}

	err := MapsStrToStructs(ms, &us)
	if err != nil {
		fmt.Println(err)
	} else {
		for _, u := range us {
			fmt.Println(u)
		}
	}
}
