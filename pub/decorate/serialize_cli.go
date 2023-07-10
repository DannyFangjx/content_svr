package decorate

import (
	"content_svr/pub/errors"
	"content_svr/pub/logger"
	"content_svr/pub/middleware"
	"content_svr/pub/requestid"
	"content_svr/pub/utils"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"net/http"
	"reflect"
	"strings"
	"time"
)

// 给客户端用。返回的结构是 errors.CliError
func SerializeDecoratorCli(handler interface{}) func(ctx *gin.Context) {
	hType := reflect.TypeOf(handler)
	hValue := reflect.ValueOf(handler)
	realFc := func(ctx *gin.Context) {}
	realBody := func(args []reflect.Value) (results []reflect.Value) {
		ctx := args[0].Interface().(*gin.Context)
		var realIN []reflect.Value
		realIN = append(realIN, args[0])
		if hType.NumIn() == 2 {
			param := hType.In(1)
			if param.Kind() == reflect.Ptr {
				param = param.Elem()
			}
			val := reflect.New(param)
			if ctx.Request.Method == http.MethodGet {
				if err := ctx.ShouldBindQuery(val.Interface()); err != nil {
					logger.Errorf(ctx, "bind to query failed. err=%v", err)
					ctx.AbortWithStatus(http.StatusInternalServerError)
					return nil
				}
			} else {
				if err := json.Unmarshal([]byte(middleware.GetRequestBody(ctx)), val.Interface()); err != nil {
					logger.Errorf(ctx, "json Unmarshal to struct failed. data=%v, err=%v",
						middleware.GetRequestBody(ctx), err)
					ctx.AbortWithStatus(http.StatusInternalServerError)
					return nil
				}
			}
			for _i := 0; _i < val.Elem().NumField(); _i++ {
				if val.Elem().Field(_i).Kind() == reflect.String {
					val.Elem().Field(_i).SetString(strings.Trim(val.Elem().Field(_i).String(), " "))
				}
			}
			realIN = append(realIN, val)
		}
		vals := hValue.Call(realIN)
		// 最后一个返回参数是error
		valNum := hType.NumOut()
		if valNum != 0 {
			statusCode := http.StatusOK
			var content interface{}
			content = http.StatusText(http.StatusOK)
			if vals[valNum-1].Interface() != nil {
				logger.Errorf(ctx, "resp failed. err=%v", utils.StructToJsonString(vals[valNum-1].Interface()))
				//应答失败
				result := errors.CliError{
					Message:   "",
					MessageId: utils.GetLocalIp() + "_" + requestid.GetRequestID(ctx),
					Status:    201, // 默认err
					Timestamp: time.Now().UnixNano() / 1e6,
				}

				//result := errors.ErrServer
				errImpl, ok := vals[valNum-1].Interface().(*errors.Error)
				if ok {
					if errImpl.UserErrCode != 0 {
						result.Status = errImpl.UserErrCode
					}
					result.Message = errImpl.UserMsg
					if errImpl.HttpCode != 0 {
						statusCode = errImpl.HttpCode
					}
				} else {
					logger.Error(ctx, "internal server", vals[valNum-1].Interface().(error))
				}
				content = result
			} else if hType.NumOut() != 1 {
				//应答成功。
				result := errors.CliError{
					Message:   "",
					MessageId: utils.GetLocalIp() + "_" + requestid.GetRequestID(ctx),
					Status:    1000,
					Timestamp: time.Now().UnixNano() / 1e6,
				}
				result.Content = vals[0].Interface()
				content = result
			}
			if !ctx.IsAborted() {
				ctx.JSON(statusCode, content)
			}
		}
		ctx.Next()
		return nil
	}
	h := reflect.MakeFunc(reflect.TypeOf(realFc), realBody)
	return h.Interface().(func(ctx *gin.Context))
}
