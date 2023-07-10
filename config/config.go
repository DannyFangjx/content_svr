package config

import (
	"content_svr/pub/utils"
	"encoding/json"
	"fmt"
	"os"
	"strconv"
)

type MysqlConfig struct {
	Host   string `json:"host"`
	Port   int    `json:"port"`
	User   string `json:"user"`
	Pwd    string `json:"pwd"`
	DBName string `json:"db_name"`
}

type RedisConfig struct {
	Addr string `json:"addr"`
	Pwd  string `json:"pwd"`
}

type MongodbConfig struct {
	Addr   string `json:"addr"`
	User   string `json:"user"`
	Pwd    string `json:"pwd"`
	DBName string `json:"db_name"`
}

type KafkaConfig struct {
	Addr []string `json:"addr"`
}

type Config struct {
	Env           string         `json:"env"`
	Port          int            `json:"port"`
	InnerHost     string         `json:"inner_host"`
	ImageHost     string         `json:"image_host"`
	CheckSvrHost  string         `json:"check_svr_host"`
	CheckSvrIp    string         `json:"check_svr_ip"`
	KafkaConfig   *KafkaConfig   `json:"kafka_config"`
	MysqlConfig   *MysqlConfig   `json:"mysql_config"`
	RedisConfig   *RedisConfig   `json:"redis_config"`
	MongodbConfig *MongodbConfig `json:"mongodb_config"`
}

var (
	ServerConfig Config
	NodeIdx      int64
)

func init() {
	wpath, err := utils.GetWorkPath()
	if err != nil {
		panic(fmt.Sprintf("get workpath failed. err:%v", err.Error()))
	}

	fileBytes, err := os.ReadFile(wpath + "/conf/config.json")
	if err != nil {
		panic(err)
	}
	err = json.Unmarshal(fileBytes, &ServerConfig)
	if err != nil {
		panic(err)
	}

	NODE_IDX_str := os.Getenv("NODE_IDX")
	NODE_IDX_int, err := strconv.Atoi(NODE_IDX_str)
	NodeIdx = int64(NODE_IDX_int)
}
