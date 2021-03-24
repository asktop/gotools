package aqywx

import (
    "fmt"
    "github.com/asktop/gotools/ajson"
    "testing"
)

func TestClient_GetDepartmentList(t *testing.T) {
    corpid := "wwd7cff2c33f1c771f"
    corpsecret := ""
    token := ""

    client := NewClient(corpid, corpsecret).GetBase()
    if token != "" {
        client.SetAccessToken(token)
    }
    token, err := client.GetAccessToken()
    if err != nil {
        fmt.Println(err)
        return
    }
    fmt.Println(token)

    ds, err := client.GetDepartmentList()
    if err != nil {
        fmt.Println(err)
        return
    }
    fmt.Println(ajson.Encode(ds))
}
