package service_test

import (
	"testing"

	"food-store-calculator/internal/model"
	"food-store-calculator/internal/service"
)

func TestCalculateTotal(t *testing.T) {
	calcService := service.NewCalculatorService()

	tests := []struct {
		name          string
		request       model.OrderRequest
		expectedTotal float64
	}{
		{
			name: "Desk #1: Red 1 + Green 1 with Member Card",
			request: model.OrderRequest{
				Items:         map[model.Item]int{model.Red: 1, model.Green: 1},
				HasMemberCard: true,
			},
			expectedTotal: 81.0,
		},
		{
			name: "Orange 5 items without Member Card (2 pairs 5% discount)",
			request: model.OrderRequest{
				Items:         map[model.Item]int{model.Orange: 5},
				HasMemberCard: false,
			},
			expectedTotal: 576.0,
		},
		{
			name: "Pink 3 items with Member Card (1 pair 5% discount + Member 10%)",
			request: model.OrderRequest{
				Items:         map[model.Item]int{model.Pink: 3},
				HasMemberCard: true,
			},
			expectedTotal: 208.8,
		},
		{
			name: "Non-discount items (Blue + Yellow + Purple) without Member Card",
			request: model.OrderRequest{
				Items:         map[model.Item]int{model.Blue: 2, model.Yellow: 1, model.Purple: 1},
				HasMemberCard: false,
			},
			expectedTotal: 200.0,
		},
		{
			name: "Invalid item and negative quantity should be ignored",
			request: model.OrderRequest{
				Items: map[model.Item]int{
					model.Item("UnknownItem"): 2,
					model.Red:                 -5,
					model.Green:               0,
				},
				HasMemberCard: false,
			},
			expectedTotal: 0.0,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := calcService.CalculateTotal(tt.request)
			if got != tt.expectedTotal {
				t.Errorf("CalculateTotal() = %v, want %v", got, tt.expectedTotal)
			}
		})
	}
}
