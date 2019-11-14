package avalid

import (
	"fmt"
	"github.com/asktop/gotools/astring"
)

//必须为银行卡号
type isBankCard struct {
	title    string
	value    interface{}
	valueStr string
	msgs     []string
}

func (c *isBankCard) Check() (msg string, ok bool) {
	if c.valueStr == "" {
		return "", true
	}
	if len(c.msgs) > 0 {
		msg = c.msgs[0]
	}
	if !astring.IsBankCard(c.valueStr) {
		if len(c.msgs) == 0 {
			msg = fmt.Sprintf("%s不是正确的银行卡号格式", c.title)
		}
		return msg, false
	}
	return "", true
}
