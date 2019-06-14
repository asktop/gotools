package avalid

import (
	"fmt"
	"github.com/asktop/gotools/astring"
)

//必须为整数
type isInt struct {
	title    string
	value    interface{}
	valueStr string
	msgs     []string
}

func (c *isInt) Check() (msg string, ok bool) {
	msg = fmt.Sprintf("%s必须为整数", c.title)
	if len(c.msgs) > 0 {
		msg = c.msgs[0]
	}
	if astring.IsAllInt(c.valueStr) {
		return "", true
	}
	return
}
