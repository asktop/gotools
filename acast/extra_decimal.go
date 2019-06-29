package acast

import (
	big "github.com/asktop/decimal"
	"github.com/shopspring/decimal"
)

func ToBig(i interface{}, defaultVal ...string) *big.Big {
	v := new(big.Big)
	_, ok := v.SetString(ToString(i))
	if len(defaultVal) > 0 && !ok {
		v.SetString(defaultVal[0])
	}
	return v
}

func ToDecimal(i interface{}, defaultVal ...string) decimal.Decimal {
	v, err := decimal.NewFromString(ToString(i))
	if len(defaultVal) > 0 && err != nil {
		v, _ = decimal.NewFromString(defaultVal[0])
	}
	return v
}
