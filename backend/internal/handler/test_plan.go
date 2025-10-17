package handler

import (
	"net/http"
	"strconv"
	"time"

	"github.com/AntVerkh/test-management-system/internal/domain"
	"github.com/AntVerkh/test-management-system/internal/service"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type TestPlanHandler struct {
	testPlanService service.TestPlanService
}

func NewTestPlanHandler(testPlanService service.TestPlanService) *TestPlanHandler {
	return &TestPlanHandler{testPlanService: testPlanService}
}

type CreateTestPlanRequest struct {
	ProjectID   uuid.UUID `json:"project_id" binding:"required"`
	Name        string    `json:"name" binding:"required"`
	Description string    `json:"description"`
	Deadline    string    `json:"deadline"`
}

func (h *TestPlanHandler) CreateTestPlan(c *gin.Context) {
	var req CreateTestPlanRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	userID, exists := c.Get("userID")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "unauthorized"})
		return
	}

	var deadline time.Time
	if req.Deadline != "" {
		var err error
		deadline, err = time.Parse(time.RFC3339, req.Deadline)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "invalid deadline format"})
			return
		}
	}

	plan := &domain.TestPlan{
		ProjectID:   req.ProjectID,
		Name:        req.Name,
		Description: req.Description,
		Deadline:    deadline,
		CreatedBy:   userID.(uuid.UUID),
	}

	if err := h.testPlanService.CreateTestPlan(c.Request.Context(), plan); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, plan)
}

func (h *TestPlanHandler) GetTestPlan(c *gin.Context) {
	id, err := uuid.Parse(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid test plan ID"})
		return
	}

	plan, err := h.testPlanService.GetTestPlan(c.Request.Context(), id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "test plan not found"})
		return
	}

	c.JSON(http.StatusOK, plan)
}

func (h *TestPlanHandler) ListTestPlans(c *gin.Context) {
	projectID, err := uuid.Parse(c.Query("project_id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid project ID"})
		return
	}

	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	size, _ := strconv.Atoi(c.DefaultQuery("size", "20"))

	plans, total, err := h.testPlanService.ListTestPlans(c.Request.Context(), projectID, page, size)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data":  plans,
		"total": total,
		"page":  page,
		"size":  size,
	})
}

type UpdateTestPlanRequest struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	Deadline    string `json:"deadline"`
	Status      string `json:"status"`
}

func (h *TestPlanHandler) UpdateTestPlan(c *gin.Context) {
	id, err := uuid.Parse(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid test plan ID"})
		return
	}

	var req UpdateTestPlanRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	plan, err := h.testPlanService.GetTestPlan(c.Request.Context(), id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "test plan not found"})
		return
	}

	if req.Name != "" {
		plan.Name = req.Name
	}
	if req.Description != "" {
		plan.Description = req.Description
	}
	if req.Deadline != "" {
		deadline, err := time.Parse(time.RFC3339, req.Deadline)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "invalid deadline format"})
			return
		}
		plan.Deadline = deadline
	}
	if req.Status != "" {
		plan.Status = req.Status
	}

	if err := h.testPlanService.UpdateTestPlan(c.Request.Context(), plan); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, plan)
}

type AddTestCaseRequest struct {
	TestCaseID uuid.UUID `json:"test_case_id" binding:"required"`
}

func (h *TestPlanHandler) AddTestCase(c *gin.Context) {
	planID, err := uuid.Parse(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid test plan ID"})
		return
	}

	var req AddTestCaseRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := h.testPlanService.AddTestCaseToPlan(c.Request.Context(), planID, req.TestCaseID); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Test case added to plan successfully"})
}
