package aunique

import (
	"fmt"
	"github.com/asktop/gotools/astring"
	"github.com/asktop/gotools/atime"
	"strconv"
	"strings"
	"sync"
	"sync/atomic"
)

var prefixLast = map[string]*last{}

type last struct {
	mu        sync.RWMutex
	timestamp string //当前时间戳
	number    int64  //当前时间戳内序号
}

// 唯一序号
// @param length 序号数字部分长度，20到23位（不包括前缀，12位年月日时分秒+n位纳秒时间戳+4位同时间戳内序号+1位验证码）
// @param prefix 序号前缀
func UniqueNo(numLength int, prefix ...string) string {
	//序号前缀
	var prefixStr string
	if len(prefix) > 0 {
		prefixStr = strings.TrimSpace(prefix[0])
	}
	key := prefixStr + strconv.Itoa(numLength)
	//序号数字部分长度
	if numLength < 20 || numLength > 23 {
		panic("UniqueNo numLength must gte 20 and lte 23")
	}
	//同时间戳内序号长度
	sortLength := 4

	//生成当前时间戳 12位年月日时分秒+n位纳秒时间戳
	nanosecond := fmt.Sprintf("%d", atime.Now().Nanosecond())
	timestamp := atime.Now().Format("060102150405") + astring.Substr(nanosecond, 0, numLength-12-sortLength-1)
	var sort int64 = 1

	//上一个序号
	if _, ok := prefixLast[key]; !ok {
		prefixLast[key] = &last{mu: sync.RWMutex{}}
	}
	lastNo := prefixLast[key]
	//更新上一个序号为当前序号
	lastNo.mu.Lock()
	if lastNo.timestamp != timestamp {
		lastNo.timestamp = timestamp
		atomic.StoreInt64(&lastNo.number, 1)
	} else {
		atomic.AddInt64(&lastNo.number, 1)
		sort = atomic.LoadInt64(&lastNo.number)
	}
	lastNo.mu.Unlock()

	//生成4位同时间戳内序号
	sortStr := astring.IntToStr(int(sort), sortLength)
	if len(sortStr) > sortLength {
		panic("UniqueNo 同时间戳内序号长度超长")
	}
	uniqueNo := prefixStr + timestamp + sortStr
	//生成最后1位校验码
	uniqueNo += getChechNo(uniqueNo)
	return uniqueNo
}

// 校验唯一序号是否合法
func CheckUniqueNo(uniqueNo string) bool {
	source := astring.Substr(uniqueNo, 0, len(uniqueNo)-1)
	checkNo := astring.Substr(uniqueNo, 0, -1)
	return checkNo == getChechNo(source)
}

// 生成最后1位校验码
func getChechNo(source string) string {
	var sum int
	for _, c := range source {
		sum += int(c)
	}
	return strconv.Itoa(sum % 10)
}
