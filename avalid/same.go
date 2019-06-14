package avalid

import "fmt"

//相同
type same struct {
	title    string
	value    interface{}
	valueStr string
	msgs     []string
	sameVal  interface{}
}

func (c *same) Check() (msg string, ok bool) {
	msg = fmt.Sprintf("%s不能为空", c.title)
	if c.value != nil {
		if c.value == c.sameVal {
			return "", true
		} else {
			msg = fmt.Sprintf("%s与规定不相同", c.title)
		}
		if len(c.msgs) > 0 {
			msg = c.msgs[0]
		}
	}
	return
}
