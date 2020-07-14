package server

import (
    "github.com/asktop/gotools/grpc/protobuf"
    "github.com/asktop/gotools/log"
    "google.golang.org/grpc"
    "net"
    "os"
)

//启动 grpc 服务
//git clone https://github.com/grpc/grpc-go.git --depth=1 $GOPATH/src/google.golang.org/grpc
//git clone https://github.com/googleapis/go-genproto.git --depth=1 $GOPATH/src/google.golang.org/genproto
//import "google.golang.org/grpc"
func StartGrpc(port string) {
    //监听端口，启动服务
    log.Info("--- 启动 grpc 服务 ---", "端口:", port)

    //创建 gRPC Server
    server := grpc.NewServer()

    //注册应用（方法）
    pb.RegisterDemoServer(server, Demo{})

    listen, err := net.Listen("tcp", ":"+port)
    if err != nil {
        log.Error("--- grpc 服务启动失败 ---", "err:", err)
        os.Exit(0)
    }
    err = server.Serve(listen)
    if err != nil {
        log.Error("--- grpc 服务启动失败 ---", "err:", err)
        os.Exit(0)
    }
}
