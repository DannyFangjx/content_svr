package app_comm

import (
	"content_svr/internal/data_cache"
	"content_svr/internal/mg_model"
	"content_svr/internal/model"

	go_cache "github.com/fanjindong/go-cache"
	rdsV8 "github.com/go-redis/redis/v8"
	"go.mongodb.org/mongo-driver/mongo"
	"gorm.io/gorm"
)

func BuildDataCache(db *gorm.DB,
	mgDb *mongo.Database,
	redis_cli *rdsV8.Client,
	localCache go_cache.ICache) data_cache.IDataCacheMng {

	// 初始化 mgdb层

	IpAddressMgModel := mg_model.NewIpAddressMgModelImpl(mgDb)

	// 初始化model层
	userDbImpl := model.NewUserinfoDbModelImpl(db)

	dataCache := data_cache.NewDataCacheMng(
		IpAddressMgModel,
		userDbImpl,
		redis_cli,
		localCache)
	return dataCache
}
