package web

type DroppedFraction struct {
	AmountDropped float64 `json:"amount_dropped"`
	FractionType  string  `json:"fraction_type"`
}

type CalculatePriceRequest struct {
	Date             string            `json:"date"`
	DroppedFractions []DroppedFraction `json:"dropped_fractions"`
	PersonID         string            `json:"person_id"`
	VisitID          string            `json:"visit_id"`
}

type CalculatePriceResponse struct {
	PriceAmount   float64 `json:"price_amount"`
	PersonID      string  `json:"person_id"`
	VisitID       string  `json:"visit_id"`
	PriceCurrency string  `json:"price_currency"`
}
