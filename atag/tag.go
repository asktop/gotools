package atag

import (
	"bufio"
	"fmt"
	"strings"
)

//为结构添加tag标签， 支持"json", "form", "db", "gorm"等
func New(in string, tags ...string) (string) {
	//获取处理要添加的tag标签，默认json
	if len(tags) == 0 {
		tags = append(tags, "json")
	}

	var result string
	scanner := bufio.NewScanner(strings.NewReader(in))
	for scanner.Scan() {
		oldLineTmp := strings.Trim(scanner.Text(), " ")
		if oldLineTmp == "" {
			continue
		}
		lineTmp := ""
		noteTmp := ""
		if strings.Contains(oldLineTmp, "//") && strings.Index(oldLineTmp, "//") == 0 {
			//只有注释时原样输出
			result = result + oldLineTmp + "\n"
			continue
		} else if strings.Contains(oldLineTmp, "//") && strings.Index(oldLineTmp, "//") > 0 {
			//分割属性和注释
			lineTmp = tagSubstr(oldLineTmp, 0, strings.Index(oldLineTmp, "//"))
			noteTmp = "\t" + tagSubstr(oldLineTmp, strings.Index(oldLineTmp, "//"), -1)
		} else {
			lineTmp = oldLineTmp
		}

		//结构体原样输出
		if strings.Contains(lineTmp, "{") || strings.Contains(lineTmp, "}") {
			result = result + lineTmp + noteTmp + "\n"
			continue
		}

		seperateArr := tagSplit(lineTmp, " ")
		//继承的父类不参与tag
		if len(seperateArr) == 1 {
			result = result + lineTmp + noteTmp + "\n"
			continue
		}

		//组装tag
		propertyTmp := tagHumpToUnderLine(seperateArr[0])
		allTagTmp := ""
		for _, tag := range tags {
			allTagTmp = allTagTmp + fmt.Sprintf("%s:\"%s\" ", tag, propertyTmp)
		}
		allTagTmp = strings.Trim(allTagTmp, " ")
		tagTmp := fmt.Sprintf("\t`%s`", allTagTmp)
		result = result + lineTmp + tagTmp + noteTmp + "\n"
	}
	return result
}

//截取字符串
// @param length 负数：截取全部
func tagSubstr(s string, start int, length int) string {
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

// 增强型split，对  a,,,,,,,b,,c     以","进行切割成[a,b,c]
func tagSplit(s string, sub string) []string {
	var rs = make([]string, 0, 20)
	tmp := ""
	tagSplit2(s, sub, &tmp, &rs)
	return rs
}

// 附属于Split，可独立使用
func tagSplit2(s string, sub string, tmp *string, rs *[]string) {
	s = strings.Trim(s, sub)
	if !strings.Contains(s, sub) {
		*tmp = s
		*rs = append(*rs, *tmp)
		return
	}
	for i, _ := range s {
		if string(s[i]) == sub {
			*tmp = s[:i]
			*rs = append(*rs, *tmp)
			s = s[i+1:]
			tagSplit2(s, sub, tmp, rs)
			return
		}
	}
}

// 驼峰转下划线
func tagHumpToUnderLine(s string) string {
	if s == "ID" {
		return "id"
	}
	var rs string
	elements := tagFindUpperElement(s)
	for _, e := range elements {
		s = strings.Replace(s, e, "_"+strings.ToLower(e), -1)
	}
	rs = strings.Trim(s, " ")
	rs = strings.Trim(rs, "\t")
	return strings.Trim(rs, "_")
}

// 找到字符串中大写字母的列表,附属于HumpToUnderLine
func tagFindUpperElement(s string) []string {
	var rs = make([]string, 0, 10)
	for i := range s {
		if s[i] >= 65 && s[i] <= 90 {
			rs = append(rs, string(s[i]))
		}
	}
	return rs
}
