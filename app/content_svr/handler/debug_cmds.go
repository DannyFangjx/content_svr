package handler

import (
	"content_svr/protobuf/pbapi"
	"github.com/gin-gonic/gin"
)

func (p *AdminHandler) Debug(ctx *gin.Context) (*pbapi.BaseResp, error) {
	p.ContentMng.Debug(ctx)
	return &pbapi.BaseResp{}, nil
}
