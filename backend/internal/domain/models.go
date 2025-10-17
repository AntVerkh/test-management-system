package domain

import (
	"time"

	"github.com/google/uuid"
	_ "gorm.io/gorm"
)

type UserRole string

const (
	RoleAdmin UserRole = "admin"
	RoleUser  UserRole = "user"
	RoleGuest UserRole = "guest"
)

type User struct {
	ID        uuid.UUID `gorm:"type:uuid;primary_key" json:"id"`
	Email     string    `gorm:"uniqueIndex;not null" json:"email"`
	Password  string    `gorm:"not null" json:"-"`
	Role      UserRole  `gorm:"type:varchar(20);not null" json:"role"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type Project struct {
	ID          uuid.UUID `gorm:"type:uuid;primary_key" json:"id"`
	Name        string    `gorm:"not null" json:"name"`
	Description string    `json:"description"`
	CreatedBy   uuid.UUID `gorm:"type:uuid" json:"created_by"`
	CreatedAt   time.Time `json:"created_at"`
}

type TestPlan struct {
	ID          uuid.UUID `gorm:"type:uuid;primary_key" json:"id"`
	ProjectID   uuid.UUID `gorm:"type:uuid;not null" json:"project_id"`
	Name        string    `gorm:"not null" json:"name"`
	Description string    `json:"description"`
	Deadline    time.Time `json:"deadline"`
	Status      string    `gorm:"default:'draft'" json:"status"`
	CreatedBy   uuid.UUID `gorm:"type:uuid" json:"created_by"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`

	Checklists []Checklist `gorm:"many2many:test_plan_checklists;" json:"checklists,omitempty"`
	TestCases  []TestCase  `gorm:"many2many:test_plan_cases;" json:"test_cases,omitempty"`
	History    []History   `gorm:"foreignKey:EntityID" json:"history,omitempty"`
	Comments   []Comment   `gorm:"foreignKey:EntityID" json:"comments,omitempty"`
}

type TestStrategy struct {
	ID          uuid.UUID `gorm:"type:uuid;primary_key" json:"id"`
	ProjectID   uuid.UUID `gorm:"type:uuid;not null" json:"project_id"`
	Name        string    `gorm:"not null" json:"name"`
	Description string    `json:"description"`
	Content     string    `gorm:"type:text" json:"content"`
	CreatedBy   uuid.UUID `gorm:"type:uuid" json:"created_by"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`

	History  []History `gorm:"foreignKey:EntityID" json:"history,omitempty"`
	Comments []Comment `gorm:"foreignKey:EntityID" json:"comments,omitempty"`
}

type Checklist struct {
	ID          uuid.UUID       `gorm:"type:uuid;primary_key" json:"id"`
	ProjectID   uuid.UUID       `gorm:"type:uuid;not null" json:"project_id"`
	Name        string          `gorm:"not null" json:"name"`
	Description string          `json:"description"`
	Items       []ChecklistItem `gorm:"foreignKey:ChecklistID" json:"items"`
	CreatedBy   uuid.UUID       `gorm:"type:uuid" json:"created_by"`
	CreatedAt   time.Time       `json:"created_at"`
	UpdatedAt   time.Time       `json:"updated_at"`

	History  []History `gorm:"foreignKey:EntityID" json:"history,omitempty"`
	Comments []Comment `gorm:"foreignKey:EntityID" json:"comments,omitempty"`
}

type ChecklistItem struct {
	ID             uuid.UUID `gorm:"type:uuid;primary_key" json:"id"`
	ChecklistID    uuid.UUID `gorm:"type:uuid;not null" json:"checklist_id"`
	Description    string    `gorm:"not null" json:"description"`
	ExpectedResult string    `json:"expected_result"`
	Order          int       `json:"order"`
}

type TestCase struct {
	ID             uuid.UUID  `gorm:"type:uuid;primary_key" json:"id"`
	ProjectID      uuid.UUID  `gorm:"type:uuid;not null" json:"project_id"`
	Title          string     `gorm:"not null" json:"title"`
	Description    string     `json:"description"`
	PreSteps       string     `gorm:"type:text" json:"pre_steps"`
	Steps          []TestStep `gorm:"foreignKey:TestCaseID" json:"steps"`
	ExpectedResult string     `gorm:"type:text" json:"expected_result"`
	CreatedBy      uuid.UUID  `gorm:"type:uuid" json:"created_by"`
	CreatedAt      time.Time  `json:"created_at"`
	UpdatedAt      time.Time  `json:"updated_at"`

	Attachments []Attachment `gorm:"foreignKey:TestCaseID" json:"attachments,omitempty"`
	History     []History    `gorm:"foreignKey:EntityID" json:"history,omitempty"`
	Comments    []Comment    `gorm:"foreignKey:EntityID" json:"comments,omitempty"`
}

type TestStep struct {
	ID             uuid.UUID `gorm:"type:uuid;primary_key" json:"id"`
	TestCaseID     uuid.UUID `gorm:"type:uuid;not null" json:"test_case_id"`
	Description    string    `gorm:"not null" json:"description"`
	ExpectedResult string    `json:"expected_result"`
	Order          int       `json:"order"`
	CreatedAt      time.Time `json:"created_at"`
}

type TestRun struct {
	ID          uuid.UUID  `gorm:"type:uuid;primary_key" json:"id"`
	TestPlanID  uuid.UUID  `gorm:"type:uuid;not null" json:"test_plan_id"`
	Name        string     `gorm:"not null" json:"name"`
	StartedBy   uuid.UUID  `gorm:"type:uuid" json:"started_by"`
	StartedAt   time.Time  `json:"started_at"`
	CompletedAt *time.Time `json:"completed_at,omitempty"`

	Results  []TestResult `gorm:"foreignKey:TestRunID" json:"results"`
	History  []History    `gorm:"foreignKey:EntityID" json:"history,omitempty"`
	Comments []Comment    `gorm:"foreignKey:EntityID" json:"comments,omitempty"`
}

type TestResult struct {
	ID              uuid.UUID  `gorm:"type:uuid;primary_key" json:"id"`
	TestRunID       uuid.UUID  `gorm:"type:uuid;not null" json:"test_run_id"`
	TestCaseID      *uuid.UUID `gorm:"type:uuid" json:"test_case_id,omitempty"`
	ChecklistItemID *uuid.UUID `gorm:"type:uuid" json:"checklist_item_id,omitempty"`
	Status          string     `gorm:"not null" json:"status"` // pass, fail, blocked, skipped
	Comments        string     `json:"comments"`
	ExecutedBy      uuid.UUID  `gorm:"type:uuid" json:"executed_by"`
	ExecutedAt      time.Time  `json:"executed_at"`

	TestCase      *TestCase      `gorm:"foreignKey:TestCaseID" json:"test_case,omitempty"`
	ChecklistItem *ChecklistItem `gorm:"foreignKey:ChecklistItemID" json:"checklist_item,omitempty"`
}

type Attachment struct {
	ID         uuid.UUID `gorm:"type:uuid;primary_key" json:"id"`
	TestCaseID uuid.UUID `gorm:"type:uuid;not null" json:"test_case_id"`
	FileName   string    `gorm:"not null" json:"file_name"`
	FilePath   string    `gorm:"not null" json:"file_path"`
	FileSize   int64     `json:"file_size"`
	MimeType   string    `json:"mime_type"`
	UploadedBy uuid.UUID `gorm:"type:uuid" json:"uploaded_by"`
	UploadedAt time.Time `json:"uploaded_at"`
}

type Comment struct {
	ID         uuid.UUID `gorm:"type:uuid;primary_key" json:"id"`
	EntityID   uuid.UUID `gorm:"type:uuid;not null" json:"entity_id"`
	EntityType string    `gorm:"not null" json:"entity_type"` // test_plan, test_case, etc.
	Content    string    `gorm:"type:text;not null" json:"content"`
	CreatedBy  uuid.UUID `gorm:"type:uuid" json:"created_by"`
	CreatedAt  time.Time `json:"created_at"`

	User User `gorm:"foreignKey:CreatedBy" json:"user,omitempty"`
}

type History struct {
	ID         uuid.UUID `gorm:"type:uuid;primary_key" json:"id"`
	EntityID   uuid.UUID `gorm:"type:uuid;not null" json:"entity_id"`
	EntityType string    `gorm:"not null" json:"entity_type"`
	Action     string    `gorm:"not null" json:"action"` // created, updated, deleted
	Changes    string    `gorm:"type:jsonb" json:"changes"`
	ChangedBy  uuid.UUID `gorm:"type:uuid" json:"changed_by"`
	ChangedAt  time.Time `json:"changed_at"`

	User User `gorm:"foreignKey:ChangedBy" json:"user,omitempty"`
}
