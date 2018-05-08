package main

import (
	"fmt"
	"strings"
	"app/lib/redis"
	"app/core/config"
)

func main() {
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
}
