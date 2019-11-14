package avalid

import (
	"fmt"
	"strings"
)

//不能有空格
type hasNoBlank struct {
	title    string
	value    interface{}
	valueStr string
	msgs     []string
}

func (c *hasNoBlank) Check() (msg string, ok bool) {
	if c.valueStr == "" {
		return "", true
	}
	if len(c.msgs) > 0 {
		msg = c.msgs[0]
	} else {
		msg = fmt.Sprintf("%s不能有空格", c.title)
	}
	if strings.Contains(c.valueStr, " ") {
		return msg, false
	}
	return "", true
}
