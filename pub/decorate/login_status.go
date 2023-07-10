package decorate

//
//import (
//	"net/http"
//
//	"pms/internal/p_common/errors"
//
//	"pms/internal/p_common/constant"
//	"pms/internal/p_common/middleware"
//	"github.com/gin-gonic/gin"
//)
//
//func LoginStatusDecorator() gin.HandlerFunc {
//	return func(ctx *gin.Context) {
//		if ctx.Request.Method == http.MethodOptions && constant.CurrentEnv != constant.EnvLive {
//			ctx.Next()
//			return
//		}
//		loginUser, err := middleware.SessionLoginUser(ctx)
//		if err != nil {
//			ctx.AbortWithStatus(http.StatusInternalServerError)
//			return
//		}
//		if loginUser == nil {
//			ctx.AbortWithStatusJSON(http.StatusForbidden, errors.ErrNotLogin)
//			return
//		}
//		ctx.Next()
//	}
//}
