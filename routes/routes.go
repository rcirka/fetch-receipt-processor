package routes

import (
	"github.com/gin-gonic/gin"
	"receipt-processor/handlers"
)

// SetupRouter initializes the Gin router and configures all routes
func SetupRouter() *gin.Engine {
	router := gin.Default()

	// Group routes under /receipts
	receiptsGroup := router.Group("/receipts")
	{
		receiptsGroup.POST("/process", handlers.ProcessReceipt)
		receiptsGroup.GET("/:id/points", handlers.GetPoints)
	}

	return router
} 