package cache_const

import (
	"content_svr/config"
	"strings"
)

type cacheDef struct {
	KeyFmt string //key
	Expire int64  //本地重新读取的过期时间。单位秒
	//RedisExpire int64  // redis超时时间, 单位秒
}

func getRdsEnvPrefix() string {
	if strings.ToLower(config.ServerConfig.Env) == "prod" {
		return "prod"
	} else {
		return "test"
	}
}

// 新管理平台登录态。
var AdminTokenRcache = cacheDef{
	KeyFmt: "platform:" + getRdsEnvPrefix() + ":soul_soup:AdminToken:%v", // token ~ op.email
	Expire: 1 * 86400,
}
