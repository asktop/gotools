package ajwt

import (
    "fmt"
    "github.com/asktop/gotools/atime"
    "testing"
    "time"
)

var (
    privateKey = `
-----BEGIN PRIVATE KEY-----
MIICdwIBADANBgkqhkiG9w0BAQEFAASCAmEwggJdAgEAAoGBAKmOMrDNDEMgapVC
HM+0hK36OPWVCzw4lRtPtxOSobV9C5h5y+RGSEbRwoZvHXntl19iuTwBMu9PXaNA
sE+WFvkpSRy0K0PcSsQ0p+hnqbumHvz6eAxnXLW1jZBi/3HE7xcZWDQZHk03jPVk
s2lMdsu2PG5TVUtAocq2ThEUDT+LAgMBAAECgYAQS+2NVrf1/7iezfLs98HE4wb0
e7XPvPR/4oKLLA3E3tbtec9iCmtJ+0FCII3puS9SaK+7F7Zoj+1FEqOfkqJ0eh+Z
Hbs3kSGrLTnUqhx6hfbbZ7KmOb5ER9YoXjMLiJa6GjrI3JSvTTX4Gq4gtVOmmgKo
Sk1SGw4hDeIRcjqf0QJBAN++0JIbsPfnhp9LC8oeyyh4sJlnXNRdgg35VF2pntLG
4w8MM+tulXz+hoh/yjipiynq+f5FMrUILxLj+Yg7M7kCQQDB/4m7Yx7OZB9E7uQD
tZZKJDsf6jOaITe27HVkVoaYfdjiJdWxgTqOu+XfUlSamDfS8T577ozYe+UJQbjC
J7djAkAdV/Yco1sTOB8Utw/lwyIbvbBTfhXTmCvdT0y8N+mndQQETjJk9wzN5seY
EUJhKgMhTTraGRMsYNVjodhxAGqpAkEAp2hGiKfua5hGy4uCxBitCmndg5rs0LKd
i4LdgqqQF4Nx5yVTFIw55fYLVf2L+KQmeCeqDfQ7Io03hhdqwlDXlQJBAISJqBbK
a/CqrACq9IsQgUaN+OSrVouJWieEYwRYCSHuHFN6nF3q4mOOeQ/qs03NKEL2y5Ad
ZbT6qUiSgz/zHdY=
-----END PRIVATE KEY-----
`
    publicKey = `
-----BEGIN PUBLIC KEY-----
MIGfMA0GCSqGSIb3DQEBAQUAA4GNADCBiQKBgQCpjjKwzQxDIGqVQhzPtISt+jj1
lQs8OJUbT7cTkqG1fQuYecvkRkhG0cKGbx157ZdfYrk8ATLvT12jQLBPlhb5KUkc
tCtD3ErENKfoZ6m7ph78+ngMZ1y1tY2QYv9xxO8XGVg0GR5NN4z1ZLNpTHbLtjxu
U1VLQKHKtk4RFA0/iwIDAQAB
-----END PUBLIC KEY-----
`
)

func TestNewToken(t *testing.T) {
    //生成token
    var exp int64
    //exp = atime.Now().Unix()
    exp = atime.Now().Add(time.Second * 60 * 60 * 24).Unix()
    token, err := Encrypt(map[string]interface{}{"user_id": 123}, exp, "asktop")
    if err != nil {
        fmt.Println(err)
        return
    } else {
        fmt.Println(token)
    }

    //token := `eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE1NjAwODg4MjEsInVzZXJpZCI6MTIzfQ.N5HT1gpwA2tXip9V9-47iwd9fWwHAY5waUZVKleMIkQ`

    fmt.Println("--------------------")
    time.Sleep(time.Second)

    //解析token
    info2, err := Decrypt(token, "asktop")
    fmt.Println(IsExpired(err))
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
    err = DecryptObj(token, &info, "asktop")
    fmt.Println(IsExpired(err))
    if err != nil {
        fmt.Println(err)
    } else {
        fmt.Println(info)
        fmt.Println(info.UserId)
    }

}

func TestNewRsaToken(t *testing.T) {
    //生成token
    var exp int64
    //exp = atime.Now().Unix()
    //exp = atime.Now().Add(time.Second * 60 * 60 * 24).Unix()
    token, err := RsaEncrypt(map[string]interface{}{"user_id": 123}, exp, privateKey)
    if err != nil {
        fmt.Println(err)
        return
    } else {
        fmt.Println(token)
    }

    //token := `eyJhbGciOiJSUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE1NjAwODg5MzksInVzZXJpZCI6MTIzfQ.FSt7-tyhYbhumFIoC02sTikUQvX4zWq9EW5id2nV4F6tbuLq3E3Y1GMqeGzrNcpcFNtVvKUk2CSB-UoqQqKHWl7UxNeL3kCsxuZ_2XBS3y4Br3qaoOPEJR8hJ03d1z4hsJct62uPjXGXGkshuXJGJILZwj0MzfDKuJrcgVfZL5I`

    fmt.Println("--------------------")
    time.Sleep(time.Second)

    //解析token
    info2, err := RsaDecrypt(token, publicKey)
    fmt.Println(IsExpired(err))
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
    err = RsaDecryptObj(token, &info, publicKey)
    fmt.Println(IsExpired(err))
    if err != nil {
        fmt.Println(err)
    } else {
        fmt.Println(info)
        fmt.Println(info.UserId)
    }

}
