package repository

import (
	"context"
	"time"

	"github.com/AntVerkh/test-management-system/internal/domain"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type testRunRepository struct {
	db *gorm.DB
}

func NewTestRunRepository(db *gorm.DB) TestRunRepository {
	return &testRunRepository{db: db}
}

func (r *testRunRepository) Create(ctx context.Context, testRun *domain.TestRun) error {
	return r.db.WithContext(ctx).Create(testRun).Error
}

func (r *testRunRepository) GetByID(ctx context.Context, id uuid.UUID) (*domain.TestRun, error) {
	var testRun domain.TestRun
	err := r.db.WithContext(ctx).
		Preload("Results").
		Preload("Results.TestCase").
		Preload("Results.ChecklistItem").
		First(&testRun, "id = ?", id).Error
	return &testRun, err
}

func (r *testRunRepository) Update(ctx context.Context, testRun *domain.TestRun) error {
	return r.db.WithContext(ctx).Save(testRun).Error
}

func (r *testRunRepository) List(ctx context.Context, testPlanID uuid.UUID, page, size int) ([]domain.TestRun, int64, error) {
	var testRuns []domain.TestRun
	var total int64

	offset := (page - 1) * size

	query := r.db.WithContext(ctx).Where("test_plan_id = ?", testPlanID)

	err := query.Model(&domain.TestRun{}).Count(&total).Error
	if err != nil {
		return nil, 0, err
	}

	err = query.Offset(offset).Limit(size).Order("started_at DESC").Find(&testRuns).Error
	return testRuns, total, err
}

func (r *testRunRepository) Complete(ctx context.Context, id uuid.UUID) error {
	completedAt := time.Now()
	return r.db.WithContext(ctx).Model(&domain.TestRun{}).
		Where("id = ?", id).
		Update("completed_at", completedAt).Error
}
