package repository

import (
	"context"
	"github.com/AntVerkh/test-management-system/internal/domain"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type checklistRepository struct {
	db *gorm.DB
}

func NewChecklistRepository(db *gorm.DB) ChecklistRepository {
	return &checklistRepository{db: db}
}

func (r *checklistRepository) Create(ctx context.Context, checklist *domain.Checklist) error {
	return r.db.WithContext(ctx).Create(checklist).Error
}

func (r *checklistRepository) GetByID(ctx context.Context, id uuid.UUID) (*domain.Checklist, error) {
	var checklist domain.Checklist
	err := r.db.WithContext(ctx).
		Preload("Items").
		First(&checklist, "id = ?", id).Error
	return &checklist, err
}

func (r *checklistRepository) Update(ctx context.Context, checklist *domain.Checklist) error {
	return r.db.WithContext(ctx).Save(checklist).Error
}

func (r *checklistRepository) List(ctx context.Context, projectID uuid.UUID, page, size int) ([]domain.Checklist, int64, error) {
	var checklists []domain.Checklist
	var total int64

	offset := (page - 1) * size

	query := r.db.WithContext(ctx).Where("project_id = ?", projectID)

	err := query.Model(&domain.Checklist{}).Count(&total).Error
	if err != nil {
		return nil, 0, err
	}

	err = query.Offset(offset).Limit(size).Order("created_at DESC").Find(&checklists).Error
	return checklists, total, err
}
