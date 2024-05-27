package infra_repository

import (
	"fmt"
	"sync"

	"github.com/salvovitale/dddeu24-tact-patterns-ws/internal/domain"
)

type repoKey struct {
	City         domain.City
	VisitorType  domain.VisitorType
	FractionType domain.FractionType
}

type FractionPriceInMemoryRepository struct {
	prices map[repoKey]domain.Price
	mu     sync.RWMutex
}

func NewFractionPriceInMemoryRepository() *FractionPriceInMemoryRepository {
	return &FractionPriceInMemoryRepository{
		prices: initPrices(),
		mu:     sync.RWMutex{},
	}
}

func (r *FractionPriceInMemoryRepository) Get(city domain.City, visitorType domain.VisitorType, fractionType domain.FractionType) (*domain.FractionPrice, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()
	key := repoKey{
		City:         city,
		VisitorType:  visitorType,
		FractionType: fractionType,
	}

	fmt.Println("key", key)
	p, ok := r.prices[key]
	if !ok {
		return nil, domain.ErrFractionPriceNotFound
	}
	return &domain.FractionPrice{
		City:         city,
		VisitorType:  visitorType,
		FractionType: fractionType,
		PricePerKg:   p,
	}, nil
}

func (r *FractionPriceInMemoryRepository) Save(frPrice *domain.FractionPrice) error {
	r.mu.Lock()
	defer r.mu.Unlock()
	key := repoKey{
		City:         frPrice.City,
		VisitorType:  frPrice.VisitorType,
		FractionType: frPrice.FractionType,
	}
	r.prices[key] = frPrice.PricePerKg
	return nil
}

func initPrices() map[repoKey]domain.Price {
	return map[repoKey]domain.Price{
		{City: domain.CityPineville, VisitorType: domain.VisitorTypePrivate, FractionType: domain.FractionTypeGreenWaste}:         domain.Price{Value: 0.1},
		{City: domain.CityPineville, VisitorType: domain.VisitorTypePrivate, FractionType: domain.FractionTypeConstructionWaste}:  domain.Price{Value: 0.15},
		{City: domain.CityOakCity, VisitorType: domain.VisitorTypePrivate, FractionType: domain.FractionTypeGreenWaste}:           domain.Price{Value: 0.08},
		{City: domain.CityOakCity, VisitorType: domain.VisitorTypePrivate, FractionType: domain.FractionTypeConstructionWaste}:    domain.Price{Value: 0.19},
		{City: domain.CityPineville, VisitorType: domain.VisitorTypeBusiness, FractionType: domain.FractionTypeGreenWaste}:        domain.Price{Value: 0.12},
		{City: domain.CityPineville, VisitorType: domain.VisitorTypeBusiness, FractionType: domain.FractionTypeConstructionWaste}: domain.Price{Value: 0.13},
		{City: domain.CityOakCity, VisitorType: domain.VisitorTypeBusiness, FractionType: domain.FractionTypeGreenWaste}:          domain.Price{Value: 0.08},
		{City: domain.CityOakCity, VisitorType: domain.VisitorTypeBusiness, FractionType: domain.FractionTypeConstructionWaste}:   domain.Price{Value: 0.19},
	}
}
