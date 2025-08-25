package infra

import (
	"context"
	"fmt"

	"github.com/line/line-bot-sdk-go/v7/linebot"

	"cj-store-line-api/internal/domain"
)

// LineService implements the domain.LineService interface
type LineService struct {
	client *linebot.Client
}

// NewLineService creates a new LINE service
func NewLineService(channelSecret, channelAccessToken string) (*LineService, error) {
	client, err := linebot.New(channelSecret, channelAccessToken)
	if err != nil {
		return nil, fmt.Errorf("failed to create LINE bot client: %w", err)
	}

	return &LineService{
		client: client,
	}, nil
}

// SendTextMessage sends a text message to a LINE user
func (s *LineService) SendTextMessage(ctx context.Context, userID, message string) error {
	textMessage := linebot.NewTextMessage(message)
	_, err := s.client.PushMessage(userID, textMessage).Do()
	if err != nil {
		return fmt.Errorf("failed to send text message: %w", domain.ErrLineMessageFailed)
	}
	return nil
}

// SendFlexMessage sends a flex message to a LINE user
func (s *LineService) SendFlexMessage(ctx context.Context, userID string, flexMessage interface{}) error {
	// Convert flexMessage to linebot.FlexContainer
	// This is a simplified implementation - you would need to properly convert based on your flex message structure
	flexContainer, ok := flexMessage.(linebot.FlexContainer)
	if !ok {
		return fmt.Errorf("invalid flex message format: %w", domain.ErrInvalidWebhookEvent)
	}

	message := linebot.NewFlexMessage("Flex Message", flexContainer)
	_, err := s.client.PushMessage(userID, message).Do()
	if err != nil {
		return fmt.Errorf("failed to send flex message: %w", domain.ErrLineMessageFailed)
	}
	return nil
}

// SendImageMessage sends an image message to a LINE user
func (s *LineService) SendImageMessage(ctx context.Context, userID, originalURL, previewURL string) error {
	imageMessage := linebot.NewImageMessage(originalURL, previewURL)
	_, err := s.client.PushMessage(userID, imageMessage).Do()
	if err != nil {
		return fmt.Errorf("failed to send image message: %w", domain.ErrLineMessageFailed)
	}
	return nil
}
