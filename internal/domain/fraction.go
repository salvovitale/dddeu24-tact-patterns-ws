package domain

import (
	"fmt"
	"strings"
	"time"
)

type City string

const (
	CityPineville City = "Pineville"
	CityOakCity   City = "Oak City"
)

func ParseCity(s string) (City, error) {
	trimmedLower := strings.ToLower(strings.TrimSpace(s))
	replaced := strings.Replace(trimmedLower, " ", "_", -1)
	switch replaced {
	case "pineville":
		return CityPineville, nil
	case "oak_city":
		return CityOakCity, nil
	default:
		return "", fmt.Errorf("unknown city: %s", s)
	}
}

type FractionPriceModel func(Visitor) Price

type FractionType string

const (
	FractionTypeGreenWaste        FractionType = "green_waste"
	FractionTypeConstructionWaste FractionType = "construction_waste"
)

func ParseFractionType(s string) (FractionType, error) {
	trimmedLower := strings.ToLower(strings.TrimSpace(s))
	replaced := strings.Replace(trimmedLower, " ", "_", -1)

	switch replaced {
	case "green_waste":
		return FractionTypeGreenWaste, nil
	case "construction_waste":
		return FractionTypeConstructionWaste, nil
	default:
		return "", fmt.Errorf("unknown fraction type: %s", s)
	}
}

type Fraction struct {
	Type FractionType
	Kg   FractionWeight
}

type FractionWeight float64

func (w FractionWeight) Float64() float64 {
	return float64(w)
}

func ParseWeight(f float64) (FractionWeight, error) {
	if f < 0 {
		return 0, fmt.Errorf("weight must be positive")
	}
	return FractionWeight(f), nil
}

type FractionPrice struct {
	FractionType       FractionType
	City               City
	VisitorType        VisitorType
	FractionPriceModel FractionPriceModel
}

var (
	ErrFractionPriceNotFound = fmt.Errorf("fraction price not found")
)

type FractionPriceRepository interface {
	Get(city City, fractionType FractionType, visitorType VisitorType) (*FractionPrice, error)
}

func ExtractChargeAfterVisit(pricePerKg Price, visitCount uint16) FractionPriceModel {
	return func(v Visitor) Price {
		if v.VisitCounter.Counter >= visitCount {
			return Price{
				Value:    pricePerKg.Value * 1.05,
				Currency: EUR,
			}
		}
		return pricePerKg
	}
}

func CumulatedWeightOverPeriod(weightThreshold FractionWeight, p1, p2 Price, fractionType FractionType, period time.Duration) FractionPriceModel {
	return func(v Visitor) Price {
		totalWeight := v.CumulativeWeightOverPeriod(period, fractionType).Float64()
		lastVisit := v.LastVisit()
		lastVisitWeight := lastVisit.WeightOf(fractionType).Float64()
		totalWeightBeforeLast := totalWeight - lastVisitWeight

		if totalWeight <= weightThreshold.Float64() {
			return Price{
				Value:    p1.Value,
				Currency: EUR,
			}
		}
		if totalWeightBeforeLast > weightThreshold.Float64() {
			return Price{
				Value:    p2.Value,
				Currency: EUR,
			}
		}
		dV1 := weightThreshold.Float64() - totalWeightBeforeLast
		dV2 := totalWeight - weightThreshold.Float64()
		return Price{
			Value:    (p1.Value*dV1 + p2.Value*dV2) / lastVisitWeight,
			Currency: EUR,
		}
	}
}
