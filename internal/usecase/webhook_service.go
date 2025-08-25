package usecase

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	"cj-store-line-api/internal/domain"
)

// WebhookRequest represents the incoming webhook request
type WebhookRequest struct {
	Source    string      `json:"source" binding:"required"`
	EventType string      `json:"event_type" binding:"required"`
	UserID    string      `json:"user_id" binding:"required"`
	Message   string      `json:"message,omitempty"`
	Payload   interface{} `json:"payload,omitempty"`
}

// WebhookService handles webhook business logic
type WebhookService struct {
	webhookRepo domain.WebhookEventRepository
	lineService domain.LineService
}

// NewWebhookService creates a new webhook service
func NewWebhookService(webhookRepo domain.WebhookEventRepository, lineService domain.LineService) *WebhookService {
	return &WebhookService{
		webhookRepo: webhookRepo,
		lineService: lineService,
	}
}

// ProcessWebhook processes incoming webhook events
func (s *WebhookService) ProcessWebhook(ctx context.Context, req WebhookRequest) error {
	// Validate request
	if err := s.validateWebhookRequest(req); err != nil {
		return fmt.Errorf("validation failed: %w", err)
	}

	// Convert payload to JSON string
	payloadJSON, err := json.Marshal(req.Payload)
	if err != nil {
		return fmt.Errorf("failed to marshal payload: %w", err)
	}

	// Create webhook event
	event := &domain.WebhookEvent{
		Source:    req.Source,
		EventType: req.EventType,
		UserID:    req.UserID,
		Payload:   string(payloadJSON),
		Status:    "pending",
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	// Save webhook event
	if err := s.webhookRepo.Save(ctx, event); err != nil {
		return fmt.Errorf("failed to save webhook event: %w", err)
	}

	// Process the webhook event
	if err := s.handleWebhookEvent(ctx, event, req); err != nil {
		// Update status to failed
		if updateErr := s.webhookRepo.UpdateStatus(ctx, event.ID, "failed"); updateErr != nil {
			return fmt.Errorf("failed to update status and process event: %v, %w", updateErr, err)
		}
		return fmt.Errorf("failed to process webhook event: %w", err)
	}

	// Update status to processed
	if err := s.webhookRepo.UpdateStatus(ctx, event.ID, "processed"); err != nil {
		return fmt.Errorf("failed to update status to processed: %w", err)
	}

	return nil
}

// handleWebhookEvent processes the webhook event based on its type
func (s *WebhookService) handleWebhookEvent(ctx context.Context, event *domain.WebhookEvent, req WebhookRequest) error {
	switch req.EventType {
	case "message":
		return s.handleMessageEvent(ctx, req)
	case "order_update":
		return s.handleOrderUpdateEvent(ctx, req)
	case "promotion":
		return s.handlePromotionEvent(ctx, req)
	default:
		// For unknown event types, just send a generic message
		return s.lineService.SendTextMessage(ctx, req.UserID, "Thank you for your interaction!")
	}
}

// handleMessageEvent handles message-type webhook events
func (s *WebhookService) handleMessageEvent(ctx context.Context, req WebhookRequest) error {
	message := "Hello! We received your message."
	if req.Message != "" {
		message = fmt.Sprintf("Thank you for your message: %s", req.Message)
	}

	return s.lineService.SendTextMessage(ctx, req.UserID, message)
}

// handleOrderUpdateEvent handles order update webhook events
func (s *WebhookService) handleOrderUpdateEvent(ctx context.Context, req WebhookRequest) error {
	message := "Your order has been updated! Please check your order status."
	return s.lineService.SendTextMessage(ctx, req.UserID, message)
}

// handlePromotionEvent handles promotion webhook events
func (s *WebhookService) handlePromotionEvent(ctx context.Context, req WebhookRequest) error {
	message := "🎉 New promotion available! Check out our latest offers."
	return s.lineService.SendTextMessage(ctx, req.UserID, message)
}

// validateWebhookRequest validates the incoming webhook request
func (s *WebhookService) validateWebhookRequest(req WebhookRequest) error {
	if req.Source == "" {
		return domain.ErrInvalidWebhookEvent
	}
	if req.EventType == "" {
		return domain.ErrInvalidWebhookEvent
	}
	if req.UserID == "" {
		return domain.ErrInvalidUserID
	}
	return nil
}
