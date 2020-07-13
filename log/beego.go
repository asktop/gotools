package log

import (
    "github.com/astaxie/beego"
    "github.com/astaxie/beego/logs"
    "strings"
)

func startBeegoLog(config Config) {
    //确定输出位置
    outputs := config.Outputs
    if len(outputs) == 0 {
        outputs = append(outputs, "console")
    }
    outputsStr := strings.Join(outputs, ",")
    if strings.Contains(outputsStr, "console") {
        logs.SetLogger(logs.AdapterConsole, ``)
    }
    if strings.Contains(outputsStr, "file") {
        logfile := config.Filepath
        logs.SetLogger(logs.AdapterFile, `{"filename":"`+logfile+`"}`)
        Info("--- 加载 log 日志 ---", "beegoLogFile:", logfile)
    }

    //确定输出级别
    var level int
    switch config.Level {
    case "info":
        level = logs.LevelInfo
    case "warn":
        level = logs.LevelWarn
    case "error":
        level = logs.LevelError
    default:
        level = logs.LevelDebug
    }
    beego.SetLevel(level)
}
