package types

type PriceResponse struct {
	Tricker string  `json:"ticker"`
	Price   float64 `json:"price"`
}
