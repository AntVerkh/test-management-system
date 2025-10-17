package handler

import (
	"net/http"
	"strconv"

	"github.com/AntVerkh/test-management-system/internal/domain"
	"github.com/AntVerkh/test-management-system/internal/service"
	"github.com/gin-gonic/gin"
)

type ExportHandler struct {
	exportService service.ExportService
}

func NewExportHandler(exportService service.ExportService) *ExportHandler {
	return &ExportHandler{exportService: exportService}
}

type ExportRequest struct {
	EntityType      string              `json:"entity_type" binding:"required"`
	EntityID        string              `json:"entity_id" binding:"required"`
	Format          domain.ExportFormat `json:"format" binding:"required"`
	IncludeHistory  bool                `json:"include_history"`
	IncludeComments bool                `json:"include_comments"`
}

func (h *ExportHandler) Export(c *gin.Context) {
	var req ExportRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	content, filename, err := h.exportService.ExportEntity(c.Request.Context(), &domain.ExportRequest{
		EntityType:      req.EntityType,
		EntityID:        req.EntityID,
		Format:          req.Format,
		IncludeHistory:  req.IncludeHistory,
		IncludeComments: req.IncludeComments,
	})

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Set headers for file download
	c.Header("Content-Type", "text/markdown; charset=utf-8")
	c.Header("Content-Disposition", "attachment; filename="+filename)
	c.Header("Content-Length", strconv.Itoa(len(content)))

	c.String(http.StatusOK, content)
}

// Convenience endpoints for direct export
func (h *ExportHandler) ExportTestPlan(c *gin.Context) {
	h.exportEntity(c, "test_plan")
}

func (h *ExportHandler) ExportTestCase(c *gin.Context) {
	h.exportEntity(c, "test_case")
}

func (h *ExportHandler) ExportChecklist(c *gin.Context) {
	h.exportEntity(c, "checklist")
}

func (h *ExportHandler) ExportTestStrategy(c *gin.Context) {
	h.exportEntity(c, "test_strategy")
}

func (h *ExportHandler) ExportTestRun(c *gin.Context) {
	h.exportEntity(c, "test_run")
}

func (h *ExportHandler) exportEntity(c *gin.Context, entityType string) {
	entityID := c.Param("id")
	format := domain.ExportFormat(c.DefaultQuery("format", "markdown"))
	includeHistory := c.DefaultQuery("include_history", "false") == "true"
	includeComments := c.DefaultQuery("include_comments", "false") == "true"

	content, filename, err := h.exportService.ExportEntity(c.Request.Context(), &domain.ExportRequest{
		EntityType:      entityType,
		EntityID:        entityID,
		Format:          format,
		IncludeHistory:  includeHistory,
		IncludeComments: includeComments,
	})

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Set headers for file download
	c.Header("Content-Type", "text/markdown; charset=utf-8")
	c.Header("Content-Disposition", "attachment; filename="+filename)
	c.Header("Content-Length", strconv.Itoa(len(content)))

	c.String(http.StatusOK, content)
}
