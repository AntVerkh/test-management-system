package repository

import (
	"context"

	"github.com/AntVerkh/test-management-system/internal/domain"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type testCaseRepository struct {
	db *gorm.DB
}

func NewTestCaseRepository(db *gorm.DB) TestCaseRepository {
	return &testCaseRepository{db: db}
}

func (r *testCaseRepository) Create(ctx context.Context, testCase *domain.TestCase) error {
	return r.db.WithContext(ctx).Create(testCase).Error
}

func (r *testCaseRepository) GetByID(ctx context.Context, id uuid.UUID) (*domain.TestCase, error) {
	var testCase domain.TestCase
	err := r.db.WithContext(ctx).
		Preload("Steps").
		Preload("Attachments").
		First(&testCase, "id = ?", id).Error
	return &testCase, err
}

func (r *testCaseRepository) Update(ctx context.Context, testCase *domain.TestCase) error {
	return r.db.WithContext(ctx).Save(testCase).Error
}

func (r *testCaseRepository) List(ctx context.Context, projectID uuid.UUID, page, size int) ([]domain.TestCase, int64, error) {
	var testCases []domain.TestCase
	var total int64

	offset := (page - 1) * size

	query := r.db.WithContext(ctx).Where("project_id = ?", projectID)

	err := query.Model(&domain.TestCase{}).Count(&total).Error
	if err != nil {
		return nil, 0, err
	}

	err = query.Offset(offset).Limit(size).Order("created_at DESC").Find(&testCases).Error
	return testCases, total, err
}
