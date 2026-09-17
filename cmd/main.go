package main

import (
	"fmt"

	"food-store-calculator/internal/handler"
	"food-store-calculator/internal/model"
	"food-store-calculator/internal/service"
)

func main() {
	// Initialize Dependencies
	calcService := service.NewCalculatorService()
	calcHandler := handler.NewCalculatorHandler(calcService)

	// Test Case 1: Red + Green with Member Card
	req1 := model.OrderRequest{
		Items: map[model.Item]int{
			model.Red:   1,
			model.Green: 1,
		},
		HasMemberCard: true,
	}
	res1 := calcHandler.Checkout(req1)
	fmt.Printf("Desk #1 Total: %.2f THB\n", res1.Total)

	// Test Case 2: Orange 5 items
	req2 := model.OrderRequest{
		Items: map[model.Item]int{
			model.Orange: 5,
		},
		HasMemberCard: false,
	}
	res2 := calcHandler.Checkout(req2)
	fmt.Printf("Orange 5 items Total: %.2f THB\n", res2.Total)
}
