package atime

import (
	"fmt"
	"testing"
	"time"
)

func TestTime(t *testing.T) {
	fmt.Println("-----初始-----")

	now := time.Now()
	fmt.Println(now)
	fmt.Println(now.Unix())
	nowstr := now.Format("2006-01-02 15:04:05")
	fmt.Println(nowstr)

	fmt.Println("-----错误-----", "time.Parse", "没有转换为本地时区，转成了UTC时区")

	now2, _ := time.Parse("2006-01-02 15:04:05", nowstr)
	fmt.Println(now2)
	fmt.Println(now2.Local())
	fmt.Println(now2.Unix())
	now2str := now2.Format("2006-01-02 15:04:05")
	fmt.Println(now2str)

	fmt.Println("-----正常-----")

	now3, _ := Parse("2006-01-02 15:04:05", nowstr)
	fmt.Println(now3)
	fmt.Println(now3.Local())
	fmt.Println(now3.Unix())
	now3str := now3.Format("2006-01-02 15:04:05")
	fmt.Println(now3str)

	fmt.Println("----------")

	now4str := FormatTimestamp("2006-01-02 15:04:05", now.Unix())
	fmt.Println(now4str)
}

func TestOffset(t *testing.T) {
	now := Now()
	fmt.Println(now)
	fmt.Println(now.UnixNano())

	Offset(time.Second*60)
	now = Now()
	fmt.Println(now)
	fmt.Println(now.UnixNano())

	Offset(0)
	now = Now()
	fmt.Println(now)
	fmt.Println(now.UnixNano())
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
			fmt.Println("start")
			for {
				select {
				case <-ticker.C:
					fmt.Println(a)
				}
			}
			fmt.Println("stop")
		}(i)
		time.Sleep(time.Second * 10)
	}
	fmt.Println("for end")
	time.Sleep(time.Second * 10)
	//关闭定时任务
	if ticker != nil {
		ticker.Stop()
	}
	fmt.Println("stop all 1")
	time.Sleep(time.Second * 10)
	fmt.Println("stop all 2")
}

//获取当前时间
func TestGetNow(t *testing.T) {
	//获取本地时区 +0800 CST 当前时间（同一时间）
	fmt.Println(time.Now())
	//获取时间的年、月、日、时、分、秒、纳秒
	fmt.Println(time.Now().Year())
	fmt.Println(time.Now().Minute())
	//获取UTC时区 +0000 UTC 当前时间（同一时间）
	fmt.Println(time.Now().UTC())
}

//获取指定的时间
func TestGetTime(t *testing.T) {
	//指定本地时区 +0800 CST 时间（非同一时间）
	fmt.Println(time.Date(2018, 10, 1, 18, 32, 40, 1000*1000*23, time.Local))
	//指定UTC时区 +0000 UTC 时间（非同一时间）
	fmt.Println(time.Date(2018, 10, 1, 18, 32, 40, 1000*1000*23, time.UTC))
}

//时间戳转时间
func TestInt642Time(t *testing.T) {
	//转本地时区 +0800 CST 时间（是同一时间）
	fmt.Println(time.Unix(1535510340, 0))
	fmt.Println(time.Unix(1535510340, 0).Unix())
	//nsec为纳秒值
	fmt.Println(time.Unix(1535510340, 1000))
	//转UTC时区 +0000 UTC 时间（是同一时间）
	fmt.Println(time.Unix(1535510340, 0).UTC())
	fmt.Println(time.Unix(1535510340, 0).UTC().Unix())
}

//字符串转时间
func TestString2Time(t *testing.T) {
	//转本地时区 +0800 CST 时间（非同一时间）
	time1, _ := time.Parse("2006-01-02 15:04:05 MST", "2017-03-04 13:21:45 CST")
	fmt.Println(time1)
	fmt.Println(time1.Unix())
	//转UTC时区 +0000 UTC 时间（非同一时间）
	time2, _ := time.Parse("2006-01-02 15:04:05", "2017-03-04 13:21:45")
	fmt.Println(time2)
	fmt.Println(time2.Unix())
}

//时间转时间戳
func TestTime2Int64(t *testing.T) {
	//本地时区时间（本地时区与UTC时区只是显示时间不同，其时间戳是相同的）
	fmt.Println(time.Now())                  //2018-08-29 10:39:00.7937526 +0800 CST m=+0.062003601
	fmt.Println(time.Now().Unix())           //时间戳（秒，10位）	1535510340
	fmt.Println(time.Now().UnixNano() / 1e6) //时间戳（毫秒，13位）	1535510340793
	fmt.Println(time.Now().UnixNano())       //时间戳（纳秒，19位）	1535510340793752600
	//UTC时区时间
	fmt.Println(time.Now().UTC())        //2018-08-29 02:39:00.7937526 +0000 UTC
	fmt.Println(time.Now().UTC().Unix()) //时间戳（秒，10位）   1535510340
}

//时间转字符串
func TestTime2String(t *testing.T) {
	fmt.Println(time.Now().Format("2006-01-02 15:04:05.999999999"))
	fmt.Println(time.Now().Format("2006-01-02 15:04:05.999999"))
	fmt.Println(time.Now().Format("2006-01-02 15:04:05.999"))
	fmt.Println(time.Now().Format("2006-01-02 15:04:05"))
	fmt.Println(time.Now().Format("2006-01-02 15:00:00"))
	fmt.Println(time.Now().Format("2006-01-02"))
}

//时间工具
func TestTimeTool(t *testing.T) {
	fmt.Println("--- 时间比较 ---")
	time1 := time.Date(2018, 10, 1, 18, 32, 40, 1000*1000*23, time.Local)
	time2 := time.Date(2018, 10, 1, 18, 33, 40, 1000*1000*23, time.Local)
	fmt.Println(time1.Before(time2))
	fmt.Println(time1.Equal(time2))
	fmt.Println(time1.After(time2))

	fmt.Println("--- 时间加减、时间间距 ---")
	fmt.Println(time1.Add(time.Minute * 2))
	fmt.Println(time1.Add(time.Hour * -2))
	fmt.Println(time2.Sub(time1))

	fmt.Println("--- 距今时间、时间间距 ---")
	fmt.Println(time.Since(time1))
	fmt.Println(time.Since(time2))
	fmt.Println(time.Since(time1) - time.Since(time2))

	fmt.Println("--- 时间段 ---")
	times := time.Minute * 3
	fmt.Println(times)
	fmt.Println(int64(times))
	fmt.Println(int64(times.Seconds()))

	fmt.Println("--- 时间暂停 ---")
	time.Sleep(time.Second * 3)
	fmt.Println("--- 暂停了3秒 ---")
}
