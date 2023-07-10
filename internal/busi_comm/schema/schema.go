package schema

import "content_svr/pub/validator"

var PersonalTalkCheckReqSchema = map[string]interface{}{ //
	"required": []string{"content", "interact", "from_user_id", "to_user_id"},
	"type":     "object",
	"properties": map[string]interface{}{
		"content":      validator.StringSchema,
		"interact":     validator.Uint64Schema,
		"from_user_id": validator.Uint64Schema,
		"to_user_id":   validator.Uint64Schema,
	},
}

var SendMsgReqSchema = map[string]interface{}{ //
	"required": []string{"toUserId"},
	"type":     "object",
	"properties": map[string]interface{}{
		"toUserId":    validator.Uint64NotEmptySchema,
		"content":     validator.Str1000Schema,
		"workId":      validator.Int64Schema,          // 可能不传，比如互关的好友
		"messageType": validator.Uint32Schema,         // 可能不传？ todo 待前端确认。
		"width":       validator.Uint32Schema,         // 可能为空，只有图片和表情时候，才有值。
		"high":        validator.Uint32Schema,         // 可能为空，只有图片和表情时候，才有值。
		"duration":    validator.Uint32Schema,         // 可能为空，只有语音时候，才有值。
		"objectId":    validator.Str1000Schema,        // 可能为空，只有图片和语音时候，才有值。
		"clientId":    validator.Str1000Schema,        // todo待确认。
		"memeId":      validator.Uint64NotEmptySchema, // 可能为空，之有表情时 才有值
		"sessionType": validator.Uint32Schema,         // 可能为空，现在只有 0 -单聊
		"longitude":   validator.FloatSchema,          // 经度
		"latitude":    validator.FloatSchema,          // 纬度
	},
}

// todo
var PushChitchatReqSchema = map[string]interface{}{ //
	"required": []string{},
	"type":     "object",
	"properties": map[string]interface{}{
		"content":     validator.Str1000Schema,
		"workId":      validator.Int64Schema,          // 可能不传，比如互关的好友
		"messageType": validator.Uint32Schema,         // 可能不传？ todo 待前端确认。
		"width":       validator.Uint32Schema,         // 可能为空，只有图片和表情时候，才有值。
		"high":        validator.Uint32Schema,         // 可能为空，只有图片和表情时候，才有值。
		"duration":    validator.Uint32Schema,         // 可能为空，只有语音时候，才有值。
		"objectId":    validator.Str1000Schema,        // 可能为空，只有图片和语音时候，才有值。
		"clientId":    validator.Str1000Schema,        // todo待确认。
		"memeId":      validator.Uint64NotEmptySchema, // 可能为空，之有表情时 才有值
		"sessionType": validator.Uint32Schema,         // 可能为空，现在只有 0 -单聊
		"longitude":   validator.FloatSchema,          // 经度
		"latitude":    validator.FloatSchema,          // 纬度
	},
}

var KeywordsAddSchema = map[string]interface{}{ //
	"required": []string{"key_words", "kw_type_id", "operator"},
	"type":     "object",
	"properties": map[string]interface{}{
		"key_words": map[string]interface{}{
			"type":     "array",
			"items":    validator.StringSchema,
			"minItems": 0,    // >=0
			"maxItems": 5000, // <=3
		},
		"kw_type_id": validator.Uint64Schema,
		"operator":   validator.StringSchema,
	},
}

var KeywordsDelSchema = map[string]interface{}{ //
	"required": []string{"key_words", "kw_type_id", "operator"},
	"type":     "object",
	"properties": map[string]interface{}{
		"key_words": map[string]interface{}{
			"type":     "array",
			"items":    validator.StringSchema,
			"minItems": 0,    // >=0
			"maxItems": 5000, // <=3
		},
		"kw_type_id": validator.Uint64Schema,
		"operator":   validator.StringSchema,
	},
}

var SetStarSignBirthSchema = map[string]interface{}{ //
	"required": []string{"month", "day"},
	"type":     "object",
	"properties": map[string]interface{}{
		"month": validator.Uint64NotEmptySchema,
		"day":   validator.Uint64NotEmptySchema,
	},
}

var OpAddSchema = map[string]interface{}{ //
	"required": []string{"email", "pwd", "op_type", "name", "phone"},
	"type":     "object",
	"properties": map[string]interface{}{
		"email":   validator.EmailSchema,
		"pwd":     validator.Str256Schema,
		"op_type": validator.Uint32Schema,
		"name":    validator.Str256Schema,
		"phone":   validator.Str256Schema,
	},
}
