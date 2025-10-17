package repository

import (
	"context"
	"github.com/AntVerkh/test-management-system/internal/domain"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type testStrategyRepository struct {
	db *gorm.DB
}

func NewTestStrategyRepository(db *gorm.DB) TestStrategyRepository {
	return &testStrategyRepository{db: db}
}

func (r *testStrategyRepository) Create(ctx context.Context, strategy *domain.TestStrategy) error {
	return r.db.WithContext(ctx).Create(strategy).Error
}

func (r *testStrategyRepository) GetByID(ctx context.Context, id uuid.UUID) (*domain.TestStrategy, error) {
	var strategy domain.TestStrategy
	err := r.db.WithContext(ctx).First(&strategy, "id = ?", id).Error
	return &strategy, err
}

func (r *testStrategyRepository) Update(ctx context.Context, strategy *domain.TestStrategy) error {
	return r.db.WithContext(ctx).Save(strategy).Error
}

func (r *testStrategyRepository) List(ctx context.Context, projectID uuid.UUID, page, size int) ([]domain.TestStrategy, int64, error) {
	var strategies []domain.TestStrategy
	var total int64

	offset := (page - 1) * size

	query := r.db.WithContext(ctx).Where("project_id = ?", projectID)

	err := query.Model(&domain.TestStrategy{}).Count(&total).Error
	if err != nil {
		return nil, 0, err
	}

	err = query.Offset(offset).Limit(size).Order("created_at DESC").Find(&strategies).Error
	return strategies, total, err
}
