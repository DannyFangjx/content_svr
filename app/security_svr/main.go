package main

import (
	"content_svr/app/app_comm"
	"content_svr/app/security_svr/handler"
	"content_svr/config"
	"content_svr/internal/inner_mng"
	"content_svr/internal/kafka_proxy"
	"content_svr/internal/mg_model"
	"content_svr/internal/model"
	"content_svr/internal/security_mng"
	"content_svr/pub/logger"
	"content_svr/pub/router"
	"content_svr/pub/snow_flake"
	"context"
	"fmt"
	go_cache "github.com/fanjindong/go-cache"
	rdsV8 "github.com/go-redis/redis/v8"
	"time"
)

func main() {
	//
	snow_flake.SetMachineID(config.NodeIdx)

	localCache := go_cache.NewMemCache()

	//初始化 redis_cli
	redis_cli := InitRdsClient(config.ServerConfig.RedisConfig.Addr, config.ServerConfig.RedisConfig.Pwd)

	// 初始化kafka
	kafkaProxy := kafka_proxy.NewKafkaProxyImpl(config.ServerConfig.KafkaConfig)

	//初始化DB
	db, err := model.NewDBInstance(config.ServerConfig.MysqlConfig)
	if err != nil {
		logger.Error(context.Background(), "new db instansce failed", err)
		return
	}

	// 初始化mongoDB
	mgDb, err := mg_model.NewMongoDBInstance(config.ServerConfig.MongodbConfig)
	if err != nil {
		logger.Error(context.Background(), "new mongo db instansce failed", err)
		return
	}

	// 初始化innerProxy
	innerProxy := inner_mng.NewInnerProxyImpl(redis_cli)

	// 初始化 data_cache层
	dataCache := app_comm.BuildDataCache(db, mgDb, redis_cli, localCache)

	kwMng := security_mng.NewSecurityMng(dataCache, innerProxy, kafkaProxy)

	// 初始化handle层
	adminServerHandle := handler.NewAdminHandler(kwMng)
	appRouter := router.NewRouter(nil)
	handler.SetUrls(appRouter.Engine, adminServerHandle)

	//运行服务
	logger.Infof(context.Background(), "server starting, port=%v, node_idx=%v....",
		config.ServerConfig.Port, config.NodeIdx)
	appRouter.Run(config.ServerConfig.Port)

}

func InitRdsClient(addr, pwd string) *rdsV8.Client {
	client := rdsV8.NewClient(&rdsV8.Options{
		Addr:         addr,
		Password:     pwd,
		DB:           0,
		ReadTimeout:  2 * time.Minute,
		WriteTimeout: 1 * time.Minute,
		PoolTimeout:  2 * time.Minute,
		IdleTimeout:  10 * time.Minute,
		PoolSize:     1000,
	})
	_, err := client.Ping(context.Background()).Result()
	if err != nil {
		panic(fmt.Sprintf("redis init failed. err:%v", err.Error()))
	}
	return client
}
