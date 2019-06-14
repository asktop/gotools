package avalid

import (
	"fmt"
	"github.com/asktop/gotools/acast"
	"github.com/ericlagergren/decimal"
)

//数值相等
type equal struct {
	title    string
	value    interface{}
	valueStr string
	msgs     []string
	equalVal interface{}
}

func (c *equal) Check() (msg string, ok bool) {
	msg = fmt.Sprintf("%s不能为空", c.title)
	if c.valueStr != "" {
		val, _ := new(decimal.Big).SetString(c.valueStr)
		equalVal, _ := new(decimal.Big).SetString(acast.ToString(c.equalVal))
		if val.Cmp(equalVal) == 0 {
			return "", true
		} else {
			msg = fmt.Sprintf("%s与规定不相等", c.title)
		}
		if len(c.msgs) > 0 {
			msg = c.msgs[0]
		}
	}
	return
}
