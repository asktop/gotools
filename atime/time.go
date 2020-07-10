package atime

import (
    "fmt"
    "github.com/asktop/gotools/acast"
    "sync"
    "time"
)

const (
	DATETIME = "2006-01-02 15:04:05"
	DATE     = "2006-01-02"
	TIME     = "15:04:05"
	MONTH    = "2006-01"
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

func toInt64(timestamp interface{}) int64 {
	return acast.ToInt64(timestamp)
}

//秒级时间戳 转换为 毫秒级时间戳
func UnixToMilli(timestamp interface{}) int64 {
	return toInt64(timestamp) * 1e3
}

//秒级时间戳 转换为 纳秒级时间戳
func UnixToNano(timestamp interface{}) int64 {
	return toInt64(timestamp) * 1e9
}

//毫秒级时间戳 转换为 秒级时间戳
func MilliToUnix(timestamp interface{}) int64 {
	return toInt64(timestamp) / 1e3
}

//毫秒级时间戳 转换为 纳秒时间戳
func MilliToNano(timestamp interface{}) int64 {
	return toInt64(timestamp) * 1e6
}

//纳秒级时间戳 转换为 秒级时间戳
func NanoToUnix(timestamp interface{}) int64 {
	return toInt64(timestamp) / 1e9
}

//纳秒级时间戳 转换为 毫秒级时间戳
func NanoToMilli(timestamp interface{}) int64 {
	return toInt64(timestamp) / 1e6
}

//将 时间戳 转换成 本地时区时间
func ParseTimestamp(timestamp interface{}) (time.Time, error) {
	var err error
	fn := Now()
    sec, err := acast.ToInt64E(timestamp)
	if err != nil {
		return fn, err
	}
	return time.Unix(sec, 0), nil
}

//将 当前时间戳 转换成 指定格式的时间字符串
func FormatNow(format string) string {
	return Now().Format(format)
}

//将 时间戳 转换成 指定格式的时间字符串
func FormatTimestamp(format string, timestamp interface{}) string {
	fn, err := ParseTimestamp(timestamp)
	if err != nil {
		if err.Error() != "" {
			fmt.Println("github.com/asktop/gotools/atime ParseTimestamp", "timestamp:", timestamp, "err:", err)
		}
		return ""
	}
	return fn.Format(format)
}

//将 时间戳 转换成 指定格式的时间字符串 格式："2006-01-02 15:04:05"
func FormatDateTime(timestamp interface{}) string {
	return FormatTimestamp(DATETIME, timestamp)
}

//将 时间戳 转换成 指定格式的时间字符串 格式："2006-01-02"
func FormatDate(timestamp interface{}) string {
	return FormatTimestamp(DATE, timestamp)
}

//将 时间戳 转换成 指定格式的时间字符串 格式："15:04:05"
func FormatTime(timestamp interface{}) string {
	return FormatTimestamp(TIME, timestamp)
}

//将 时间戳 转换成 指定格式的时间字符串 格式："2006-01"
func FormatMonth(timestamp interface{}) string {
	return FormatTimestamp(MONTH, timestamp)
}

//将 指定格式的时间字符串 转换成 纳秒级时间戳
func UnFormatUnixNano(format string, timeStr string) int64 {
	t, err := time.ParseInLocation(format, timeStr, time.Local)
	if err != nil {
		fmt.Println("github.com/asktop/gotools/atime UnFormatUnixNano", "format:", format, "timeStr:", timeStr, "err:", err)
		return 0
	}
	return t.UnixNano()
}

//将 指定格式的时间字符串 转换成 秒级时间戳
func UnFormat(format string, timeStr string) int64 {
	t, err := time.ParseInLocation(format, timeStr, time.Local)
	if err != nil {
		fmt.Println("github.com/asktop/gotools/atime UnFormat", "format:", format, "timeStr:", timeStr, "err:", err)
		return 0
	}
	return t.Unix()
}

//将 指定格式的时间字符串 转换成 秒级时间戳
func UnFormatDateTime(timeStr string) int64 {
	return UnFormat(DATETIME, timeStr)
}

//将 指定格式的时间字符串 转换成 秒级时间戳
func UnFormatDate(timeStr string) int64 {
	return UnFormat(DATE, timeStr)
}

//将 指定格式的时间字符串 转换成 秒级时间戳
func UnFormatTime(timeStr string) int64 {
	return UnFormat(TIME, timeStr)
}

//将 指定格式的时间字符串 转换成 秒级时间戳
func UnFormatMonth(timeStr string) int64 {
	return UnFormat(MONTH, timeStr)
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
