package avalid

import (
	"fmt"
	"github.com/asktop/gotools/acast"
	"github.com/asktop/gotools/astring"
)

//必须为数值
type isDecimal struct {
	title    string
	value    interface{}
	valueStr string
	msgs     []string
	length   interface{}
}

func (c *isDecimal) Check() (msg string, ok bool) {
	if c.length != nil {
		l := acast.ToInt(c.length)
		msg = fmt.Sprintf("%s必须为数值，且小数位数必须为 %d", c.title, l)
		if astring.IsAllDecimal(c.valueStr, l) {
			return "", true
		}
	} else {
		msg = fmt.Sprintf("%s必须为数值", c.title)
		if astring.IsAllDecimal(c.valueStr) {
			return "", true
		}
	}
	if len(c.msgs) > 0 {
		msg = c.msgs[0]
	}
	return
}
