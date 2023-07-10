package model

import (
	"content_svr/protobuf/pbapi"
	"content_svr/pub/errors"
	"context"
	"github.com/gogo/protobuf/proto"
	"gorm.io/gorm"
	"gorm.io/plugin/dbresolver"
)

type IUserinfoDbModelModel interface {
	// 根据业务而不同。
	CreateItem(ctx context.Context, model *pbapi.UserinfoDbModel) (*pbapi.UserinfoDbModel, error)
	Create(ctx context.Context, userId int64, userName string) (*pbapi.UserinfoDbModel, error)

	// 更新，根据业务而不同。
	//UpdateItem(ctx context.Context, model *pbapi.UserinfoDbModel) (*pbapi.UserinfoDbModel, error)

	GetByUserId(ctx context.Context, userId int64) (*pbapi.UserinfoDbModel, error)
	ListItemsByCondition(ctx context.Context, condition map[string]interface{},
		page, size int32) ([]*pbapi.UserinfoDbModel, error)
	CountItemsByCondition(ctx context.Context, condition map[string]interface{}) (total int64, _ error)
	DictByUserId(ctx context.Context, userIds []int64) (map[int64]*pbapi.UserinfoDbModel, error)
}

type UserinfoDbModelImpl struct {
	DB *gorm.DB
}

func NewUserinfoDbModelImpl(db *gorm.DB) IUserinfoDbModelModel {
	return &UserinfoDbModelImpl{DB: db}
}

func (impl *UserinfoDbModelImpl) table() string {
	return "user_info"
}
func (impl *UserinfoDbModelImpl) GetByUserId(ctx context.Context, userId int64) (*pbapi.UserinfoDbModel, error) {
	model := pbapi.UserinfoDbModel{}
	result := impl.DB.WithContext(ctx).Clauses(dbresolver.Write).Table("user_info").
		Where(&pbapi.UserinfoDbModel{UserId: proto.Int64(userId)}).First(&model)
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return nil, nil
	}
	return &model, errors.Wrap(result.Error)
}

func (impl *UserinfoDbModelImpl) CreateItem(ctx context.Context, model *pbapi.UserinfoDbModel) (*pbapi.UserinfoDbModel, error) {
	result := impl.DB.WithContext(ctx).Table(impl.table()).Debug().Create(model)
	return model, errors.Wrap(result.Error)
}

func (impl *UserinfoDbModelImpl) Create(ctx context.Context, userId int64, userName string) (*pbapi.UserinfoDbModel, error) {
	item := &pbapi.UserinfoDbModel{
		UserId:   proto.Int64(userId),
		NickName: proto.String(userName),
	}
	impl.CreateItem(ctx, item)
	return nil, nil
}

func (impl *UserinfoDbModelImpl) ListItemsByCondition(ctx context.Context, condition map[string]interface{},
	page, size int32) ([]*pbapi.UserinfoDbModel, error) {
	offset := (page - 1) * size
	var items []*pbapi.UserinfoDbModel
	result := impl.DB.WithContext(ctx).Table(impl.table()).Limit(int(size)).Offset(int(offset)).
		Where(condition).Order("id desc").Find(&items)
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return items, nil
	}
	return items, errors.Wrap(result.Error)
}

func (impl *UserinfoDbModelImpl) CountItemsByCondition(ctx context.Context, condition map[string]interface{}) (total int64, _ error) {
	result := impl.DB.WithContext(ctx).Table(impl.table()).Model(&pbapi.UserinfoDbModel{}).
		Where(condition).Count(&total)
	return total, errors.Wrap(result.Error)
}

func (impl *UserinfoDbModelImpl) DictByUserId(ctx context.Context, userIds []int64) (map[int64]*pbapi.UserinfoDbModel, error) {
	var items []*pbapi.UserinfoDbModel
	result := impl.DB.WithContext(ctx).Table(impl.table()).Model(&pbapi.UserinfoDbModel{}).
		Where("user_id in ?", userIds).Find(&items)

	resp := make(map[int64]*pbapi.UserinfoDbModel)
	for _, item := range items {
		resp[item.GetUserId()] = item
	}
	return resp, errors.Wrap(result.Error)
}
