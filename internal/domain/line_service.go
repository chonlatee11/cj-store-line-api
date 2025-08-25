package domain

import "context"

// LineService defines the interface for LINE Official Account integration
type LineService interface {
	SendTextMessage(ctx context.Context, userID, message string) error
	SendFlexMessage(ctx context.Context, userID string, flexMessage interface{}) error
	SendImageMessage(ctx context.Context, userID, originalURL, previewURL string) error
}
