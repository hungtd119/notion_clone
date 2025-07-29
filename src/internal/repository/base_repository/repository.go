package base_repository

import (
	"context"
	"notion/src/config"

	"gorm.io/gorm"
)

const keyDB = "db"

func GetDBFromContext(ctx context.Context) *gorm.DB {
	val, ok := ctx.Value(keyDB).(*gorm.DB)

	if !ok {
		return config.GetInstance().GetDB()
	}

	return val
}

func SetDBWithContext(ctx context.Context, db *gorm.DB) context.Context {
	return context.WithValue(ctx, keyDB, db)
}

type BaseRepository struct {
	model interface{}
}

func (r *BaseRepository) GetDB(ctx context.Context) *gorm.DB {
	return GetDBFromContext(ctx)
}

func (r *BaseRepository) BeginTransaction(ctx context.Context) *gorm.DB {
	db := r.GetDB(ctx)
	return db.Begin()
}

func (r *BaseRepository) BeginTransactionWithContext(ctx context.Context) context.Context {
	db := r.BeginTransaction(ctx)
	return SetDBWithContext(ctx, db)
}

func (r *BaseRepository) Commit(ctx context.Context) *gorm.DB {
	db := r.GetDB(ctx)
	return db.Commit()
}

func (r *BaseRepository) Rollback(ctx context.Context) *gorm.DB {
	db := r.GetDB(ctx)
	return db.Rollback()
}

type IBaseRepository interface {
	BeginTransaction(ctx context.Context) *gorm.DB
	Commit(ctx context.Context) *gorm.DB
	Rollback(ctx context.Context) *gorm.DB
	BeginTransactionWithContext(ctx context.Context) context.Context
}
