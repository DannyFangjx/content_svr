package errors

// 返回给客户端的应答形式。 Status=1000 成功
type CliError struct {
	Content   interface{} `json:"content"`
	Message   string      `json:"message,omitempty"`   //应答备注
	MessageId string      `json:"messageId,omitempty"` //消息id
	Status    int32       `json:"status"`
	Timestamp int64       `json:"timestamp,omitempty"`
}

// 猫爪内容服务，目前admin端应答. Status=200 成功
type MzError struct {
	Data      interface{} `json:"data"`
	Message   string      `json:"message,omitempty"`    //应答备注
	MessageId string      `json:"message_id,omitempty"` //消息id
	Status    int32       `json:"status"`
	Timestamp int64       `json:"timestamp,omitempty"`
}
