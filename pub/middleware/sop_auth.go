package middleware

import (
	"content_svr/internal/busi_comm/cache_const"
	"content_svr/pub/logger"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis/v8"
	"net/http"
	"strings"
)

const (
	appReqHeaderToken = "Token"
)

func AppTokenMiddleware(redis_cli *redis.Client) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		uri := ctx.Request.RequestURI
		if uri == "/test/" || uri == "/metrics" {
			ctx.Next()
			return
		}
		token := ctx.Request.Header.Get(appReqHeaderToken)
		if token != "" {
			//rdsKey := fmt.Sprintf(cache_const.UserTokenRcache.KeyFmt, token)
			//dateBytes, err := redis_cli.Get(ctx, rdsKey).Result()
			//if err == nil && len(dateBytes) != 0 {
			//	resp := &pbapi.LoginTokenRedis{}
			//	json.Unmarshal([]byte(dateBytes), resp)
			//	SetUserID(ctx, resp.GetUserId())
			//}
		}
		ctx.Next()
	}
}

func AdminAuthMiddleware(redis_cli *redis.Client) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		uri := ctx.Request.RequestURI
		if uri == "/test/" || uri == "/metrics" {
			ctx.Next()
			return
		}
		token := ctx.Request.Header.Get(appReqHeaderToken)
		// 除了登录外，其他接口都必须传入token
		if token == "" && !strings.Contains(uri, "admin/mz/api/operator/login_in") {
			logger.Errorf(ctx, "token not find")
			ctx.AbortWithStatus(http.StatusForbidden)
			return
		}

		if token != "" {
			rdsKey := fmt.Sprintf(cache_const.AdminTokenRcache.KeyFmt, token)
			emailBytes, err := redis_cli.Get(ctx, rdsKey).Bytes()
			if err != nil {
				logger.Errorf(ctx, "token not find, rdsKey=%v, err=%v", rdsKey, err)
				ctx.AbortWithStatus(http.StatusForbidden)
				return
			}
			SetLoginUser(ctx, string(emailBytes))
		}
		ctx.Next()
	}
}
