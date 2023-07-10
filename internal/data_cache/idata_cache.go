package data_cache

import (
	"content_svr/internal/mg_model"
	"content_svr/internal/model"
	go_cache "github.com/fanjindong/go-cache"
	rdsV8 "github.com/go-redis/redis/v8"
)

type DataCacheMng struct {
	//mg_db
	IpAddressMgModel mg_model.IIpAddressMgModel

	// mysql db
	UserInfoModel model.IUserinfoDbModelModel

	//
	LocalCache go_cache.ICache
	RedisCli   *rdsV8.Client
}

var mDataCacheMng *DataCacheMng = nil

type IDataCacheMng interface {
	// 获取对象
	GetImpl() *DataCacheMng
}

func NewDataCacheMng(
	IpAddressMgModel mg_model.IIpAddressMgModel,

	userDB model.IUserinfoDbModelModel,
	redis *rdsV8.Client,
	localCache go_cache.ICache) IDataCacheMng {
	iMng := &DataCacheMng{
		// mgDb
		IpAddressMgModel: IpAddressMgModel,

		// DB
		UserInfoModel: userDB,

		RedisCli:   redis,
		LocalCache: localCache,
	}
	mDataCacheMng = iMng
	return iMng
}

func (p *DataCacheMng) GetImpl() *DataCacheMng {
	return mDataCacheMng
}
