package amath

import (
	"github.com/shopspring/decimal"
	"testing"
)

func TestDecimalCeil(t *testing.T) {
	a1, _ := decimal.NewFromString("123.1453")
	a2 := a1
	a3 := a1
	t.Log(DecimalCeil(a1, 0).String())
	t.Log(DecimalCeil(a2, 2).String())
	t.Log(DecimalCeil(a3, 8).String())

	b1, _ := decimal.NewFromString("-123.1453")
	b2 := b1
	b3 := b1
	t.Log(DecimalCeil(b1, 0).String())
	t.Log(DecimalCeil(b2, 2).String())
	t.Log(DecimalCeil(b3, 8).String())
}

func TestDecimalFloor(t *testing.T) {
	a1, _ := decimal.NewFromString("123.1453")
	a2 := a1
	a3 := a1
	t.Log(DecimalFloor(a1, 0).String())
	t.Log(DecimalFloor(a2, 2).String())
	t.Log(DecimalFloor(a3, 8).String())

	b1, _ := decimal.NewFromString("-123.1453")
	b2 := b1
	b3 := b1
	t.Log(DecimalFloor(b1, 0).String())
	t.Log(DecimalFloor(b2, 2).String())
	t.Log(DecimalFloor(b3, 8).String())
}

func TestDecimalRound(t *testing.T) {
	a1, _ := decimal.NewFromString("123.1453")
	a2 := a1
	a3 := a1
	t.Log(DecimalRound(a1, 0).String())
	t.Log(DecimalRound(a2, 2).String())
	t.Log(DecimalRound(a3, 8).String())

	b1, _ := decimal.NewFromString("-123.1453")
	b2 := b1
	b3 := b1
	t.Log(DecimalRound(b1, 0).String())
	t.Log(DecimalRound(b2, 2).String())
	t.Log(DecimalRound(b3, 8).String())
}

func TestDecimalScale(t *testing.T) {
	a1, _ := decimal.NewFromString("123.1453")
	a2 := a1
	a3 := a1
	t.Log(DecimalScale(a1, 0).String())
	t.Log(DecimalScale(a2, 2).String())
	t.Log(DecimalScale(a3, 8).String())

	b1, _ := decimal.NewFromString("-123.1453")
	b2 := b1
	b3 := b1
	t.Log(DecimalScale(b1, 0).String())
	t.Log(DecimalScale(b2, 2).String())
	t.Log(DecimalScale(b3, 8).String())
}
