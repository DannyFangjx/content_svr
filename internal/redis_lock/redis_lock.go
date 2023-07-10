package redis_lock

import (
	"github.com/go-redis/redis/v8"
	"time"
)

type DistributedLock struct {
	client *redis.Client
	key    string
	val    string
	ttl    int // 锁的过期时间，单位为秒
}

func NewDistributedLock(key, val string, ttl int) *DistributedLock {
	return &DistributedLock{
		client: redis.NewClient(&redis.Options{
			Addr:     "localhost:6379", // Redis 服务器地址和端口
			Password: "",               // Redis 认证密码
			DB:       0,                // Redis 数据库编号
		}),
		key: key,
		val: val,
		ttl: ttl,
	}
}

func (dl *DistributedLock) Acquire() bool {
	success, err := dl.client.SetNX(nil, dl.key, dl.val, time.Second*time.Duration(dl.ttl)).Result() // TODO
	if err != nil {
		return false
	}
	return success
}

func (dl *DistributedLock) Release() error {
	_, err := dl.client.Del(nil, dl.key).Result() // TODO
	return err
}
