package client

import (
    "context"
    "fmt"
    "github.com/asktop/gotools/grpc/protobuf"
    "testing"
    "time"
)

func TestCallGrpc(t *testing.T) {
    //创建客户端
    client, conn, err := NewGrpcDemoClient()
    if conn == nil {
        fmt.Printf("grpc请求失败，err:%s\n", err.Error())
        return
    }
    defer conn.Close()

    //调用方法
    in := pb.GetUsersIn{
        Name: "abc",
    }

    //方式1：直接请求
    //out, err := client.GetUsers(context.Background(), &in)

    //方式2：设置请求超时时间
    ctx, cancel := context.WithTimeout(context.Background(), time.Minute*5)
    defer cancel()
    out, err := client.GetUsers(ctx, &in)

    if err != nil {
        fmt.Printf("grpc请求失败，err:%s\n", err.Error())
        return
    }
    fmt.Println(out)
    fmt.Println(out.Status)
    fmt.Println(out.UserMap)
    fmt.Println(out.Users)
}
