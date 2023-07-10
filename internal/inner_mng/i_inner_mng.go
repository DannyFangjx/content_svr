package inner_mng

import (
	"bytes"
	"content_svr/internal/busi_comm/cache_const"
	"content_svr/pub/utils"
	"context"
	"fmt"
	rdsV8 "github.com/go-redis/redis/v8"
	"net/http"
	"time"
)

type IInnerProxy interface {
}

type InnerProxyImpl struct {
	RedisCli *rdsV8.Client
}

func NewInnerProxyImpl(redisCli *rdsV8.Client) IInnerProxy {
	return &InnerProxyImpl{
		RedisCli: redisCli,
	}
}

func (p *InnerProxyImpl) postRequest(ctx context.Context, url string, body []byte) (*http.Response, error) {
	req, err := http.NewRequest(http.MethodPost, url, bytes.NewBuffer(body))
	if err != nil {
		return nil, err
	}
	req.Header.Add("Content-Type", "application/json")
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (p *InnerProxyImpl) getRequest(ctx context.Context, url string, body []byte) (*http.Response, error) {
	req, err := http.NewRequest(http.MethodGet, url, bytes.NewBuffer(body))
	if err != nil {
		return nil, err
	}
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (p *InnerProxyImpl) sendPutRequest(ctx context.Context, url string, body []byte) (*http.Response, error) {
	req, err := http.NewRequest(http.MethodPut, url, bytes.NewBuffer(body))
	if err != nil {
		return nil, err
	}
	req.Header.Add("Content-Type", "application/json")
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (p *InnerProxyImpl) setRdsSign(ctx context.Context) (int64, error) {
	sign := utils.GetCurTsMs()
	rdsKey := fmt.Sprintf(cache_const.InnerHttpRedisSign.KeyFmt, sign)
	expire := cache_const.InnerHttpRedisSign.Expire
	p.RedisCli.SetEX(ctx, rdsKey, sign, time.Duration(expire)*time.Second).Result()
	return sign, nil
}
