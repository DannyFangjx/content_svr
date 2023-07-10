package decorate

//
//import (
//	"net/http"
//
//	"github.com/gin-gonic/gin"
//
//	"pms/internal/p_common/errors"
//	"pms/internal/p_common/middleware"
//)
//
//func SopAuthDecorator(authCodesNeed ...string) gin.HandlerFunc {
//	return func(ctx *gin.Context) {
//		allAuth := middleware.GetAllAuth(ctx)
//		authCodes := middleware.GetAuthCodes(ctx)
//		if allAuth {
//			ctx.Next()
//			return
//		}
//		authMap := make(map[string]bool, len(authCodes))
//		for _, _code := range authCodes {
//			authMap[_code] = true
//		}
//		for _, code := range authCodesNeed {
//			if _, ok := authMap[code]; ok {
//				ctx.Next()
//				return
//			}
//		}
//		ctx.AbortWithStatusJSON(http.StatusForbidden, errors.ErrAuth)
//	}
//}
