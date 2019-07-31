package amath

import (
	"github.com/asktop/decimal"
	"testing"
)

func TestBigCeil(t *testing.T) {
	a1, _ := new(decimal.Big).SetString("123.1453")
	a2 := new(decimal.Big).Copy(a1)
	a3 := new(decimal.Big).Copy(a1)
	t.Log(BigCeil(a1, 0).String())
	t.Log(BigCeil(a2, 2).String())
	t.Log(BigCeil(a3, 8).String())

	b1, _ := new(decimal.Big).SetString("-123.1453")
	b2 := new(decimal.Big).Copy(b1)
	b3 := new(decimal.Big).Copy(b1)
	t.Log(BigCeil(b1, 0).String())
	t.Log(BigCeil(b2, 2).String())
	t.Log(BigCeil(b3, 8).String())
}

func TestBigFloor(t *testing.T) {
	a1, _ := new(decimal.Big).SetString("123.1453")
	a2 := new(decimal.Big).Copy(a1)
	a3 := new(decimal.Big).Copy(a1)
	t.Log(BigFloor(a1, 0).String())
	t.Log(BigFloor(a2, 2).String())
	t.Log(BigFloor(a3, 8).String())

	b1, _ := new(decimal.Big).SetString("-123.1453")
	b2 := new(decimal.Big).Copy(b1)
	b3 := new(decimal.Big).Copy(b1)
	t.Log(BigFloor(b1, 0).String())
	t.Log(BigFloor(b2, 2).String())
	t.Log(BigFloor(b3, 8).String())
}

func TestBigRound(t *testing.T) {
	a1, _ := new(decimal.Big).SetString("123.1453")
	a2 := new(decimal.Big).Copy(a1)
	a3 := new(decimal.Big).Copy(a1)
	t.Log(BigRound(a1, 0).String())
	t.Log(BigRound(a2, 2).String())
	t.Log(BigRound(a3, 8).String())

	b1, _ := new(decimal.Big).SetString("-123.1453")
	b2 := new(decimal.Big).Copy(b1)
	b3 := new(decimal.Big).Copy(b1)
	t.Log(BigRound(b1, 0).String())
	t.Log(BigRound(b2, 2).String())
	t.Log(BigRound(b3, 8).String())
}

func TestBigScale(t *testing.T) {
	a1, _ := new(decimal.Big).SetString("123.1453")
	a2 := new(decimal.Big).Copy(a1)
	a3 := new(decimal.Big).Copy(a1)
	t.Log(BigScale(a1, 0).String())
	t.Log(BigScale(a2, 2).String())
	t.Log(BigScale(a3, 8).String())

	b1, _ := new(decimal.Big).SetString("-123.1453")
	b2 := new(decimal.Big).Copy(b1)
	b3 := new(decimal.Big).Copy(b1)
	t.Log(BigScale(b1, 0).String())
	t.Log(BigScale(b2, 2).String())
	t.Log(BigScale(b3, 8).String())
}
