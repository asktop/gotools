package alimit

import (
    "container/list"
    "sync"
    "time"
)

const (
    Level_Close Level = 0 //关闭校验
    Level_Weak  Level = 1 //弱校验（当限速后，新请求不再加入请求队列）
    Level_High  Level = 2 //强校验（当限速后，新请求依旧加入请求队列）
)

var defaultLimit = NewLimit(1)

type Level int

type Limit struct {
    mu     sync.RWMutex
    level  Level //校验级别
    apiMap map[string]*apiWrap
}

type apiWrap struct {
    mu       sync.RWMutex
    limit    int        //限速频率
    seconds  int64      //限速单位（多少秒）
    timeList *list.List //时间戳队列
}

//API接口访问频次限制
func NewLimit(level Level) *Limit {
    if level > 2 || level < 0 {
        level = 0
    }
    limit := &Limit{mu: sync.RWMutex{}, level: level, apiMap: make(map[string]*apiWrap)}
    limit.cleanTask()
    return limit
}

//设置校验级别
func (o *Limit) SetLevel(level Level) {
    if level > 2 || level < 0 {
        level = 0
    }
    o.mu.Lock()
    o.level = level
    o.mu.Unlock()
}

//是否开启了校验
func (o *Limit) IsOpen() bool {
    o.mu.RLock()
    level := o.level
    o.mu.RUnlock()
    if level == 1 || level == 2 {
        return true
    } else {
        return false
    }
}

//判断接口访问频次是否超频
// @param apiUniqueKey 	当前接口访问唯一标识
// @param limit    		限速频率（小于等于0时关闭该接口验证）
// @param seconds 		限速单位（多少秒）
func (o *Limit) Check(apiUniqueKey string, limit int, seconds int64) (checked bool, times int) {
    //默认验证通过
    checked = true

    //关闭验证，验证通过
    o.mu.RLock()
    level := o.level
    o.mu.RUnlock()
    if level == 0 || limit <= 0 || seconds <= 0 || apiUniqueKey == "" {
        return checked, times
    }

    //获取当前系统时间戳（毫秒值）
    now := time.Now().UnixNano() / 1e6
    //获取校验起始时间戳（毫秒值）
    start := now - 1000*seconds

    o.mu.Lock()
    if api, ok := o.apiMap[apiUniqueKey]; !ok {
        apiTemp := &apiWrap{mu: sync.RWMutex{}, limit: limit, seconds: seconds, timeList: list.New()}
        apiTemp.timeList.PushBack(now)
        o.apiMap[apiUniqueKey] = apiTemp
        times = 1
        o.mu.Unlock()
    } else {
        o.mu.Unlock()
        api.mu.Lock()
        for e := api.timeList.Front(); e != nil; e = e.Next() {
            etime := e.Value.(int64)
            if etime <= start {
                api.timeList.Remove(e)
            } else {
                break
            }
        }
        times = api.timeList.Len() + 1
        if times > limit {
            o.mu.RLock()
            if o.level == 2 {
                api.timeList.PushBack(now)
            }
            o.mu.RUnlock()
            checked = false
        } else {
            api.timeList.PushBack(now)
        }
        api.mu.Unlock()
    }
    return checked, times
}

//定时清除内存中超时的接口访问频次
func (o *Limit) cleanTask() {
    go func() {
        ticker := time.NewTicker(time.Hour)
        for {
            select {
            case <-ticker.C:
                o.mu.Lock()
                if o.level != 0 {
                    for k, api := range o.apiMap {
                        api.mu.Lock()
                        //判断时间单元个数为0，则删除
                        if api.timeList.Len() == 0 {
                            delete(o.apiMap, k)
                        } else {
                            e := api.timeList.Back()
                            etime := e.Value.(int64)
                            //获取校验起始时间戳（毫秒值）
                            start := time.Now().UnixNano()/1e6 - 1000*api.seconds
                            //判断最后插入的单元时间验证时间，则删除
                            if etime <= start {
                                delete(o.apiMap, k)
                            }
                        }
                        api.mu.Unlock()
                    }
                }
                o.mu.Unlock()
            }
        }
    }()
}

//设置默认
func SetDefaultLimit(level Level) {
    if level > 2 || level < 0 {
        level = 0
    }
    defaultLimit.level = level
}

func CheckDefaultLimit(apiUniqueKey string, limit int, seconds int64) (checked bool, times int) {
    return defaultLimit.Check(apiUniqueKey, limit, seconds)
}
