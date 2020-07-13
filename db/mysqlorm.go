package db

import (
    "github.com/asktop/gotools/log"
    "github.com/astaxie/beego/orm"
    _ "github.com/go-sql-driver/mysql"
)

func StartMysqlOrm(config Config, readConfig ...Config) error {
    dbConfig := config.GetConfig()
    maxIdleConns := config.MaxIdleConns
    maxOpenConns := config.MaxOpenConns
    log.Info("--- 连接 mysql 主库（写库） ---", "config:", dbConfig)
    err := orm.RegisterDataBase("default", "mysql", dbConfig, maxIdleConns, maxOpenConns)
    if err != nil {
        log.Error("--- 连接 mysql 主库（写库）出错 ---", "err:", err)
        return err
    }
    return nil
}

func Orm() orm.Ormer {
    return orm.NewOrm()
}

//sql创建工具
func OrmSql() orm.QueryBuilder {
    qb, _ := orm.NewQueryBuilder("mysql")
    return qb
}
