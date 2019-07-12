package avalid

import (
	"fmt"
	"github.com/asktop/gotools/astring"
)

//检查密码
// level: 密码强度级别
// 	1：包含数字、字母
// 	2：包含数字、字母、下划线
// 	3：包含数字、字母、特殊字符
// 	4：包含数字、大小写字母
// 	5：包含数字、大小写字母、下划线
// 	6：包含数字、大小写字母、特殊字符
type isPwd struct {
	title    string
	value    interface{}
	valueStr string
	msgs     []string
	level    uint
	length   []uint
}

func (c *isPwd) Check() (msg string, ok bool) {
	if len(c.msgs) > 0 {
		msg = c.msgs[0]
	}
	if len(c.length) == 0 {
		if !astring.IsPwd(c.valueStr, c.level) {
			if len(c.msgs) == 0 {
				msg = fmt.Sprintf("%s不符合要求", c.title)
			}
			return msg, false
		}
	} else {
		if !astring.IsPwd(c.valueStr, c.level, c.length...) {
			if len(c.msgs) == 0 {
				var lenStr string
				if len(c.length) == 1 {
					lenStr = fmt.Sprintf("%d", c.length[0])
				} else {
					lenStr = fmt.Sprintf("%d 至 %d", c.length[0], c.length[1])
				}
				msg = fmt.Sprintf("%s不符合要求，且长度必须为 %s", c.title, lenStr)
			}
			return msg, false
		}
	}
	return "", true
}
