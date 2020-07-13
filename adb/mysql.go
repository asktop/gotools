package db

import (
    "database/sql"
    "github.com/asktop/gotools/alog"
    _ "github.com/go-sql-driver/mysql"
    "time"
)

var (
    db     *sql.DB //写库连接池
    dbRead *sql.DB //读库连接池
)

//初始化连接mysql
func StartMysql(config Config, readConfig ...Config) error {
    //写数据库连接
    dbConfig := config.GetConfig()
    alog.Info("--- 连接 mysql 主库（写库） ---", "config:", dbConfig)
    var err error
    db, err = sql.Open("mysql", dbConfig)
    if err != nil {
        alog.Error("--- 连接 mysql 主库（写库）出错 ---", "err:", err)
        return err
    }
    maxIdleConns := config.MaxIdleConns
    maxOpenConns := config.MaxOpenConns
    connMaxLifetime := time.Duration(config.ConnMaxLifetime) * time.Second //时间单位：秒
    db.SetMaxIdleConns(maxIdleConns)
    db.SetMaxOpenConns(maxOpenConns)
    db.SetConnMaxLifetime(connMaxLifetime)

    //读数据库连接（若没有单独配置读库，则与写库相同）
    rConfig := config
    if len(readConfig) > 0 {
        rConfig = readConfig[0]
    }
    dbConfig = rConfig.GetConfig()
    alog.Info("--- 连接 mysql 从库（读库） ---", "config:", dbConfig)
    dbRead, err = sql.Open("mysql", dbConfig)
    if err != nil {
        alog.Error("--- 连接 mysql 从库（读库）出错 ---", "err:", err)
        dbRead = db
        return nil
    }
    maxIdleConns = rConfig.MaxIdleConns
    maxOpenConns = rConfig.MaxOpenConns
    connMaxLifetime = time.Duration(rConfig.ConnMaxLifetime) * time.Second //时间单位：秒
    dbRead.SetMaxIdleConns(maxIdleConns)
    dbRead.SetMaxOpenConns(maxOpenConns)
    dbRead.SetConnMaxLifetime(connMaxLifetime)
    return nil
}

func Db() *sql.DB {
    return db
}

func DbRead() *sql.DB {
    return dbRead
}
