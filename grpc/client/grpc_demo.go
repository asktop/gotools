package client

import (
    "github.com/asktop/gotools/grpc/protobuf"
    "google.golang.org/grpc"
)

var grpcAddr = "127.0.0.1:8883" //不可以加scheme

//创建grpc客户端
//git clone https://github.com/grpc/grpc-go.git --depth=1 $GOPATH/src/google.golang.org/grpc
//git clone https://github.com/googleapis/go-genproto.git --depth=1 $GOPATH/src/google.golang.org/genproto
//import "google.golang.org/grpc
func NewGrpcDemoClient(host ...string) (client pb.DemoClient, conn *grpc.ClientConn, err error) {
    addr := grpcAddr
    if len(host) > 0 && host[0] != "" {
        addr = host[0]
    }
    //连接grpc服务器
    conn, err = grpc.Dial(addr, grpc.WithInsecure())
    if err != nil {
        return client, nil, err
    }
    client = pb.NewDemoClient(conn)
    return client, conn, nil
}
