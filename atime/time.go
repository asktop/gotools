package atime

import (
	"fmt"
	"github.com/spf13/cast"
	"strings"
	"time"
)

const (
	DATETIME = "2006-01-02 15:04:05"
	DATE     = "2006-01-02"
	TIME     = "15:04:05"
)

var (
	localFlag  string        //本地时区 标识 CST：北京时区
	offsetTime time.Duration //时间偏移量
)

func init() {
	localFlag = strings.Split(time.Now().Local().String(), " ")[3]
}

//设置时间偏量（改变当前时间）
func Offset(offset time.Duration) {
	offsetTime = offset
}

//当前时间
func Now() time.Time {
	now := time.Now()
	if offsetTime != 0 {
		now = now.Add(offsetTime)
	}
	return now
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

//格式化时间戳 格式指定
func FormatTimestamp(format string, timestamp interface{}) string {
	fn, err := ParseTimestamp(timestamp)
	if err != nil {
		fmt.Println(err)
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

//将 本地时区的字符串时间 转换成 本地时区时间（默认会转换成UTC时区的时间）
func Parse(format string, timeStr string) (time.Time, error) {
	return time.Parse(fmt.Sprintf("%s MST", format), fmt.Sprintf("%s %s", timeStr, localFlag))
}

//将 时间戳 转换成 本地时区时间（默认会转换成UTC时区的时间）
func ParseTimestamp(timestamp interface{}) (time.Time, error) {
	var err error
	var sec, nsec int64
	fn := Now()
	tsStr, err := cast.ToStringE(timestamp)
	if err != nil {
		return fn, err
	}
	tsLen := len(tsStr)
	if tsLen <= 10 {
		sec, err = cast.ToInt64E(tsStr)
		if err != nil {
			return fn, err
		}
	} else if tsLen <= 19 {
		if tsLen < 19 {
			tsStr += strings.Repeat("0", 19-tsLen)
		}
		secStr := tsStr[0:10]
		nsecStr := tsStr[11:]
		sec, err = cast.ToInt64E(secStr)
		if err != nil {
			return fn, err
		}
		nsec, err = cast.ToInt64E(nsecStr)
		if err != nil {
			return fn, err
		}
	} else {
		return fn, err
	}
	return time.Unix(sec, nsec), nil
}
