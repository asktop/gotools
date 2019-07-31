package amath

import (
	big "github.com/asktop/decimal"
	"github.com/shopspring/decimal"
	"testing"
)

func TestStrCeil(t *testing.T) {
	a1 := "123.1453"
	t.Log(StrCeil(a1, 0))
	a2, _ := new(big.Big).SetString(a1)
	t.Log(StrCeil(a2, 2))
	a3, _ := decimal.NewFromString(a1)
	t.Log(StrCeil(a3, 8))

	b1 := "-123.1453"
	t.Log(StrCeil(b1, 0))
	b2, _ := new(big.Big).SetString(b1)
	t.Log(StrCeil(b2, 2))
	b3, _ := decimal.NewFromString(b1)
	t.Log(StrCeil(b3, 8))
}

func TestStrFloor(t *testing.T) {
	a1 := "123.1453"
	t.Log(StrFloor(a1, 0))
	a2, _ := new(big.Big).SetString(a1)
	t.Log(StrFloor(a2, 2))
	a3, _ := decimal.NewFromString(a1)
	t.Log(StrFloor(a3, 8))

	b1 := "-123.1453"
	t.Log(StrFloor(b1, 0))
	b2, _ := new(big.Big).SetString(b1)
	t.Log(StrFloor(b2, 2))
	b3, _ := decimal.NewFromString(b1)
	t.Log(StrFloor(b3, 8))
}

func TestStrRound(t *testing.T) {
	a1 := "123.1453"
	t.Log(StrRound(a1, 0))
	a2, _ := new(big.Big).SetString(a1)
	t.Log(StrRound(a2, 2))
	a3, _ := decimal.NewFromString(a1)
	t.Log(StrRound(a3, 8))

	b1 := "-123.1453"
	t.Log(StrRound(b1, 0))
	b2, _ := new(big.Big).SetString(b1)
	t.Log(StrRound(b2, 2))
	b3, _ := decimal.NewFromString(b1)
	t.Log(StrRound(b3, 8))
}

func TestStrScale(t *testing.T) {
	a1 := "123.1453"
	t.Log(StrScale(a1, 0))
	a2, _ := new(big.Big).SetString(a1)
	t.Log(StrScale(a2, 2))
	a3, _ := decimal.NewFromString(a1)
	t.Log(StrScale(a3, 8))

	b1 := "-123.1453"
	t.Log(StrScale(b1, 0))
	b2, _ := new(big.Big).SetString(b1)
	t.Log(StrScale(b2, 2))
	b3, _ := decimal.NewFromString(b1)
	t.Log(StrScale(b3, 8))
}
