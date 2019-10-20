package avalid

import (
	"fmt"
	"strings"
)

//必需
type required struct {
	title    string
	value    interface{}
	valueStr string
	msgs     []string
}

func (c *required) Check() (msg string, ok bool) {
	if len(c.msgs) > 0 {
		msg = c.msgs[0]
	} else {
		msg = fmt.Sprintf("%s不能为空", c.title)
	}
	if c.value != nil {
		if strings.Trim(c.valueStr, " ") != "" {
			return "", true
		}
	}
	return msg, false
}
