package utils

//
//import (
//	"content_svr/pub/logger"
//	"context"
//	"fmt"
//	"time"
//)
//
//func Retry(ctx context.Context, retryTimes int, sleep time.Duration, fn func(ctx context.Context) error) error {
//	if err := fn(ctx); err != nil {
//		if retryTimes--; retryTimes > 0 {
//			logger.Error(ctx, fmt.Sprintf("retry func error retryTimes #%d after %s.", retryTimes, sleep), err)
//			if sleep > 0 {
//				time.Sleep(sleep)
//			}
//			return Retry(ctx, retryTimes, sleep, fn)
//		}
//		return err
//	}
//	return nil
//}
