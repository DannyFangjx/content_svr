package middleware

import (
	"context"
	"github.com/gin-gonic/gin"
)

type (
	ContextValueKey string
)

const (
	CtxReqBody        = "req_body"
	CtxUserID         = "ctx_user_id"
	CtxLoginUser      = "ctx_login_user"
	CtxSOPCountry     = "ctx_sop_country"
	CtxSOPAuthCodes   = "ctx_sop_auth_codes"
	CtxSOPAllAuth     = "ctx_sop_all_auth"
	CtxBDCenterRegion = "bd-region"
)

func GetCountry(ctx *gin.Context) string {
	return ctx.GetString(CtxSOPCountry)
}

func SetCountry(ctx *gin.Context, country string) {
	ctx.Set(CtxSOPCountry, country)
}

func GetLoginUser(ctx *gin.Context) string {
	return ctx.GetString(CtxLoginUser)
}

func SetLoginUser(ctx *gin.Context, currentUser string) {
	ctx.Set(CtxLoginUser, currentUser)
}

func GetUserID(ctx *gin.Context) (i int64) {
	if val, ok := ctx.Get(CtxUserID); ok && val != nil {
		i, _ = val.(int64)
	}
	return
}

func SetUserID(ctx *gin.Context, currentUser int64) {
	ctx.Set(CtxUserID, currentUser)
}

func GetAuthCodes(ctx *gin.Context) []string {
	return ctx.GetStringSlice(CtxSOPAuthCodes)
}

func SetAuthCodes(ctx *gin.Context, codes []string) {
	ctx.Set(CtxSOPAuthCodes, codes)
}

func GetRequestBody(ctx *gin.Context) string {
	return ctx.GetString(CtxReqBody)
}

func SetRequestBody(ctx *gin.Context, body string) {
	ctx.Set(CtxReqBody, body)
}

func SetAllAuth(ctx *gin.Context, allAuth bool) {
	ctx.Set(CtxSOPAllAuth, allAuth)
}

func GetAllAuth(ctx *gin.Context) bool {
	return ctx.GetBool(CtxSOPAllAuth)
}

func GetBDCenterRegion(ctx *gin.Context) string {
	return ctx.GetString(CtxBDCenterRegion)
}

func SetBDCenterRegion(ctx *gin.Context, region string) {
	ctx.Set(CtxBDCenterRegion, region)
}

func GetOpEmail(ctx context.Context) string {
	return ctx.Value(CtxLoginUser).(string)
}
