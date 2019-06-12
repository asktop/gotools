package axml

import (
	"fmt"
	"github.com/asktop/gotools/acharset"
	"github.com/clbanning/mxj"
	"regexp"
	"strings"
)

// 将XML内容解析为map变量
func Decode(content []byte) (map[string]interface{}, error) {
	res, err := convert(content)
	if err != nil {
		return nil, err
	}
	return mxj.NewMapXml(res)
}

// 将map变量解析为XML格式内容
func Encode(v map[string]interface{}, rootTag ...string) ([]byte, error) {
	return mxj.Map(v).Xml(rootTag...)
}

func EncodeWithIndent(v map[string]interface{}, rootTag ...string) ([]byte, error) {
	return mxj.Map(v).XmlIndent("", "\t", rootTag...)
}

// XML格式内容直接转换为JSON格式内容
func ToJson(content []byte) ([]byte, error) {
	res, err := convert(content)
	if err != nil {
		fmt.Println("convert error. ", err)
		return nil, err
	}

	mv, err := mxj.NewMapXml(res)
	if err == nil {
		return mv.Json()
	} else {
		return nil, err
	}
}

// XML字符集预处理
func convert(xml []byte) (res []byte, err error) {
	patten := `<\?xml.*encoding\s*=\s*['|"](.*?)['|"].*\?>`
	reg, err := regexp.Compile(patten)
	if err != nil {
		return nil, err
	}
	matchStr := reg.FindStringSubmatch(string(xml))
	xmlEncode := "UTF-8"
	if len(matchStr) == 2 {
		xmlEncode = matchStr[1]
	}
	xmlEncode = strings.ToUpper(xmlEncode)
	s := acharset.GetCharset(xmlEncode)
	if s == false {
		return nil, fmt.Errorf("not support charset:%s\n", xmlEncode)
	}

	res = reg.ReplaceAll(xml, []byte(""))
	if xmlEncode != "UTF-8" && xmlEncode != "UTF8" {
		dst, err := acharset.Convert("UTF-8", xmlEncode, string(res))
		if err != nil {
			return nil, err
		}
		res = []byte(dst)
	}
	return res, nil
}
