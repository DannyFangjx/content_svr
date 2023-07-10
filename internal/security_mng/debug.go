package security_mng

import (
	"content_svr/pub/logger"
	"context"
)

func (p *SecurityMng) Debug(ctx context.Context) error {
	logger.Infof(ctx, "GetUserWorkLikeMgDB")
	return nil
}
