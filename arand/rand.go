package arand

import (
	"encoding/base32"
	"encoding/base64"
	"encoding/hex"
	"math/rand"
	"strings"
)

//生成在 min 和 max 之间的随机数（包含 min 和 max）
func Rand(min int64, max int64) int64 {
	if min >= max {
		return min
	}
	return rand.Int63n(max-min+1) + min
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
	rs := make([]rune, length)
	var source []rune
	if len(sources) > 0 {
		typ := sources[0]
		switch typ {
		case "0":
			source = []rune("0123456789")
		case "A":
			source = []rune("ABCDEFGHIJKLMNOPQRSTUVWXYZ")
		case "a":
			source = []rune("abcdefghijklmnopqrstuvwxyz")
		case "Aa":
			source = []rune("ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz")
		case "aA":
			source = []rune("ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz")
		case "_":
			source = []rune("0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz_")
		default:
			source = []rune(strings.Join(sources, ""))
		}
	} else {
		source = []rune("0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz")
	}
	sourceLen := len(source)
	if sourceLen == 0 {
		return ""
	}

	for i := range rs {
		rs[i] = source[rand.Intn(sourceLen)]
	}

	return string(rs)
}

//随机md5字符串 32位 52fdfc072182654f163f5f0f9a621d72
func RandMd5() string {
	data := make([]byte, 16)
	rand.Read(data)
	return hex.EncodeToString(data)
}

//随机base32字符串 32位 SVTMOTIQAN6E2653AQD5DYWGJE======
func RandBase32() string {
	data := make([]byte, 16)
	rand.Read(data)
	return base32.StdEncoding.EncodeToString(data)
}

//随机base32字符串 26位 SVTMOTIQAN6E2653AQD5DYWGJE
func RandBase32Trim() string {
	data := make([]byte, 16)
	rand.Read(data)
	randStr := base32.StdEncoding.EncodeToString(data)
	return strings.TrimRight(randStr, "=")
}

//随机base64字符串 24位 gYVa2GgdDYbR6R4AFnk5yw==
func RandBase64() string {
	data := make([]byte, 16)
	rand.Read(data)
	return base64.StdEncoding.EncodeToString(data)
}

//随机base64字符串 22位 gYVa2GgdDYbR6R4AFnk5yw
func RandBase64Trim() string {
	data := make([]byte, 16)
	rand.Read(data)
	randStr := base64.StdEncoding.EncodeToString(data)
	return strings.TrimRight(randStr, "=")
}
