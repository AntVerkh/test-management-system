package domain

import (
	"fmt"
	"strings"
)

type ExportFormat string

const (
	ExportFormatMarkdown ExportFormat = "markdown"
	ExportFormatHTML     ExportFormat = "html"
	ExportFormatPDF      ExportFormat = "pdf"
)

type ExportRequest struct {
	EntityType      string       `json:"entity_type" binding:"required"`
	EntityID        string       `json:"entity_id" binding:"required"`
	Format          ExportFormat `json:"format" binding:"required"`
	IncludeHistory  bool         `json:"include_history"`
	IncludeComments bool         `json:"include_comments"`
}

type Exporter interface {
	ExportTestPlan(plan *TestPlan, includeHistory, includeComments bool) (string, error)
	ExportTestCase(testCase *TestCase, includeHistory, includeComments bool) (string, error)
	ExportChecklist(checklist *Checklist, includeHistory, includeComments bool) (string, error)
	ExportTestStrategy(strategy *TestStrategy, includeHistory, includeComments bool) (string, error)
	ExportTestRun(testRun *TestRun, includeHistory, includeComments bool) (string, error)
}

type MarkdownExporter struct{}

func NewMarkdownExporter() *MarkdownExporter {
	return &MarkdownExporter{}
}

func (e *MarkdownExporter) ExportTestPlan(plan *TestPlan, includeHistory, includeComments bool) (string, error) {
	var sb strings.Builder

	// Header
	sb.WriteString(fmt.Sprintf("# Test Plan: %s\n\n", plan.Name))
	sb.WriteString(fmt.Sprintf("**ID:** %s\n", plan.ID))
	sb.WriteString(fmt.Sprintf("**Project ID:** %s\n", plan.ProjectID))
	sb.WriteString(fmt.Sprintf("**Status:** %s\n", plan.Status))
	if !plan.Deadline.IsZero() {
		sb.WriteString(fmt.Sprintf("**Deadline:** %s\n", plan.Deadline.Format("2006-01-02 15:04")))
	}
	sb.WriteString(fmt.Sprintf("**Created:** %s\n", plan.CreatedAt.Format("2006-01-02 15:04")))
	sb.WriteString(fmt.Sprintf("**Last Updated:** %s\n\n", plan.UpdatedAt.Format("2006-01-02 15:04")))

	// Description
	if plan.Description != "" {
		sb.WriteString("## Description\n\n")
		sb.WriteString(plan.Description + "\n\n")
	}

	// Test Cases
	if len(plan.TestCases) > 0 {
		sb.WriteString("## Test Cases\n\n")
		for i, testCase := range plan.TestCases {
			sb.WriteString(fmt.Sprintf("### %d. %s\n", i+1, testCase.Title))
			sb.WriteString(fmt.Sprintf("**ID:** %s\n", testCase.ID))
			if testCase.Description != "" {
				sb.WriteString("**Description:** " + testCase.Description + "\n")
			}
			if testCase.PreSteps != "" {
				sb.WriteString("**Pre-Steps:**\n" + testCase.PreSteps + "\n")
			}
			if testCase.ExpectedResult != "" {
				sb.WriteString("**Expected Result:** " + testCase.ExpectedResult + "\n")
			}

			if len(testCase.Steps) > 0 {
				sb.WriteString("**Test Steps:**\n")
				for j, step := range testCase.Steps {
					sb.WriteString(fmt.Sprintf("%d. %s\n", j+1, step.Description))
					if step.ExpectedResult != "" {
						sb.WriteString(fmt.Sprintf("   *Expected:* %s\n", step.ExpectedResult))
					}
				}
			}
			sb.WriteString("\n")
		}
	}

	// Checklists
	if len(plan.Checklists) > 0 {
		sb.WriteString("## Checklists\n\n")
		for i, checklist := range plan.Checklists {
			sb.WriteString(fmt.Sprintf("### %d. %s\n", i+1, checklist.Name))
			if checklist.Description != "" {
				sb.WriteString("**Description:** " + checklist.Description + "\n")
			}

			if len(checklist.Items) > 0 {
				sb.WriteString("**Checklist Items:**\n")
				for _, item := range checklist.Items {
					sb.WriteString(fmt.Sprintf("- [ ] %s\n", item.Description))
					if item.ExpectedResult != "" {
						sb.WriteString(fmt.Sprintf("  *Expected:* %s\n", item.ExpectedResult))
					}
				}
			}
			sb.WriteString("\n")
		}
	}

	return sb.String(), nil
}

func (e *MarkdownExporter) ExportTestCase(testCase *TestCase, includeHistory, includeComments bool) (string, error) {
	var sb strings.Builder

	// Header
	sb.WriteString(fmt.Sprintf("# Test Case: %s\n\n", testCase.Title))
	sb.WriteString(fmt.Sprintf("**ID:** %s\n", testCase.ID))
	sb.WriteString(fmt.Sprintf("**Project ID:** %s\n", testCase.ProjectID))
	sb.WriteString(fmt.Sprintf("**Created:** %s\n", testCase.CreatedAt.Format("2006-01-02 15:04")))
	sb.WriteString(fmt.Sprintf("**Last Updated:** %s\n\n", testCase.UpdatedAt.Format("2006-01-02 15:04")))

	// Description
	if testCase.Description != "" {
		sb.WriteString("## Description\n\n")
		sb.WriteString(testCase.Description + "\n\n")
	}

	// Pre-Steps
	if testCase.PreSteps != "" {
		sb.WriteString("## Pre-Steps\n\n")
		sb.WriteString(testCase.PreSteps + "\n\n")
	}

	// Test Steps
	if len(testCase.Steps) > 0 {
		sb.WriteString("## Test Steps\n\n")
		for i, step := range testCase.Steps {
			sb.WriteString(fmt.Sprintf("### Step %d\n", i+1))
			sb.WriteString(fmt.Sprintf("**Action:** %s\n", step.Description))
			if step.ExpectedResult != "" {
				sb.WriteString(fmt.Sprintf("**Expected Result:** %s\n", step.ExpectedResult))
			}
			sb.WriteString("\n")
		}
	}

	// Expected Result
	if testCase.ExpectedResult != "" {
		sb.WriteString("## Expected Result\n\n")
		sb.WriteString(testCase.ExpectedResult + "\n\n")
	}

	// Attachments
	if len(testCase.Attachments) > 0 {
		sb.WriteString("## Attachments\n\n")
		for _, attachment := range testCase.Attachments {
			sb.WriteString(fmt.Sprintf("- **%s** (%s, %d bytes)\n",
				attachment.FileName,
				attachment.MimeType,
				attachment.FileSize))
		}
		sb.WriteString("\n")
	}

	return sb.String(), nil
}

func (e *MarkdownExporter) ExportChecklist(checklist *Checklist, includeHistory, includeComments bool) (string, error) {
	var sb strings.Builder

	// Header
	sb.WriteString(fmt.Sprintf("# Checklist: %s\n\n", checklist.Name))
	sb.WriteString(fmt.Sprintf("**ID:** %s\n", checklist.ID))
	sb.WriteString(fmt.Sprintf("**Project ID:** %s\n", checklist.ProjectID))
	sb.WriteString(fmt.Sprintf("**Created:** %s\n", checklist.CreatedAt.Format("2006-01-02 15:04")))
	sb.WriteString(fmt.Sprintf("**Last Updated:** %s\n\n", checklist.UpdatedAt.Format("2006-01-02 15:04")))

	// Description
	if checklist.Description != "" {
		sb.WriteString("## Description\n\n")
		sb.WriteString(checklist.Description + "\n\n")
	}

	// Checklist Items
	if len(checklist.Items) > 0 {
		sb.WriteString("## Checklist Items\n\n")
		for i, item := range checklist.Items {
			sb.WriteString(fmt.Sprintf("%d. [ ] %s\n", i+1, item.Description))
			if item.ExpectedResult != "" {
				sb.WriteString(fmt.Sprintf("   *Expected:* %s\n", item.ExpectedResult))
			}
			sb.WriteString("\n")
		}
	}

	return sb.String(), nil
}

func (e *MarkdownExporter) ExportTestStrategy(strategy *TestStrategy, includeHistory, includeComments bool) (string, error) {
	var sb strings.Builder

	// Header
	sb.WriteString(fmt.Sprintf("# Test Strategy: %s\n\n", strategy.Name))
	sb.WriteString(fmt.Sprintf("**ID:** %s\n", strategy.ID))
	sb.WriteString(fmt.Sprintf("**Project ID:** %s\n", strategy.ProjectID))
	sb.WriteString(fmt.Sprintf("**Created:** %s\n", strategy.CreatedAt.Format("2006-01-02 15:04")))
	sb.WriteString(fmt.Sprintf("**Last Updated:** %s\n\n", strategy.UpdatedAt.Format("2006-01-02 15:04")))

	// Description
	if strategy.Description != "" {
		sb.WriteString("## Description\n\n")
		sb.WriteString(strategy.Description + "\n\n")
	}

	// Content
	if strategy.Content != "" {
		sb.WriteString("## Strategy Content\n\n")
		sb.WriteString(strategy.Content + "\n\n")
	}

	return sb.String(), nil
}

func (e *MarkdownExporter) ExportTestRun(testRun *TestRun, includeHistory, includeComments bool) (string, error) {
	var sb strings.Builder

	// Header
	sb.WriteString(fmt.Sprintf("# Test Run: %s\n\n", testRun.Name))
	sb.WriteString(fmt.Sprintf("**ID:** %s\n", testRun.ID))
	sb.WriteString(fmt.Sprintf("**Test Plan ID:** %s\n", testRun.TestPlanID))
	sb.WriteString(fmt.Sprintf("**Started:** %s\n", testRun.StartedAt.Format("2006-01-02 15:04")))
	if testRun.CompletedAt != nil {
		sb.WriteString(fmt.Sprintf("**Completed:** %s\n", testRun.CompletedAt.Format("2006-01-02 15:04")))
	}
	sb.WriteString("\n")

	// Test Results Summary
	if len(testRun.Results) > 0 {
		passed := 0
		failed := 0
		blocked := 0
		skipped := 0

		for _, result := range testRun.Results {
			switch result.Status {
			case "pass":
				passed++
			case "fail":
				failed++
			case "blocked":
				blocked++
			case "skipped":
				skipped++
			}
		}

		total := passed + failed + blocked + skipped
		sb.WriteString("## Test Results Summary\n\n")
		sb.WriteString(fmt.Sprintf("- **Total:** %d\n", total))
		sb.WriteString(fmt.Sprintf("- **✅ Passed:** %d\n", passed))
		sb.WriteString(fmt.Sprintf("- **❌ Failed:** %d\n", failed))
		sb.WriteString(fmt.Sprintf("- **🚫 Blocked:** %d\n", blocked))
		sb.WriteString(fmt.Sprintf("- **⏭️ Skipped:** %d\n", skipped))

		if total > 0 {
			passRate := float64(passed) / float64(total) * 100
			sb.WriteString(fmt.Sprintf("- **📊 Pass Rate:** %.1f%%\n", passRate))
		}
		sb.WriteString("\n")

		// Detailed Results
		sb.WriteString("## Detailed Results\n\n")
		for i, result := range testRun.Results {
			statusIcon := "❓"
			switch result.Status {
			case "pass":
				statusIcon = "✅"
			case "fail":
				statusIcon = "❌"
			case "blocked":
				statusIcon = "🚫"
			case "skipped":
				statusIcon = "⏭️"
			}

			entityName := e.getEntityName(&result)
			sb.WriteString(fmt.Sprintf("### %d. %s %s\n", i+1, statusIcon, entityName))
			sb.WriteString(fmt.Sprintf("**Status:** %s\n", result.Status))
			sb.WriteString(fmt.Sprintf("**Executed By:** %s\n", result.ExecutedBy))
			sb.WriteString(fmt.Sprintf("**Executed At:** %s\n", result.ExecutedAt.Format("2006-01-02 15:04")))

			if result.Comments != "" {
				sb.WriteString("**Comments:** " + result.Comments + "\n")
			}
			sb.WriteString("\n")
		}
	}

	return sb.String(), nil
}

func (e *MarkdownExporter) getEntityName(result *TestResult) string {
	if result.TestCaseID != nil {
		return "Test Case"
	} else if result.ChecklistItemID != nil {
		return "Checklist Item"
	}
	return "Unknown Entity"
}
