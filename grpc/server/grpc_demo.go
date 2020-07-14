package server

import (
    "context"
    "fmt"
    "github.com/asktop/gotools/grpc/protobuf"
)

//定义grpc应用
type Demo struct{}

//实现grpc方法
func (Demo) GetUsers(ctx context.Context, in *pb.GetUsersIn) (*pb.GetUsersOut, error) {
    name := in.GetName()
    fmt.Println(name)
    userMap := map[int64]string{
        1: name,
        2: name,
    }
    users := []*pb.User{
        {Id: 1, Name: name},
        {Id: 2, Name: name},
    }
    return &pb.GetUsersOut{Status: 1, UserMap: userMap, Users: users}, nil
}
