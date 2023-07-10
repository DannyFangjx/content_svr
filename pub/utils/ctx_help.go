package utils

import (
	"content_svr/protobuf/pbapi"
	"github.com/gin-gonic/gin"
	"strconv"
)

func GetCtxHeadersSession(c *gin.Context) (*pbapi.HttpHeaderInfo, error) {
	resp := &pbapi.HttpHeaderInfo{}
	resp.Apptype = GetCtxHeadersStrVal(c, "Apptype")
	resp.Appname = GetCtxHeadersStrVal(c, "Appname")
	resp.Token = GetCtxHeadersStrVal(c, "Token")
	resp.Versioncode = GetCtxHeadersStrVal(c, "Versioncode")
	resp.Uk = GetCtxHeadersStrVal(c, "Uk")
	resp.Ip = GetCtxHeadersStrVal(c, "X-Forwarded-For")
	resp.Platform = GetCtxHeadersStrVal(c, "Platform")
	strUserId := GetCtxHeadersStrVal(c, "Debuguserid")
	userId, _ := strconv.Atoi(strUserId)
	resp.Debuguserid = int64(userId)
	return resp, nil
}

func GetCtxHeadersStrVal(c *gin.Context, key string) string {
	values, exist := c.Request.Header[key]
	if !exist {
		return ""
	}
	return values[0]
}

func GetCtxHeadersInt(c *gin.Context, key string) (int32, error) {
	val := GetCtxHeadersStrVal(c, key)
	valInt, err := strconv.Atoi(val)
	if err != nil {
		return 0, err
	}
	return int32(valInt), nil
}
