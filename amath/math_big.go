package amath

import (
	"github.com/ericlagergren/decimal"
	"github.com/ericlagergren/decimal/math"
	"strings"
)

// 四舍五入，正数时数值舍变小入变大，负数时数值舍变大入变小
func BigRound(x *decimal.Big, num int) *decimal.Big {
	z := new(decimal.Big)
	if x != nil {
		z.Copy(x)
	}
	return z.Quantize(num)
}

// 直接舍去，正数时数值变小，负数时数值变大
func BigScale(x *decimal.Big, num int) *decimal.Big {
	z := new(decimal.Big)
	if x != nil {
		z.Copy(x)
	}
	scale := decimal.New(1, num)
	z.Quo(z, scale)
	if z.Cmp(decimal.New(0, 0)) >= 0 {
		math.Floor(z, z)
	} else {
		math.Ceil(z, z)
	}
	z.Mul(z, scale)
	//位数补全
	if num > 0 && z.Scale() >= 0 && z.Scale() < num {
		y := z.String()
		if z.Scale() == 0 {
			y = y + "." + strings.Repeat("0", num)
		} else {
			y = y + strings.Repeat("0", num-z.Scale())
		}
		z.SetString(y)
	}
	return z
}

// 天花板数，数值变大
func BigCeil(x *decimal.Big, num int) *decimal.Big {
	z := new(decimal.Big)
	if x != nil {
		z.Copy(x)
	}
	scale := decimal.New(1, num)
	z.Quo(z, scale)
	math.Ceil(z, z)
	z.Mul(z, scale)
	//位数补全
	if num > 0 && z.Scale() >= 0 && z.Scale() < num {
		y := z.String()
		if z.Scale() == 0 {
			y = y + "." + strings.Repeat("0", num)
		} else {
			y = y + strings.Repeat("0", num-z.Scale())
		}
		z.SetString(y)
	}
	return z
}

// 地板数，数值变小
func BigFloor(x *decimal.Big, num int) *decimal.Big {
	z := new(decimal.Big)
	if x != nil {
		z.Copy(x)
	}
	scale := decimal.New(1, num)
	z.Quo(z, scale)
	math.Floor(z, z)
	z.Mul(z, scale)
	//位数补全
	if num > 0 && z.Scale() >= 0 && z.Scale() < num {
		y := z.String()
		if z.Scale() == 0 {
			y = y + "." + strings.Repeat("0", num)
		} else {
			y = y + strings.Repeat("0", num-z.Scale())
		}
		z.SetString(y)
	}
	return z
}
