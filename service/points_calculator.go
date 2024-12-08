package service

import (
	"math"
	"receipt-processor/models"
	"regexp"
	"strconv"
	"strings"
	"time"
)

// CalculatePoints calculates the points for a receipt based on the rules
func CalculatePoints(receipt models.Receipt) int {
	points := 0

	// Rule 1: One point for every alphanumeric character in the retailer name
	points += countAlphanumeric(receipt.Retailer)

	// Rule 2: 50 points if the total is a round dollar amount with no cents
	if isRoundDollarAmount(receipt.Total) {
		points += 50
	}

	// Rule 3: 25 points if the total is a multiple of 0.25
	if isMultipleOfQuarter(receipt.Total) {
		points += 25
	}

	// Rule 4: 5 points for every two items on the receipt
	points += (len(receipt.Items) / 2) * 5

	// Rule 5: Points for item descriptions
	points += calculateItemDescriptionPoints(receipt.Items)

	// Rule 6: 6 points if the purchase date is odd
	if isOddDay(receipt.PurchaseDate) {
		points += 6
	}

	// Rule 7: 10 points if the purchase time is between 2:00pm and 4:00pm
	if isTimeBetween14And16(receipt.PurchaseTime) {
		points += 10
	}

	return points
}

// Helper functions

func countAlphanumeric(s string) int {
	alphanumeric := regexp.MustCompile(`[a-zA-Z0-9]`)
	return len(alphanumeric.FindAllString(s, -1))
}

func isRoundDollarAmount(total string) bool {
	if value, err := strconv.ParseFloat(total, 64); err == nil {
		return math.Mod(value, 1.0) == 0
	}
	return false
}

func isMultipleOfQuarter(total string) bool {
	if value, err := strconv.ParseFloat(total, 64); err == nil {
		return math.Mod(value*100, 25) == 0
	}
	return false
}

func calculateItemDescriptionPoints(items []models.Item) int {
	points := 0
	for _, item := range items {
		trimmed := strings.TrimSpace(item.ShortDescription)
		if len(trimmed)%3 == 0 {
			price, _ := strconv.ParseFloat(item.Price, 64)
			points += int(math.Ceil(price * 0.2))
		}
	}
	return points
}

func isOddDay(purchaseDate string) bool {
	date, err := time.Parse("2006-01-02", purchaseDate)
	if err != nil {
		return false
	}
	return date.Day()%2 == 1
}

func isTimeBetween14And16(purchaseTime string) bool {
	t, err := time.Parse("15:04", purchaseTime)
	if err != nil {
		return false
	}
	hour := t.Hour()
	return hour >= 14 && hour < 16
} 