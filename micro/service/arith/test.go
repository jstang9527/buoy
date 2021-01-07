package main

import (
	"context"
	"fmt"

	"github.com/e421083458/golang_common/lib"
	pb "github.com/jstang9527/buoy/micro/proto/grpc"
	"github.com/micro/go-micro/client/selector"
	"github.com/micro/go-micro/registry"
	"github.com/micro/go-plugins/registry/consul"
	"google.golang.org/grpc"
)

// User ...
type User struct {
	ServiceName     string
	RegistryAddress string
	Services        []*registry.Service
}

// InitUser ...
func InitUser(name, addr string) *User {
	return &User{
		ServiceName:     name,
		RegistryAddress: addr,
	}
}

// RegistryConn 连接注册中心
func (u *User) RegistryConn() error {
	reg := consul.NewRegistry(registry.Addrs(lib.GetStringConf(u.RegistryAddress)))
	Services, err := reg.GetService(u.ServiceName)
	if err != nil {
		return err
	}
	u.Services = Services
	return nil
}

// GetNodeByRamdom 获取节点
func (u *User) GetNodeByRamdom() (*registry.Node, error) {
	next := selector.Random(u.Services)
	node, err := next()
	if err != nil {
		return nil, err
	}
	return node, nil
}

func main() {
	// 1.连接注册中心获取服务节点
	user := InitUser("test_python3", "172.31.50.249:8500")
	if err := user.RegistryConn(); err != nil {
		fmt.Printf("Failed conn Registry center, info: %v\n", err)
		return
	}
	node, err := user.GetNodeByRamdom()
	if err != nil {
		fmt.Println("Failed get node, info: ", err)
		return
	}
	fmt.Printf("%#v\n", node)
	// 2.连接服务节点调用方法
	conn, err := grpc.Dial(node.Address, grpc.WithInsecure())
	if err != nil {
		fmt.Printf("Failed connect service node, info: %v\n", err)
		return
	}
	defer conn.Close()

	call := pb.NewArithClient(conn)
	resp, err := call.XiangJia(context.Background(), &pb.ArithRequest{Num1: 11, Num2: 22})
	if err != nil {
		fmt.Printf("Failed call method(XiangJia), info: %#v\n", err)
		return
	}
	fmt.Printf("add_result=%v\n", resp.Result)
}
