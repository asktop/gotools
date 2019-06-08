package astring

import (
	"strconv"
	"strings"
	"unicode/utf8"
)

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

// int 转换成指定长度的 string
func IntToStr(num int, length int) string {
	if length <= 0 {
		return "0"
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
