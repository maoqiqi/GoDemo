package redis

import (
	"github.com/go-redis/redis"
	"strings"
	"fmt"
)

var (
	redisList map[string]*redis.Client
	errs      [] string
)

type Config struct {
	Addr     string
	Password string
	DB       int
}

func Connect(configs map[string]*Config) {
	defer func() {
		if len(errs) > 0 {
			panic("[redis] " + strings.Join(errs, "\n"))
		}
	}()

	redisList = make(map[string]*redis.Client)
	for name, config := range configs {
		client := redis.NewClient(&redis.Options{
			Addr:     config.Addr,
			Password: config.Password,
			DB:       config.DB,
		})

		_, err := client.Ping().Result()
		if err != nil {
			fmt.Println("[redis] ping", err.Error())
			errs = append(errs, err.Error())
		}

		redisList[name] = client
	}
}

func Client(name ... string) (*redis.Client, bool) {
	if redisList == nil {
		fmt.Println("Please call the redis.Connect method first")
		return nil, false
	}

	key := "default"
	if len(name) > 0 {
		key = name[0]
	}
	client, ok := redisList[key]
	return client, ok
}
