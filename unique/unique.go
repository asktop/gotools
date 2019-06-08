package unique

import (
	"encoding/hex"
	"github.com/asktop/gotools/astring"
	"github.com/asktop/gotools/atime"
	"math/rand"
	"strings"
	"sync"
)

var (
	emptyLast  = &last{mu: sync.RWMutex{}}
	prefixLast = map[string]*last{}
)

type last struct {
	mu     sync.RWMutex
	second string //秒
	number int    //当前秒内序号
}

// 唯一序号
// @param length 序号长度，不少于15（不包括前缀）
// @param prefix 序号前缀
func UniqueNo(length int, prefix ...string) string {
	if length < 15 {
		panic("length must gte 15")
	}
	//序号前缀
	var prefixStr string
	if len(prefix) > 0 {
		prefixStr = strings.TrimSpace(prefix[0])
	}
	//上一个序号
	var lastNo *last
	if prefixStr == "" {
		lastNo = emptyLast
	} else {
		if LastTemp, ok := prefixLast[prefixStr]; ok {
			lastNo = LastTemp
		} else {
			lastNo = &last{mu: sync.RWMutex{}}
			prefixLast[prefixStr] = lastNo
		}
	}
	//当前序号
	lastNo.mu.Lock()
	defer lastNo.mu.Unlock()
	second := atime.Now().Format("060102150405")
	if lastNo.second == second {
		lastNo.number++
	} else {
		lastNo.second = second
		lastNo.number = 1
	}

	return prefixStr + second + astring.IntToStr(lastNo.number, length-12)
}

//随机字符串 指定长度
// @param sources 数据源
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

//随机md5字符串 32位
func RandMd5() string {
	data := make([]byte, 16)
	rand.Read(data)
	return hex.EncodeToString(data)
}
