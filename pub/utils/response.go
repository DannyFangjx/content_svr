package utils

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func RenderFile(ctx *gin.Context, filename string, content []byte) error {
	// write header before write body
	ctx.Writer.Header().Add("Content-type", "application/octet-stream")
	ctx.Writer.Header().Add("Content-Disposition", fmt.Sprintf("attachment;filename=%s", filename))
	if _, err := ctx.Writer.Write(content); err != nil {
		return err
	}
	ctx.Writer.Header().Add("Content-Length", strconv.Itoa(len(content)))
	ctx.Writer.WriteHeader(http.StatusOK)
	ctx.Abort()
	return nil
}
