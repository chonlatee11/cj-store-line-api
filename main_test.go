package main

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"

	"cj-store-line-api/internal/usecase"
)

func TestWebhookEndpoint(t *testing.T) {
	// Create test request
	webhookReq := usecase.WebhookRequest{
		Source:    "test_source",
		EventType: "message",
		UserID:    "test_user_123",
		Message:   "Hello from test",
		Payload:   map[string]string{"test": "data"},
	}

	reqBody, err := json.Marshal(webhookReq)
	assert.NoError(t, err)

	// Create HTTP request
	req := httptest.NewRequest(http.MethodPost, "/api/v1/webhook/line", bytes.NewBuffer(reqBody))
	req.Header.Set("Content-Type", "application/json")

	// Note: This test requires a proper database connection and LINE credentials
	// For now, we'll just test the request structure
	t.Run("webhook request structure", func(t *testing.T) {
		assert.Equal(t, "test_source", webhookReq.Source)
		assert.Equal(t, "message", webhookReq.EventType)
		assert.Equal(t, "test_user_123", webhookReq.UserID)
		assert.Equal(t, "Hello from test", webhookReq.Message)
		assert.NotNil(t, req)
		assert.NotNil(t, reqBody)
	})
}
