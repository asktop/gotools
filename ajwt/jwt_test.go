package ajwt

import (
    "fmt"
    "testing"
    "time"
)

func TestNewToken(t *testing.T) {
    //生成token
    token, _, err := NewToken(map[string]interface{}{"user_id": 123})
    if err != nil {
        fmt.Println(err)
    } else {
        fmt.Println(token)
    }
    //eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE1NjAwODg4MjEsInVzZXJpZCI6MTIzfQ.N5HT1gpwA2tXip9V9-47iwd9fWwHAY5waUZVKleMIkQ

    fmt.Println("--------------------")
    time.Sleep(time.Second * 2)

    //解析token
    info2, err := GetInfo(token)
    if err != nil {
        fmt.Println(err)
    } else {
        fmt.Println(info2)
        fmt.Println(info2["user_id"])
    }

    fmt.Println("--------------------")

    //解析token
    type Info struct {
        UserId int64 `json:"user_id"`
    }
    info := Info{}
    err = GetInfoObj(token, &info)
    if err != nil {
        fmt.Println(err)
    } else {
        fmt.Println(info)
        fmt.Println(info.UserId)
    }

}

func TestNewRsaToken(t *testing.T) {
    //生成token
    token, _, err := NewRsaToken(map[string]interface{}{"user_id": 123})
    if err != nil {
        fmt.Println(err)
    } else {
        fmt.Println(token)
    }
    //eyJhbGciOiJSUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE1NjAwODg5MzksInVzZXJpZCI6MTIzfQ.FSt7-tyhYbhumFIoC02sTikUQvX4zWq9EW5id2nV4F6tbuLq3E3Y1GMqeGzrNcpcFNtVvKUk2CSB-UoqQqKHWl7UxNeL3kCsxuZ_2XBS3y4Br3qaoOPEJR8hJ03d1z4hsJct62uPjXGXGkshuXJGJILZwj0MzfDKuJrcgVfZL5I

    fmt.Println("--------------------")
    time.Sleep(time.Second * 2)

    //解析token
    info2, err := GetRsaInfo(token)
    if err != nil {
        fmt.Println(err)
    } else {
        fmt.Println(info2)
        fmt.Println(info2["user_id"])
    }

    fmt.Println("--------------------")

    //解析token
    type Info struct {
        UserId int64 `json:"user_id"`
    }
    info := Info{}
    err = GetRsaInfoObj(token, &info)
    if err != nil {
        fmt.Println(err)
    } else {
        fmt.Println(info)
        fmt.Println(info.UserId)
    }

}
