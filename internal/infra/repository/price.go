package infra_repository

import (
	"sync"
	"time"

	"github.com/salvovitale/dddeu24-tact-patterns-ws/internal/domain"
)

type repoKey struct {
	City         domain.City
	FractionType domain.FractionType
	VisitorType  domain.VisitorType
}

type FractionPriceInMemoryRepository struct {
	prices map[repoKey]domain.FractionPriceModel
	mu     sync.RWMutex
}

func NewFractionPriceInMemoryRepository() *FractionPriceInMemoryRepository {
	return &FractionPriceInMemoryRepository{
		prices: initPrices(),
		mu:     sync.RWMutex{},
	}
}

func (r *FractionPriceInMemoryRepository) Get(city domain.City, fractionType domain.FractionType, visitorType domain.VisitorType) (*domain.FractionPrice, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()
	key := repoKey{
		City:         city,
		FractionType: fractionType,
		VisitorType:  visitorType,
	}
	p, ok := r.prices[key]
	if !ok {
		return nil, domain.ErrFractionPriceNotFound
	}
	return &domain.FractionPrice{
		City:               city,
		FractionType:       fractionType,
		VisitorType:        visitorType,
		FractionPriceModel: p,
	}, nil
}

func (r *FractionPriceInMemoryRepository) Save(frPrice *domain.FractionPrice) error {
	r.mu.Lock()
	defer r.mu.Unlock()
	key := repoKey{
		City:         frPrice.City,
		FractionType: frPrice.FractionType,
		VisitorType:  frPrice.VisitorType,
	}
	r.prices[key] = frPrice.FractionPriceModel
	return nil
}

/*
Price lists
* Pineville:
	* Private visitor:
  		* Green waste: 0.10 // extract charge after 3 visits
  		* Construction waste: 0.15 // extract charge after 3 visits
	* Business visitor:
		* Green waste: 0.12
		* Construction waste: 0.13
* Oak City:
	* Private visitor:
		* Green waste: 0.08 // extract charge after 3 visits
		* Construction waste: 0.19 // extract charge after 3 visits
	* Business visitor:
		* Green waste: 0.08
		* Construction waste:  <= 1000kg  0.21
		* Construction waste:  > 1000kg  0.29
*/

func initPrices() map[repoKey]domain.FractionPriceModel {
	return map[repoKey]domain.FractionPriceModel{
		{City: domain.CityPineville, FractionType: domain.FractionTypeGreenWaste, VisitorType: domain.VisitorTypePrivate}:        domain.ExtractChargeAfterVisit(domain.Price{Value: 0.1, Currency: domain.EUR}, 3),
		{City: domain.CityPineville, FractionType: domain.FractionTypeConstructionWaste, VisitorType: domain.VisitorTypePrivate}: domain.ExtractChargeAfterVisit(domain.Price{Value: 0.15, Currency: domain.EUR}, 3),
		{City: domain.CityPineville, FractionType: domain.FractionTypeGreenWaste, VisitorType: domain.VisitorTypeBusiness}: func(v domain.Visitor) domain.Price {
			return domain.Price{Value: 0.12, Currency: domain.EUR}
		},
		{City: domain.CityPineville, FractionType: domain.FractionTypeConstructionWaste, VisitorType: domain.VisitorTypeBusiness}: func(v domain.Visitor) domain.Price {
			return domain.Price{Value: 0.13, Currency: domain.EUR}
		},
		{City: domain.CityOakCity, FractionType: domain.FractionTypeGreenWaste, VisitorType: domain.VisitorTypePrivate}:        domain.ExtractChargeAfterVisit(domain.Price{Value: 0.08, Currency: domain.EUR}, 3),
		{City: domain.CityOakCity, FractionType: domain.FractionTypeConstructionWaste, VisitorType: domain.VisitorTypePrivate}: domain.ExtractChargeAfterVisit(domain.Price{Value: 0.19, Currency: domain.EUR}, 3),
		{City: domain.CityOakCity, FractionType: domain.FractionTypeGreenWaste, VisitorType: domain.VisitorTypeBusiness}: func(v domain.Visitor) domain.Price {
			return domain.Price{Value: 0.08, Currency: domain.EUR}
		},
		{City: domain.CityOakCity, FractionType: domain.FractionTypeConstructionWaste, VisitorType: domain.VisitorTypeBusiness}: domain.CumulatedWeightOverPeriod(domain.FractionWeight(1000), domain.Price{Value: 0.21, Currency: domain.EUR}, domain.Price{Value: 0.29, Currency: domain.EUR}, domain.FractionTypeConstructionWaste, time.Duration(24*time.Hour*356)),
	}
}
