package service

import (
	"context"

	"github.com/AntVerkh/test-management-system/internal/domain"
	"github.com/google/uuid"
)

// AuthService interface
type AuthService interface {
	Login(ctx context.Context, email, password string) (string, *domain.User, error)
	Register(ctx context.Context, user *domain.User) error
	ValidateToken(ctx context.Context, token string) (*domain.User, error)
}

// UserService interface
type UserService interface {
	GetUserByID(ctx context.Context, id uuid.UUID) (*domain.User, error)
	ListUsers(ctx context.Context) ([]domain.User, error)
	UpdateUserRole(ctx context.Context, userID uuid.UUID, role domain.UserRole) error
}

// TestPlanService interface
type TestPlanService interface {
	CreateTestPlan(ctx context.Context, plan *domain.TestPlan) error
	GetTestPlan(ctx context.Context, id uuid.UUID) (*domain.TestPlan, error)
	UpdateTestPlan(ctx context.Context, plan *domain.TestPlan) error
	ListTestPlans(ctx context.Context, projectID uuid.UUID, page, size int) ([]domain.TestPlan, int64, error)
	AddTestCaseToPlan(ctx context.Context, planID, testCaseID uuid.UUID) error
	AddChecklistToPlan(ctx context.Context, planID, checklistID uuid.UUID) error
}

// TestCaseService interface
type TestCaseService interface {
	CreateTestCase(ctx context.Context, testCase *domain.TestCase) error
	GetTestCase(ctx context.Context, id uuid.UUID) (*domain.TestCase, error)
	UpdateTestCase(ctx context.Context, testCase *domain.TestCase) error
	ListTestCases(ctx context.Context, projectID uuid.UUID, page, size int) ([]domain.TestCase, int64, error)
}

// TestRunService interface
type TestRunService interface {
	StartTestRun(ctx context.Context, testRun *domain.TestRun) error
	RecordTestResult(ctx context.Context, result *domain.TestResult) error
	GetTestRun(ctx context.Context, id uuid.UUID) (*domain.TestRun, error)
	CompleteTestRun(ctx context.Context, id uuid.UUID) error
}

// JWTService interface
type JWTService interface {
	GenerateToken(user *domain.User) (string, error)
	ValidateToken(token string) (*domain.User, error)
}

// ExportService interface
type ExportService interface {
	ExportEntity(ctx context.Context, req *domain.ExportRequest) (string, string, error)
	ExportTestPlan(ctx context.Context, planID uuid.UUID, format domain.ExportFormat, includeHistory, includeComments bool) (string, string, error)
	ExportTestCase(ctx context.Context, testCaseID uuid.UUID, format domain.ExportFormat, includeHistory, includeComments bool) (string, string, error)
	ExportChecklist(ctx context.Context, checklistID uuid.UUID, format domain.ExportFormat, includeHistory, includeComments bool) (string, string, error)
	ExportTestStrategy(ctx context.Context, strategyID uuid.UUID, format domain.ExportFormat, includeHistory, includeComments bool) (string, string, error)
	ExportTestRun(ctx context.Context, testRunID uuid.UUID, format domain.ExportFormat, includeHistory, includeComments bool) (string, string, error)
}
