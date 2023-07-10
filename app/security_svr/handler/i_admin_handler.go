package handler

import (
	"content_svr/internal/security_mng"
	"content_svr/pub/decorate"
	"github.com/gin-gonic/gin"
)

type AdminHandler struct {
	SecurityMng security_mng.ISecurityMng
}

func NewAdminHandler(mng security_mng.ISecurityMng) *AdminHandler {
	return &AdminHandler{
		SecurityMng: mng,
	}
}

func SetUrls(router *gin.Engine, adminHandler *AdminHandler, wares ...gin.HandlerFunc) {
	baseAPI := router.Group("platform/security/")

	// debug
	baseAPI.GET("debug/common_debug",
		decorate.SerializeDecorator(adminHandler.Debug))
}
