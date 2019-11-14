package avalid

import (
	"fmt"
	"github.com/asktop/gotools/astring"
)

var accountTip = "必须以字母开头，由数字字母下划线组成"

//检查账号（字母开头，数字字母下划线）
type isAccount struct {
	title    string
	value    interface{}
	valueStr string
	msgs     []string
	length   []uint
}

func (c *isAccount) Check() (msg string, ok bool) {
	if c.valueStr == "" {
		return "", true
	}
	if len(c.msgs) > 0 {
		msg = c.msgs[0]
	}
	if len(c.length) == 0 {
		if !astring.IsAccount(c.valueStr) {
			if len(c.msgs) == 0 {
				msg = fmt.Sprintf("%s%s", c.title, accountTip)
			}
			return msg, false
		}
	} else {
		if !astring.IsAccount(c.valueStr, c.length...) {
			if len(c.msgs) == 0 {
				var lenStr string
				if len(c.length) == 1 {
					lenStr = fmt.Sprintf("%d", c.length[0])
				} else {
					lenStr = fmt.Sprintf("%d 至 %d", c.length[0], c.length[1])
				}
				msg = fmt.Sprintf("%s%s，且长度必须为 %s", c.title, accountTip, lenStr)
			}
			return msg, false
		}
	}
	return "", true
}
