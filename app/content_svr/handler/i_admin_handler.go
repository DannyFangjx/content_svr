package handler

import (
	"content_svr/internal/content_mng"
	"content_svr/pub/decorate"
	"github.com/gin-gonic/gin"
)

type AdminHandler struct {
	ContentMng content_mng.IContentMng
}

func NewAdminHandler(
	iKeyWordMng content_mng.IContentMng) *AdminHandler {
	return &AdminHandler{
		ContentMng: iKeyWordMng,
	}
}

func SetUrls(router *gin.Engine, adminHandler *AdminHandler, wares ...gin.HandlerFunc) {
	baseAPI := router.Group("platform/")
	// debug
	baseAPI.GET("debug/common_debug",
		decorate.SerializeDecoratorCli(adminHandler.Debug))
}
