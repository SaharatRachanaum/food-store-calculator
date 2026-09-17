package model

type Item string

const (
	Red    Item = "Red"
	Green  Item = "Green"
	Blue   Item = "Blue"
	Yellow Item = "Yellow"
	Pink   Item = "Pink"
	Purple Item = "Purple"
	Orange Item = "Orange"
)

// MenuPrices เก็บราคาสินค้าตามโจทย์
var MenuPrices = map[Item]float64{
	Red:    50.0,
	Green:  40.0,
	Blue:   30.0,
	Yellow: 50.0,
	Pink:   80.0,
	Purple: 90.0,
	Orange: 120.0,
}

type OrderRequest struct {
	Items         map[Item]int `json:"items"`
	HasMemberCard bool         `json:"has_member_card"`
}

type OrderResponse struct {
	Total float64 `json:"total"`
}
