package amath

import (
	"github.com/asktop/decimal"
	"github.com/asktop/gotools/acast"
	"strings"
)

// 天花板数，数值变大
func StrCeil(numStr interface{}, n int) string {
	num, ok := new(decimal.Big).SetString(acast.ToString(numStr))
	if !ok {
		return "0." + strings.Repeat("0", n)
	}
	return BigCeil(num, n).String()
}

// 地板数，数值变小
func StrFloor(numStr interface{}, n int) string {
	num, ok := new(decimal.Big).SetString(acast.ToString(numStr))
	if !ok {
		return "0." + strings.Repeat("0", n)
	}
	return BigFloor(num, n).String()
}

// 四舍五入，正数时数值舍变小入变大，负数时数值舍变大入变小
func StrRound(numStr interface{}, n int) string {
	num, ok := new(decimal.Big).SetString(acast.ToString(numStr))
	if !ok {
		return "0." + strings.Repeat("0", n)
	}
	return BigRound(num, n).String()
}

// 直接舍去，正数时数值变小，负数时数值变大
func StrScale(numStr interface{}, n int) string {
	num, ok := new(decimal.Big).SetString(acast.ToString(numStr))
	if !ok {
		return "0." + strings.Repeat("0", n)
	}
	return BigScale(num, n).String()
}
