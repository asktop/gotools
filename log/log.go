package log

import (
    "github.com/astaxie/beego/logs"
    "os"
)

var logDriver string

//日志配置
type Config struct {
    Driver   string   `json:"driver" yaml:"driver"`     //驱动：beego,zap，默认beego
    Outputs  []string `json:"outputs" yaml:"outputs"`   //输出位置：console;file，默认console
    Filepath string   `json:"filepath" yaml:"filepath"` //日志文件绝对路径和文件名，例如 D:/log/app.log
    Level    string   `json:"level" yaml:"level"`       //输出级别：debug,info,...，默认debug
    Format   string   `json:"format" yaml:"format"`     //输出格式：text,json，默认text
}

//启动log日志
func StartLog(config Config) {
    if config.Driver == "beego" || config.Driver == "zap" {
        logDriver = config.Driver
    } else {
        logDriver = "beego"
    }

    if logDriver == "beego" {
        startBeegoLog(config)
    }
    if logDriver == "zap" {
        startZapLog(config)
    }
}

func GetDriver() string {
    return logDriver
}

func Debug(args ...interface{}) {
    if logDriver == "" {
        StartLog(Config{Driver: "beego"})
    }
    if logDriver == "beego" {
        logs.Debug("", args...)
    }
    if logDriver == "zap" {
        Zap.Debug(appendBlank(args)...)
    }
}

func Debugf(format string, args ...interface{}) {
    if logDriver == "" {
        StartLog(Config{Driver: "beego"})
    }
    if logDriver == "beego" {
        logs.Debug(format, args...)
    }
    if logDriver == "zap" {
        Zap.Debugf(format, args...)
    }
}

func Info(args ...interface{}) {
    if logDriver == "" {
        StartLog(Config{Driver: "beego"})
    }
    if logDriver == "beego" {
        logs.Info("", args...)
    }
    if logDriver == "zap" {
        Zap.Info(appendBlank(args)...)
    }
}

func Infof(format string, args ...interface{}) {
    if logDriver == "" {
        StartLog(Config{Driver: "beego"})
    }
    if logDriver == "beego" {
        logs.Info(format, args...)
    }
    if logDriver == "zap" {
        Zap.Infof(format, args...)
    }
}

func Warn(args ...interface{}) {
    if logDriver == "" {
        StartLog(Config{Driver: "beego"})
    }
    if logDriver == "beego" {
        logs.Warn("", args...)
    }
    if logDriver == "zap" {
        Zap.Warn(appendBlank(args)...)
    }
}

func Warnf(format string, args ...interface{}) {
    if logDriver == "" {
        StartLog(Config{Driver: "beego"})
    }
    if logDriver == "beego" {
        logs.Warn(format, args...)
    }
    if logDriver == "zap" {
        Zap.Warnf(format, args...)
    }
}

func Error(args ...interface{}) {
    if logDriver == "" {
        StartLog(Config{Driver: "beego"})
    }
    if logDriver == "beego" {
        logs.Error("", args...)
    }
    if logDriver == "zap" {
        Zap.Error(appendBlank(args)...)
    }
}

func Errorf(format string, args ...interface{}) {
    if logDriver == "" {
        StartLog(Config{Driver: "beego"})
    }
    if logDriver == "beego" {
        logs.Error(format, args...)
    }
    if logDriver == "zap" {
        Zap.Errorf(format, args...)
    }
}

func Fatal(args ...interface{}) {
    if logDriver == "" {
        StartLog(Config{Driver: "beego"})
    }
    if logDriver == "beego" {
        logs.Error("", args...)
    }
    if logDriver == "zap" {
        Zap.Fatal(appendBlank(args)...)
    }
    os.Exit(1)
}

func Fatalf(format string, args ...interface{}) {
    if logDriver == "" {
        StartLog(Config{Driver: "beego"})
    }
    if logDriver == "beego" {
        logs.Error(format, args...)
    }
    if logDriver == "zap" {
        Zap.Fatalf(format, args...)
    }
    os.Exit(1)
}

func appendBlank(args []interface{}) []interface{} {
    argsTemp := []interface{}{}
    for i, arg := range args {
        if i < len(args)-1 {
            argsTemp = append(argsTemp, arg, " ")
        } else {
            argsTemp = append(argsTemp, arg)
        }
    }
    return argsTemp
}
