package model

import (
	"content_svr/config"
	"content_svr/pub/errors"
	"content_svr/pub/logger"
	"context"
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	ormLogger "gorm.io/gorm/logger"
	"gorm.io/plugin/dbresolver"
	"time"
)

const (
	StatusDelete      = -1
	StatusValid       = 1
	StatusInValid     = 2
	ctxTransactionKey = "ctx_transaction"
)

func NewDBInstance(config *config.MysqlConfig) (*gorm.DB, error) {
	//masterURL := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&timeout=10s&readTimeout=5s&writeTimeout=5s",
	//	config.User, config.Pwd, config.Host, config.Port, config.DBName)
	masterURL := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&timeout=10s&readTimeout=5s&writeTimeout=5s",
		config.User, config.Pwd, config.Host, config.Port, config.DBName)
	db, err := gorm.Open(mysql.Open(masterURL), &gorm.Config{
		SkipDefaultTransaction: true,
		PrepareStmt:            true,
		Logger:                 ormLogger.Default.LogMode(ormLogger.Silent),
	})
	if err != nil {
		logger.Info(context.Background(), "open mysql fail")
		panic(err)
	}
	err = db.Use(dbresolver.Register(dbresolver.Config{
		Sources: []gorm.Dialector{mysql.Open(masterURL)},
		//Replicas: []gorm.Dialector{mysql.Open(slaveURL)},
		Policy: dbresolver.RandomPolicy{}}).
		// should use go1.15
		// SetConnMaxIdleTime(time.Hour).
		SetConnMaxLifetime(4 * time.Hour).
		SetMaxIdleConns(8).
		SetMaxOpenConns(58))
	if err != nil {
		logger.Error(context.Background(), "open mysql fail", err)
		return nil, err
	}
	return db, nil
}

type DbImpl struct {
	DB *gorm.DB
}

type IDbImpl interface {
	UpdateItem(ctx context.Context, model interface{}) error
	CreateItem(ctx context.Context, model interface{}) error
}

func (impl *DbImpl) table() string {
	return "impl_tab"
}

func (impl *DbImpl) UpdateItem(ctx context.Context, model interface{}) error {
	result := impl.DB.WithContext(ctx).Table(impl.table()).Updates(model)
	return errors.Wrap(result.Error)
}

func (impl *DbImpl) CreateItem(ctx context.Context, model interface{}) error {
	result := impl.DB.WithContext(ctx).Table(impl.table()).Debug().Create(model)
	return errors.Wrap(result.Error)
}
