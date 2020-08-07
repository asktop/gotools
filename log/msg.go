package log

import (
    "github.com/asktop/gotools/astring"
    "github.com/astaxie/beego/logs"
)

type Msg struct {
    msgs []string
}

func NewMsg() *Msg {
    return new(Msg)
}

func (m *Msg) Info(title string, args ...interface{}) {
    msg := astring.Join(args...)
    m.msgs = append(m.msgs, msg)

    logargs := []interface{}{}
    logargs = append(logargs, title, msg)
    if logDriver == "beego" {
        logs.Info("", logargs...)
    }
    if logDriver == "zap" {
        Zap.Info(AppendBlank(logargs)...)
    }
}

func (m *Msg) Warn(title string, args ...interface{}) {
    msg := astring.Join(args...)
    m.msgs = append(m.msgs, msg)

    logargs := []interface{}{}
    logargs = append(logargs, title, msg)
    if logDriver == "beego" {
        logs.Warn("", logargs...)
    }
    if logDriver == "zap" {
        Zap.Warn(AppendBlank(logargs)...)
    }
}

func (m *Msg) Error(title string, args ...interface{}) {
    msg := astring.Join(args...)
    m.msgs = append(m.msgs, msg)

    logargs := []interface{}{}
    logargs = append(logargs, title, msg)
    if logDriver == "beego" {
        logs.Error("", logargs...)
    }
    if logDriver == "zap" {
        Zap.Error(AppendBlank(logargs)...)
    }
}

func (m *Msg) Append(msgs []string) {
    m.msgs = append(m.msgs, msgs...)
}

func (m *Msg) Msgs() []string {
    return m.msgs
}
