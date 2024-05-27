package domain

import (
	"fmt"
	"log/slog"
)

type PriceSvc struct {
	visitorRepo VisitorRepository
}

type PriceService interface {
	CalculatePrice(visit Visit) (float64, error)
}

func NewPriceSvc(visitorRepo VisitorRepository) *PriceSvc {
	return &PriceSvc{
		visitorRepo: visitorRepo,
	}
}

func (p *PriceSvc) CalculatePrice(visit Visit) (float64, error) {
	visitor, err := p.visitorRepo.Get(visit.VisitorID)
	if err != nil {
		if err == ErrVisitorNotFound {
			visitor, err = NewVisitor(visit)
			if err != nil {
				return 0, fmt.Errorf("error creating visitor: %w", err)
			}
		} else {
			return 0, fmt.Errorf("error getting visitor: %w", err)
		}
	} else {
		visitor.RegisterVisit(visit.Date)
	}

	var price float64
	for _, f := range visit.Fractions {
		pricePerCity, err := f.Type.PricePerKg(visit.City)
		if err != nil {
			return 0, fmt.Errorf("error getting price per kg: %w", err)
		}
		price += f.Kg.Float64() * pricePerCity.Float64()
	}

	slog.Info("price before surcharge", slog.Any("price", price))

	price *= surChargePolicy(visitor)

	err = p.visitorRepo.Save(visitor)
	if err != nil {
		return 0, fmt.Errorf("error saving visitor: %w", err)
	}
	return price, nil
}

func surChargePolicy(visitor Visitor) float64 {
	if visitor.VisitCounter.Counter >= 3 {
		return 1.1
	}
	return 1.0
}
