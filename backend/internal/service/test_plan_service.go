package service

import (
	"context"
	"time"

	"github.com/AntVerkh/test-management-system/internal/domain"
	"github.com/AntVerkh/test-management-system/internal/repository"
	"github.com/google/uuid"
)

type testPlanService struct {
	repo repository.TestPlanRepository
}

func NewTestPlanService(repo repository.TestPlanRepository) TestPlanService {
	return &testPlanService{repo: repo}
}

func (s *testPlanService) CreateTestPlan(ctx context.Context, plan *domain.TestPlan) error {
	plan.ID = uuid.New()
	plan.CreatedAt = time.Now()
	plan.UpdatedAt = time.Now()

	return s.repo.Create(ctx, plan)
}

func (s *testPlanService) GetTestPlan(ctx context.Context, id uuid.UUID) (*domain.TestPlan, error) {
	return s.repo.GetByID(ctx, id)
}

func (s *testPlanService) UpdateTestPlan(ctx context.Context, plan *domain.TestPlan) error {
	plan.UpdatedAt = time.Now()
	return s.repo.Update(ctx, plan)
}

func (s *testPlanService) ListTestPlans(ctx context.Context, projectID uuid.UUID, page, size int) ([]domain.TestPlan, int64, error) {
	return s.repo.List(ctx, projectID, page, size)
}

func (s *testPlanService) AddTestCaseToPlan(ctx context.Context, planID, testCaseID uuid.UUID) error {
	return s.repo.AddTestCase(ctx, planID, testCaseID)
}

func (s *testPlanService) AddChecklistToPlan(ctx context.Context, planID, checklistID uuid.UUID) error {
	return s.repo.AddChecklist(ctx, planID, checklistID)
}
