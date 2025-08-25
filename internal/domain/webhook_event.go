package domain

import (
	"context"
	"time"
)

// WebhookEvent represents a webhook event entity
type WebhookEvent struct {
	ID        uint      `json:"id"`
	Source    string    `json:"source"`
	EventType string    `json:"event_type"`
	UserID    string    `json:"user_id"`
	Payload   string    `json:"payload"`
	Status    string    `json:"status"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

// WebhookEventRepository defines the interface for webhook event storage
type WebhookEventRepository interface {
	Save(ctx context.Context, event *WebhookEvent) error
	FindByID(ctx context.Context, id uint) (*WebhookEvent, error)
	FindByUserID(ctx context.Context, userID string) ([]*WebhookEvent, error)
	UpdateStatus(ctx context.Context, id uint, status string) error
}
