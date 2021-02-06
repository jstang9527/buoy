# buoy

# Require
- go-micro
- consul
- hydra

# Funcs
- 资产木马C/S 服务
- 资产木马CURD API
- 木马即时通讯 API

# Run
```shell
go mod tidy
swag init
go run main.go
```

# Todo
一次性资产扫描并汇报

-> 扫描主机端口并识别服务
-> 有服务则对应服务, 无服务则对应端口, 否则该端口使用所有poc进行扫描(支持用户自定义渗透模式(验证、利用、建立代理))
-> 


# tomorrow
先把任务资产那些检索先搞定吧