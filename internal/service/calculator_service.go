package service

import (
	"math"

	"food-store-calculator/internal/model"
)

type CalculatorService interface {
	CalculateTotal(req model.OrderRequest) float64
}

type calculatorService struct{}

func NewCalculatorService() CalculatorService {
	return &calculatorService{}
}

func (s *calculatorService) CalculateTotal(req model.OrderRequest) float64 {
	total := 0.0

	for item, qty := range req.Items {
		price, exists := model.MenuPrices[item]
		if !exists || qty <= 0 {
			continue
		}

		if s.isBundleItem(item) {
			// ทุกๆ 2 ชิ้น คิดลด 5% เฉพาะคู่นั้น
			pairs := qty / 2
			remaining := qty % 2

			discountedPairPrice := (price * 2) * 0.95
			total += (float64(pairs) * discountedPairPrice) + (float64(remaining) * price)
		} else {
			total += price * float64(qty)
		}
	}

	// ส่วนลด Member Card 10% จาก Total
	if req.HasMemberCard {
		total = total * 0.90
	}

	return math.Round(total*100) / 100
}

func (s *calculatorService) isBundleItem(item model.Item) bool {
	return item == model.Orange || item == model.Pink || item == model.Green
}
