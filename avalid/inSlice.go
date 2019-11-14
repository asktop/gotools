package avalid

import (
	"fmt"
	"github.com/asktop/gotools/aslice"
)

//在切片中
type inSlice struct {
	title    string
	value    interface{}
	valueStr string
	msgs     []string
	slice    []string
}

func (c *inSlice) Check() (msg string, ok bool) {
	if c.valueStr == "" {
		return "", true
	}
	if len(c.msgs) > 0 {
		msg = c.msgs[0]
	}
	if !aslice.ContainString(c.slice, c.valueStr) {
		if len(c.msgs) == 0 {
			msg = fmt.Sprintf("%s不在规定范围内", c.title)
		}
		return msg, false
	}
	return "", true
}
