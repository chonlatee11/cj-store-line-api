package http

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// HealthHandler handles health check endpoints
type HealthHandler struct{}

// NewHealthHandler creates a new health handler
func NewHealthHandler() *HealthHandler {
	return &HealthHandler{}
}

// Healthz handles liveness probe
func (h *HealthHandler) Healthz(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"status": "ok",
		"probe":  "liveness",
	})
}

// Readyz handles readiness probe
func (h *HealthHandler) Readyz(c *gin.Context) {
	// TODO: Add actual readiness checks (database connection, external services, etc.)
	c.JSON(http.StatusOK, gin.H{
		"status": "ok",
		"probe":  "readiness",
	})
}
