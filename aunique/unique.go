package aunique

import (
	"fmt"
	"github.com/asktop/gotools/acast"
	"github.com/asktop/gotools/amap"
	"github.com/asktop/gotools/astring"
	"github.com/asktop/gotools/atime"
	"strconv"
	"strings"
	"sync"
)

var prefixLast = amap.NewStrIfaceMap()

type last struct {
	mu        sync.RWMutex
	prefix    string
	length    int
	timestamp string //时间戳
	number    int    //当前秒内序号
}

// 唯一序号
// @param length 序号长度，不少于16（不包括前缀）
// @param prefix 序号前缀
func UniqueNo(length int, prefix ...string) string {
	if length < 16 || length > 20 {
		panic("UniqueNo length must gte 16 and lte 20")
	}
	//序号前缀
	var prefixStr string
	if len(prefix) > 0 {
		prefixStr = strings.TrimSpace(prefix[0])
	}
	key := prefixStr + strconv.Itoa(length)

	//上一个序号
	var lastNo *last
	lastNo = prefixLast.SetOrGet(key, &last{mu: sync.RWMutex{}, prefix: prefixStr, length: length}).(*last)
	//当前序号
	lastNo.mu.Lock()
	defer lastNo.mu.Unlock()
	nanosecond := atime.Now().Format("060102150405") + fmt.Sprintf("%d", atime.Now().Nanosecond())
	timestamp := astring.Substr(nanosecond, 0, lastNo.length-4)
	if lastNo.timestamp == timestamp {
		lastNo.number++
	} else {
		lastNo.timestamp = timestamp
		lastNo.number = 1
	}
	number := astring.IntToStr(lastNo.number, 3)
	if len(number) > 3 {
		tag := acast.ToUint8(astring.SubstrByEnd(number, 0, -2))
		if tag <= 36 {
			tag += 55 //A:64
		} else {
			tag += 87 //a:97
		}
		number = string(tag) + astring.Substr(number, 0, -2)
	}
	uniqueNo := lastNo.prefix + timestamp + number
	var sum int
	for _, c := range uniqueNo {
		sum += int(c)
	}
	uniqueNo += strconv.Itoa(sum % 10)
	return uniqueNo
}
