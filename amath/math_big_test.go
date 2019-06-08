package amath

import (
	"fmt"
	"github.com/ericlagergren/decimal"
	"testing"
)

func TestBigMath(t *testing.T) {
	x := new(decimal.Big)
	x.SetString("123.1453")
	fmt.Println(x)
	_, ok := x.SetString("")
	fmt.Println(ok)
	fmt.Println(x)
}

func TestBigRound(t *testing.T) {
	x := new(decimal.Big)
	x.SetString("123.1453")
	fmt.Println(BigRound(x, 0).String())
	fmt.Println(BigRound(x, 2).String())
	fmt.Println(BigRound(x, 8).String())
	x.SetString("-123.1453")
	fmt.Println(BigRound(x, 2).String())
}

func TestBigScale(t *testing.T) {
	x := new(decimal.Big)
	x.SetString("123.1453")
	fmt.Println(BigScale(x, 0).String())
	fmt.Println(BigScale(x, 2).String())
	fmt.Println(BigScale(x, 8).String())
	x.SetString("-123.1453")
	fmt.Println(BigScale(x, 2).String())
}

func TestBigCeil(t *testing.T) {
	x := new(decimal.Big)
	x.SetString("123.1453")
	fmt.Println(BigCeil(x, 0).String())
	fmt.Println(BigCeil(x, 2).String())
	fmt.Println(BigCeil(x, 8).String())
	x.SetString("-123.1453")
	fmt.Println(BigCeil(x, 2).String())
}

func TestBigFloor(t *testing.T) {
	x := new(decimal.Big)
	x.SetString("123.1453")
	fmt.Println(BigFloor(x, 0).String())
	fmt.Println(BigFloor(x, 2).String())
	fmt.Println(BigFloor(x, 8).String())
	x.SetString("-123.1453")
	fmt.Println(BigFloor(x, 2).String())
}
