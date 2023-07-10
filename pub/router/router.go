package router

import (
	"content_svr/internal/kafka_proxy"
	"content_svr/pub/graceful.v1"
	"content_svr/pub/middleware"
	"fmt"
	"github.com/gin-contrib/pprof"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"time"
)

type Router struct {
	*gin.Engine
}

func NewRouter(kfkProxy kafka_proxy.IKafkaProxy) *Router {
	gin.SetMode(gin.ReleaseMode)
	router := gin.New()
	pprof.Register(router)
	router.Use(
		//middleware.MonitorMiddleware(),
		middleware.CORS(),
		middleware.AccessLogMiddleware(kfkProxy),
		middleware.Recovery(),
		//logger.LogerMiddleware(),
	)
	//router.GET("/metrics", gin.WrapH(promhttp.Handler()))
	router.GET("/test", func(ctx *gin.Context) {
		ctx.String(http.StatusOK, "ok")
	})
	router.GET("/", func(ctx *gin.Context) {
		ctx.String(http.StatusOK, "ok")
	})
	return &Router{Engine: router}
}

func NewAdminRouter(kfkProxy kafka_proxy.IKafkaProxy) *Router {
	gin.SetMode(gin.ReleaseMode)
	router := gin.New()
	pprof.Register(router)
	router.Use(
		//middleware.MonitorMiddleware(),
		middleware.CORS(),
		middleware.AccessLogMiddleware(kfkProxy),
		middleware.Recovery(),
		//logger.LogerMiddleware(),
	)
	//router.GET("/metrics", gin.WrapH(promhttp.Handler()))
	router.GET("/test", func(ctx *gin.Context) {
		ctx.String(http.StatusOK, "ok")
	})
	router.GET("/", func(ctx *gin.Context) {
		ctx.String(http.StatusOK, "ok")
	})
	return &Router{Engine: router}
}

func (r *Router) Run(port int) {
	log.Println("Router starting")
	graceful.Run(fmt.Sprintf(":%v", port), 30*time.Second, r.Engine)
}
