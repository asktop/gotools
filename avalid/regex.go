package avalid

import (
	"fmt"
	"regexp"
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
	msg = fmt.Sprintf("%s验证不合法", c.title)
	if len(c.msgs) > 0 {
		msg = c.msgs[0]
	}
	ok, err := regexp.MatchString(c.exp, c.valueStr)
	if err != nil {
		return fmt.Sprintf("%s正则验证表达式不合法", c.title), false
	} else if ok {
		msg = ""
	}
	return
}
