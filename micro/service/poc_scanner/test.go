package main

import (
	"context"
	"fmt"

	pb "github.com/jstang9527/buoy/micro/proto/grpc"
	"google.golang.org/grpc"
)

const (
	// Address gRPC服务地址
	Address = "127.0.0.1:50051"
)

func main() {
	// 连接
	conn, err := grpc.Dial(Address, grpc.WithInsecure())
	if err != nil {
		fmt.Println(err)
		return
	}
	defer conn.Close()

	// 初始化客户端
	// c := pb.NewHelloClient(conn)
	// c := pb.NewArithClient(conn)
	c := pb.NewPocScanClient(conn)

	// 调用相加方法
	req := &pb.VerifyRequest{Cmd: "jstang"}
	res1, err := c.VerifyPoc(context.Background(), req)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("schedule add method: %v = %v\n", req.Cmd, res1.Msg)
	// // 调用相加方法
	// res2, err := c.XiangJian(context.Background(), req)
	// if err != nil {
	// 	fmt.Println(err)
	// 	return
	// }
	// fmt.Printf("schedule subtract method: %v - %v = %v\n", req.Num1, req.Num2, res2.Result)
}
