package avalid

import (
	"fmt"
	"github.com/asktop/gotools/astring"
)

//必须为数值
type isDecimal struct {
	title    string
	value    interface{}
	valueStr string
	msgs     []string
	length   []uint
}

func (c *isDecimal) Check() (msg string, ok bool) {
	if len(c.msgs) > 0 {
		msg = c.msgs[0]
	}
	if len(c.length) == 0 {
		if !astring.IsDecimal(c.valueStr) {
			if len(c.msgs) == 0 {
				msg = fmt.Sprintf("%s必须为数值", c.title)
			}
			return msg, false
		}
	} else {
		if !astring.IsDecimal(c.valueStr, c.length...) {
			if len(c.msgs) == 0 {
				var lenStr string
				if len(c.length) == 1 {
					lenStr = fmt.Sprintf("%d", c.length[0])
				} else {
					lenStr = fmt.Sprintf("%d 至 %d", c.length[0], c.length[1])
				}
				msg = fmt.Sprintf("%s必须为数值，且小数位数必须为 %s", c.title, lenStr)
			}
			return msg, false
		}
	}
	return "", true
}
