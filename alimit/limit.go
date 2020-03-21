package alimit

import (
	"sync"
	"time"
)

type APIRateLimit struct {
	open           bool //是否开启频率验证
	limitPerMinute int  //每分钟访问频次
	mu             sync.RWMutex
	timeListMap    map[string][]int64
}

//API接口访问频次限制
func NewAPIRateLimit(limitPerMinute int) *APIRateLimit {
	open := true
	if limitPerMinute <= 0 {
		open = false
	}
	object := &APIRateLimit{open: open, limitPerMinute: limitPerMinute, mu: sync.RWMutex{}, timeListMap: make(map[string][]int64)}
	object.cleanTask()
	return object
}

//判断接口访问频次是否超频
func (o *APIRateLimit) Check(apiUniqueKey string) bool {
	if !o.open {
		return true
	}
	check := true
	//获取当前系统时间戳（毫秒值）
	currentTime := time.Now().UnixNano() / 1e6
	//获取1分钟前的毫秒值
	checkTime := currentTime - 1000*60
	o.mu.RLock()
	if timelist, find := o.timeListMap[apiUniqueKey]; !find {
		o.mu.RUnlock()
		o.mu.Lock()
		if timelist, find = o.timeListMap[apiUniqueKey]; !find {
			o.timeListMap[apiUniqueKey] = []int64{currentTime}
			check = true
		}
		o.mu.Unlock()
	} else {
		o.mu.RUnlock()
		o.mu.Lock()
		if timelist, find = o.timeListMap[apiUniqueKey]; find {
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
			if len(timelist) >= o.limitPerMinute {
				check = false
			} else {
				timelist = append(timelist, currentTime)
				check = true
			}
			o.timeListMap[apiUniqueKey] = timelist
		}
		o.mu.Unlock()
	}
	return check
}

//定时清除内存中超时的接口访问频次
func (o *APIRateLimit) cleanTask() {
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
					for k, v := range o.timeListMap {
						//判断时间单元个数为0，则删除
						if len(v) == 0 {
							delete(o.timeListMap, k)
						} else {
							//判断最后插入的单元时间验证时间，则删除
							tempTime := v[len(v)-1]
							if tempTime < checkTime {
								delete(o.timeListMap, k)
							}
						}
					}
				}
			}
		}()
	}
}
