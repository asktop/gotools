package atime

import "time"

//开始时间工具
type startTime struct {
	fn time.Time
}

//获取 指定时间 指定精度的 开始时间
func NewStartTime(timestamp ...interface{}) *startTime {
	st := new(startTime)
	fn := Now()
	if len(timestamp) > 0 {
		fn, _ = ParseTimestamp(timestamp[0])
	}
	st.fn = fn
	return st
}

//获取当天的开始时间
func (st *startTime) Day() time.Time {
	return time.Date(st.fn.Year(), st.fn.Month(), st.fn.Day(), 0, 0, 0, 0, time.Local)
}

//获取当前小时的开始时间
func (st *startTime) Hour() time.Time {
	return time.Date(st.fn.Year(), st.fn.Month(), st.fn.Day(), st.fn.Hour(), 0, 0, 0, time.Local)
}
