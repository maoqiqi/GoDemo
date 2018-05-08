package config

import (
	"os"
	"io/ioutil"
	"encoding/json"
	"log"
	"app/lib/redis"
)

var (
	configPath  string
	RedisConfig map[string]*redis.Config
)

func init() {
	// 获取go path
	goPath := os.Getenv("GOPATH")

	// 获取运行环境
	appEnv := os.Getenv("APP_ENV")
	if appEnv == "" {
		appEnv = "local"
	}

	if goPath == "" {
		configPath = "./config/" + appEnv + "/"
	} else {
		configPath = goPath + "/src/app/config/" + appEnv + "/"
	}
	log.Print("configPath=", configPath)

	LoadConfig(&RedisConfig, "redis")
}

// 加载配置文件
func LoadConfig(v interface{}, name string) error {
	file := configPath + name + ".json"
	buffer, err := ioutil.ReadFile(file)
	if err != nil {
		return err
	}
	return json.Unmarshal(buffer, v)
}
