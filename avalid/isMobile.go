package avalid

import (
	"fmt"
	"github.com/asktop/gotools/astring"
)

//必须为手机号
type isMobile struct {
	title    string
	value    interface{}
	valueStr string
	msgs     []string
}

func (c *isMobile) Check() (msg string, ok bool) {
	if len(c.msgs) > 0 {
		msg = c.msgs[0]
	}
	if !astring.IsMobile(c.valueStr) {
		if len(c.msgs) == 0 {
			msg = fmt.Sprintf("%s 手机号格式不正确", c.title)
		}
		return msg, false
	}
	return "", true
}
