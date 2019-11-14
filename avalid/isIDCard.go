package avalid

import (
	"fmt"
	"github.com/asktop/gotools/astring"
)

//必须为身份证号码
type isIDCard struct {
	title    string
	value    interface{}
	valueStr string
	msgs     []string
}

func (c *isIDCard) Check() (msg string, ok bool) {
	if c.valueStr == "" {
		return "", true
	}
	if len(c.msgs) > 0 {
		msg = c.msgs[0]
	}
	if !astring.IsIDCard(c.valueStr) {
		if len(c.msgs) == 0 {
			msg = fmt.Sprintf("%s不是正确的身份证号码格式", c.title)
		}
		return msg, false
	}
	return "", true
}
