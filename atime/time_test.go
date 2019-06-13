package atime

import (
	"testing"
	"time"
)

//时间偏移
func TestOffset(t *testing.T) {
	now := Now()
	t.Log(now)
	t.Log(now.UnixNano())

	//时间偏移设置
	Offset(time.Second * 60)
	now = Now()
	t.Log(now)
	t.Log(now.UnixNano())

	//时间偏移恢复
	Offset(0)
	now = Now()
	t.Log(now)
	t.Log(now.UnixNano())
}

//定时器
var ticker *time.Ticker

//测试定时任务
func TestTicker(t *testing.T) {
	for i := 0; i < 3; i++ {
		//关闭定时任务
		if ticker != nil {
			ticker.Stop()
		}
		if i == 0 {
			ticker = time.NewTicker(time.Second * 1)
		}
		if i == 1 {
			ticker = time.NewTicker(time.Second * 3)
		}
		if i == 2 {
			ticker = time.NewTicker(time.Second * 5)
		}
		//开启定时任务
		go func(a int) {
			t.Log("start")
			for {
				select {
				case <-ticker.C:
					t.Log(a)
				}
			}
			t.Log("stop")
		}(i)
		time.Sleep(time.Second * 10)
	}
	t.Log("for end")
	time.Sleep(time.Second * 10)
	//关闭定时任务
	if ticker != nil {
		ticker.Stop()
	}
	t.Log("stop all 1")
	time.Sleep(time.Second * 10)
	t.Log("stop all 2")
}

//获取当前时间
func TestGetNow(t *testing.T) {
	//获取本地时区 +0800 CST 当前时间（同一时间）
	t.Log(time.Now())
	//获取时间的年、月、日、时、分、秒、纳秒
	t.Log(time.Now().Year())
	t.Log(time.Now().Minute())
	//获取UTC时区 +0000 UTC 当前时间（同一时间）
	t.Log(time.Now().UTC())
}

//获取指定的时间
func TestGetTime(t *testing.T) {
	//指定本地时区 +0800 CST 时间（非同一时间）
	t.Log(time.Date(2018, 10, 1, 18, 32, 40, 1000*1000*23, time.Local))
	//指定UTC时区 +0000 UTC 时间（非同一时间）
	t.Log(time.Date(2018, 10, 1, 18, 32, 40, 1000*1000*23, time.UTC))
}

//时间戳转时间
func TestInt642Time(t *testing.T) {
	//转本地时区 +0800 CST 时间（是同一时间）
	t.Log(time.Unix(1535510340, 0))
	t.Log(time.Unix(1535510340, 0).Unix())
	//nsec为纳秒值
	t.Log(time.Unix(1535510340, 1000))
	//转UTC时区 +0000 UTC 时间（是同一时间）
	t.Log(time.Unix(1535510340, 0).UTC())
	t.Log(time.Unix(1535510340, 0).UTC().Unix())
}

//字符串转时间
func TestString2Time(t *testing.T) {
	t.Log("-----初始-----")

	now := time.Now()
	t.Log(now)
	t.Log(now.Unix())
	nowstr := now.Format("2006-01-02 15:04:05")
	t.Log(nowstr)

	t.Log("-----错误-----", "time.Parse", "没有转换为本地时区，转成了UTC时区")

	now2, _ := time.Parse("2006-01-02 15:04:05", nowstr)
	t.Log(now2)
	t.Log(now2.Local())
	t.Log(now2.Unix())
	now2str := now2.Format("2006-01-02 15:04:05")
	t.Log(now2str)

	t.Log("-----正常-----")

	now21, _ := time.ParseInLocation("2006-01-02 15:04:05", nowstr, time.Local)
	t.Log(now21)
	t.Log(now21.Local())
	t.Log(now21.Unix())
	now21str := now21.Format("2006-01-02 15:04:05")
	t.Log(now21str)
}

//时间转时间戳
func TestTime2Int64(t *testing.T) {
	//本地时区时间（本地时区与UTC时区只是显示时间不同，其时间戳是相同的）
	t.Log(time.Now())                  //2018-08-29 10:39:00.7937526 +0800 CST m=+0.062003601
	t.Log(time.Now().Unix())           //时间戳（秒，10位）	1535510340
	t.Log(time.Now().UnixNano() / 1e6) //时间戳（毫秒，13位）	1535510340793
	t.Log(time.Now().UnixNano())       //时间戳（纳秒，19位）	1535510340793752600
	//UTC时区时间
	t.Log(time.Now().UTC())        //2018-08-29 02:39:00.7937526 +0000 UTC
	t.Log(time.Now().UTC().Unix()) //时间戳（秒，10位）   1535510340
}

//时间转字符串
func TestTime2String(t *testing.T) {
	t.Log(time.Now().Format("2006-01-02 15:04:05.999999999"))
	t.Log(time.Now().Format("2006-01-02 15:04:05.999999"))
	t.Log(time.Now().Format("2006-01-02 15:04:05.999"))
	t.Log(time.Now().Format("2006-01-02 15:04:05"))
	t.Log(time.Now().Format("2006-01-02 15:00:00")) //整点
	t.Log(time.Now().Format("2006-01-02"))
	t.Log(time.Now().Format("15:04:05"))
}

//时间工具
func TestTimeTool(t *testing.T) {
	t.Log("--- 时间比较 ---")
	time1 := time.Date(2018, 10, 1, 18, 32, 40, 1000*1000*23, time.Local)
	time2 := time.Date(2018, 10, 1, 18, 33, 40, 1000*1000*23, time.Local)
	t.Log(time1.Before(time2))
	t.Log(time1.Equal(time2))
	t.Log(time1.After(time2))

	t.Log("--- 时间加减、时间间距 ---")
	t.Log(time1.Add(time.Minute * 2))
	t.Log(time1.Add(time.Hour * -2))
	t.Log(time2.Sub(time1))

	t.Log("--- 距今时间、时间间距 ---")
	t.Log(time.Since(time1))
	t.Log(time.Since(time2))
	t.Log(time.Since(time1) - time.Since(time2))

	t.Log("--- 时间段 ---")
	times := time.Minute * 3
	t.Log(times)
	t.Log(int64(times))
	t.Log(int64(times.Seconds()))

	t.Log("--- 时间暂停 ---")
	time.Sleep(time.Second * 3)
	t.Log("--- 暂停了3秒 ---")
}
