package aunique

import (
	"fmt"
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
	prefix    string //序号前缀
	length    int    //数字序号长度
	timestamp string //时间戳
	number    int    //当前时间戳内序号
}

// 唯一序号
// @param length 数字序号长度，20到23位（不包括前缀，12位年月日时分秒+n位时间戳+4位同时间戳内序号+1位验证码）
// @param prefix 序号前缀
func UniqueNo(numLength int, prefix ...string) string {
	numberLen := 4 //同时间戳内序号长度
	if numLength < 20 || numLength > 23 {
		panic("UniqueNo numLength must gte 20 and lte 23")
	}
	//序号前缀
	var prefixStr string
	if len(prefix) > 0 {
		prefixStr = strings.TrimSpace(prefix[0])
	}
	key := prefixStr + strconv.Itoa(numLength)

	//上一个序号
	lastNo := prefixLast.SetOrGet(key, &last{mu: sync.RWMutex{}, prefix: prefixStr, length: numLength}).(*last)
	//当前序号
	lastNo.mu.Lock()
	defer lastNo.mu.Unlock()
	//生成当前时间戳 12位年月日时分秒+n位时间戳
	nanosecond := atime.Now().Format("060102150405") + fmt.Sprintf("%d", atime.Now().Nanosecond())
	timestamp := astring.Substr(nanosecond, 0, lastNo.length-numberLen-1)
	if lastNo.timestamp == timestamp {
		lastNo.number++
	} else {
		lastNo.timestamp = timestamp
		lastNo.number = 1
	}
	//生成4位同时间戳内序号
	number := astring.IntToStr(lastNo.number, numberLen)
	if len(number) > numberLen {
		panic("UniqueNo 同时间戳内序号长度超长")
	}
	uniqueNo := lastNo.prefix + timestamp + number
	//生成最后1位校验码
	var sum int
	for _, c := range uniqueNo {
		sum += int(c)
	}
	uniqueNo += strconv.Itoa(sum % 10)
	return uniqueNo
}
