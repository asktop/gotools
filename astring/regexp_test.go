package astring

import (
	"fmt"
	"regexp"
	"strconv"
	"testing"
)

func TestIsNum_EN(t *testing.T) {
	fmt.Println(IsNum_EN("jasdf_2934_"))
	fmt.Println(IsNum_EN("jasdf_29.34_"))
}

func TestIsAllDecimal(t *testing.T) {
	fmt.Println(IsAllDecimal("213123_2934"))
	fmt.Println(IsAllDecimal("-213123.2934"))
	fmt.Println(IsAllDecimal("213123.293400", 8))
	fmt.Println(IsAllDecimal("-213123.293400", 6))
	fmt.Println(IsAllDecimal("213123.293400", 3, 8))
	fmt.Println(IsAllDecimal("-213123.293400", 3, 6))
	fmt.Println(IsAllDecimal("-213123", 0))
	fmt.Println(IsAllDecimal("-213123", 0, 1))
	fmt.Println(IsAllDecimal("-213123.", 0, 1))
	fmt.Println(IsAllDecimal("213123"))
}

func TestIsAllDateFormat(t *testing.T) {
	data1 := "2018-2-1"
	data2 := "1918.01.1"
	data3 := "2018年12月30日"
	data4 := "2018/10/1"
	fmt.Println(IsAllDateFormat(data1, "-"))
	fmt.Println(IsAllDateFormat(data2, "."))
	fmt.Println(IsAllDateFormat(data3, "/"))
	fmt.Println(IsAllDateFormat(data4, "/"))
}

//匹配并替换
func TestRegReplace(t *testing.T) {
	data := "John: 2578.34 William: 4567.23 Steve: 5632.18"
	pattern := "[0-9]+.[0-9]+"
	repl := "##.#"

	//将匹配到的浮点数替换为"##.#"
	str := regexp.MustCompile(pattern).ReplaceAllString(data, repl)
	fmt.Println(str)

	//将匹配到的浮点数替换为乘以2的浮点数
	f := func(s string) string {
		v, _ := strconv.ParseFloat(s, 32)
		return strconv.FormatFloat(v*2, 'f', 2, 32)
	}
	str2 := regexp.MustCompile(pattern).ReplaceAllStringFunc(data, f)
	fmt.Println(str2)
}

//匹配并获取第一条匹配结果
func TestFindString(t *testing.T) {
	pattern := "共(\\d+)页"
	data := `<td height="2">共12页&nbsp;1300条&nbsp;共23页<u>首页</u>&nbsp;<u>上一页</u>&nbsp;&nbsp;<a href=ClassList-42-1.html><u><font color=red>1</u></font></a>&nbsp;&nbsp;<a href=ClassList-42-2.html><u>2</u></a>&nbsp;&nbsp;<a href=ClassList-42-3.html><u>3</u></a>&nbsp;&nbsp;<a href=ClassList-42-4.html><u>4</u></a>&nbsp;&nbsp;<a href=ClassList-42-5.html><u>5</u></a>&nbsp;&nbsp;<a href=ClassList-42-6.html><u>6</u></a>&nbsp;&nbsp;<a href=ClassList-42-7.html><u>7</u></a>&nbsp;&nbsp;<a href=ClassList-42-8.html><u>8</u></a>&nbsp;&nbsp;<a href=ClassList-42-9.html><u>9</u></a>&nbsp;&nbsp;<a href=ClassList-42-10.html><u>10</u></a>&nbsp;&nbsp;<a href=ClassList-42-2.html><u>下一页</u></a>&nbsp;<a href=ClassList-42-12.html><u>尾页</u></a>&nbsp;114条/页&nbsp;</td>`
	result1 := regexp.MustCompile(pattern).FindString(data)
	fmt.Println(result1)

	//查看与上面的区别，若想提取出正则匹配结果，必须加括号
	result2 := regexp.MustCompile(pattern).FindStringSubmatch(data)
	fmt.Println(result2)
	fmt.Println(result2[1])
}

//匹配并获取所有匹配结果
func TestFindAllString(t *testing.T) {
	pattern := "共(\\d+)页"
	data := `<td height="2">共12页&nbsp;1300条&nbsp;共23页<u>首页</u>&nbsp;<u>上一页</u>&nbsp;&nbsp;<a href=ClassList-42-1.html><u><font color=red>1</u></font></a>&nbsp;&nbsp;<a href=ClassList-42-2.html><u>2</u></a>&nbsp;&nbsp;<a href=ClassList-42-3.html><u>3</u></a>&nbsp;&nbsp;<a href=ClassList-42-4.html><u>4</u></a>&nbsp;&nbsp;<a href=ClassList-42-5.html><u>5</u></a>&nbsp;&nbsp;<a href=ClassList-42-6.html><u>6</u></a>&nbsp;&nbsp;<a href=ClassList-42-7.html><u>7</u></a>&nbsp;&nbsp;<a href=ClassList-42-8.html><u>8</u></a>&nbsp;&nbsp;<a href=ClassList-42-9.html><u>9</u></a>&nbsp;&nbsp;<a href=ClassList-42-10.html><u>10</u></a>&nbsp;&nbsp;<a href=ClassList-42-2.html><u>下一页</u></a>&nbsp;<a href=ClassList-42-12.html><u>尾页</u></a>&nbsp;114条/页&nbsp;</td>`
	//匹配并获取所有匹配结果，n为获取匹配的个数，n为-1时获取所有匹配结果
	result1 := regexp.MustCompile(pattern).FindAllString(data, 1)
	for i, result := range result1 {
		fmt.Println(i, ":", result)
	}

	//查看与上面的区别，若想提取出正则匹配结果，必须加括号
	result2 := regexp.MustCompile(pattern).FindAllStringSubmatch(data, -1)
	for i, result := range result2 {
		fmt.Println(i, ":", result)
		fmt.Println(result[1])
	}
}
