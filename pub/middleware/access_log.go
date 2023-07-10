package middleware

import (
	"bytes"
	"content_svr/internal/busi_comm/constant/cm_const"
	"content_svr/internal/kafka_proxy"
	"content_svr/pub/logger"
	"content_svr/pub/requestid"
	"content_svr/pub/utils"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"net/http"
	"time"
)

const (
	CtxClientIp = "ctx_client_ip"
)

func AccessLogMiddleware(kfkProxy kafka_proxy.IKafkaProxy) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		ctx.Set(CtxClientIp, ctx.Request.Header.Get("X-REAL-IP"))
		requestID := ctx.GetHeader(requestid.CtxRequestID)
		if requestID == "" {
			requestID = requestid.GenerateRequestID()
		}
		ctx.Writer.Header().Set(requestid.CtxRequestID, requestID)
		ctx.Set(requestid.CtxRequestID, requestID)
		path := ctx.Request.URL.Path
		method := ctx.Request.Method
		if path == "/test" || path == "/metrics" {
			return
		}
		start := time.Now()
		var reqBuf []byte
		if method == http.MethodPost ||
			method == http.MethodPut ||
			method == http.MethodDelete {
			reqBuf, _ = ioutil.ReadAll(ctx.Request.Body)
			ctx.Request.Body = ioutil.NopCloser(bytes.NewBuffer(reqBuf))
		} else if method == http.MethodGet {
			m := map[string]interface{}{}
			for k, v := range ctx.Request.URL.Query() {
				m[k] = v[0]
			}
			reqBuf, _ = json.Marshal(m)
		}
		SetRequestBody(ctx, string(reqBuf))
		respBuf := utils.GetBytesBuffer()
		respBuf.Reset()
		accessW := &accessWriter{ctx.Writer, respBuf}
		ctx.Writer = accessW
		ctx.Next() //执行后面的中间件和handler，完了再实行下面的。
		respBody := respBuf.Bytes()
		utils.PutBytesBuffer(respBuf)
		if len(respBody) > cm_const.MaxBodySize {
			respBody = respBody[:cm_const.MaxBodySize]
		}
		latency := time.Since(start) / time.Millisecond
		if path != "/" { //心跳检查产生的日志太多了
			logger.Access(ctx, fmt.Sprintf("%d##%s##%s##%d##%s##%s##%s",
				latency, path, method, accessW.ResponseWriter.Status(),
				string(reqBuf), string(respBody), ctx.Request.RemoteAddr))
			//日志上报
			if kfkProxy != nil {
				go kfkProxy.SendLog(ctx, start.UnixNano()/1e6, time.Now().UnixNano()/1e6, string(reqBuf), string(respBody))
			}
		}
	}
}

type accessWriter struct {
	gin.ResponseWriter
	buf *bytes.Buffer
}

func (a accessWriter) WriteString(s string) (int, error) {
	a.buf.WriteString(s)
	return a.ResponseWriter.Write([]byte(s))
}

func (a accessWriter) Write(data []byte) (int, error) {
	a.buf.Write(data)
	return a.ResponseWriter.Write(data)
}

func (a accessWriter) WriteHeader(code int) {
	a.ResponseWriter.WriteHeader(code)
}
