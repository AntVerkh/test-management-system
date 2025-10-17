package service

import (
	"context"
	"time"

	"github.com/AntVerkh/test-management-system/internal/domain"
	"github.com/AntVerkh/test-management-system/internal/repository"
	"github.com/google/uuid"
)

type testCaseService struct {
	repo repository.TestCaseRepository
}

func NewTestCaseService(repo repository.TestCaseRepository) TestCaseService {
	return &testCaseService{repo: repo}
}

func (s *testCaseService) CreateTestCase(ctx context.Context, testCase *domain.TestCase) error {
	testCase.ID = uuid.New()
	testCase.CreatedAt = time.Now()
	testCase.UpdatedAt = time.Now()

	// Generate IDs for steps
	for i := range testCase.Steps {
		testCase.Steps[i].ID = uuid.New()
		testCase.Steps[i].CreatedAt = time.Now()
	}

	return s.repo.Create(ctx, testCase)
}

func (s *testCaseService) GetTestCase(ctx context.Context, id uuid.UUID) (*domain.TestCase, error) {
	return s.repo.GetByID(ctx, id)
}

func (s *testCaseService) UpdateTestCase(ctx context.Context, testCase *domain.TestCase) error {
	testCase.UpdatedAt = time.Now()
	return s.repo.Update(ctx, testCase)
}

func (s *testCaseService) ListTestCases(ctx context.Context, projectID uuid.UUID, page, size int) ([]domain.TestCase, int64, error) {
	return s.repo.List(ctx, projectID, page, size)
}
