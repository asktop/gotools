package astring

import (
	"fmt"
	"regexp"
)

func MatchString(pattern string, str string) bool {
	ok, err := regexp.MatchString(pattern, str)
	if err != nil {
		fmt.Println(err)
		return false
	}
	return ok
}

/*
	公民身份证号
	xxxxxx yyyy MM dd 375 0     十八位
	xxxxxx   yy MM dd  75 0     十五位

	地区：[1-9]\d{5}
	年的前两位：(18|19|([23]\d))      1800-2399
	年的后两位：\d{2}
	月份：((0[1-9])|(10|11|12))
	天数：(([0-2][1-9])|10|20|30|31) 闰年不能禁止29+

	三位顺序码：\d{3}
	两位顺序码：\d{2}
	校验码：   [0-9Xx]

	十八位：^[1-9]\d{5}(18|19|([23]\d))\d{2}((0[1-9])|(10|11|12))(([0-2][1-9])|10|20|30|31)\d{3}[0-9Xx]$
	十五位：^[1-9]\d{5}\d{2}((0[1-9])|(10|11|12))(([0-2][1-9])|10|20|30|31)\d{3}$

	总：
	(^[1-9]\d{5}(18|19|([23]\d))\d{2}((0[1-9])|(10|11|12))(([0-2][1-9])|10|20|30|31)\d{3}[0-9Xx]$)|(^[1-9]\d{5}\d{2}((0[1-9])|(10|11|12))(([0-2][1-9])|10|20|30|31)\d{3}$)
 */
func IsIDCard(data string) bool {
	return MatchString(`(^[1-9]\d{5}(18|19|([23]\d))\d{2}((0[1-9])|(10|11|12))(([0-2][1-9])|10|20|30|31)\d{3}[0-9Xx]$)|(^[1-9]\d{5}\d{2}((0[1-9])|(10|11|12))(([0-2][1-9])|10|20|30|31)\d{2}$)`, data)
}

/*
	验证所给手机号码是否符合手机号的格式.
	移动: 134、135、136、137、138、139、150、151、152、157、158、159、182、183、184、187、188、178(4G)、147(上网卡)；
	联通: 130、131、132、155、156、185、186、176(4G)、145(上网卡)、175；
	电信: 133、153、180、181、189 、177(4G)；
	卫星通信:  1349
	虚拟运营商: 170、173
	2018新增: 16x, 19x
 */
func IsPhone(data string) bool {
	return MatchString(`^13[\d]{9}$|^14[5,7]{1}\d{8}$|^15[^4]{1}\d{8}$|^16[\d]{9}$|^17[0,3,5,6,7,8]{1}\d{8}$|^18[\d]{9}$|^19[\d]{9}$`, data)
}

//国内座机电话号码："XXXX-XXXXXXX"、"XXXX-XXXXXXXX"、"XXX-XXXXXXX"、"XXX-XXXXXXXX"、"XXXXXXX"、"XXXXXXXX"
func IsTel(data string) bool {
	return MatchString(`^((\d{3,4})|\d{3,4}-)?\d{7,8}$`, data)
}

//手机或电话号码
func IsTelOrPhone(data string) bool {
	if IsTel(data) || IsPhone(data) {
		return true
	}
	return false
}

//Email地址
func IsEmail(data string) bool {
	//return MatchString(`^\w+([-+.]\w+)*@\w+([-.]\w+)*\.\w+([-.]\w+)*$`, data)
	return MatchString(`^[a-zA-Z0-9_\-\.]+@[a-zA-Z0-9_\-]+(\.[a-zA-Z0-9_\-]+)+$`, data)
}

//URL地址
func IsUrl(data string) bool {
	return MatchString(`(https?|ftp|file)://[-A-Za-z0-9+&@#/%?=~_|!:,.;]+[-A-Za-z0-9+&@#/%=~_|]`, data)
}

//Mac地址
func IsMac(data string) bool {
	return MatchString(`^([0-9A-Fa-f]{2}[\-:]){5}[0-9A-Fa-f]{2}$`, data)
}

//腾讯QQ号，从10000开始
func IsQQ(data string) bool {
	return MatchString(`^[1-9][0-9]{4,}$`, data)
}

//邮政编码
func IsPostCode(data string) bool {
	return MatchString(`^\d{6}$`, data)
}

//账号/密码（字母开头，数字字母下划线）
func IsAccount(data string, length ...int) bool {
	if len(length) >= 2 && length[0] > 0 && length[1] > 0 && length[0] < length[1] {
		return MatchString(fmt.Sprintf(`^[A-Za-z]{1}[0-9A-Za-z_]{%d,%d}$`, length[0]-1, length[1]-1), data)
	} else if len(length) >= 1 && length[0] > 0 {
		return MatchString(fmt.Sprintf(`^[A-Za-z]{1}[0-9A-Za-z_]{%d}$`, length[0]-1), data)
	} else {
		return MatchString("^[A-Za-z]{1}[0-9A-Za-z_]+$", data)
	}
}

//是数字字母下划线
// @param length 长度验证： 1个值时为指定长度；2个值时分别为 min 和 max
func IsNum_EN(data string, length ...int) bool {
	if len(length) >= 2 && length[0] > 0 && length[1] > 0 && length[0] < length[1] {
		return MatchString(fmt.Sprintf(`^[0-9A-Za-z_]{%d,%d}$`, length[0], length[1]), data)
	} else if len(length) >= 1 && length[0] > 0 {
		return MatchString(fmt.Sprintf(`^[0-9A-Za-z_]{%d}$`, length[0]), data)
	} else {
		return MatchString("^[0-9A-Za-z_]+$", data)
	}
}

//包含数字字母下划线
func HasNum_EN(data string) bool {
	return HasNumber(data) && HasEN(data) && MatchString("[_]", data)
}

//是全中文汉字
func IsCN(data string) bool {
	return MatchString("^[\u4e00-\u9fa5]+$", data)
}

//包含中文汉字
func HasCN(data string) bool {
	return MatchString("[\u4e00-\u9fa5]", data)
}

//是全英文字母
// @param length 长度验证： 1个值时为指定长度；2个值时分别为 min 和 max
func IsEN(data string, length ...int) bool {
	if len(length) >= 2 && length[0] > 0 && length[1] > 0 && length[0] < length[1] {
		return MatchString(fmt.Sprintf(`^[A-Za-z]{%d,%d}$`, length[0], length[1]), data)
	} else if len(length) >= 1 && length[0] > 0 {
		return MatchString(fmt.Sprintf(`^[A-Za-z]{%d}$`, length[0]), data)
	} else {
		return MatchString("^[A-Za-z]+$", data)
	}
}

//包含英文字母
func HasEN(data string) bool {
	return MatchString("[A-Za-z]", data)
}

//是大写字母
func IsUpperChar(char string) bool {
	return MatchString("^[A-Z]$", char)
}

//包含大写字母
func HasUpperChar(str string) bool {
	return MatchString("[A-Z]", str)
}

//是小写字母
func IsLowerChar(char string) bool {
	return MatchString("^[a-z]$", char)
}

//包含小写字母
func HasLowerChar(str string) bool {
	return MatchString("[a-z]", str)
}

//包含标点字符
func HasChar(data string) bool {
	return MatchString(`[\.~!@#$%^&*()\-=_+:;,?]`, data)
}

//是数字
// @param length 长度验证： 1个值时为指定长度；2个值时分别为 min 和 max
func IsNumber(data string, length ...int) bool {
	if len(length) >= 2 && length[0] > 0 && length[1] > 0 && length[0] < length[1] {
		return MatchString(fmt.Sprintf(`^[0-9]{%d,%d}$`, length[0], length[1]), data)
	} else if len(length) >= 1 && length[0] > 0 {
		return MatchString(fmt.Sprintf(`^[0-9]{%d}$`, length[0]), data)
	} else {
		return MatchString(`^[0-9]+$`, data)
	}
}

//包含数字
func HasNumber(data string) bool {
	return MatchString(`[0-9]`, data)
}

//是实数
// @param scale 小数位长度验证： 1个值时为指定长度；2个值时分别为 min 和 max
func IsAllDecimal(data string, scale ...int) bool {
	if len(scale) >= 2 && scale[0] >= 0 && scale[1] > 0 && scale[0] < scale[1] {
		if scale[0] == 0 {
			return MatchString(fmt.Sprintf(`^-?(([1-9]\d*)|0)([.]\d{1,%d})?$`, scale[1]), data)
		} else {
			return MatchString(fmt.Sprintf(`^-?(([1-9]\d*)|0)[.]\d{%d,%d}$`, scale[0], scale[1]), data)
		}
	} else if len(scale) >= 1 && scale[0] > 0 {
		return MatchString(fmt.Sprintf(`^-?(([1-9]\d*)|0)[.]\d{%d}$`, scale[0]), data)
	} else {
		return MatchString(`^-?(([1-9]\d*)|0)([.]\d+)?$`, data)
	}
}

//是非负实数
// @param scale 小数位长度验证： 1个值时为指定长度；2个值时分别为 min 和 max
func IsDecimal(data string, scale ...int) bool {
	if len(scale) >= 2 && scale[0] >= 0 && scale[1] > 0 && scale[0] < scale[1] {
		if scale[0] == 0 {
			return MatchString(fmt.Sprintf(`^(([1-9]\d*)|0)([.]\d{1,%d})?$`, scale[1]), data)
		} else {
			return MatchString(fmt.Sprintf(`^(([1-9]\d*)|0)[.]\d{%d,%d}$`, scale[0], scale[1]), data)
		}
	} else if len(scale) >= 1 && scale[0] > 0 {
		return MatchString(fmt.Sprintf(`^(([1-9]\d*)|0)[.]\d{%d}$`, scale[0]), data)
	} else {
		return MatchString(`^(([1-9]\d*)|0)([.]\d+)?$`, data)
	}
}

//是整数
func IsAllInt(data string) bool {
	return MatchString(`^((-?([1-9]\d*))|0)$`, data)
}

//是非负整数
func IsInt(data string) bool {
	return MatchString(`^(([1-9]\d*)|0)$`, data)
}

//是全部日期，日月格式可以为1或01
func IsAllDateFormat(data string, sep string) bool {
	var pattern string
	if sep == "" {
		pattern = `^(18|19|2\d|3\d)\d{2}[-]((0?[1-9])|10|11|12)[-](0?[1-9])|([12][0-9])|30|31)$`
	} else if sep == "年" {
		pattern = `^(18|19|2\d|3\d)\d{2}[年]((0?[1-9])|10|11|12)[月]((0?[1-9])|([12][0-9])|30|31)[日]$`
	} else {
		pattern = `^(18|19|2\d|3\d)\d{2}[` + sep + `]((0?[1-9])|10|11|12)[` + sep + `]((0?[1-9])|([12][0-9])|30|31)$`
	}
	return MatchString(pattern, data)
}

//是日期，日月格式必须为01
func IsDateFormat(data string, sep string) bool {
	var pattern string
	if sep == "" {
		pattern = `^(18|19|2\d|3\d)\d{2}((0[1-9])|10|11|12)((0[1-9])|([12][0-9])|30|31)$`
	} else if sep == "年" {
		pattern = `^(18|19|2\d|3\d)\d{2}[年]((0[1-9])|10|11|12)[月]((0[1-9])|([12][0-9])|30|31)[日]$`
	} else {
		pattern = `^(18|19|2\d|3\d)\d{2}[` + sep + `]((0[1-9])|10|11|12)[` + sep + `]((0[1-9])|([12][0-9])|30|31)$`
	}
	return MatchString(pattern, data)
}
