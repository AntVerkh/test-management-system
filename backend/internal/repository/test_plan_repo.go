package repository

import (
	"context"
	"github.com/AntVerkh/test-management-system/internal/domain"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type TestPlanRepository interface {
	Create(ctx context.Context, plan *domain.TestPlan) error
	GetByID(ctx context.Context, id uuid.UUID) (*domain.TestPlan, error)
	Update(ctx context.Context, plan *domain.TestPlan) error
	List(ctx context.Context, projectID uuid.UUID, page, size int) ([]domain.TestPlan, int64, error)
	AddTestCase(ctx context.Context, planID, testCaseID uuid.UUID) error
	AddChecklist(ctx context.Context, planID, checklistID uuid.UUID) error
}

type testPlanRepository struct {
	db *gorm.DB
}

func NewTestPlanRepository(db *gorm.DB) TestPlanRepository {
	return &testPlanRepository{db: db}
}

func (r *testPlanRepository) Create(ctx context.Context, plan *domain.TestPlan) error {
	return r.db.WithContext(ctx).Create(plan).Error
}

func (r *testPlanRepository) GetByID(ctx context.Context, id uuid.UUID) (*domain.TestPlan, error) {
	var plan domain.TestPlan
	err := r.db.WithContext(ctx).
		Preload("Checklists").
		Preload("TestCases").
		Preload("TestCases.Steps").
		First(&plan, "id = ?", id).Error
	return &plan, err
}

func (r *testPlanRepository) Update(ctx context.Context, plan *domain.TestPlan) error {
	return r.db.WithContext(ctx).Save(plan).Error
}

func (r *testPlanRepository) List(ctx context.Context, projectID uuid.UUID, page, size int) ([]domain.TestPlan, int64, error) {
	var plans []domain.TestPlan
	var total int64

	offset := (page - 1) * size

	query := r.db.WithContext(ctx).Where("project_id = ?", projectID)

	err := query.Model(&domain.TestPlan{}).Count(&total).Error
	if err != nil {
		return nil, 0, err
	}

	err = query.Offset(offset).Limit(size).Order("created_at DESC").Find(&plans).Error
	return plans, total, err
}

func (r *testPlanRepository) AddTestCase(ctx context.Context, planID, testCaseID uuid.UUID) error {
	return r.db.WithContext(ctx).Exec(
		"INSERT INTO test_plan_cases (test_plan_id, test_case_id) VALUES (?, ?)",
		planID, testCaseID,
	).Error
}

func (r *testPlanRepository) AddChecklist(ctx context.Context, planID, checklistID uuid.UUID) error {
	return r.db.WithContext(ctx).Exec(
		"INSERT INTO test_plan_checklists (test_plan_id, checklist_id) VALUES (?, ?)",
		planID, checklistID,
	).Error
}
