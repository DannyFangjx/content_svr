package middleware

import (
	"content_svr/pub/errors"
	"content_svr/pub/logger"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func Recovery() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		defer func() {
			if err := recover(); err != nil {
				if str, ok := err.(string); ok {
					logger.Error(ctx, fmt.Sprintf("error convert to string %v", str), nil)
				} else if _e, ok := err.(error); ok {
					logger.Error(ctx, "recover", _e)
				} else {
					logger.Error(ctx, fmt.Sprintf("recover unknow type %v", err), nil)
				}
				ctx.JSON(http.StatusInternalServerError, errors.ErrServer)
			}
		}()
		ctx.Next()
	}
}
