package repository

import (
	"context"

	"github.com/AntVerkh/test-management-system/internal/domain"
	"github.com/google/uuid"
)

// Add these repository interfaces
type TestCaseRepository interface {
	Create(ctx context.Context, testCase *domain.TestCase) error
	GetByID(ctx context.Context, id uuid.UUID) (*domain.TestCase, error)
	Update(ctx context.Context, testCase *domain.TestCase) error
	List(ctx context.Context, projectID uuid.UUID, page, size int) ([]domain.TestCase, int64, error)
}

type ChecklistRepository interface {
	Create(ctx context.Context, checklist *domain.Checklist) error
	GetByID(ctx context.Context, id uuid.UUID) (*domain.Checklist, error)
	Update(ctx context.Context, checklist *domain.Checklist) error
	List(ctx context.Context, projectID uuid.UUID, page, size int) ([]domain.Checklist, int64, error)
}

type TestStrategyRepository interface {
	Create(ctx context.Context, strategy *domain.TestStrategy) error
	GetByID(ctx context.Context, id uuid.UUID) (*domain.TestStrategy, error)
	Update(ctx context.Context, strategy *domain.TestStrategy) error
	List(ctx context.Context, projectID uuid.UUID, page, size int) ([]domain.TestStrategy, int64, error)
}

type TestRunRepository interface {
	Create(ctx context.Context, testRun *domain.TestRun) error
	GetByID(ctx context.Context, id uuid.UUID) (*domain.TestRun, error)
	Update(ctx context.Context, testRun *domain.TestRun) error
	List(ctx context.Context, testPlanID uuid.UUID, page, size int) ([]domain.TestRun, int64, error)
	Complete(ctx context.Context, id uuid.UUID) error
}
