package main

import (
	"context"
	"fmt"
	"regexp"
	"time"

	"github.com/jstang9527/buoy/micro/proto/model"
	"github.com/jstang9527/buoy/micro/proto/rpcapi"
	"github.com/micro/go-micro"
	"github.com/micro/go-micro/registry"
	"github.com/micro/go-plugins/registry/consul"
)

// TrojanClient 远程客户端交互
func TrojanClient() {
	reg := consul.NewRegistry(registry.Addrs("172.31.50.178:8500"))
	service := micro.NewService(micro.Name("victims.dtboCI"), micro.Registry(reg))
	service.Init()
	controller := rpcapi.NewVictimService("victims.dtboCI", service.Client())
	rsp, err := controller.Communication(context.Background(), &model.Request{Cmd: "pwd"})
	if err != nil {
		panic(err)
	}
	fmt.Println(rsp)
}

func reMethod() {
	// 5fb8e 3dc11 109ff b5e8c dc3e
	// 5fb78 266db 210b4 6b329 77cb
	text := "5fb8e3dc11109ffb5e8cdc3e.weblogic.172.31.50.178:6379.62"
	assetRe, _ := regexp.Compile("^\\w{24}")
	// re, _ := regexp.Compile("(\\d{1,3}\\.){3}\\d{1,3}:\\d{2,5}")
	fmt.Println(string(assetRe.Find([]byte(text))))
}

// TimeOutMethod ...
func TimeOutMethod(ctx context.Context) {
	for i := 0; i < 20; i++ {
		select {
		case <-ctx.Done(): // 10秒后执行
			fmt.Println("ctx over")
			return // beak是没有用的
		default:
			fmt.Printf("exec method-(%v)\n", i)
			time.Sleep(time.Second)
		}

	}
}

// TimeOutMethodTest ...
func TimeOutMethodTest() {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()
	TimeOutMethod(ctx)
	fmt.Println("finished.")
}

func mapTest() {
	val := make(map[string]string, 10)
	val["aa"] = "aaaa"
	fmt.Printf("%#v", val["aaa"])
}

func main() {
	reMethod()
	// TimeOutMethodTest()
	// mapTest()
}
