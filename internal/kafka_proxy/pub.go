package kafka_proxy

import (
	"content_svr/config"
	"strings"
)

var (
	// topics
	topicSecretLogGather = "secret_log_gather"
)

type constParamDef struct {
	redPacket       string
	normalRedPacket string
	singleTalk      string
	groupTalk       string
	systemMessage   string
}

var ConstParamDef = &constParamDef{
	redPacket:       "redPacket",
	normalRedPacket: "normalRedPacket",
	singleTalk:      "singleTalk",
	groupTalk:       "groupTalk",
	systemMessage:   "systemMessage",
}

func getTopic(topic string) string {
	if strings.ToLower(config.ServerConfig.Env) == "prod" {
		return topic
	} else {
		return "test_" + topic
	}
}
