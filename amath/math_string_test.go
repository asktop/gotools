package amath

import (
	"github.com/ericlagergren/decimal"
	"testing"
)

func TestStrScale(t *testing.T) {
	numStr := "0.1234500"
	num, _ := new(decimal.Big).SetString(numStr)
	t.Log(num.Scale())
	t.Log(StrScale(numStr, 2))

	numStr2 := "-0.12"
	num2, _ := new(decimal.Big).SetString(numStr2)
	t.Log(num2.Scale())
	t.Log(StrScale(numStr2, 4))
}
