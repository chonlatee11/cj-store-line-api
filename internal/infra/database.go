package infra

import (
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	"cj-store-line-api/internal/adapters/repository"
)

// NewDatabase creates a new database connection
func NewDatabase(databaseURL string) (*gorm.DB, error) {
	db, err := gorm.Open(postgres.Open(databaseURL), &gorm.Config{})
	if err != nil {
		return nil, fmt.Errorf("failed to connect to database: %w", err)
	}

	return db, nil
}

// AutoMigrate runs database migrations
func AutoMigrate(db *gorm.DB) error {
	return db.AutoMigrate(
		&repository.WebhookEventModel{},
	)
}
