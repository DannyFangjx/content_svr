package security_mng

import (
	"content_svr/internal/data_cache"
	"content_svr/internal/inner_mng"
	"content_svr/internal/kafka_proxy"
	"context"
)

type SecurityMng struct {
	DataCache  data_cache.IDataCacheMng
	InnerProxy inner_mng.IInnerProxy
	KafkaProxy kafka_proxy.IKafkaProxy
}

type ISecurityMng interface {
	// debug func
	Debug(ctx context.Context) error
}

func NewSecurityMng(
	dataCache data_cache.IDataCacheMng,
	innerProxy inner_mng.IInnerProxy,
	kafkaProxy kafka_proxy.IKafkaProxy) ISecurityMng {
	return &SecurityMng{
		DataCache:  dataCache,
		InnerProxy: innerProxy,
		KafkaProxy: kafkaProxy,
	}
}
