package infra

import (
	"log"

	"github.com/gin-gonic/gin"

	"cj-store-line-api/internal/adapters/http"
	"cj-store-line-api/internal/adapters/repository"
	"cj-store-line-api/internal/usecase"
)

// Router sets up the HTTP routes
func Router(config *Config) *gin.Engine {
	// Set Gin mode
	gin.SetMode(config.GinMode)

	// Initialize database
	db, err := NewDatabase(config.DatabaseURL)
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	// Run migrations
	if err := AutoMigrate(db); err != nil {
		log.Fatalf("Failed to run migrations: %v", err)
	}

	// Initialize LINE service
	lineService, err := NewLineService(config.LineChannelSecret, config.LineChannelId)
	if err != nil {
		log.Fatalf("Failed to initialize LINE service: %v", err)
	}

	// Initialize repositories
	webhookRepo := repository.NewWebhookEventRepository(db)

	// Initialize services
	webhookService := usecase.NewWebhookService(webhookRepo, lineService)

	// Initialize handlers
	healthHandler := http.NewHealthHandler()
	webhookHandler := http.NewWebhookHandler(webhookService)

	// Setup router
	router := gin.New()

	// Add middleware
	router.Use(http.LoggingMiddleware())
	router.Use(http.CorrelationIDMiddleware())
	router.Use(http.CORSMiddleware())
	router.Use(gin.Recovery())

	// Health check routes
	router.GET("/healthz", healthHandler.Healthz)
	router.GET("/readyz", healthHandler.Readyz)

	// API routes
	api := router.Group("/api/v1")
	{
		// Webhook routes
		webhook := api.Group("/webhook")
		{
			webhook.POST("/line", webhookHandler.HandleWebhook)
		}
	}

	return router
}
