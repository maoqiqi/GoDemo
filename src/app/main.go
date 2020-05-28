package main

import (
	"fmt"
	"strings"
	"app/lib/redis"
	"app/core/config"
	"app/lib/httpclient"
	"io/ioutil"
	"time"
)

func main1() {
	// strings.Join 用法
	arr := []string{"a", "b", "c", "d"}
	str := strings.Join(arr, "--")
	// 输出a--b--c--d
	fmt.Println(arr, str)

	// append
	var errs [] string
	errs = append(errs, "error", "error")
	fmt.Println(errs)

	redis.Connect(config.RedisConfig)
	baseRedis := redis.BaseRedis{}
	baseRedis.Set("key", "val")
	val, err := baseRedis.Get("key")
	fmt.Println(val, err)

	resp, err := httpclient.Client().Get("https://www.baidu.com")
	if err != nil {
		fmt.Println("Http error:", err)
	}

	respBytes, err := ioutil.ReadAll(resp.Body)
	resp.Body.Close()

	if err != nil {
		fmt.Println("Body Reader error:", err)
	}

	fmt.Println(string(respBytes))
}

const FORMAT_TIME = "2006-01-02 15:04:05"

func main2() {
	now := time.Now()
	fmt.Println("当前时间:", now)

	// 1s之前的时间
	negative1, _ := time.ParseDuration("-1s")
	time1 := now.Add(negative1).Format(FORMAT_TIME)
	fmt.Println("1s之前的时间:", time1)

	// 3s之前的时间
	negative3, _ := time.ParseDuration("-3s")
	time3 := now.Add(negative3).Format(FORMAT_TIME)
	fmt.Println("3s之前的时间:", time3)
}
