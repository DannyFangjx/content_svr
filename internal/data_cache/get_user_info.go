package data_cache

import (
	"content_svr/protobuf/pbapi"
	"content_svr/pub/logger"
	"context"
	"fmt"
	go_cache "github.com/fanjindong/go-cache"
	"github.com/gogo/protobuf/proto"
	"time"
)

type UserInfoLocal struct {
	UserInfoDbModel *pbapi.UserinfoDbModel
}

func (p *DataCacheMng) GetUserBasicInfo(ctx context.Context, userId int64, bSkipCache bool) (*pbapi.UserinfoDbModel, error) {
	item := p.getUserInfoModelLD(ctx, userId, bSkipCache)
	if item == nil {
		logger.Error(ctx, fmt.Sprintf("getUserInfoModelLD ret nil. userid=%v", userId), nil)
		return nil, nil
	}
	return item, nil
}

func (p *DataCacheMng) GetUserInfoLocal(ctx context.Context, header *pbapi.HttpHeaderInfo,
	userId int64, bSkipCache bool) (*UserInfoLocal, error) {
	resp := &UserInfoLocal{}
	item := p.getUserInfoModelLD(ctx, userId, bSkipCache)
	if item == nil {
		logger.Error(ctx, fmt.Sprintf("getUserInfoModelLD ret nil. userid=%v", userId), nil)
		return nil, nil
	}

	resp.UserInfoDbModel = item // lcache +db
	//lcache+mgDb
	return resp, nil
}

var localCacheKeyFmtUserinfoDbModel = "UserinfoDbModel:%v" // user_id
func (p *DataCacheMng) setUserInfoModelLocalCache(ctx context.Context, uInfoDbModel *pbapi.UserinfoDbModel) {
	if uInfoDbModel == nil || uInfoDbModel.GetUserId() == 0 {
		return
	}
	lKey := fmt.Sprintf(localCacheKeyFmtUserinfoDbModel, uInfoDbModel.GetUserId())
	bret := p.LocalCache.Set(lKey, uInfoDbModel, go_cache.WithEx(time.Duration(5*60)*time.Second))
	if !bret {
		logger.Error(ctx, fmt.Sprintf("setUserInfoModelLocalCache failed, user_id=%v", uInfoDbModel.GetUserId()), nil)
	}
	return
}

func (p *DataCacheMng) getUserInfoModelLocalCache(userId int64) *pbapi.UserinfoDbModel {
	if userId == 0 {
		return nil
	}
	lKey := fmt.Sprintf(localCacheKeyFmtUserinfoDbModel, userId)
	cResp, exist := p.LocalCache.Get(lKey)
	if exist && cResp != nil {
		lResp, ok := cResp.(*pbapi.UserinfoDbModel)
		if ok && lResp != nil {
			return lResp
		}
	}
	return nil
}

func (p *DataCacheMng) getUserInfoModelLD(ctx context.Context, userId int64, bSkipCache bool) *pbapi.UserinfoDbModel {
	var err error
	// 从内存取
	if bSkipCache == false {
		item := p.getUserInfoModelLocalCache(userId)
		if item != nil {
			return item
		}
	}

	// 从db取
	item, err := p.UserInfoModel.GetByUserId(ctx, userId)
	if err != nil {
		return nil
	}
	if item == nil {
		return nil
	}
	if item.GetNickName() == "" {
		item.NickName = proto.String(item.GetOpenNickName())
	}
	if item.GetPhoto() == "" {
		item.Photo = proto.String(item.GetOpenPhoto())
	}
	// 返回非空，则设置到内存缓存
	if bSkipCache == false {
		p.setUserInfoModelLocalCache(ctx, item)
	}
	return item
}
