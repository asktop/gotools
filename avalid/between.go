package avalid

import (
	"fmt"
	"github.com/asktop/gotools/acast"
	"github.com/asktop/decimal"
)

//数值的范围
type between struct {
	title    string
	value    interface{}
	valueStr string
	msgs     []string
	min      interface{}
	max      interface{}
}

func (c *between) Check() (msg string, ok bool) {
	if c.valueStr == "" {
		return "", true
	}
	if len(c.msgs) > 0 {
		msg = c.msgs[0]
	}
	val, _ := new(decimal.Big).SetString(c.valueStr)
	if c.min != nil {
		mi, _ := new(decimal.Big).SetString(acast.ToString(c.min))
		if val.Cmp(mi) < 0 {
			if len(c.msgs) == 0 {
				msg = fmt.Sprintf("%s必须大于等于 %s", c.title, mi.String())
			}
			return msg, false
		}
	}
	if c.max != nil {
		ma, _ := new(decimal.Big).SetString(acast.ToString(c.max))
		if val.Cmp(ma) > 0 {
			if len(c.msgs) == 0 {
				msg = fmt.Sprintf("%s必须小于等于 %s", c.title, ma.String())
			}
			return msg, false
		}
	}
	return "", true
}
