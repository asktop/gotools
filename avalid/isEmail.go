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
	if c.valueStr == "" {
		return "", true
	}
	if len(c.msgs) > 0 {
		msg = c.msgs[0]
	}
	if !astring.IsEmail(c.valueStr) {
		if len(c.msgs) == 0 {
			msg = fmt.Sprintf("%s不是正确的Email格式", c.title)
		}
		return msg, false
	}
	return "", true
}
