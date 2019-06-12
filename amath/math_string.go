package amath

import (
	"github.com/asktop/gotools/acast"
	"github.com/ericlagergren/decimal"
	"strings"
)

// 直接舍去，正数时数值变小，负数时数值变大
func StrScale(numStr interface{}, scale int) string {
	num, ok := new(decimal.Big).SetString(acast.ToString(numStr))
	if !ok {
		return "0." + strings.Repeat("0", scale)
	}
	return BigScale(num, scale).String()
}
