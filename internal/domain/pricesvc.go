package domain

import (
	"fmt"
	"log/slog"
)

type PriceSvc struct {
	fracPriceRepo FractionPriceRepository
}

type PriceService interface {
	CalculatePrice(visit Visit, visitor Visitor) (float64, error)
}

func NewPriceSvc(fracPriceRepo FractionPriceRepository) *PriceSvc {
	return &PriceSvc{
		fracPriceRepo: fracPriceRepo,
	}
}

func (p *PriceSvc) CalculatePrice(visit Visit, visitor Visitor) (float64, error) {
	var price float64
	for _, f := range visit.Fractions {
		fracPrice, err := p.fracPriceRepo.Get(visitor.City, visitor.Type, f.Type)
		if err != nil {
			return 0, fmt.Errorf("error getting price per kg: %w", err)
		}
		price += f.Kg.Float64() * fracPrice.PricePerKg.Value
	}

	slog.Info("price before surcharge", slog.Any("price", price))

	price *= surChargePolicy(visitor)
	return price, nil
}

func surChargePolicy(visitor Visitor) float64 {
	if visitor.VisitCounter.Counter >= 3 {
		return 1.05
	}
	return 1.0
}
