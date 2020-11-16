package astress

import (
    "encoding/json"
    "fmt"
    "log"
    "strconv"
    "sync"
    "sync/atomic"
    "time"
)

//压测配置
type Config struct {
    StartNumber int64 `json:"start_number"` //初始并发数（必需）
    EndSecond   int64 `json:"end_second"`   //总压测秒数（必需，到达时终止压测）

    StepSecond int64 `json:"step_second"` //每多少秒增长一次并发数
    StepNumber int64 `json:"step_number"` //每次增长多少并发数
    MaxNumber  int64 `json:"max_number"`  //最大并发数
}

//快捷压测（每秒执行一次）
func QuickStress(config Config, stressFunc func() error) {
    Stress(config, stressFunc, time.Second)
}

//压测
//config 压测配置
//stressFunc 压测方法
//duration 压测方法执行频率
func Stress(config Config, stressFunc func() error, duration time.Duration) {
    if config.StartNumber <= 0 || config.EndSecond <= 0 {
        log.Println("请设置 初始并发数(StartNumber) 和 总压测秒数(EndSecond)")
        return
    }
    if config.StepSecond > 0 || config.StepNumber > 0 || config.MaxNumber > 0 {
        if config.StepSecond <= 0 || config.StepNumber <= 0 || config.MaxNumber <= 0 {
            log.Println("请设置 每多少秒增长一次并发数(StepSecond) 和 每次增长多少并发数(StepNumber) 和 最大并发数(MaxNumber)")
            return
        }
        if config.MaxNumber < config.StartNumber {
            log.Println("最大并发数(MaxNumber) 不能小于 初始并发数(StartNumber)")
            return
        }
    }

    data, _ := json.Marshal(&config)
    configStr := string(data)
    log.Println("=== 开始压测 ===", configStr)

    var startTime int64                  //全局开始时间
    var currentNumber int64              //当前并发数
    var stepUseSecond int64              //当前梯度用时
    var totalNumber, totalOkNumber int64 //总压测次数，总成功次数
    startTime = time.Now().UnixNano() / 1e6
    currentNumber = config.StartNumber

    dSecond := duration / time.Second
    ticker := time.NewTicker(duration)
    for {
        select {
        case <-ticker.C:

            perStartTime := time.Now().UnixNano() / 1e6 //单次压测开始时间
            var perOkNumber int64                       //单次压测成功数

            wg := sync.WaitGroup{}
            for i := 0; i < int(currentNumber); i++ {
                wg.Add(1)
                go func(wg *sync.WaitGroup, perOkNumber *int64) {
                    err := stressFunc()
                    if err == nil {
                        atomic.AddInt64(perOkNumber, 1)
                    }
                    wg.Done()
                }(&wg, &perOkNumber)
            }
            wg.Wait()

            perEndTime := time.Now().UnixNano() / 1e6 //单次压测结束时间

            perUseTime := perEndTime - perStartTime //单次压测用时
            useTime := perEndTime - startTime       //全局压测用时
            totalNumber += currentNumber
            totalOkNumber += perOkNumber

            log.Println("--- 本次压测 ---",
                fmt.Sprintf("用时：%s", getStrFromMilliSecond(perUseTime)),
                fmt.Sprintf("成功率：%s", percentStr(perOkNumber, currentNumber)),
                fmt.Sprintf("并发数：%d", currentNumber),
                fmt.Sprintf("成功数：%d", perOkNumber),
            )

            if useTime/1000 >= config.EndSecond {
                log.Println("=== 结束压测 ===",
                    fmt.Sprintf("用时：%s", getStrFromMilliSecond(useTime)),
                    fmt.Sprintf("成功率：%s", percentStr(totalOkNumber, totalNumber)),
                    fmt.Sprintf("压测数：%d", totalNumber),
                    fmt.Sprintf("成功数：%d", totalOkNumber),
                    configStr)
                time.Sleep(time.Second * 3)
                return
            }

            perSecond := perUseTime/1000 + int64(dSecond)
            if perSecond == 0 {
                perSecond = 1
            }
            stepUseSecond += perSecond
            if stepUseSecond >= config.StepSecond {
                stepUseSecond = 0
                currentNumber += config.StepNumber

                if currentNumber > config.MaxNumber {
                    currentNumber = config.MaxNumber
                }

            }
        }
    }
}

func getStrFromMilliSecond(milliSecond int64) string {
    m := milliSecond / 1000 / 60
    s := milliSecond / 1000 % 60
    ms := milliSecond % 1000
    if m > 0 {
        return fmt.Sprintf("%dm%ds%dms", m, s, ms)
    } else {
        if s > 0 {
            return fmt.Sprintf("%ds%dms", s, ms)
        } else {
            return fmt.Sprintf("%dms", ms)
        }
    }
}

func percent(v1 int64, v2 int64) float64 {
    v := float64(v1) / float64(v2) * 100
    v, _ = strconv.ParseFloat(fmt.Sprintf("%.2f", v), 64)
    return v
}

func percentStr(v1 int64, v2 int64) string {
    v := float64(v1) / float64(v2) * 100
    return fmt.Sprintf("%.2f%s", v, "%")
}
