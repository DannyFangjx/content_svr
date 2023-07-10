package decorate

//
//import (
//	"encoding/json"
//	"net/http"
//	"strings"
//
//	"github.com/mohae/deepcopy"
//
//	"github.com/gin-gonic/gin"
//	"github.com/xeipuuv/gojsonschema"
//
//	"pms/internal/p_common/errors"
//	"content_svr/pub/logger"
//	"pms/internal/p_common/middleware"
//)
//
//func SchemaValidateDecorate(schemaMap interface{}) func(*gin.Context) {
//	return func(ctx *gin.Context) {
//		source := ctx.GetString(middleware.CtxReqBody)
//		if ctx.Request.Method == http.MethodGet {
//			_s, _ := schemaMap.(map[string]interface{})
//			_p, _ := _s["properties"].(map[string]interface{})
//			for _f, _s := range _p {
//				_s, _ := _s.(map[string]interface{})
//				if _s["type"] == "integer" {
//					_p[_f] = map[string]interface{}{
//						"type":    "string",
//						"pattern": "^\\d+$",
//					}
//				}
//			}
//			_s["properties"] = _p
//			schemaMap = _s
//			params := make(map[string]interface{})
//			for _k, _v := range ctx.Request.URL.Query() {
//				params[_k] = strings.Trim(_v[0], " ")
//			}
//			body, err := json.Marshal(params)
//			if err != nil {
//				ctx.AbortWithStatus(http.StatusInternalServerError)
//				return
//			}
//			source = string(body)
//		}
//		doc := gojsonschema.NewStringLoader(source)
//		goLoader := gojsonschema.NewGoLoader(schemaMap)
//		schema, err := gojsonschema.NewSchema(goLoader)
//		if err != nil {
//			ctx.AbortWithStatus(http.StatusInternalServerError)
//			return
//		}
//		res, err := schema.Validate(doc)
//		if err != nil {
//			logger.Error(ctx, "", err)
//			ctx.AbortWithStatus(http.StatusInternalServerError)
//			return
//		}
//		if !res.Valid() {
//			var errMsgs []string
//			for _, err := range res.Errors() {
//				errStr := err.String()
//				errStr = strings.TrimSuffix(errStr, "/1")
//				errStr = strings.TrimPrefix(errStr, "(root): ")
//				errMsgs = append(errMsgs, errStr)
//			}
//			msg := strings.Join(errMsgs, "; ")
//			logger.Error(ctx, "schema valid error_code, "+msg, nil)
//			errParam := deepcopy.Copy(errors.ErrParam).(*errors.Error)
//			errParam.ErrDetail = msg
//			ctx.AbortWithStatusJSON(http.StatusBadRequest, errParam)
//			return
//		}
//		ctx.Next()
//	}
//}
