package schema

import "content_svr/pub/validator"

var FinishUgPtRecordReqSchema = map[string]interface{}{ //
	"required": []string{"ids"},
	"type":     "object",
	"properties": map[string]interface{}{
		"ids": validator.Int64ArrSchemaNotBlank,
	},
}

var AddUgAccountidReqSchema = map[string]interface{}{ //
	"required": []string{"accountid", "accountType", "userId"},
	"type":     "object",
	"properties": map[string]interface{}{
		"accountid":   validator.StringSchema,
		"accountType": validator.Uint32Schema,
		"userId":      validator.Int64Schema,
		"id":          validator.Int64Schema,
	},
}
