package handlers

import (
	"net/http"
	"receipt-processor/models"
	"receipt-processor/service"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

// Store receipts in memory using a map
var receiptStore = make(map[string]models.Receipt)

// ProcessReceipt handles the POST request to process a new receipt
func ProcessReceipt(c *gin.Context) {
	var receipt models.Receipt
	if err := c.ShouldBindJSON(&receipt); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid receipt format"})
		return
	}

	// Generate a unique ID for the receipt
	id := uuid.New().String()
	
	// Store the receipt
	receiptStore[id] = receipt

	c.JSON(http.StatusOK, gin.H{"id": id})
}

// GetPoints handles the GET request to retrieve points for a receipt
func GetPoints(c *gin.Context) {
	id := c.Param("id")
	
	receipt, exists := receiptStore[id]
	if !exists {
		c.JSON(http.StatusNotFound, gin.H{"error": "Receipt not found"})
		return
	}

	
	points := service.CalculatePoints(receipt)
	c.JSON(http.StatusOK, gin.H{"points": points})
} 