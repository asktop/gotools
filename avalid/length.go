package avalid

import (
	"fmt"
	"github.com/asktop/gotools/acast"
)

//字符串长度范围
type length struct {
	title    string
	value    interface{}
	valueStr string
	msgs     []string
	min      interface{}
	max      interface{}
}

func (c *length) Check() (msg string, ok bool) {
	val := []rune(c.valueStr)
	if c.min != nil {
		mi := acast.ToInt(c.min)
		if len(val) < mi {
			if len(c.msgs) > 0 {
				msg = c.msgs[0]
			} else {
				msg = fmt.Sprintf("%s长度必须大于等于 %d", c.title, mi)
			}
			return msg, false
		}
	}
	if c.max != nil {
		ma := acast.ToInt(c.max)
		if len(val) > ma {
			if len(c.msgs) > 0 {
				msg = c.msgs[0]
			} else {
				msg = fmt.Sprintf("%s长度必须小于等于 %d", c.title, ma)
			}
			return msg, false
		}
	}
	return "", true
}
