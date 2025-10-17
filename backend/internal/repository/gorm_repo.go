package repository

import (
	"context"

	"gorm.io/gorm"
)

type GormRepository struct {
	db *gorm.DB
}

func NewGormRepository(db *gorm.DB) *GormRepository {
	return &GormRepository{db: db}
}

func (r *GormRepository) Create(ctx context.Context, entity interface{}) error {
	return r.db.WithContext(ctx).Create(entity).Error
}

func (r *GormRepository) Update(ctx context.Context, entity interface{}) error {
	return r.db.WithContext(ctx).Save(entity).Error
}

func (r *GormRepository) Delete(ctx context.Context, entity interface{}) error {
	return r.db.WithContext(ctx).Delete(entity).Error
}

func (r *GormRepository) FindByID(ctx context.Context, id uint, entity interface{}) error {
	return r.db.WithContext(ctx).First(entity, id).Error
}
