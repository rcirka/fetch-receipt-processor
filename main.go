package main

import (
	"log"
	"receipt-processor/routes"
)

func main() {
	// Initialize and start the HTTP server
	router := routes.SetupRouter()
	
	if err := router.Run(":8080"); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
