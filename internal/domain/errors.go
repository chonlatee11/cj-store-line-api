package domain

import "errors"

// Domain errors
var (
	ErrWebhookEventNotFound = errors.New("webhook event not found")
	ErrInvalidWebhookEvent  = errors.New("invalid webhook event")
	ErrLineMessageFailed    = errors.New("failed to send LINE message")
	ErrInvalidUserID        = errors.New("invalid user ID")
)
