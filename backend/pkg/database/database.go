package database

import (
	"github.com/AntVerkh/test-management-system/internal/domain"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func NewPostgresDB(dsn string) (*gorm.DB, error) {
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	return db, nil
}

func Migrate(db *gorm.DB) error {
	return db.AutoMigrate(
		&domain.User{},
		&domain.Project{},
		&domain.TestPlan{},
		&domain.TestStrategy{},
		&domain.Checklist{},
		&domain.ChecklistItem{},
		&domain.TestCase{},
		&domain.TestStep{},
		&domain.TestRun{},
		&domain.TestResult{},
		&domain.Attachment{},
		&domain.Comment{},
		&domain.History{},
	)
}
