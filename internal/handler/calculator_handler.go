package handler

import (
	"food-store-calculator/internal/model"
	"food-store-calculator/internal/service"
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
