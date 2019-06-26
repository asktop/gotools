package main

import (
	"fmt"
	"github.com/asktop/gotools/acast"
	"github.com/asktop/gotools/ajson"
	"github.com/ericlagergren/decimal"
	"testing"
)

type User struct {
	Name string
}

func TestGotools(t *testing.T) {

}

func TestToString(t *testing.T) {
	fmt.Println("----- 1 -----")

	fmt.Println(nil)
	fmt.Println(acast.ToString(nil))
	fmt.Println(acast.ToStringForce(nil))
	fmt.Println(ajson.Encode(nil))

	fmt.Println("----- 2 -----")

	a := decimal.New(27000, 2)
	fmt.Println(a)
	fmt.Println(acast.ToString(a))
	fmt.Println(acast.ToStringForce(a))
	fmt.Println(ajson.Encode(a))

	fmt.Println("----- 3 -----")

	b := errors.New("errrror")
	fmt.Println(b)
	fmt.Println(acast.ToString(b))
	fmt.Println(acast.ToStringForce(b))
	fmt.Println(ajson.Encode(b))

	fmt.Println("----- 4 -----")

	c := User{}
	fmt.Println(c)
	fmt.Println(acast.ToString(c))
	fmt.Println(acast.ToStringForce(c))
	fmt.Println(ajson.Encode(c))

	fmt.Println("----- 5 -----")

	d := new(User)
	fmt.Println(d)
	fmt.Println(acast.ToString(d))
	fmt.Println(acast.ToStringForce(d))
	fmt.Println(ajson.Encode(d))

	fmt.Println("----- 6 -----")

	e := new(User)
	e = nil
	fmt.Println(e)
	fmt.Println(acast.ToString(e))
	fmt.Println(acast.ToStringForce(e))
	fmt.Println(ajson.Encode(e))
}
