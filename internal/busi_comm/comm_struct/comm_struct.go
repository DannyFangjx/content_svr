package comm_struct

import "content_svr/protobuf/pbkfk"

// web socket消息结构
type WsMessageKafkaDto struct {
	UserIds        []int64                `json:"userIds"  structs:"userIds"`
	Types          []string               `json:"types" structs:"types"`
	Data           map[string]interface{} `json:"data" structs:"data"`
	MinimumVersion int32                  `json:"minimumVersion" structs:"minimumVersion"`
}

type LoggerKfkMsg struct {
	MessageId string           `json:"messageId"` // 日志id
	Method    string           `json:"method"`    //请求方式GET,POST,PUT,DEL。。。对应RequestMethod枚举
	Status    int32            `json:"status"`    //请求状态
	Uri       string           `json:"uri"`       // 请求uri
	UserId    int64            `json:"userId"`    // 	用户id
	AppType   string           `json:"appType"`   //客户端类型，对应appType枚举
	Header    *pbkfk.CtxHeader `json:"header"`    //header信息
	Request   interface{}      `json:"request"`   //请求参数
	Response  interface{}      `json:"response"`  //返回参数
	Start     int64            `json:"start"`     //请求开始时间
	End       int64            `json:"end"`       //请求结束时间
}
