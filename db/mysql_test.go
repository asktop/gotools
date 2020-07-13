package db

import (
    "fmt"
    "github.com/asktop/gotools/atime"
    "testing"
)

func TInitDb() {
    err := StartMysql(Config{Host: "127.0.0.1", Username: "root", Password: "kf123456", Database: "asktop", SqlLogLevel: 1})
    if err != nil {
        panic(err)
    }
}

func TestDbSelect(t *testing.T) {
    TInitDb()
    rows, err := Db().Query("SELECT name, value FROM config WHERE id >= ? AND id < ?", 1, 5)
    if err != nil {
        fmt.Println(err)
    } else {
        defer rows.Close()
        for rows.Next() {
            var key, value string
            rows.Scan(&key, &value)
            fmt.Println(key, ":", value)
        }
    }
}

func TestDbInsert(t *testing.T) {
    TInitDb()
    rs, err := Db().Exec("INSERT INTO config(name, value) VALUES ('key1', 'value1')")
    if err != nil {
        fmt.Println(err)
    } else {
        fmt.Println(rs.LastInsertId())
    }
}

func TestDbUpdate(t *testing.T) {
    TInitDb()
    rs, err := Db().Exec("UPDATE config SET value = ?, update_time = ? WHERE name = ?", "value1 u", atime.Now().Unix(), "key1")
    if err != nil {
        fmt.Println(err)
    } else {
        fmt.Println(rs.RowsAffected())
    }
}

func TestDbDelete(t *testing.T) {
    TInitDb()
    rs, err := Db().Exec("DELETE FROM config WHERE name = ?", "key1")
    if err != nil {
        fmt.Println(err)
    } else {
        fmt.Println(rs.RowsAffected())
    }
}
