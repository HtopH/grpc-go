package main

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	pb "google.golang.org/grpc/examples/mytest/mytest"
	"time"
)

const (
	address = "118.89.66.195:50051"
)

func main() {
	conn, err := grpc.Dial(address, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		fmt.Println(err)
		return
	}
	defer conn.Close()
	c := pb.NewGreeterClient(conn)
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	r, err := c.UserInfo(ctx, &pb.UserInfoRequest{
		People: "张小海",
		Age:    28,
	})
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(r.Message)

	res, err := c.Add(ctx, &pb.AddRequest{
		One:   1,
		Two:   2,
		Three: 3,
	})
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(res.Res)
}
