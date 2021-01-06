package main

import (
	"fmt"
	"strconv"
	"time"

	"github.com/gomodule/redigo/redis"
)

// SMQ 肉鸡执行队列
const SMQ string = "execMQ"

// RMQ 返回执行结果队列
const RMQ string = "dataMQ"

func producer() {
	// redisConn, err := redis.Dial("tcp", "127.0.0.1:6379", redis.DialPassword("hdiot"))
	redisConn, err := redis.Dial("tcp", "127.0.0.1:6379")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer redisConn.Close()

	var i = 1
	for i < 10 {
		_, err = redisConn.Do("rpush", RMQ, strconv.Itoa(i))
		if err != nil {
			fmt.Println("produce error")
			continue
		}
		fmt.Printf("produce element:%d\n", i)
		time.Sleep(time.Second)
		i++
	}
}

func consumer() {
	// redisConn, err := redis.Dial("tcp", "127.0.0.1:6379", redis.DialPassword("hdiot"))
	redisConn, err := redis.Dial("tcp", "127.0.0.1:6379")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer redisConn.Close()

	for {
		ele, err := redis.String(redisConn.Do("lpop", RMQ))
		if err != nil {
			fmt.Println("no msg.sleep now")
			time.Sleep(time.Second * 2)
		} else {
			fmt.Printf("cosume element:%s\n", ele)
		}
	}
}

func main() {
	go producer()
	consumer()
}
