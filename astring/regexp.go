package astring

import (
	"fmt"
	"regexp"
	"strconv"
)

//是身份证号码
func IsIDCard(data string) bool {
	ok, err := regexp.MatchString(`(^[1-9]\d{5}(18|19|([23]\d))\d{2}((0[1-9])|(10|11|12))(([0-2][1-9])|10|20|30|31)\d{3}[0-9Xx]$)|(^[1-9]\d{5}\d{2}((0[1-9])|(10|11|12))(([0-2][1-9])|10|20|30|31)\d{2}$)`, data)
	if err != nil {
		fmt.Println(err)
		return false
	}
	return ok
}

//是手机号码
func IsPhone(data string) bool {
	ok, err := regexp.MatchString(`^1[34578]\d{9}$`, data)
	if err != nil {
		fmt.Println(err)
		return false
	}
	return ok
}

//是电话号码
func IsTel(data string) bool {
	ok, err := regexp.MatchString(`^0\d{2,3}-\d{7,8}$`, data)
	if err != nil {
		fmt.Println(err)
		return false
	}
	return ok
}

//是手机或电话号码
func IsTelOrPhone(data string) bool {
	ok, err := regexp.MatchString(`^((0\d{2,3}-\d{7,8})|(1[34578]\d{9}))$`, data)
	if err != nil {
		fmt.Println(err)
		return false
	}
	return ok
}

//是Email地址
func IsEmail(data string) bool {
	ok, err := regexp.MatchString(`^\w+([-+.]\w+)*@\w+([-.]\w+)*\.\w+([-.]\w+)*$`, data)
	if err != nil {
		fmt.Println(err)
		return false
	}
	return ok
}

//是数字字母下划线
// @param length 长度验证： 1个值时为指定长度；2个值时分别为 min 和 max
func IsNum_EN(data string, length ...int) bool {
	if len(length) >= 2 && length[0] > 0 && length[1] > 0 && length[0] < length[1] {
		ok, err := regexp.MatchString(`^[0-9A-Za-z_]{`+strconv.Itoa(length[0])+`,`+strconv.Itoa(length[1])+`}$`, data)
		if err != nil {
			fmt.Println(err)
			return false
		}
		return ok
	} else if len(length) >= 1 && length[0] > 0 {
		ok, err := regexp.MatchString(`^[0-9A-Za-z_]{`+strconv.Itoa(length[0])+`}$`, data)
		if err != nil {
			fmt.Println(err)
			return false
		}
		return ok
	} else {
		ok, err := regexp.MatchString("^[0-9A-Za-z_]+$", data)
		if err != nil {
			fmt.Println(err)
			return false
		}
		return ok
	}
}

//是全中文汉字
func IsCN(data string) bool {
	ok, err := regexp.MatchString("^[\u4e00-\u9fa5]+$", data)
	if err != nil {
		fmt.Println(err)
		return false
	}
	return ok
}

//包含中文汉字
func HasCN(data string) bool {
	ok, err := regexp.MatchString("[\u4e00-\u9fa5]", data)
	if err != nil {
		fmt.Println(err)
		return false
	}
	return ok
}

//是全英文字母
// @param length 长度验证： 1个值时为指定长度；2个值时分别为 min 和 max
func IsEN(data string, length ...int) bool {
	if len(length) >= 2 && length[0] > 0 && length[1] > 0 && length[0] < length[1] {
		ok, err := regexp.MatchString(`^[A-Za-z]{`+strconv.Itoa(length[0])+`,`+strconv.Itoa(length[1])+`}$`, data)
		if err != nil {
			fmt.Println(err)
			return false
		}
		return ok
	} else if len(length) >= 1 && length[0] > 0 {
		ok, err := regexp.MatchString(`^[A-Za-z]{`+strconv.Itoa(length[0])+`}$`, data)
		if err != nil {
			fmt.Println(err)
			return false
		}
		return ok
	} else {
		ok, err := regexp.MatchString("^[A-Za-z]+$", data)
		if err != nil {
			fmt.Println(err)
			return false
		}
		return ok
	}
}

//包含英文字母
func HasEN(data string) bool {
	ok, err := regexp.MatchString("[A-Za-z]", data)
	if err != nil {
		fmt.Println(err)
		return false
	}
	return ok
}

//包含标点字符
func HasChar(data string) bool {
	ok, err := regexp.MatchString(`[\.~!@#$%^&*()\-=_+:;,?]`, data)
	if err != nil {
		fmt.Println(err)
		return false
	}
	return ok
}

//是数字
// @param length 长度验证： 1个值时为指定长度；2个值时分别为 min 和 max
func IsNumber(data string, length ...int) bool {
	if len(length) >= 2 && length[0] > 0 && length[1] > 0 && length[0] < length[1] {
		ok, err := regexp.MatchString(`^[0-9]{`+strconv.Itoa(length[0])+`,`+strconv.Itoa(length[1])+`}$`, data)
		if err != nil {
			fmt.Println(err)
			return false
		}
		return ok
	} else if len(length) >= 1 && length[0] > 0 {
		ok, err := regexp.MatchString(`^[0-9]{`+strconv.Itoa(length[0])+`}$`, data)
		if err != nil {
			fmt.Println(err)
			return false
		}
		return ok
	} else {
		ok, err := regexp.MatchString(`^[0-9]+$`, data)
		if err != nil {
			fmt.Println(err)
			return false
		}
		return ok
	}
}

//包含数字
func HasNumber(data string) bool {
	ok, err := regexp.MatchString(`[0-9]`, data)
	if err != nil {
		fmt.Println(err)
		return false
	}
	return ok
}

//是实数
// @param scale 小数位长度验证： 1个值时为指定长度；2个值时分别为 min 和 max
func IsAllDecimal(data string, scale ...int) bool {
	if len(scale) >= 2 && scale[0] >= 0 && scale[1] > 0 && scale[0] < scale[1] {
		if scale[0] == 0 {
			ok, err := regexp.MatchString(`^-?(([1-9]\d*)|0)([.]\d{1,`+strconv.Itoa(scale[1])+`})?$`, data)
			if err != nil {
				return false
			}
			return ok
		} else {
			ok, err := regexp.MatchString(`^-?(([1-9]\d*)|0)[.]\d{`+strconv.Itoa(scale[0])+`,`+strconv.Itoa(scale[1])+`}$`, data)
			if err != nil {
				return false
			}
			return ok
		}
	} else if len(scale) >= 1 && scale[0] > 0 {
		ok, err := regexp.MatchString(`^-?(([1-9]\d*)|0)[.]\d{`+strconv.Itoa(scale[0])+`}$`, data)
		if err != nil {
			return false
		}
		return ok
	} else {
		ok, err := regexp.MatchString(`^-?(([1-9]\d*)|0)([.]\d+)?$`, data)
		if err != nil {
			fmt.Println(err)
			return false
		}
		return ok
	}
}

//是非负实数
// @param scale 小数位长度验证： 1个值时为指定长度；2个值时分别为 min 和 max
func IsDecimal(data string, scale ...int) bool {
	if len(scale) >= 2 && scale[0] >= 0 && scale[1] > 0 && scale[0] < scale[1] {
		if scale[0] == 0 {
			ok, err := regexp.MatchString(`^(([1-9]\d*)|0)([.]\d{1,`+strconv.Itoa(scale[1])+`})?$`, data)
			if err != nil {
				fmt.Println(err)
				return false
			}
			return ok
		} else {
			ok, err := regexp.MatchString(`^(([1-9]\d*)|0)[.]\d{`+strconv.Itoa(scale[0])+`,`+strconv.Itoa(scale[1])+`}$`, data)
			if err != nil {
				fmt.Println(err)
				return false
			}
			return ok
		}
	} else if len(scale) >= 1 && scale[0] > 0 {
		ok, err := regexp.MatchString(`^(([1-9]\d*)|0)[.]\d{`+strconv.Itoa(scale[0])+`}$`, data)
		if err != nil {
			fmt.Println(err)
			return false
		}
		return ok
	} else {
		ok, err := regexp.MatchString(`^(([1-9]\d*)|0)([.]\d+)?$`, data)
		if err != nil {
			fmt.Println(err)
			return false
		}
		return ok
	}
}

//是整数
func IsAllInt(data string) bool {
	ok, err := regexp.MatchString(`^((-?([1-9]\d*))|0)$`, data)
	if err != nil {
		fmt.Println(err)
		return false
	}
	return ok
}

//是非负整数
func IsInt(data string) bool {
	ok, err := regexp.MatchString(`^(([1-9]\d*)|0)$`, data)
	if err != nil {
		fmt.Println(err)
		return false
	}
	return ok
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
	ok, err := regexp.MatchString(pattern, data)
	if err != nil {
		fmt.Println(err)
		return false
	}
	return ok
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
	ok, err := regexp.MatchString(pattern, data)
	if err != nil {
		fmt.Println(err)
		return false
	}
	return ok
}
