package validator

import (
	"content_svr/internal/busi_comm/constant/cm_const"
	"content_svr/pub/errors"
	"content_svr/pub/logger"
	"content_svr/pub/utils"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/xeipuuv/gojsonschema"
	"net/http"
	"reflect"
	"strconv"
	"strings"
)

func GenValidateHandler(schemaMap interface{}) func(*gin.Context) {
	goLoader := gojsonschema.NewGoLoader(schemaMap)
	schema, err := gojsonschema.NewSchema(goLoader)
	if err != nil {
		panic(err)
	}
	return func(ctx *gin.Context) {
		body := ctx.GetString(cm_const.CtxReqBody)
		if ctx.Request.Method == http.MethodGet {
			_json := make(map[string]interface{})
			if err = json.Unmarshal([]byte(body), &_json); err != nil {
				logger.Error(ctx, "json unmarshal error", err)
			}
			convert := convertBody(schemaMap, _json)
			if convert {
				bodyB, _ := json.Marshal(_json)
				body = string(bodyB)
				ctx.Set(cm_const.CtxReqBody, body)
			}
		} else if strings.Contains(ctx.Request.Header.Get("Content-Type"), "form-data") {
			err := ctx.Request.ParseMultipartForm(128)
			if err != nil {
				return
			} //保存表单缓存的内存大小128M
			data := ctx.Request.Form
			_json := make(map[string]interface{})
			for k, v := range data {
				if len(v) == 0 {
					continue
				}
				_json[k] = v[0]
			}
			convertBody(schemaMap, _json)
			bodyB, _ := json.Marshal(_json)
			body = string(bodyB)
			ctx.Set(cm_const.CtxReqBody, body)
		}
		doc := gojsonschema.NewStringLoader(body)
		res, err := schema.Validate(doc)
		if err != nil {
			logger.Error(ctx, fmt.Sprintf("req.header=%v", utils.StructToJsonString(ctx.Request.Header)), err)
			ctx.JSON(http.StatusBadRequest, errors.ErrParam)
			ctx.Abort()
			return
		}
		if !res.Valid() {
			var errMsgs []string
			for _, e := range res.Errors() {
				errStr := e.String()
				errStr = strings.TrimSuffix(errStr, "/1")
				errStr = strings.TrimPrefix(errStr, "(root): ")
				errMsgs = append(errMsgs, errStr)
			}
			msg := strings.Join(errMsgs, "; ")
			logger.Error(ctx, fmt.Sprintf("schema valid error_code=%v, req.header=%v", msg, utils.StructToJsonString(ctx.Request.Header)), nil)
			errParam := errors.Error{ErrMsg: "error_param", ErrDetail: msg}
			ctx.JSON(http.StatusBadRequest, errParam)
			ctx.Abort()
			return
		}
	}
}

func convertBody(schemaMap interface{}, _json map[string]interface{}) bool {
	_s, _ := schemaMap.(map[string]interface{})
	_p, _ := _s["properties"].(map[string]interface{})
	convert := false
	for _f, _s := range _p {
		_s, _ := _s.(map[string]interface{})
		if _, ok := _json[_f]; ok && _s["type"] == "integer" {
			_v, err := strconv.Atoi(_json[_f].(string))
			if err == nil {
				_json[_f] = _v
				convert = true
			}
		}
	}
	return convert
}

func SerializeFormDecorator(handler interface{}) func(ctx *gin.Context) {
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
			err := ctx.ShouldBind(val.Interface())
			if err != nil {
				ctx.JSON(http.StatusBadRequest, errors.ErrParam)
				ctx.Abort()
				return nil
			}
			realIN = append(realIN, val)
		}
		vals := hValue.Call(realIN)
		if hType.NumOut() == 1 {
			if vals[0].Interface() != nil {
				if err, ok := vals[0].Interface().(*errors.Error); ok {
					logger.Error(ctx, "", err)
					ctx.JSON(http.StatusInternalServerError, vals[0].Interface())
				} else {
					logger.Error(ctx, "", vals[0].Interface().(error))
					ctx.JSON(http.StatusInternalServerError, errors.ErrServer)
				}
			} else {
				ctx.String(http.StatusOK, "ok")
			}
		} else if hType.NumOut() == 2 {
			if vals[1].Interface() != nil {
				if err, ok := vals[1].Interface().(*errors.Error); ok {
					logger.Error(ctx, "", err)
					ctx.JSON(http.StatusInternalServerError, vals[1].Interface())
				} else {
					logger.Error(ctx, "", err)
					ctx.JSON(http.StatusInternalServerError, errors.ErrServer)
				}
			} else {
				ctx.JSON(http.StatusOK, vals[0].Interface())
			}
		}
		ctx.Next()
		return nil
	}
	h := reflect.MakeFunc(reflect.TypeOf(realFc), realBody)
	return h.Interface().(func(ctx *gin.Context))
}

func TempalteSerializeDecorator(handler interface{}) func(ctx *gin.Context) {
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
			err := ctx.ShouldBind(val.Interface())
			if err != nil {
				ctx.JSON(http.StatusBadRequest, errors.ErrParam)
				ctx.Abort()
				return nil
			}
			realIN = append(realIN, val)
		}
		_ = hValue.Call(realIN)
		ctx.Next()
		return nil
	}
	h := reflect.MakeFunc(reflect.TypeOf(realFc), realBody)
	return h.Interface().(func(ctx *gin.Context))
}
