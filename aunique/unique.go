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

var (
    lastMu   = new(sync.RWMutex)
    lastPool = map[string]*last{}
)

type last struct {
    mu        sync.RWMutex
    timestamp string //当前时间戳
    number    int64  //当前时间戳内序号
}

// 唯一序号
// @param numLength 序号数字部分长度，20到26位（不包括前缀，14位年月日时分秒+(numLength-20)位纳秒时间戳+5位同时间戳内序号+1位验证码）
// @param prefix 序号前缀
func UniqueNo(numLength int, prefix ...string) string {
    //序号数字部分长度校验
    if numLength < 20 || numLength > 26 {
        panic("UniqueNo numLength must gte 20 and lte 26")
    }
    //序号前缀
    var prefixStr string
    if len(prefix) > 0 {
        prefixStr = strings.TrimSpace(prefix[0])
    }
    //同时间戳内序号长度
    sortLength := 5

    //获取上一个序号
    lastMu.Lock()
    lastKey := prefixStr + strconv.Itoa(numLength)
    if _, ok := lastPool[lastKey]; !ok {
        lastPool[lastKey] = &last{mu: sync.RWMutex{}}
    }
    lastNo := lastPool[lastKey]
    lastMu.Unlock()

    lastNo.mu.Lock()

    //生成当前时间戳 14位年月日时分秒+n位纳秒时间戳
    now := atime.Now()
    nanosecond := astring.Substr(fmt.Sprintf("%d", now.UnixNano()), 10, numLength-14-1-sortLength)
    timestamp := now.Format("20060102150405") + nanosecond

    //更新上一个序号为当前序号
    if lastNo.timestamp < timestamp {
        lastNo.timestamp = timestamp
        atomic.StoreInt64(&lastNo.number, 1)
    } else {
        atomic.AddInt64(&lastNo.number, 1)
    }
    sort := lastNo.number
    lastNo.mu.Unlock()

    //生成4位同时间戳内序号
    sortStr := astring.Int64ToStr(sort, sortLength)
    if len(sortStr) > sortLength {
        return UniqueNo(numLength, prefix...)
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
