package atime

import (
	"fmt"
	"github.com/asktop/gotools/acast"
	"strings"
	"sync"
	"time"
)

const (
	DATETIME = "2006-01-02 15:04:05"
	DATE     = "2006-01-02"
	TIME     = "15:04:05"
	Day      = time.Duration(time.Hour * 24)
)

var mu sync.RWMutex
var offsetTime time.Duration //时间偏移量
var fixedTime *time.Time     //固定时间

//设置时间偏量（改变当前时间）
func Offset(offset time.Duration) {
	mu.Lock()
	defer mu.Unlock()
	offsetTime = offset
}

//设置固定时间（改变当前时间）
func Fixed(fixed *time.Time) {
	mu.Lock()
	defer mu.Unlock()
	fixedTime = fixed
}

//当前时间
func Now() time.Time {
	mu.Lock()
	defer mu.Unlock()
	if fixedTime != nil {
		return *fixedTime
	}
	return time.Now().Add(offsetTime)
}

//获取当前时间 秒级时间戳
func NowUnix() int64 {
	return Now().Unix()
}

//获取当前时间 毫秒级时间戳
func NowMilli() int64 {
	return Now().UnixNano() / 1e6
}

//获取当前时间 纳秒级时间戳
func NowNano() int64 {
	return Now().UnixNano()
}

//秒级时间戳 转换为 毫秒级时间戳
func UnixToMilli(timestamp int64) int64 {
	return timestamp * 1e3
}

//秒级时间戳 转换为 纳秒级时间戳
func UnixToNano(timestamp int64) int64 {
	return timestamp * 1e9
}

//毫秒级时间戳 转换为 秒级时间戳
func MilliToUnix(timestamp int64) int64 {
	return timestamp / 1e3
}

//毫秒级时间戳 转换为 纳秒时间戳
func MilliToNano(timestamp int64) int64 {
	return timestamp * 1e6
}

//纳秒级时间戳 转换为 秒级时间戳
func NanoToUnix(timestamp int64) int64 {
	return timestamp / 1e9
}

//纳秒级时间戳 转换为 毫秒级时间戳
func NanoToMilli(timestamp int64) int64 {
	return timestamp / 1e6
}

//将 时间戳 转换成 本地时区时间
func ParseTimestamp(timestamp interface{}) (time.Time, error) {
	var err error
	var sec, nsec int64
	fn := Now()
	tsStr, err := acast.ToStringE(timestamp)
	if err != nil {
		return fn, err
	}
	tsLen := len(tsStr)
	if tsLen <= 10 {
		sec, err = acast.ToInt64E(tsStr)
		if err != nil {
			return fn, err
		}
	} else if tsLen <= 19 {
		if tsLen < 19 {
			tsStr += strings.Repeat("0", 19-tsLen)
		}
		secStr := tsStr[0:10]
		nsecStr := tsStr[10:]
		sec, err = acast.ToInt64E(secStr)
		if err != nil {
			return fn, err
		}
		nsecStr = strings.TrimLeft(nsecStr, "0")
		if nsecStr == "" {
			nsecStr = "0"
		}
		nsec, err = acast.ToInt64E(nsecStr)
		if err != nil {
			return fn, err
		}
	} else {
		return fn, err
	}
	return time.Unix(sec, nsec), nil
}

//格式化时间戳 格式指定
func FormatTimestamp(format string, timestamp interface{}) string {
	fn, err := ParseTimestamp(timestamp)
	if err != nil {
		fmt.Println("github.com/asktop/gotools/atime ParseTimestamp", "timestamp:", timestamp, "err:", err)
		return ""
	}
	return fn.Format(format)
}

//格式化时间戳 格式："2006-01-02"
func FormatDateT(timestamp interface{}) string {
	return FormatTimestamp(DATE, timestamp)
}

//格式化时间戳 格式："15:04:05"
func FormatTimeT(timestamp interface{}) string {
	return FormatTimestamp(TIME, timestamp)
}

//格式化时间戳 格式："2006-01-02 15:04:05"
func FormatDateTimeT(timestamp interface{}) string {
	return FormatTimestamp(DATETIME, timestamp)
}

//格式化时间 格式指定
func Format(format string, t ...time.Time) string {
	fn := Now()
	if len(t) > 0 {
		fn = t[0]
	}
	return fn.Format(format)
}

//格式化时间 格式："2006-01-02"
func FormatDate(t ...time.Time) string {
	return Format(DATE, t...)
}

//格式化时间 格式："15:04:05"
func FormatTime(t ...time.Time) string {
	return Format(TIME, t...)
}

//格式化时间 格式："2006-01-02 15:04:05"
func FormatDateTime(t ...time.Time) string {
	return Format(DATETIME, t...)
}

func startTime(timestamp ...interface{}) time.Time {
	fn := Now()
	if len(timestamp) > 0 {
		var err error
		fn, err = ParseTimestamp(timestamp[0])
		if err != nil {
			fmt.Println("github.com/asktop/gotools/atime ParseTimestamp", "timestamp:", timestamp, "err:", err)
		}
	}
	return fn
}

//获取 当前时间 或 指定时间戳 的 当前月开始时间
func StartMonth(timestamp ...interface{}) time.Time {
	fn := startTime(timestamp...)
	return time.Date(fn.Year(), fn.Month(), 1, 0, 0, 0, 0, time.Local)
}

//获取 当前时间 或 指定时间戳 的 当天开始时间
func StartDay(timestamp ...interface{}) time.Time {
	fn := startTime(timestamp...)
	return time.Date(fn.Year(), fn.Month(), fn.Day(), 0, 0, 0, 0, time.Local)
}

//获取 当前时间 或 指定时间戳 的 当前小时开始时间
func StartHour(timestamp ...interface{}) time.Time {
	fn := startTime(timestamp...)
	return time.Date(fn.Year(), fn.Month(), fn.Day(), fn.Hour(), 0, 0, 0, time.Local)
}

//获取 当前时间 或 指定时间戳 的 当前分钟开始时间
func StartMinute(timestamp ...interface{}) time.Time {
	fn := startTime(timestamp...)
	return time.Date(fn.Year(), fn.Month(), fn.Day(), fn.Hour(), fn.Minute(), 0, 0, time.Local)
}
