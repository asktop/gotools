package amath

import (
	"fmt"
	"github.com/ericlagergren/decimal"
	"testing"
)

func TestStrScale(t *testing.T) {
	numStr := "0.1234500"
	num, _ := new(decimal.Big).SetString(numStr)
	fmt.Println(num.Scale())
	fmt.Println(StrScale(numStr, 2))

	numStr2 := "-0.12"
	num2, _ := new(decimal.Big).SetString(numStr2)
	fmt.Println(num2.Scale())
	fmt.Println(StrScale(numStr2, 4))
}
