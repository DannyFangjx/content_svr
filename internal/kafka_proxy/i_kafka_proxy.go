package kafka_proxy

import (
	"content_svr/config"
	"content_svr/pub/logger"
	"context"
	"fmt"
	"github.com/Shopify/sarama"
	"github.com/gin-gonic/gin"
)

type KafkaProxyImpl struct {
	Producer sarama.SyncProducer
}

type IKafkaProxy interface {
	SendLog(ctx *gin.Context, startTs, endTs int64, reqStr, respStr string) error
}

func NewKafkaProxyImpl(cfg *config.KafkaConfig) IKafkaProxy {
	config := sarama.NewConfig()
	config.Producer.Return.Successes = true
	config.Producer.RequiredAcks = sarama.WaitForAll
	config.Producer.Partitioner = sarama.NewRandomPartitioner
	config.Version = sarama.V0_10_2_1

	producer, err := sarama.NewSyncProducer(cfg.Addr, config)
	if err != nil {
		panic(fmt.Sprintf("new kafka proxy producer failed. err=%v", err))
	}

	return &KafkaProxyImpl{
		Producer: producer,
	}
}

func (p *KafkaProxyImpl) produce(ctx context.Context, topic string, body []byte) error {
	if len(topic) == 0 || len(body) == 0 {
		return nil
	}

	msg := &sarama.ProducerMessage{
		Topic:     topic,
		Partition: int32(-1),
		Key:       sarama.StringEncoder("key"),
		Value:     sarama.ByteEncoder(body),
	}
	//paritition, offset, err := p.Producer.SendMessage(msg)
	_, _, err := p.Producer.SendMessage(msg)
	if err != nil {
		logger.Error(ctx, "kafka sendmessage failed.", err)
		return err
	}
	//logger.Infof(ctx, "kafka send message suc. topic=%v, body=%v, paritition=%v, offset=%v",
	//	topic, string(body), paritition, offset)
	return nil
}
