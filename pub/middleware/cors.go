package middleware

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func CORS() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		method := ctx.Request.Method
		ctx.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		ctx.Header("Access-Control-Allow-Origin", "*")
		ctx.Header("Access-Control-Allow-Headers", "Crossorigin,Origin,"+
			"access-control-allow-methods, access-control-allow-origin,access-control-allow-credentials,"+
			"X-Requested-With,Accept,Content-Type,AccessToken,X-CSRF-Token,Authorization,Token,withcredentials")
		ctx.Header("Access-Control-Allow-Methods", "POST, PUT, GET, DELETE, OPTIONS")
		ctx.Header("Access-Control-Expose-Headers", "Content-Length,Access-Control-Allow-Origin,Access-Control-Allow-Headers,Content-Type")
		ctx.Header("Access-Control-Allow-Credentials", "true")
		if method == http.MethodOptions {
			ctx.AbortWithStatus(http.StatusNoContent)
		}
		ctx.Next()
	}
}

//func IsLive() bool {
//	env, ok := os.LookupEnv(constant.EnvKey)
//	return ok && env == string(constant.EnvLive)
//}
//
//func IsTest() bool {
//	env, ok := os.LookupEnv(constant.EnvKey)
//	return ok && env == string(constant.EnvTest)
//}
