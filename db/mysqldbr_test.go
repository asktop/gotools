package db

import (
    "fmt"
    "github.com/asktop/dbr"
    big "github.com/asktop/decimal"
    "github.com/asktop/gotools/cache"
    "github.com/asktop/gotools/atime"
    "github.com/shopspring/decimal"
    "testing"
)

func TInitDbr() {
    err := StartMysqlDbr(Config{Host: "127.0.0.1", Username: "root", Password: "kf123456", Database: "asktop", SqlLogLevel: 1})
    if err != nil {
        panic(err)
    }
}

type Demo struct {
    Id         int64
    User_id    int64               //数据库字段名：user_id		下划线法，首字母大写
    UserName   string              //数据库字段名：user_name	驼峰法，首字母和下划线后首字母大写
    CreateTime int64 `db:"w_time"` //数据库字段名：w_time		标签法，自由指定
}

// config 系统配置表
type DemoConfig struct {
    Id         int64           `json:"id" db:"id"`
    Module     string          `json:"module" db:"module"`           //模块名
    Name       string          `json:"name" db:"name"`               //参数名
    Value      string          `json:"value" db:"value"`             //参数值
    Remark     string          `json:"remark" db:"remark"`           //备注
    Status     int64           `json:"status" db:"status"`           //状态[0：禁用，1：启用]
    CreateTime *big.Big        `json:"create_time" db:"create_time"` //创建时间
    UpdateTime decimal.Decimal `json:"update_time" db:"update_time"` //修改时间
}

func TestDbrSelect(t *testing.T) {
    TInitDbr()
    // where 拼接条件 in 的值可以是一个数组
    ids := []int{7, 8, 9}
    rows, err := Session().Select("id", "name", "value").From("config").Where("id >= ? AND id < ? OR id in ?", 1, 5, ids).Rows()
    if err != nil {
        fmt.Println(err)
    } else {
        defer rows.Close()
        for rows.Next() {
            var id int64
            var key, value string
            rows.Scan(&id, &key, &value)
            fmt.Println(id, key, value)
        }
    }

    fmt.Println("--------------------")

    // where dbr条件 eq 的值可以是一个数组
    var configs []DemoConfig
    where := dbr.And(dbr.Gte("id", 1), dbr.Lt("id", 5))
    where = dbr.Or(where, dbr.Eq("id", ids))
    count, err := Session().Select("id", "name", "value", "create_time", "update_time").From("config").Where(where).Load(&configs)
    if err != nil {
        fmt.Println(err)
    } else {
        for _, config := range configs {
            fmt.Println(config)
        }
        fmt.Println(count)
    }
}

func TestDbrSelectMaps(t *testing.T) {
    TInitDbr()
    //map值的类型必须指定，不能是interface
    var maps []map[string]string
    _, err := Session().Select("id", "name", "value").From("config").Where("id >= ? AND id < ?", 1, 5).Load(&maps)
    if err != nil {
        fmt.Println(err)
    } else {
        for _, row := range maps {
            fmt.Println(row["id"], ":", row["name"], ":", row["value"])
        }
    }
}

func TestDbrSelectMap(t *testing.T) {
    TInitDbr()
    //map值的类型必须指定，不能是interface
    var maps map[string]string
    _, err := Session().Select("name", "value").From("config").Where("id >= ? AND id < ?", 1, 5).Load(&maps)
    if err != nil {
        fmt.Println(err)
    } else {
        for key, value := range maps {
            fmt.Println(key, ":", value)
        }
    }
}

//关联查询
func TestDbrSelectExpr(t *testing.T) {
    TInitDbr()
    var maps []map[string]string
    count, err := Session().Select("name", "start_time", "msg").From("task_log").
        Where("name in (select name from task where id in ?)", []int64{2}).
        //Where(dbr.Expr("name in (select name from task where id in ?)", []int64{2})). //表达式方式
        Load(&maps)
    if err != nil {
        fmt.Println(err)
    } else {
        fmt.Println(count)
        for _, data := range maps {
            fmt.Println(data)
        }
    }
}

func TestDbrSelectBetween(t *testing.T) {
    TInitDbr()
    ids := []int{7, 8, 9}
    stmt := Session().Select("id", "name", "value").From("config").
        Where(dbr.Or(dbr.Between("id", 1, 5), dbr.Eq("id", ids)))
    fmt.Println(stmt.Count())
    sql, _ := stmt.GetSQL()
    fmt.Println(sql)
    rows, err := stmt.Rows()
    if err != nil {
        fmt.Println(err)
    } else {
        defer rows.Close()
        for rows.Next() {
            var id int64
            var key, value string
            rows.Scan(&id, &key, &value)
            fmt.Println(id, key, value)
        }
    }
}

func TestDbrInsert(t *testing.T) {
    TInitDbr()
    config3 := DemoConfig{
        Name:  "key3",
        Value: "value3",
    }
    config4 := DemoConfig{
        Name:  "key4",
        Value: "value4",
    }
    stmt := Session().InsertInto("config").
        Columns("name", "value").
        Values("key1", "value1"). //批量插入
        Values("key2", "value2").
        Record(&config3).
        Record(&config4)
    sql, _ := stmt.GetSQL()
    fmt.Println(sql)
    rs, err := stmt.Exec()
    if err != nil {
        fmt.Println(err)
    } else {
        fmt.Println(rs.LastInsertId()) //批量插入第一条的id
        fmt.Println(rs.RowsAffected())
    }
}

func TestDbrInsertMap(t *testing.T) {
    TInitDbr()
    kv := map[string]interface{}{
        "name":  "key5",
        "value": "value5",
    }
    rs, err := Session().InsertInto("config").Map(kv).Exec()
    if err != nil {
        fmt.Println(err)
    } else {
        fmt.Println(rs.LastInsertId())
    }
}

func TestDbrUpdate(t *testing.T) {
    TInitDbr()
    rs, err := Session().Update("config").
        Set("value", "value1 u").
        Set("status", dbr.Expr("status - ?", "1")). //值为表达式
        Set("update_time", atime.Now().Unix()).
        Where("name = 'key1'").Exec()
    if err != nil {
        fmt.Println(err)
    } else {
        fmt.Println(rs.RowsAffected())
    }

    fmt.Println("--------------------")

    setMap := map[string]interface{}{
        "value":       "value2 u",
        "status":      dbr.Expr("status + ?", "1"), //值为表达式
        "update_time": atime.Now().Unix(),
    }
    rs, err = Session().Update("config").
        SetMap(setMap).
        Where("name = 'key2'").Exec()
    if err != nil {
        fmt.Println(err)
    } else {
        fmt.Println(rs.RowsAffected())
    }
}

func TestDbrDelete(t *testing.T) {
    TInitDbr()
    names := []string{
        "key1",
        "key2",
        "key3",
        "key4",
        "key5",
    }
    rs, err := Session().DeleteFrom("config").Where("name in ?", names).Exec()
    if err != nil {
        fmt.Println(err)
    } else {
        fmt.Println(rs.RowsAffected())
    }
}

func TestCaseUpdate(t *testing.T) {
    TInitDbr()
    stmt := Session().CaseUpdate("user").Columns("name", "age", "info").
        Values("a", dbr.Expr("age + ?", 10), "aaa").
        Values("b", dbr.Expr("age - ?", 1), "bbb").
        Values("c", 31, "ccc").
        Values("d", 41, "ddd")
    fmt.Println(stmt.GetSQL())
    //err := stmt.SetRunLen(3).Exec()
    //fmt.Println(err)

    /*
    UPDATE `user`
    SET `age` =
        CASE `name`
    WHEN 'a' THEN
    age + 10
    WHEN 'b' THEN
    age - 1
    WHEN 'c' THEN
    31
    WHEN 'd' THEN
    41
    END,
    `info` =
    CASE `name`
    WHEN 'a' THEN
    'aaa'
    WHEN 'b' THEN
        'bbb'
    WHEN 'c' THEN
        'ccc'
    WHEN 'd' THEN
        'ddd'
    END
    WHERE
    `name` IN ('a', 'b', 'c', 'd')
    */
}

type DemoUser struct {
    UserId   int64
    Mobile   string
    Username string
}

func TestCache(t *testing.T) {
    TInitDbr()
    //此处循环调用，不能再此处测试
    err := cache.StartRedis(cache.Config{Host: "127.0.0.1", Password: "kf123456", Select: 0})
    if err != nil {
        panic(err)
    }
    redisconn := cache.NewRedis(true)
    defer redisconn.Close()
    stmt := Session().Select("user_id", "mobile", "username").From("user").Where(dbr.Eq("user_id", []int64{1, 2, 3})).Cache(redisconn, 60)
    sql, err := stmt.GetSQL()
    if err != nil {
        fmt.Println("GetSQL:", err)
    } else {
        fmt.Println("GetSQL:", sql)
    }
    count, err := stmt.Count()
    if err != nil {
        fmt.Println("Count:", err)
    } else {
        fmt.Println("Count:", count)
    }
    users := []DemoUser{}
    count, err = stmt.Load(&users)
    if err != nil {
        fmt.Println("Return:", err)
    } else {
        fmt.Println("Return:", count)
        fmt.Println("Return:", users)
    }
}
