package http

import (
	"context"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"

	"cj-store-line-api/internal/usecase"
)

// WebhookHandler handles HTTP requests for webhook endpoints
type WebhookHandler struct {
	service *usecase.WebhookService
}

// NewWebhookHandler creates a new webhook handler
func NewWebhookHandler(service *usecase.WebhookService) *WebhookHandler {
	return &WebhookHandler{
		service: service,
	}
}

// HandleWebhook handles incoming webhook requests
func (h *WebhookHandler) HandleWebhook(c *gin.Context) {
	var req usecase.WebhookRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":         "invalid request",
			"details":       err.Error(),
			"correlationId": c.GetString("correlationId"),
		})
		return
	}

	ctx, cancel := context.WithTimeout(c.Request.Context(), 30*time.Second)
	defer cancel()

	if err := h.service.ProcessWebhook(ctx, req); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error":         "failed to process webhook",
			"details":       err.Error(),
			"correlationId": c.GetString("correlationId"),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status":        "ok",
		"message":       "webhook processed successfully",
		"correlationId": c.GetString("correlationId"),
	})
}
