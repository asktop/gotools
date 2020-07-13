package adb

import (
    "github.com/asktop/dbr"
    "github.com/asktop/gotools/alog"
    _ "github.com/go-sql-driver/mysql"
    "time"
)

var (
    conn     *dbr.Connection
    connRead *dbr.Connection
)

func StartMysqlDbr(config Config, readConfig ...Config) error {
    defer SetDbrShowLog(config.SqlLogLevel, alog.Info)

    //写数据库连接
    dbConfig := config.GetConfig()
    alog.Info("--- 连接 mysql 主库（写库） ---", "config:", dbConfig)
    var err error
    conn, err = dbr.Open("mysql", dbConfig, nil)
    if err == nil {
        err = checkConn(conn)
    }
    if err != nil {
        alog.Error("--- 连接 mysql 主库（写库）出错 ---", "err:", err)
        return err
    }
    maxIdleConns := config.MaxIdleConns
    maxOpenConns := config.MaxOpenConns
    connMaxLifetime := time.Duration(config.ConnMaxLifetime) * time.Second //时间单位：秒
    conn.SetMaxIdleConns(maxIdleConns)
    conn.SetMaxOpenConns(maxOpenConns)
    conn.SetConnMaxLifetime(connMaxLifetime)

    //读数据库连接（若没有单独配置读库，则与写库相同）
    rConfig := config
    if len(readConfig) > 0 {
        rConfig = readConfig[0]
    }
    dbConfig = rConfig.GetConfig()
    alog.Info("--- 连接 mysql 从库（读库） ---", "config:", dbConfig)
    connRead, err = dbr.Open("mysql", dbConfig, nil)
    if err == nil {
        err = checkConn(connRead)
    }
    if err != nil {
        alog.Error("--- 连接 mysql 从库（读库）出错 ---", "err:", err)
        connRead = conn
        return nil
    }
    maxIdleConns = rConfig.MaxIdleConns
    maxOpenConns = rConfig.MaxOpenConns
    connMaxLifetime = time.Duration(rConfig.ConnMaxLifetime) * time.Second //时间单位：秒
    connRead.SetMaxIdleConns(maxIdleConns)
    connRead.SetMaxOpenConns(maxOpenConns)
    connRead.SetConnMaxLifetime(connMaxLifetime)
    return nil
}

func SetDbrShowLog(level int, logPrint func(args ...interface{})) {
    dbr.ShowSQL(level, logPrint)
}

//检查连接
func checkConn(conn *dbr.Connection) error {
    sess := conn.NewSession(nil)
    _, err := sess.SelectBySql("SELECT 1 FROM DUAL").Rows()
    return err
}

func Session() *dbr.Session {
    return conn.NewSession(nil)
}

func SessionRead() *dbr.Session {
    return connRead.NewSession(nil)
}

//判断是否是 *dbr.Tx
func IsTx(sess dbr.SessionRunner) (*dbr.Tx, bool) {
    tx, ok := sess.(*dbr.Tx)
    return tx, ok
}
