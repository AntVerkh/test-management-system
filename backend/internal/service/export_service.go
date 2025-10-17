package service

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/AntVerkh/test-management-system/internal/domain"
	"github.com/AntVerkh/test-management-system/internal/repository"
	"github.com/google/uuid"
)

type exportService struct {
	testPlanRepo     repository.TestPlanRepository
	testCaseRepo     repository.TestCaseRepository
	checklistRepo    repository.ChecklistRepository
	testStrategyRepo repository.TestStrategyRepository
	testRunRepo      repository.TestRunRepository
	exporter         domain.Exporter
}

func NewExportService(
	testPlanRepo repository.TestPlanRepository,
	testCaseRepo repository.TestCaseRepository,
	checklistRepo repository.ChecklistRepository,
	testStrategyRepo repository.TestStrategyRepository,
	testRunRepo repository.TestRunRepository,
	exporter domain.Exporter,
) ExportService {
	return &exportService{
		testPlanRepo:     testPlanRepo,
		testCaseRepo:     testCaseRepo,
		checklistRepo:    checklistRepo,
		testStrategyRepo: testStrategyRepo,
		testRunRepo:      testRunRepo,
		exporter:         exporter,
	}
}

func (s *exportService) ExportEntity(ctx context.Context, req *domain.ExportRequest) (string, string, error) {
	entityID, err := uuid.Parse(req.EntityID)
	if err != nil {
		return "", "", errors.New("invalid entity ID")
	}

	switch req.EntityType {
	case "test_plan":
		return s.ExportTestPlan(ctx, entityID, req.Format, req.IncludeHistory, req.IncludeComments)
	case "test_case":
		return s.ExportTestCase(ctx, entityID, req.Format, req.IncludeHistory, req.IncludeComments)
	case "checklist":
		return s.ExportChecklist(ctx, entityID, req.Format, req.IncludeHistory, req.IncludeComments)
	case "test_strategy":
		return s.ExportTestStrategy(ctx, entityID, req.Format, req.IncludeHistory, req.IncludeComments)
	case "test_run":
		return s.ExportTestRun(ctx, entityID, req.Format, req.IncludeHistory, req.IncludeComments)
	default:
		return "", "", errors.New("unsupported entity type")
	}
}

func (s *exportService) ExportTestPlan(ctx context.Context, planID uuid.UUID, format domain.ExportFormat, includeHistory, includeComments bool) (string, string, error) {
	plan, err := s.testPlanRepo.GetByID(ctx, planID)
	if err != nil {
		return "", "", errors.New("test plan not found")
	}

	var content string
	var filename string

	switch format {
	case domain.ExportFormatMarkdown:
		content, err = s.exporter.ExportTestPlan(plan, includeHistory, includeComments)
		filename = fmt.Sprintf("test_plan_%s_%s.md", plan.Name, time.Now().Format("20060102_150405"))
	default:
		return "", "", errors.New("unsupported export format")
	}

	if err != nil {
		return "", "", err
	}

	return content, filename, nil
}

func (s *exportService) ExportTestCase(ctx context.Context, testCaseID uuid.UUID, format domain.ExportFormat, includeHistory, includeComments bool) (string, string, error) {
	testCase, err := s.testCaseRepo.GetByID(ctx, testCaseID)
	if err != nil {
		return "", "", errors.New("test case not found")
	}

	var content string
	var filename string

	switch format {
	case domain.ExportFormatMarkdown:
		content, err = s.exporter.ExportTestCase(testCase, includeHistory, includeComments)
		filename = fmt.Sprintf("test_case_%s_%s.md", testCase.Title, time.Now().Format("20060102_150405"))
	default:
		return "", "", errors.New("unsupported export format")
	}

	if err != nil {
		return "", "", err
	}

	return content, filename, nil
}

func (s *exportService) ExportChecklist(ctx context.Context, checklistID uuid.UUID, format domain.ExportFormat, includeHistory, includeComments bool) (string, string, error) {
	checklist, err := s.checklistRepo.GetByID(ctx, checklistID)
	if err != nil {
		return "", "", errors.New("checklist not found")
	}

	var content string
	var filename string

	switch format {
	case domain.ExportFormatMarkdown:
		content, err = s.exporter.ExportChecklist(checklist, includeHistory, includeComments)
		filename = fmt.Sprintf("checklist_%s_%s.md", checklist.Name, time.Now().Format("20060102_150405"))
	default:
		return "", "", errors.New("unsupported export format")
	}

	if err != nil {
		return "", "", err
	}

	return content, filename, nil
}

func (s *exportService) ExportTestStrategy(ctx context.Context, strategyID uuid.UUID, format domain.ExportFormat, includeHistory, includeComments bool) (string, string, error) {
	strategy, err := s.testStrategyRepo.GetByID(ctx, strategyID)
	if err != nil {
		return "", "", errors.New("test strategy not found")
	}

	var content string
	var filename string

	switch format {
	case domain.ExportFormatMarkdown:
		content, err = s.exporter.ExportTestStrategy(strategy, includeHistory, includeComments)
		filename = fmt.Sprintf("test_strategy_%s_%s.md", strategy.Name, time.Now().Format("20060102_150405"))
	default:
		return "", "", errors.New("unsupported export format")
	}

	if err != nil {
		return "", "", err
	}

	return content, filename, nil
}

func (s *exportService) ExportTestRun(ctx context.Context, testRunID uuid.UUID, format domain.ExportFormat, includeHistory, includeComments bool) (string, string, error) {
	testRun, err := s.testRunRepo.GetByID(ctx, testRunID)
	if err != nil {
		return "", "", errors.New("test run not found")
	}

	var content string
	var filename string

	switch format {
	case domain.ExportFormatMarkdown:
		content, err = s.exporter.ExportTestRun(testRun, includeHistory, includeComments)
		filename = fmt.Sprintf("test_run_%s_%s.md", testRun.Name, time.Now().Format("20060102_150405"))
	default:
		return "", "", errors.New("unsupported export format")
	}

	if err != nil {
		return "", "", err
	}

	return content, filename, nil
}
