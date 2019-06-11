package arand

import (
	"encoding/hex"
	"github.com/asktop/gotools/astring"
	"math/rand"
	"strings"
)

//生成在 min 和 max 之间的随机数（包含 min 和 max）
func Random(min int64, max int64) int64 {
	if min == 0 || max == 0 || min >= max {
		return max
	}
	return rand.Int63n(max-min+1) + min
}

//随机md5字符串 32位
func RandMd5() string {
	data := make([]byte, 16)
	rand.Read(data)
	return hex.EncodeToString(data)
}

//随机字符串 指定长度
// @param sources 数据源 默认（0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz）
// 	空:数字+大小写字母
// 	0:只数字
//	a:只小写字母
//	A:只大写字母
//	Aa:大小写字母
//	_:数字+大小写字母+下划线
//	其他:自定义
func RandStr(length int, sources ...string) string {
	var source string
	if len(sources) > 0 {
		typ := sources[0]
		switch typ {
		case "0":
			source = "0123456789"
		case "A":
			source = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
		case "a":
			source = "abcdefghijklmnopqrstuvwxyz"
		case "Aa":
			source = "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz"
		case "_":
			source = "0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz_"
		default:
			source = strings.Join(sources, "")
		}
	} else {
		source = "0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz"
	}
	sourceLen := len([]rune(source))
	if sourceLen == 0 {
		return ""
	}

	result := []string{}
	for i := 0; i < length; i++ {
		result = append(result, astring.Substr(source, rand.Intn(sourceLen-1), 1))
	}
	return strings.Join(result, "")
}
