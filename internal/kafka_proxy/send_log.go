package kafka_proxy

import (
	"content_svr/config"
	"content_svr/internal/busi_comm/comm_struct"
	"content_svr/protobuf/pbkfk"
	"content_svr/pub/logger"
	"content_svr/pub/requestid"
	"content_svr/pub/utils"
	"encoding/json"
	"fmt"
	"github.com/Shopify/sarama"
	"github.com/gin-gonic/gin"
)

// sendToSander
func (p *KafkaProxyImpl) SendLog(ctx *gin.Context, startTs, endTs int64, reqStr, respStr string) error {
	topic := topicSecretLogGather

	var userId int64 = 0
	if val, ok := ctx.Get("ctx_user_id"); ok && val != nil {
		userId, _ = val.(int64)
	}

	ctxHeader := &pbkfk.CtxHeader{
		UserAgent:      utils.GetCtxHeadersStrVal(ctx, "User-Agent"),
		RemoteAddr:     utils.GetCtxHeadersStrVal(ctx, "X-Forwarded-For"),
		RemoteHost:     ctx.Request.Host,
		RemotePort:     0, // todo
		LocalAddr:      utils.GetLocalIp(),
		LocalPort:      int32(config.ServerConfig.Port),
		PhoneId:        utils.GetCtxHeadersStrVal(ctx, "phoneId"),
		PhoneBrand:     utils.GetCtxHeadersStrVal(ctx, "phoneBrand"),
		PhoneModel:     utils.GetCtxHeadersStrVal(ctx, "phoneModel"),
		PhoneOS:        utils.GetCtxHeadersStrVal(ctx, "phoneOS"),
		PhoneOSVersion: utils.GetCtxHeadersStrVal(ctx, "phoneOSVersion"),
		Token:          utils.GetCtxHeadersStrVal(ctx, "Token"),
		Uk:             utils.GetCtxHeadersStrVal(ctx, "Uk"),
		Versioncode:    utils.GetCtxHeadersStrVal(ctx, "Versioncode"),
		Channel:        utils.GetCtxHeadersStrVal(ctx, "Channel"),
		Entrance:       utils.GetCtxHeadersStrVal(ctx, "Entrance"),
		Timestamp:      utils.GetCtxHeadersStrVal(ctx, "Timestamp"),
		Appname:        utils.GetCtxHeadersStrVal(ctx, "Appname"),
		Sign:           utils.GetCtxHeadersStrVal(ctx, "Sign"),
		Apptype:        utils.GetCtxHeadersStrVal(ctx, "Apptype"),
	}

	kfkData := &comm_struct.LoggerKfkMsg{
		MessageId: requestid.GetRequestID(ctx),
		Method:    ctx.Request.Method,
		Status:    int32(ctx.Writer.Status()),
		Uri:       ctx.Request.URL.Path,
		UserId:    userId,
		AppType:   utils.GetCtxHeadersStrVal(ctx, "Apptype"),
		Header:    ctxHeader,
		Request:   reqStr,
		Response:  respStr,
		Start:     startTs,
		End:       endTs,
	}

	kfkDataBytes, _ := json.Marshal(kfkData)
	dataDto := map[string]interface{}{
		"type": "secretLog",
		"data": string(kfkDataBytes),
	}
	dataBytes, _ := json.Marshal(dataDto)
	msg := &sarama.ProducerMessage{
		Topic:     topic,
		Partition: int32(-1),
		Key:       sarama.StringEncoder(fmt.Sprintf("%v", requestid.GetRequestID(ctx))),
		Value:     sarama.ByteEncoder(dataBytes),
	}
	_, _, err := p.Producer.SendMessage(msg)
	if err != nil {
		logger.Error(ctx, fmt.Sprintf("SendLog kafka sendmessage failed. topic=%v", topic), err)
		return err
	}
	//logger.Infof(ctx, "SendLog kafka send message suc. topic=%v, body=%v, paritition=%v, offset=%v",
	//	topic, string(dataBytes), paritition, offset)
	return nil
}
