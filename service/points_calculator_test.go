package service

import (
	"receipt-processor/models"
	"testing"
)

func TestCalculatePoints(t *testing.T) {
	tests := []struct {
		name     string
		receipt  models.Receipt
		expected int
	}{
		{
			name: "One point for every alphanumeric character in the retailer name",
			receipt: models.Receipt{
				Retailer:     "Target",
				PurchaseDate: "2022-01-02",
				PurchaseTime: "00:00",
				Items:        []models.Item{},
				Total:        "0.01",
			},
			expected: 6,
		},
		{
			name: "50 points if the total is a round dollar amount with no cents / 25 points if the total is a multiple of 0.25",
			receipt: models.Receipt{
				Retailer:     "",
				PurchaseDate: "2022-01-02",
				PurchaseTime: "00:00",
				Items:        []models.Item{},
				Total:        "1",
			},
			expected: 75,
		},
		{
			name: "5 points for every two items on the receipt",
			receipt: models.Receipt{
				Retailer:     "",
				PurchaseDate: "2022-01-02",
				PurchaseTime: "00:00",
				Items: []models.Item{
					{ShortDescription: "a", Price: "6.49"},
					{ShortDescription: "b", Price: "12.25"},
				},
				Total: "1.01",
			},
			expected: 5,
		},
		{
			name: "Example Receipt 1",
			receipt: models.Receipt{
				Retailer:     "Target",
				PurchaseDate: "2022-01-01",
				PurchaseTime: "13:01",
				Items: []models.Item{
					{ShortDescription: "Mountain Dew 12PK", Price: "6.49"},
					{ShortDescription: "Emils Cheese Pizza", Price: "12.25"},
					{ShortDescription: "Knorr Creamy Chicken", Price: "1.26"},
					{ShortDescription: "Doritos Nacho Cheese", Price: "3.35"},
					{ShortDescription: "   Klarbrunn 12-PK 12 FL OZ  ", Price: "12.00"},
				},
				Total: "35.35",
			},
			expected: 28,
		}, {
			name: "Example Receipt 1",
			receipt: models.Receipt{
				Retailer:     "Target",
				PurchaseDate: "2022-01-01",
				PurchaseTime: "13:01",
				Items: []models.Item{
					{ShortDescription: "Mountain Dew 12PK", Price: "6.49"},
					{ShortDescription: "Emils Cheese Pizza", Price: "12.25"},
					{ShortDescription: "Knorr Creamy Chicken", Price: "1.26"},
					{ShortDescription: "Doritos Nacho Cheese", Price: "3.35"},
					{ShortDescription: "   Klarbrunn 12-PK 12 FL OZ  ", Price: "12.00"},
				},
				Total: "35.35",
			},
			expected: 28,
		},
		{
			name: "Example Receipt 2",
			receipt: models.Receipt{
				Retailer:     "M&M Corner Market",
				PurchaseDate: "2022-03-20",
				PurchaseTime: "14:33",
				Items: []models.Item{
					{ShortDescription: "Gatorade", Price: "2.25"},
					{ShortDescription: "Gatorade", Price: "2.25"},
					{ShortDescription: "Gatorade", Price: "2.25"},
					{ShortDescription: "Gatorade", Price: "2.25"},
				},
				Total: "9.00",
			},
			expected: 109,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			points := CalculatePoints(tt.receipt)
			if points != tt.expected {
				t.Errorf("CalculatePoints() = %v, want %v", points, tt.expected)
			}
		})
	}
}
