package redis

import (
	"time"
	"github.com/go-redis/redis"
	"fmt"
)

const (
	CACHE_MIN_TTL  = 60 * time.Second     //1分钟
	CACHE_HOUR_TTL = 3600 * time.Second   //1小时
	CACHE_DAY_TTL  = 86400 * time.Second  //1天
	CACHE_WEEK_TTL = 604800 * time.Second //7天
)

type BaseRedis struct {
}

func (*BaseRedis) GetRedisClient(name ... string) (*redis.Client, error) {
	client, ok := Client(name...)
	if !ok {
		fmt.Printf("redis client not found")
	}
	return client, nil
}

func (b *BaseRedis) Set(key, val string, db ...string) error {
	client, err := b.GetRedisClient(db...)
	if err != nil {
		return err
	}
	return client.Set(key, val, CACHE_DAY_TTL).Err()
}

func (b *BaseRedis) Get(key string, db ...string) (string, error) {
	client, err := b.GetRedisClient(db...)
	if err != nil {
		return "", err
	}
	return client.Get(key).Result()
}
