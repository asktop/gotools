package amath

import (
	"github.com/asktop/decimal"
	"testing"
)

func TestBigMath(t *testing.T) {
	x := new(decimal.Big)
	x.SetString("123.1453")
	t.Log(x)
	_, ok := x.SetString("")
	t.Log(ok)
	t.Log(x)
}

func TestBigRound(t *testing.T) {
	x := new(decimal.Big)
	x.SetString("123.1453")
	t.Log(BigRound(x, 0).String())
	t.Log(BigRound(x, 2).String())
	t.Log(BigRound(x, 8).String())
	x.SetString("-123.1453")
	t.Log(BigRound(x, 2).String())
}

func TestBigScale(t *testing.T) {
	x := new(decimal.Big)
	x.SetString("123.1453")
	t.Log(BigScale(x, 0).String())
	t.Log(BigScale(x, 2).String())
	t.Log(BigScale(x, 8).String())
	x.SetString("-123.1453")
	t.Log(BigScale(x, 2).String())
}

func TestBigCeil(t *testing.T) {
	x := new(decimal.Big)
	x.SetString("123.1453")
	t.Log(BigCeil(x, 0).String())
	t.Log(BigCeil(x, 2).String())
	t.Log(BigCeil(x, 8).String())
	x.SetString("-123.1453")
	t.Log(BigCeil(x, 2).String())
}

func TestBigFloor(t *testing.T) {
	x := new(decimal.Big)
	x.SetString("123.1453")
	t.Log(BigFloor(x, 0).String())
	t.Log(BigFloor(x, 2).String())
	t.Log(BigFloor(x, 8).String())
	x.SetString("-123.1453")
	t.Log(BigFloor(x, 2).String())
}
