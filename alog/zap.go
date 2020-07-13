package alog

import (
    "github.com/lestrrat-go/file-rotatelogs"
    "go.uber.org/zap"
    "go.uber.org/zap/zapcore"
    "os"
    "strings"
    "time"
)

var Zap *zap.SugaredLogger //正常、格式化、键值对日志

func startZapLog(config Config) {
    //日志输出到控制台
    consoleWriter := zapcore.AddSync(os.Stdout)
    //日志输出到文件
    //按大小分割 gopkg.in/natefinch/lumberjack.v2
    //hook := &lumberjack.Logger{
    //	Filename:  logfile, //组装日志文件
    //	MaxSize:   256,     //按大小进行分割归档（单位：MB）
    //	MaxAge:    1,       //按日期进行分割归档（单位：天）
    //	LocalTime: true,    //是否按本地日期进行命名归档文件，默认否
    //}
    //按日期分割 github.com/lestrrat-go/file-rotatelogs
    logfile := config.Filepath
    hook, _ := rotatelogs.New(
        logfile+".%Y-%m-%d", //实际生成的分割文件名 demo.log.YY-mm-dd
        rotatelogs.WithLinkName(logfile),
        rotatelogs.WithRotationTime(time.Hour*24), //每多久分割一次日志，每24小时(整点)分割一次日志
        rotatelogs.WithMaxAge(time.Hour*24*30),    //保存多久的历史日志，保存30天内的日志
    )
    fileWriter := zapcore.AddSync(hook)

    //确定输出位置
    outputs := config.Outputs
    if len(outputs) == 0 {
        outputs = append(outputs, "console")
    }
    outputsStr := strings.Join(outputs, ",")
    syncers := []zapcore.WriteSyncer{}
    if strings.Contains(outputsStr, "console") {
        syncers = append(syncers, consoleWriter)
    }
    if strings.Contains(outputsStr, "file") {
        syncers = append(syncers, fileWriter)
    }

    //确定输出级别
    level := zap.DebugLevel
    switch config.Level {
    case "info":
        level = zap.InfoLevel
    case "warn":
        level = zap.WarnLevel
    case "error":
        level = zap.ErrorLevel
    }

    //确定输出格式
    //日志编码参数
    encoderConfig := zap.NewDevelopmentEncoderConfig()
    encoderConfig.EncodeTime = timeEncoder //自定义行首时间输出格式
    //日志以正常格式输出 示例：2018-12-12 15:41:07.166	INFO	logutil/zap_test.go:13	zaplog	{"abc": 123, "err": "cuowu"}
    encoder := zapcore.NewConsoleEncoder(encoderConfig)
    if config.Format == "json" {
        //日志以json格式输出 示例：{"L":"INFO","T":"2018-12-12 15:41:07.166","C":"logutil/zap_test.go:13","M":"zaplog","abc":123,"err":"cuowu"}
        encoder = zapcore.NewJSONEncoder(encoderConfig)
    }

    core := zapcore.NewCore(encoder, zapcore.NewMultiWriteSyncer(syncers...), level)
    logger := zap.New(core, zap.AddCaller(), zap.AddCallerSkip(1))
    Zap = logger.Sugar()

    Info("--- 加载 log 日志 ---", "zapLogFile:", logfile)
}

//自定义行首时间输出格式
func timeEncoder(t time.Time, enc zapcore.PrimitiveArrayEncoder) {
    enc.AppendString(t.Format("2006-01-02 15:04:05.000"))
}
