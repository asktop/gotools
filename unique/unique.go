package unique

import (
	"github.com/asktop/gotools/astring"
	"github.com/asktop/gotools/atime"
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
