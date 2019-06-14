package avalid

import "fmt"

//必需
type required struct {
	title    string
	value    interface{}
	valueStr string
	msgs     []string
}

func (c *required) Check() (msg string, ok bool) {
	msg = fmt.Sprintf("%s不能为空", c.title)
	if len(c.msgs) > 0 {
		msg = c.msgs[0]
	}
	if c.value != nil {
		if c.value != "" {
			return "", true
		}
	}
	return
}
