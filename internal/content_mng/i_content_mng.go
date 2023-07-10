package content_mng

import (
	"content_svr/internal/data_cache"
	"content_svr/internal/inner_mng"
	"content_svr/internal/kafka_proxy"
	"context"
)

type ContentMng struct {
	DataCache  data_cache.IDataCacheMng
	InnerProxy inner_mng.IInnerProxy
	KafkaProxy kafka_proxy.IKafkaProxy
}

type IContentMng interface {
	Debug(ctx context.Context) error
}

func NewContentMng(
	dataCache data_cache.IDataCacheMng,
	innerProxy inner_mng.IInnerProxy,
	kafkaProxy kafka_proxy.IKafkaProxy) IContentMng {
	return &ContentMng{
		DataCache:  dataCache,
		InnerProxy: innerProxy,
		KafkaProxy: kafkaProxy,
	}
}
