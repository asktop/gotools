package aclient

import (
    "encoding/json"
    "fmt"
    "testing"
)

//测试JsonRpc客户端
func TestJsonRpc_Demo(t *testing.T) {
    url := "http://127.0.0.1:8102/rpc"
    //url := "http://192.168.0.32:8102/rpc"
    controller := "JsonRpc"
    method := "Demo" //与服务端注册的 应用结构名.结构方法名 对应

    //定义jsonRpc方法入参（参数首字必须大写，客户端与服务端必须一致）
    type JsonRpcDemoIn struct {
        Id   int         `json:"id"`
        Name string      `json:"name"`
        Age  json.Number `json:"age"`
    }

    //定义jsonRpc方法出参（参数首字必须大写，客户端与服务端必须一致）
    type JsonRpcDemoOut struct {
        Id   int         `json:"id"`
        Name string      `json:"name"`
        Age  json.Number `json:"age"`
    }

    in := JsonRpcDemoIn{
        Id:   2,
        Name: "abc",
        Age:  "18",
    }
    var out JsonRpcDemoOut

    err := NewJsonRpcClient(url, controller).Call(method, in, &out)
    if err != nil {
        fmt.Println("--- rpc error ---")
        fmt.Println(err)
    } else {
        fmt.Println("--- rpc success ---")
        fmt.Println(out)
    }

}
