package amath

import (
	"github.com/asktop/decimal"
	"github.com/asktop/decimal/math"
	"strings"
)

// 天花板数，数值变大
func BigCeil(z *decimal.Big, n int) *decimal.Big {
	scale := decimal.New(1, n)
	z.Quo(z, scale)
	math.Ceil(z, z)
	z.Mul(z, scale)
	//位数补全
	if n > 0 && z.Scale() >= 0 && z.Scale() < n {
		y := z.String()
		if z.Scale() == 0 {
			y = y + "." + strings.Repeat("0", n)
		} else {
			y = y + strings.Repeat("0", n-z.Scale())
		}
		z.SetString(y)
	}
	return z
}

// 地板数，数值变小
func BigFloor(z *decimal.Big, n int) *decimal.Big {
	scale := decimal.New(1, n)
	z.Quo(z, scale)
	math.Floor(z, z)
	z.Mul(z, scale)
	//位数补全
	if n > 0 && z.Scale() >= 0 && z.Scale() < n {
		y := z.String()
		if z.Scale() == 0 {
			y = y + "." + strings.Repeat("0", n)
		} else {
			y = y + strings.Repeat("0", n-z.Scale())
		}
		z.SetString(y)
	}
	return z
}

// 四舍五入，正数时数值舍变小入变大，负数时数值舍变大入变小
func BigRound(z *decimal.Big, n int) *decimal.Big {
	return z.Quantize(n)
}

// 直接舍去，正数时数值变小，负数时数值变大
func BigScale(z *decimal.Big, n int) *decimal.Big {
	scale := decimal.New(1, n)
	z.Quo(z, scale)
	if z.Cmp(decimal.New(0, 0)) >= 0 {
		math.Floor(z, z)
	} else {
		math.Ceil(z, z)
	}
	z.Mul(z, scale)
	//位数补全
	if n > 0 && z.Scale() >= 0 && z.Scale() < n {
		y := z.String()
		if z.Scale() == 0 {
			y = y + "." + strings.Repeat("0", n)
		} else {
			y = y + strings.Repeat("0", n-z.Scale())
		}
		z.SetString(y)
	}
	return z
}
