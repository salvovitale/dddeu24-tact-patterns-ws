package domain

import (
	"fmt"
	"log/slog"
)

type PriceSvc struct {
	fracPriceRepo FractionPriceRepository
}

type PriceService interface {
	CalculatePrice(visitor Visitor) (Price, error)
}

func NewPriceSvc(fracPriceRepo FractionPriceRepository) *PriceSvc {
	return &PriceSvc{
		fracPriceRepo: fracPriceRepo,
	}
}

func (p *PriceSvc) CalculatePrice(visitor Visitor) (Price, error) {
	var price float64
	visit := visitor.Visits[len(visitor.Visits)-1]
	for _, f := range visit.Fractions {
		fracPrice, err := p.fracPriceRepo.Get(visitor.City, f.Type, visitor.Type)
		if err != nil {
			return Price{}, fmt.Errorf("error getting price per kg: %w", err)
		}
		price += f.Kg.Float64() * fracPrice.FractionPriceModel(visitor).Value
	}
	slog.Info("price before surcharge", slog.Any("price", price))
	return Price{
		Value:    price,
		Currency: EUR,
	}, nil
}
