package astring

import (
	"strconv"
	"strings"
	"unicode/utf8"
)

// int 转换成指定长度的 string
func IntToStr(num int, length int) string {
	if length <= 0 {
		return strconv.Itoa(num)
	} else {
		if num < 0 {
			numStr := strings.Repeat("0", length) + strconv.Itoa(-num)
			return "-" + numStr[len(numStr)-length:]
		} else {
			numStr := strings.Repeat("0", length) + strconv.Itoa(num)
			return numStr[len(numStr)-length:]
		}
	}
}

//截取字符串
// @param length 负数：截取全部
func Substr(s string, start int, length int) string {
	rs := []rune(s)

	if start < 0 {
		start = 0
	}
	if start > len(rs) {
		start = start % len(rs)
	}

	var end int
	if length < 0 || start+length > len(rs) {
		end = len(rs)
	} else {
		end = start + length
	}

	return string(rs[start:end])
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

//替换字符串不区分大小写
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

//隐藏 密码
func HidePwd(s string, allHide ...bool) string {
	s = strings.TrimSpace(s)
	if (len(allHide) > 0 && allHide[0]) {
		return "******"
	} else {
		var pwd string
		rs := []rune(s)
		length := len(rs)
		switch length {
		case 0:
			pwd = "******"
		case 1:
			pwd = s + "*****"
		case 2:
			pwd = Substr(s, 0, 1) + "****" + Substr(s, 1, 1)
		default:
			pwd = Substr(s, 0, 2) + "***" + Substr(s, length-2, 1)
		}
		return pwd
	}
}

//隐藏 手机号
func HidePhone(s string) string {
	s = strings.TrimSpace(s)
	length := len(s)
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
	return HidePwd(emails[0]) + "@" + emails[1]
}
