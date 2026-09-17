package handler

import (
	"encoding/json"
	"food-store-calculator/internal/model"
	"food-store-calculator/internal/service"
	"net/http"
)

type CalculatorHandler struct {
	calcService service.CalculatorService
}

func NewCalculatorHandler(calcService service.CalculatorService) *CalculatorHandler {
	return &CalculatorHandler{
		calcService: calcService,
	}
}

// Checkout ทำหน้าที่รับ OrderRequest แล้วเรียก Service คำนวณราคากลับมา
func (h *CalculatorHandler) Checkout(req model.OrderRequest) model.OrderResponse {
	total := h.calcService.CalculateTotal(req)
	return model.OrderResponse{
		Total: total,
	}
}

// HandleCheckout เป็น HTTP Handler สำหรับรับ POST /checkout
func (h *CalculatorHandler) HandleCheckout(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var req model.OrderRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	res := h.Checkout(req)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(res)
}
