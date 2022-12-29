package main

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	pb "google.golang.org/grpc/examples/bsc/rpcProtobuf"
	"net"
	"strconv"
)

const (
	port = ":50051"
)

type server struct {
	pb.UnimplementedGreeterServer
}

func (s *server) UserInfo(ctx context.Context, in *pb.UserInfoRequest) (*pb.UserInfoReply, error) {
	age := in.Age + 2
	ageStr := strconv.Itoa(int(age))
	return &pb.UserInfoReply{Message: in.People + "后年就" + ageStr + "岁啦"}, nil
}

func (s *server) Add(ctx context.Context, param *pb.AddRequest) (*pb.AddReply, error) {
	data := &pb.AddReply{}
	data.Res = param.One + param.Two + param.Three
	return data, nil
}

func main() {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		fmt.Println(err)
		return
	}
	s := grpc.NewServer()
	pb.RegisterGreeterServer(s, &server{})
	if err = s.Serve(lis); err != nil {
		fmt.Println(err)
	}
}
