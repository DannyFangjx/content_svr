package requestid

import (
	"content_svr/pub/utils"
	"context"
)

const (
	CtxRequestID = "X-Request-Id"
)

func WithRequestID(ctx context.Context) context.Context {
	if _, ok := ctx.Value(CtxRequestID).(string); ok {
		return ctx
	}
	return context.WithValue(ctx, CtxRequestID, utils.UUID())
}

func GetRequestID(ctx context.Context) string {
	requestID, ok := ctx.Value(CtxRequestID).(string)
	if ok {
		return requestID
	}
	return ""
}

func SetRequestID(ctx context.Context, requestID string) context.Context {
	return context.WithValue(ctx, CtxRequestID, requestID)
}

func GenerateRequestID() string {
	return utils.UUID()
}
