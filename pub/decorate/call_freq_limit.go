package decorate

//
//import (
//	"encoding/json"
//	"fmt"
//	"net/http"
//
//	"pms/internal/p_common/middleware"
//
//	"pms/internal/p_common/data_cache"
//
//	"pms/internal/p_common/errors"
//	"github.com/gin-gonic/gin"
//)
//
//func CallFreqLimitDecorator(redis *data_cache.RedisClient, period int, totalTimes int, ipTimes int, paramTimes int,
//	params ...string) gin.HandlerFunc {
//	return func(ctx *gin.Context) {
//		values := make(map[string]interface{})
//		if ctx.Request.Method == http.MethodGet {
//			for _k, _v := range ctx.Request.URL.Query() {
//				values[_k] = _v[0]
//			}
//		} else {
//			if err := json.Unmarshal([]byte(middleware.GetRequestBody(ctx)), &values); err != nil {
//				ctx.AbortWithStatus(http.StatusInternalServerError)
//				return
//			}
//		}
//		key := "call_freq_limit_key_" + ctx.Request.URL.Path
//		keys := make(map[string]int)
//		if totalTimes != 0 {
//			keys[key] = totalTimes
//		}
//		if ipTimes != 0 {
//			ip := ctx.Request.Header.Get("X-REAL-IP")
//			keyIp := fmt.Sprintf("%s_ip_%s", key, ip)
//			keys[keyIp] = ipTimes
//		}
//		for _, _p := range params {
//			if _v, ok := values[_p]; ok {
//				keyParam := fmt.Sprintf("%s_%s_%v", key, _p, _v)
//				keys[keyParam] = paramTimes
//			}
//		}
//		for _k, _l := range keys {
//			cur, err := redis.Incr(ctx, _k, period)
//			if err != nil {
//				ctx.AbortWithStatus(http.StatusInternalServerError)
//				return
//			}
//			if cur > int64(_l) {
//				ctx.AbortWithStatusJSON(http.StatusBadRequest, errors.ErrCallFreq)
//				return
//			}
//		}
//		ctx.Next()
//	}
//}
