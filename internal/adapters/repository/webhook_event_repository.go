package repository

import (
	"context"
	"time"

	"gorm.io/gorm"

	"cj-store-line-api/internal/domain"
)

// WebhookEventModel represents the database model for webhook events
type WebhookEventModel struct {
	ID        uint   `gorm:"primarykey"`
	Source    string `gorm:"not null;index"`
	EventType string `gorm:"not null;index"`
	UserID    string `gorm:"not null;index"`
	Payload   string `gorm:"type:text"`
	Status    string `gorm:"not null;default:'pending';index"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}

// TableName specifies the table name for GORM
func (WebhookEventModel) TableName() string {
	return "webhook_events"
}

// WebhookEventRepository implements the domain repository interface
type WebhookEventRepository struct {
	db *gorm.DB
}

// NewWebhookEventRepository creates a new webhook event repository
func NewWebhookEventRepository(db *gorm.DB) *WebhookEventRepository {
	return &WebhookEventRepository{
		db: db,
	}
}

// Save saves a webhook event to the database
func (r *WebhookEventRepository) Save(ctx context.Context, event *domain.WebhookEvent) error {
	model := &WebhookEventModel{
		Source:    event.Source,
		EventType: event.EventType,
		UserID:    event.UserID,
		Payload:   event.Payload,
		Status:    event.Status,
		CreatedAt: event.CreatedAt,
		UpdatedAt: event.UpdatedAt,
	}

	if err := r.db.WithContext(ctx).Create(model).Error; err != nil {
		return err
	}

	// Update the domain entity with the generated ID
	event.ID = model.ID
	return nil
}

// FindByID finds a webhook event by ID
func (r *WebhookEventRepository) FindByID(ctx context.Context, id uint) (*domain.WebhookEvent, error) {
	var model WebhookEventModel
	if err := r.db.WithContext(ctx).First(&model, id).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, domain.ErrWebhookEventNotFound
		}
		return nil, err
	}

	return r.modelToDomain(&model), nil
}

// FindByUserID finds webhook events by user ID
func (r *WebhookEventRepository) FindByUserID(ctx context.Context, userID string) ([]*domain.WebhookEvent, error) {
	var models []WebhookEventModel
	if err := r.db.WithContext(ctx).Where("user_id = ?", userID).Order("created_at DESC").Find(&models).Error; err != nil {
		return nil, err
	}

	events := make([]*domain.WebhookEvent, len(models))
	for i, model := range models {
		events[i] = r.modelToDomain(&model)
	}

	return events, nil
}

// UpdateStatus updates the status of a webhook event
func (r *WebhookEventRepository) UpdateStatus(ctx context.Context, id uint, status string) error {
	return r.db.WithContext(ctx).Model(&WebhookEventModel{}).Where("id = ?", id).Update("status", status).Error
}

// modelToDomain converts a database model to a domain entity
func (r *WebhookEventRepository) modelToDomain(model *WebhookEventModel) *domain.WebhookEvent {
	return &domain.WebhookEvent{
		ID:        model.ID,
		Source:    model.Source,
		EventType: model.EventType,
		UserID:    model.UserID,
		Payload:   model.Payload,
		Status:    model.Status,
		CreatedAt: model.CreatedAt,
		UpdatedAt: model.UpdatedAt,
	}
}
