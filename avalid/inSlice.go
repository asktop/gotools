package avalid

import (
	"fmt"
	"github.com/asktop/gotools/acast"
	"github.com/asktop/gotools/aslice"
)

//在切片中
type inSlice struct {
	title    string
	value    interface{}
	valueStr string
	msgs     []string
	slice    interface{}
}

func (c *inSlice) Check() (msg string, ok bool) {
	msg = fmt.Sprintf("%s不在规定范围内", c.title)
	if len(c.msgs) > 0 {
		msg = c.msgs[0]
	}
	if c.slice != nil {
		values := acast.ToStringSlice(c.slice)
		if aslice.ContainString(values, c.valueStr) {
			return "", true
		}
	}
	return
}
