package avalid

import (
	"fmt"
	"github.com/asktop/gotools/acast"
	"github.com/asktop/gotools/astring"
)

//必须为数字字符串
type isNumber struct {
	title    string
	value    interface{}
	valueStr string
	msgs     []string
	length   interface{}
}

func (c *isNumber) Check() (msg string, ok bool) {
	if c.length != nil {
		l := acast.ToInt(c.length)
		msg = fmt.Sprintf("%s必须为数字字符串，且长度必须为 %d", c.title, l)
		if astring.IsNumber(c.valueStr, l) {
			return "", true
		}
	} else {
		msg = fmt.Sprintf("%s必须为数字字符串", c.title)
		if astring.IsNumber(c.valueStr) {
			return "", true
		}
	}
	if len(c.msgs) > 0 {
		msg = c.msgs[0]
	}
	return
}
