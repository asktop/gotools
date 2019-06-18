package astring

import (
	"github.com/asktop/gotools/acast"
	"github.com/asktop/gotools/ajson"
	"reflect"
	"regexp"
	"strconv"
	"strings"
	"unicode/utf8"
)

func Len(s string) int {
	return len([]rune(s))
}

//截取字符串
// @param length 不设置：截取全部；负数：向前截取
func Substr(s string, start int, length ...int) string {
	rs := []rune(s)
	l := len(rs)
	if len(length) > 0 {
		l = length[0]
	}
	if l > 0 {
		if start <= 0 {
			start = 0
		} else {
			if start > len(rs) {
				start = start % len(rs)
			}
		}

		end := start + l
		if start+l > len(rs) {
			end = len(rs)
		}
		return string(rs[start:end])
	} else if l < 0 {
		if start <= 0 {
			start = len(rs)
		} else {
			if start > len(rs) {
				start = start % len(rs)
			}
		}
		end := start

		start = end + l
		if end+l < 0 {
			start = 0
		}
		return string(rs[start:end])
	} else {
		return ""
	}
}

//截取字符串
// @param end 0：截取全部；负数：从后往前
func SubstrByEnd(s string, start int, end int) string {
	rs := []rune(s)

	if start < 0 {
		start = 0
	}
	if start > len(rs) {
		start = start % len(rs)
	}

	if end >= 0 {
		if end < start || end > len(rs) {
			end = len(rs)
		}
	} else {
		if len(rs)+end < start {
			end = len(rs)
		} else {
			end = len(rs) + end
		}
	}

	return string(rs[start:end])
}

//字符串是否相同（不区分大小写）
func EqualNoCase(str1 interface{}, str2 interface{}) bool {
	return strings.ToLower(acast.ToString(str1)) == strings.ToLower(acast.ToString(str2))
}

//替换字符串（不区分大小写）
func ReplaceNoCase(s string, old string, new string, n int) string {
	if n == 0 {
		return s
	}

	ls := strings.ToLower(s)
	lold := strings.ToLower(old)

	if m := strings.Count(ls, lold); m == 0 {
		return s
	} else if n < 0 || m < n {
		n = m
	}

	ns := make([]byte, len(s)+n*(len(new)-len(old)))
	w := 0
	start := 0
	for i := 0; i < n; i++ {
		j := start
		if len(old) == 0 {
			if i > 0 {
				_, wid := utf8.DecodeRuneInString(s[start:])
				j += wid
			}
		} else {
			j += strings.Index(ls[start:], lold)
		}
		w += copy(ns[w:], s[start:j])
		w += copy(ns[w:], new)
		start = j + len(old)
	}
	w += copy(ns[w:], s[start:])
	return string(ns[0:w])
}

//删除字符串两端的空格(含tab)，同时将中间多个空格(含tab)的转换为一个
func TrimSpaceToOne(s string) string {
	s = strings.TrimSpace(s)
	s = strings.Replace(s, "	", " ", -1)      //替换tab为空格
	reg, _ := regexp.Compile("\\s{2,}")          //编译正则表达式
	s2 := make([]byte, len(s))                   //定义字符数组切片
	copy(s2, s)                                  //将字符串复制到切片
	spc_index := reg.FindStringIndex(string(s2)) //在字符串中搜索
	for len(spc_index) > 0 { //找到适配项
		s2 = append(s2[:spc_index[0]+1], s2[spc_index[1]:]...) //删除多余空格
		spc_index = reg.FindStringIndex(string(s2))            //继续在字符串中搜索
	}
	return string(s2)
}

// int 转换成指定长度的 string
// @param force 强制转换，当num长度大于length时，删除前面超过的部分
func IntToStr(num int, length int, force ...bool) string {
	if length <= 0 {
		return strconv.Itoa(num)
	} else {
		if num < 0 {
			numStr := strconv.Itoa(-num)
			if len(force) > 0 && force[0] || len(numStr) < length {
				numStr = strings.Repeat("0", length) + numStr
				return "-" + numStr[len(numStr)-length:]
			} else {
				return "-" + numStr
			}
		} else {
			numStr := strconv.Itoa(num)
			if len(force) > 0 && force[0] || len(numStr) < length {
				numStr = strings.Repeat("0", length) + numStr
				return numStr[len(numStr)-length:]
			} else {
				return numStr
			}
		}
	}
}

//将多个对象拼接成字符串
func Join(args ...interface{}) string {
	var rs string
	for _, arg := range args {
		if err, ok := arg.(error); ok {
			rs += err.Error() + " "
			continue
		}
		argVal := reflect.ValueOf(arg)
		if argVal.Kind() == reflect.Ptr {
			argVal = argVal.Elem()
		}
		if argVal.Kind() == reflect.Struct || argVal.Kind() == reflect.Slice || argVal.Kind() == reflect.Map {
			rs += ajson.Encode(argVal.Interface()) + " "
		} else {
			rs += acast.ToString(argVal.Interface()) + " "
		}
	}
	return strings.TrimSpace(rs)
}

//隐藏字符串
// start：前端显示长度
// end：后端显示长度
// length：指定显示总长度，若不指定，则按原字符串长度输出
func HideNo(s string, start int, end int, length ...int) string {
	s = strings.TrimSpace(s)
	oldLen := len(s)
	newLen := oldLen
	if len(length) > 0 {
		newLen = length[0]
	}
	minLen := oldLen
	if oldLen >= newLen {
		minLen = newLen
	}
	if minLen <= 1 {
		return strings.Repeat("*", newLen)
	}
	if start >= minLen {
		start = minLen - 1
		end = 0
	} else if end >= minLen {
		start = 0
		end = minLen - 1
	} else if start+end >= minLen {
		start = minLen / 2
		end = minLen/2 - 1
	}
	rs := Substr(s, 0, start) + strings.Repeat("*", newLen-start-end) + Substr(s, 0, -end)
	return rs
}

//隐藏 手机号
func HidePhone(s string) string {
	s = strings.TrimSpace(s)
	length := len(s)
	if length == 0 {
		return ""
	}
	if strings.Contains(s, "+") {
		return Substr(s, 0, length-8) + "****" + SubstrByEnd(s, length-4, 0)
	} else {
		if strings.Contains(s, "-") || strings.Contains(s, "_") || strings.Contains(s, " ") {
			return Substr(s, 0, length-6) + "***" + SubstrByEnd(s, length-3, 0)
		} else {
			if length == 11 {
				return Substr(s, 0, 3) + "****" + SubstrByEnd(s, length-4, 0)
			} else {
				return Substr(s, 0, length-6) + "***" + SubstrByEnd(s, length-3, 0)
			}
		}
	}
}

//隐藏 邮箱
func HideEmail(s string) string {
	emails := strings.Split(s, "@")
	if len(emails) != 2 {
		return s
	}
	return HideNo(emails[0], 2, 2, 6) + "@" + emails[1]
}

//隐藏 密码
func HidePwd(s string, allHide ...bool) string {
	s = strings.TrimSpace(s)
	if (len(allHide) > 0 && allHide[0]) {
		return "******"
	} else {
		if len(s) > 0 {
			return "******"
		} else {
			return ""
		}
	}
}

//转换成 首字母大写
func ToFirstUpper(s string) string {
	s = strings.TrimSpace(s)
	if s != "" {
		s = strings.ToUpper(s[:1]) + s[1:]
	}
	return s
}

//转换成 首字母小写
func ToFirstLower(s string) string {
	s = strings.TrimSpace(s)
	if s != "" {
		s = strings.ToLower(s[:1]) + s[1:]
	}
	return s
}

//转换成 大驼峰命名（UserId）
func ToCamelCase(s string) string {
	if IsNum_EN(s) {
		var rs string
		s = strings.TrimSpace(s)
		es := strings.Split(s, "_")
		for _, e := range es {
			rs += ToFirstUpper(e)
		}
		return rs
	} else {
		return s
	}
}

//转换成 小驼峰命名（userId）
func TocamelCase(s string) string {
	return ToFirstLower(ToCamelCase(s))
}

//转换成 大下划线命名（USER_ID）
func ToUnderscoreCase(s string) string {
	return strings.ToUpper(TounderscoreCase(s))
}

//转换成 小下划线命名（user_id）
func TounderscoreCase(s string) string {
	if IsNum_EN(s) {
		var rs string
		l := len(s)
		for i := 0; i < l; i++ {
			e := s[i : i+1]
			if IsUpperChar(e) {
				e = "_" + strings.ToLower(e)
			}
			rs += e
		}
		rs = strings.TrimPrefix(rs, "_")
		rs = strings.Replace(rs, "__", "_", -1)
		return rs
	} else {
		return s
	}
}
