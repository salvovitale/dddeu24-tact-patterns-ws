package domain

import (
	"fmt"
	"strings"
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
	FractionType FractionType
	City         City
	PricePerKg   Price
	VisitorType  VisitorType
}

var (
	ErrFractionPriceNotFound = fmt.Errorf("fraction price not found")
)

type FractionPriceRepository interface {
	Get(city City, visitorType VisitorType, fractionType FractionType) (*FractionPrice, error)
}
