package main

import (
	"fmt"
	"log"

	"cj-store-line-api/internal/infra"
)

func main() {
	// Load configuration
	config, err := infra.LoadConfig()
	if err != nil {
		log.Fatalf("Failed to load config: %v", err)
	}

	// Setup router
	router := infra.Router(config)

	// Start server
	addr := fmt.Sprintf(":%s", config.Port)
	log.Printf("Starting server on %s", addr)
	if err := router.Run(addr); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
