package alimit

import (
	"sync"
	"time"
)

type APILimit struct {
	open           bool //是否开启频率验证
	limitPerMinute int  //每分钟访问频次

	mu           sync.RWMutex
	apiLimitList map[string][]int64
}

//API接口访问频次限制
// @param limitPerMinute 默认接口访问频次限制，小于等于0时关闭限制
func NewAPILimit(limitPerMinute int) *APILimit {
	open := true
	if limitPerMinute <= 0 {
		open = false
	}
	object := &APILimit{open: open, limitPerMinute: limitPerMinute, mu: sync.RWMutex{}, apiLimitList: make(map[string][]int64)}
	object.cleanTask()
	return object
}

//判断接口访问频次是否超频
// @param apiUniqueKey 当前接口访问唯一标识
// @param apiLimitPerMinute 当前接口访问频次限制，小于等于0时关闭限制
func (o *APILimit) Check(apiUniqueKey string, apiLimitPerMinute ...int) (checked bool) {
	//默认验证通过
	checked = true

	//全局关闭验证，验证通过
	if !o.open {
		return checked
	}

	limitPerMinute := o.limitPerMinute //接口验证频率
	if len(apiLimitPerMinute) > 0 {
		if apiLimitPerMinute[0] <= 0 {
			//接口关闭验证，验证通过
			return checked
		} else {
			limitPerMinute = apiLimitPerMinute[0]
		}
	}

	//获取当前系统时间戳（毫秒值）
	currentTime := time.Now().UnixNano() / 1e6
	//获取1分钟前的毫秒值
	checkTime := currentTime - 1000*60
	o.mu.RLock()
	if timelist, find := o.apiLimitList[apiUniqueKey]; !find {
		o.mu.RUnlock()
		o.mu.Lock()
		if timelist, find = o.apiLimitList[apiUniqueKey]; !find {
			o.apiLimitList[apiUniqueKey] = []int64{currentTime}
			checked = true
		}
		o.mu.Unlock()
	} else {
		o.mu.RUnlock()
		o.mu.Lock()
		if timelist, find = o.apiLimitList[apiUniqueKey]; find {
			//判断顶部时间单元是否超时，不超时则保留
			index := len(timelist)
			for k, v := range timelist {
				if v > checkTime {
					index = k
					goto INDEX_END
				}
			}
		INDEX_END:
			timelist = timelist[index:]
			//判断时间单元是否超频，超频则返回false；未超频添加当前时间，返回true
			if len(timelist) >= limitPerMinute {
				checked = false
			} else {
				timelist = append(timelist, currentTime)
				checked = true
			}
			o.apiLimitList[apiUniqueKey] = timelist
		}
		o.mu.Unlock()
	}
	return checked
}

//定时清除内存中超时的接口访问频次
func (o *APILimit) cleanTask() {
	if o.open {
		go func() {
			ticker := time.NewTicker(time.Minute * 10)
			for {
				select {
				case <-ticker.C:
					//获取当前系统时间戳（毫秒值）
					currentTime := time.Now().UnixNano() / 1e6
					//获取1分钟前的毫秒值
					checkTime := currentTime - 1000*60
					o.mu.Lock()
					defer o.mu.Unlock()
					for k, v := range o.apiLimitList {
						//判断时间单元个数为0，则删除
						if len(v) == 0 {
							delete(o.apiLimitList, k)
						} else {
							//判断最后插入的单元时间验证时间，则删除
							tempTime := v[len(v)-1]
							if tempTime < checkTime {
								delete(o.apiLimitList, k)
							}
						}
					}
				}
			}
		}()
	}
}
