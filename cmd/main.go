package main

import (
	"fmt"
	"log"
	"net/http"

	"food-store-calculator/internal/handler"
	"food-store-calculator/internal/model"
	"food-store-calculator/internal/service"
)

func main() {
	// Initialize Dependencies
	calcService := service.NewCalculatorService()
	calcHandler := handler.NewCalculatorHandler(calcService)

	// --- Console Example Runs ---
	req1 := model.OrderRequest{
		Items:         map[model.Item]int{model.Red: 1, model.Green: 1},
		HasMemberCard: true,
	}
	res1 := calcHandler.Checkout(req1)
	fmt.Printf("Desk #1 Total: %.2f THB\n", res1.Total)

	req2 := model.OrderRequest{
		Items:         map[model.Item]int{model.Orange: 5},
		HasMemberCard: false,
	}
	res2 := calcHandler.Checkout(req2)
	fmt.Printf("Orange 5 items Total: %.2f THB\n", res2.Total)

	// --- HTTP API Server Setup ---
	http.HandleFunc("/checkout", calcHandler.HandleCheckout)

	log.Println("Server is running on http://localhost:8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatalf("Server failed to start: %v", err)
	}
}
