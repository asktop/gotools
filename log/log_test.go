package log

import (
    "errors"
    "github.com/astaxie/beego"
    "github.com/astaxie/beego/logs"
    "testing"
)

func TInit() {
    //StartLog(Config{Driver: "beego"})
    //StartLog(Config{Driver: "beego", Outputs: []string{"console", "file"}, Filepath: "D://app.log", Level: "debug"})
    //StartLog(Config{Driver: "zap"})
    StartLog(Config{Driver: "zap", Outputs: []string{"console", "file"}, Filepath: "D://app.log", Level: "debug", Format: "text"})
}

func TestLog(t *testing.T) {
    TInit()
    err := errors.New("cuowu")
    //返回位置为当前文件
    Debugf("zaplog, name:%s, age:%d, err:%v", "abc", 123, err)
    Debug("debug", "abc", 123, err)
    Info("info", "abc", 123, err)
    Warn("warn", "abc", 123, err)
    Error("err", "abc", 123, err)
}

func TestBeegoLog(t *testing.T) {
    TInit()
    err := errors.New("cuowu")
    //返回位置为当前文件
    beego.Info("beegolog", "abc", 123, err)
    beego.Info("beegolog, name:%s, age:%d, err:%v", "abc", 123, err)
    //返回位置为上层调取文件
    logs.Info("beegolog", "abc", 123, err)
    logs.Info("beegolog, name:%s, age:%d, err:%v", "abc", 123, err)
}

func TestZapLog(t *testing.T) {
    TInit()
    err := errors.New("cuowu")
    //返回位置为当前文件
    Zap.Debugf("zaplog, name:%s, age:%d, err:%v", "abc", 123, err)
    Zap.Debug("debug", "abc", 123, err)
    Zap.Info("info", "abc", 123, err)
    Zap.Warn("warn", "abc", 123, err)
    Zap.Error("err", "abc", 123, err)
}
