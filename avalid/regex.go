package avalid

import (
	"fmt"
	"github.com/asktop/gotools/astring"
)

//正则表达式验证
type regex struct {
	title    string
	value    interface{}
	valueStr string
	msgs     []string
	exp      string
}

func (c *regex) Check() (msg string, ok bool) {
	if c.valueStr == "" {
		return "", true
	}
	if len(c.msgs) > 0 {
		msg = c.msgs[0]
	}
	if !astring.MatchString(c.exp, c.valueStr) {
		if len(c.msgs) > 0 {
			msg = fmt.Sprintf("%s验证不合法", c.title)
		}
		return msg, false
	}
	return "", true
}
