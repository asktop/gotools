package amath

import (
	"github.com/shopspring/decimal"
)

// 天花板数，数值变大
func DecimalCeil(z decimal.Decimal, n int) decimal.Decimal {
	z = z.Shift(int32(n))
	z = z.Ceil()
	z = z.Div(decimal.New(1, int32(n)))
	return z
}

// 地板数，数值变小
func DecimalFloor(z decimal.Decimal, n int) decimal.Decimal {
	z = z.Shift(int32(n))
	z = z.Floor()
	z = z.Div(decimal.New(1, int32(n)))
	return z
}

// 四舍五入，正数时数值舍变小入变大，负数时数值舍变大入变小
func DecimalRound(z decimal.Decimal, n int) decimal.Decimal {
	return z.Round(int32(n))
}

// 直接舍去，正数时数值变小，负数时数值变大
func DecimalScale(z decimal.Decimal, n int) decimal.Decimal {
	z = z.Shift(int32(n))
	if z.Cmp(decimal.Zero) >= 0 {
		z = z.Floor()
	} else {
		z = z.Ceil()
	}
	z = z.Div(decimal.New(1, int32(n)))
	return z
}
