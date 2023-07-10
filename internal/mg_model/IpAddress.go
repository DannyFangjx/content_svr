package mg_model

import (
	"content_svr/protobuf/pbmgdb"
	"content_svr/pub/logger"
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type IIpAddressMgModel interface {
	GetByIp(ctx context.Context, ip string) (*pbmgdb.IpAddressMgDbModel, error)
	Insert(ctx context.Context, item *pbmgdb.IpAddressMgDbModel) error
}

type IpAddressMgDbImpl struct {
	MgDB  *mongo.Database
	Table string
}

func NewIpAddressMgModelImpl(db *mongo.Database) IIpAddressMgModel {
	return &IpAddressMgDbImpl{
		MgDB:  db,
		Table: "ipAddress",
	}
}

func (impl *IpAddressMgDbImpl) GetByIp(ctx context.Context, ip string) (*pbmgdb.IpAddressMgDbModel, error) {
	retItems := make([]*pbmgdb.IpAddressMgDbModel, 0)
	collection := impl.MgDB.Collection(impl.Table)
	find, err := collection.Find(ctx, bson.M{"_id": ip})
	if err != nil {
		logger.Error(ctx, fmt.Sprintf("IpAddressMgDbModel Find failed. id=%v",
			ip), err)
		return nil, err
	}
	// 遍历查询结果
	for find.Next(ctx) {
		demo := &pbmgdb.IpAddressMgDbModel{}
		// 解码绑定数据
		err = find.Decode(demo)
		if err != nil {
			logger.Error(ctx, fmt.Sprintf("decode to IpAddressMgDbModel failed.id=%v",
				ip), err)
			return nil, err
		}
		retItems = append(retItems, demo)
	}

	if len(retItems) == 0 {
		return nil, nil
	}
	if len(retItems) > 1 {
		logger.Error(ctx, fmt.Sprintf("get IpAddressMgDbModel failed.id=%v",
			ip), err)
	}
	return retItems[0], err
}

func (impl *IpAddressMgDbImpl) Insert(ctx context.Context, item *pbmgdb.IpAddressMgDbModel) error {
	collection := impl.MgDB.Collection(impl.Table)
	result, err := collection.InsertOne(ctx, item)
	print(result, err)
	return err
}
