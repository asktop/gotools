package avalid

import (
	"fmt"
	"github.com/asktop/gotools/astring"
)

//必须为Email
type isEmail struct {
	title    string
	value    interface{}
	valueStr string
	msgs     []string
}

func (c *isEmail) Check() (msg string, ok bool) {
	msg = fmt.Sprintf("%s的Email格式不正确", c.title)
	if len(c.msgs) > 0 {
		msg = c.msgs[0]
	}
	if astring.IsEmail(c.valueStr) {
		return "", true
	}
	return
}
