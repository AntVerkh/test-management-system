package handler

import (
	"net/http"
	"strconv"

	"github.com/AntVerkh/test-management-system/internal/domain"
	"github.com/AntVerkh/test-management-system/internal/service"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type TestCaseHandler struct {
	testCaseService service.TestCaseService
}

func NewTestCaseHandler(testCaseService service.TestCaseService) *TestCaseHandler {
	return &TestCaseHandler{testCaseService: testCaseService}
}

type CreateTestCaseRequest struct {
	ProjectID      uuid.UUID         `json:"project_id" binding:"required"`
	Title          string            `json:"title" binding:"required"`
	Description    string            `json:"description"`
	PreSteps       string            `json:"pre_steps"`
	Steps          []TestStepRequest `json:"steps"`
	ExpectedResult string            `json:"expected_result"`
}

type TestStepRequest struct {
	Description    string `json:"description" binding:"required"`
	ExpectedResult string `json:"expected_result"`
	Order          int    `json:"order"`
}

func (h *TestCaseHandler) CreateTestCase(c *gin.Context) {
	var req CreateTestCaseRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	userID, exists := c.Get("userID")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "unauthorized"})
		return
	}

	testCase := &domain.TestCase{
		ProjectID:      req.ProjectID,
		Title:          req.Title,
		Description:    req.Description,
		PreSteps:       req.PreSteps,
		ExpectedResult: req.ExpectedResult,
		CreatedBy:      userID.(uuid.UUID),
	}

	// Convert steps
	for _, stepReq := range req.Steps {
		testCase.Steps = append(testCase.Steps, domain.TestStep{
			Description:    stepReq.Description,
			ExpectedResult: stepReq.ExpectedResult,
			Order:          stepReq.Order,
		})
	}

	if err := h.testCaseService.CreateTestCase(c.Request.Context(), testCase); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, testCase)
}

func (h *TestCaseHandler) GetTestCase(c *gin.Context) {
	id, err := uuid.Parse(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid test case ID"})
		return
	}

	testCase, err := h.testCaseService.GetTestCase(c.Request.Context(), id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "test case not found"})
		return
	}

	c.JSON(http.StatusOK, testCase)
}

func (h *TestCaseHandler) ListTestCases(c *gin.Context) {
	projectID, err := uuid.Parse(c.Query("project_id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid project ID"})
		return
	}

	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	size, _ := strconv.Atoi(c.DefaultQuery("size", "20"))

	testCases, total, err := h.testCaseService.ListTestCases(c.Request.Context(), projectID, page, size)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data":  testCases,
		"total": total,
		"page":  page,
		"size":  size,
	})
}

type UpdateTestCaseRequest struct {
	Title          string            `json:"title"`
	Description    string            `json:"description"`
	PreSteps       string            `json:"pre_steps"`
	Steps          []TestStepRequest `json:"steps"`
	ExpectedResult string            `json:"expected_result"`
}

func (h *TestCaseHandler) UpdateTestCase(c *gin.Context) {
	id, err := uuid.Parse(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid test case ID"})
		return
	}

	var req UpdateTestCaseRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Get existing test case
	testCase, err := h.testCaseService.GetTestCase(c.Request.Context(), id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "test case not found"})
		return
	}

	// Update fields
	if req.Title != "" {
		testCase.Title = req.Title
	}
	if req.Description != "" {
		testCase.Description = req.Description
	}
	if req.PreSteps != "" {
		testCase.PreSteps = req.PreSteps
	}
	if req.ExpectedResult != "" {
		testCase.ExpectedResult = req.ExpectedResult
	}

	// Update steps if provided
	if len(req.Steps) > 0 {
		testCase.Steps = nil
		for _, stepReq := range req.Steps {
			testCase.Steps = append(testCase.Steps, domain.TestStep{
				Description:    stepReq.Description,
				ExpectedResult: stepReq.ExpectedResult,
				Order:          stepReq.Order,
			})
		}
	}

	if err := h.testCaseService.UpdateTestCase(c.Request.Context(), testCase); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, testCase)
}
