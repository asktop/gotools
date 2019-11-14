package avalid

import (
	"fmt"
	"github.com/asktop/gotools/astring"
)

//检查密码
// level: 密码强度级别
var pwdLevelTips = map[uint]string{
	1: "必须包含数字、字母",
	2: "必须包含数字、字母、下划线",
	3: "必须包含数字、字母、特殊字符",
	4: "必须包含数字、大小写字母",
	5: "必须包含数字、大小写字母、下划线",
	6: "必须包含数字、大小写字母、特殊字符",
}

type isPwd struct {
	title    string
	value    interface{}
	valueStr string
	msgs     []string
	level    uint
	length   []uint
}

func (c *isPwd) Check() (msg string, ok bool) {
	if c.valueStr == "" {
		return "", true
	}
	if len(c.msgs) > 0 {
		msg = c.msgs[0]
	}
	levelTip := "不符合要求"
	if tip, ok := pwdLevelTips[c.level]; ok {
		levelTip = tip
	}
	if len(c.length) == 0 {
		if !astring.IsPwd(c.valueStr, c.level) {
			if len(c.msgs) == 0 {
				msg = fmt.Sprintf("%s%s", c.title, levelTip)
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
				msg = fmt.Sprintf("%s%s，且长度必须为 %s", c.title, levelTip, lenStr)
			}
			return msg, false
		}
	}
	return "", true
}
