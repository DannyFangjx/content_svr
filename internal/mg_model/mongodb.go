package mg_model

import (
	"content_svr/config"
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
	"time"
)

func NewMongoDBInstance(config *config.MongodbConfig) (*mongo.Database, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	// 建立连接
	client, err := mongo.Connect(ctx,
		options.Client().
			// 连接地址
			ApplyURI("mongodb://"+config.Addr).
			// 设置验证参数
			SetAuth(
				options.Credential{
					// 用户名
					Username: config.User,
					// 密码
					Password: config.Pwd,
				}).
			// 设置连接数
			SetMaxPoolSize(20))
	if err != nil {
		//log.Println(err)
		return nil, err
	}
	// 测试连接
	err = client.Ping(ctx, readpref.Primary())
	if err != nil {
		//log.Println(err)
		return nil, err
	}

	return client.Database(config.DBName), nil
}
