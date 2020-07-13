package adb

import (
    "fmt"
    "github.com/astaxie/beego/orm"
    "testing"
)

func TInitOrm() {
    err := StartMysqlOrm(Config{Host: "127.0.0.1", Username: "root", Password: "kf123456", Database: "asktop", SqlLogLevel: 1})
    if err != nil {
        panic(err)
    }
}

func TestOrmSelect(t *testing.T) {
    TInitOrm()
    var maps []orm.Params
    count, err := Orm().Raw("SELECT id, name,value FROM config WHERE id >= ? AND id < ?", 1, 5).Values(&maps)
    if err != nil {
        fmt.Println(err)
    } else {
        for _, row := range maps {
            fmt.Println(row["id"], ":", row["name"], ":", row["value"])
        }
        fmt.Println(count)
    }

    fmt.Println("--------------------")

    var configs []Config
    count, err = Orm().Raw("SELECT id, name, value FROM config WHERE id >= ? AND id < ?", 1, 5).QueryRows(&configs)
    if err != nil {
        fmt.Println(err)
    } else {
        for _, config := range configs {
            fmt.Println(config)
        }
        fmt.Println(count)
    }

    fmt.Println("--------------------")

    var config Config
    err = Orm().Raw("SELECT id, name, value FROM config WHERE id >= ? AND id < ?", 1, 5).QueryRow(&config)
    if err != nil {
        fmt.Println(err)
    } else {
        fmt.Println(config)
    }

    fmt.Println("--------------------")

    var kv orm.Params
    count, err = Orm().Raw("SELECT name, value FROM config WHERE id >= ? AND id < ?", 1, 5).RowsToMap(&kv, "name", "value")
    if err != nil {
        fmt.Println(err)
    } else {
        for k, v := range kv {
            fmt.Println(k, ":", v)
        }
        fmt.Println(count)
    }
}

//拼接sql，然后执行
func TestOrmSelectSql(t *testing.T) {
    TInitOrm()
    sql := OrmSql().Select("id", "name", "value").From("config").Where("id >= ? AND id < ?").String()
    fmt.Println(sql)

    var maps []orm.Params
    count, err := Orm().Raw(sql, 1, 5).Values(&maps)
    if err != nil {
        fmt.Println(err)
    } else {
        for _, row := range maps {
            fmt.Println(row["id"], ":", row["name"], ":", row["value"])
        }
        fmt.Println(count)
    }
}

func TestOrmInsert(t *testing.T) {
    TInitOrm()
    rs, err := Orm().Raw("INSERT INTO config(name, value) VALUES('key1', 'value1'),('key2', 'value2')").Exec()
    if err != nil {
        fmt.Println(err)
    } else {
        fmt.Println(rs.LastInsertId())
        fmt.Println(rs.RowsAffected())
    }

    fmt.Println("--------------------")

    rs, err = Orm().Raw("INSERT INTO config(name, value) VALUES(?, ?)", "key3", "value3").Exec()
    if err != nil {
        fmt.Println(err)
    } else {
        fmt.Println(rs.LastInsertId())
        fmt.Println(rs.RowsAffected())
    }

    fmt.Println("--------------------")

    // 预处理，批量插入
    pre, err := Orm().Raw("INSERT INTO config(name, value) VALUES(?, ?)").Prepare()
    if err != nil {
        fmt.Println(err)
    } else {
        rs, err = pre.Exec("key4", "value4")
        rs, err = pre.Exec("key5", "value5")
        pre.Close()
    }
}
